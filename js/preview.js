(function() {
	var sock = new WebSocket('ws://localhost:8080/ws');
	var lastHeartbeat = Date.now();
	function checkHeartbeat() {
		var now = Date.now();
		if (now - lastHeartbeat > 5000) {
			window.location = window.location;
		} else {
			window.setTimeout(checkHeartbeat, 5000);
		}
	};
	sock.onerror = function(err) {
		alert('Preview failed to open websocket.\nAuto-refresh will be disabled.');
	};
	sock.onopen = function() {
		window.setTimeout(checkHeartbeat, 5000);
	};
	sock.onmessage = function(e) {
		if (e.data == 'heartbeat') {
			lastHeartbeat = Date.now();
		} else if (e.data == 'refresh') {
			window.location = window.location;
		}
	};
}());