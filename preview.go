package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/websocket"
	"gopkg.in/fsnotify.v1"
)

var (
	previewScript = `<script>
new WebSocket('ws://localhost:8080/ws').onmessage = function(e) {
	window.location = window.location;
};
</script>`
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
	for {
		select {
		case <-refresh:
			c.WriteMessage(websocket.TextMessage, []byte("refresh"))
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
				log.Println(event.String())
				w.Close()
				runGenerator()
				w = initWatch()
				refresh <- true
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
	go watch(initWatch())
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ws", websocketHandler)
	log.Print("server listening at 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
