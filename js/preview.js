(function() {
	var lastHeartbeat = Date.now();
	function checkHeartbeat() {
		var now = Date.now();
		if (now - lastHeartbeat > 5000) {
			ping(function(online) {
				if (online) {
					window.location.reload(true);
				} else {
					document.getElementById('offline-label').classList.remove('invisible');
					window.setTimeout(checkHeartbeat, 5000);
				}
			});
		} else {
			window.setTimeout(checkHeartbeat, 5000);
		}
	};
	function ping(callback) {
    var uri = window.location;
    var xhr = new XMLHttpRequest();
		xhr.onreadystatechange = function() {
	    if(xhr.status == 200) {
	    	callback(true);
	    } else {
	    	callback(false);
	    }
		};
    xhr.open("GET", uri, false);
    xhr.send(null);
	};
	function connect() {
		var sock = new WebSocket('ws://localhost:8080/ws');
		sock.onerror = function(err) {
			markOffline();
		};
		sock.onopen = function() {
			window.setTimeout(checkHeartbeat, 5000);
		};
		sock.onmessage = function(e) {
			if (e.data == 'heartbeat') {
				lastHeartbeat = Date.now();
			} else if (e.data == 'refresh') {
				window.location.reload(true);
			}
		};
	};
	connect();
}());