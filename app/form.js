(function(root) {
	root.FormValues = function(e) {
		var vals = {};
		e.path.forEach(function(el) {
			if (el.tagName == 'FORM') {
				var inputs = el.querySelectorAll('input, textarea');
				for (var i = 0; i < inputs.length; i++) {
					var input = inputs[i];
					vals[input.name] = input.value;
					if (input.type == 'checkbox') {
						vals[input.name] = input.checked;
					}
				}
			}
		});
		return vals;
	};
})(this);