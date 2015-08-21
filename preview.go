package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/websocket"
	"gopkg.in/fsnotify.v1"
)

var (
	refresh chan bool
)

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	for {
		select {
		case <-refresh:
		default:
			goto UpgradeWebsocket
		}
	}
UpgradeWebsocket:
	c, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}).Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	done := make(chan bool)
	go func() {
		for {
			if _, _, err := c.NextReader(); err != nil {
				c.Close()
				done <- true
				break
			}
		}
	}()
	log.Printf("connected with websocket")
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			c.WriteMessage(websocket.TextMessage, []byte("heartbeat"))
		case <-refresh:
			c.WriteMessage(websocket.TextMessage, []byte("refresh"))
			log.Printf("refresh sent")
		case <-done:
			return
		}
	}
}

func watch(w *fsnotify.Watcher) {
	for {
		select {
		case event := <-w.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Printf("%s triggered refresh", event.Name)
				start := time.Now()
				runGenerator()
				log.Printf("generator took %v", time.Since(start))
				refresh <- true
				return
			}
		case err := <-w.Errors:
			log.Println(err)
		}
	}
}

func initWatch() *fsnotify.Watcher {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if ext == ".md" || ext == ".html" || ext == ".css" {
			w.Add(path)
		}
		return nil
	})
	return w
}

func runPreview() {
	refresh = make(chan bool, 100)
	go func() {
		for {
			w := initWatch()
			watch(w)
			go func() {
				w.Close()
			}()
		}
	}()
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ws", websocketHandler)
	log.Print("server listening at 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
