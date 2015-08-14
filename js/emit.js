var fmtre = /%([.]([0-9]*))?([dfs])/g;

function fmt(fmtString, vals) {
	if (!Array.isArray(vals)) {
		vals = [vals];
	}
	var s = fmtString;
	var i = 0;
	var matches;
	while (matches = fmtre.exec(fmtString)) {
		var value = vals[i];
		var precision = matches[2];
		var verb = matches[3];
		if (verb == "d") {
			value = Math.round(value);
		} else if (verb == "f") {
			value = new String(value);
			if (precision) {
				var sep = value.indexOf(".");
				if (sep > 0) {
					value = value.substring(0, sep+3);
				}
			}
		} else if (verb != "s") {
			value = "unknown format type '" + verb + "'"
		}
		s = s.replace(matches[0], value);
		i++;
	}
	return s;
}

function emitNode(node, fmtString, vals) {
	var el = document.createElement("code");
	var text = fmt(fmtString, vals);
	var content = document.createTextNode(text);
	el.appendChild(content);
	node.appendChild(el);
}

$(document).ready(function() {
	var blocks = document.getElementsByTagName("pre");
	for (var i = 0; i < blocks.length; i++) {
		var block = blocks[i];
		if (block.firstChild != null && block.firstChild.tagName == "CODE") {
			var script = block.firstChild.textContent;
			try {
				var emit = function(name, value, units) {
					emitNode(block, name, value, units);
				}
				eval(script);
			} catch (e) {
				console.log(e);
			}
		}
	}
});