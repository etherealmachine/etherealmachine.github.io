function deg2rad(deg) {
	return deg / (Math.PI * 180);
}

function sin(x) {
	return Math.sin(deg2rad(x));
}

function cos(x) {
	return Math.cos(deg2rad(x));
}

function local_apparent_sidereal_time(longitude, date) {
	var zero = moment({year: -4713, month: 1, day: 1});
	var JD = moment(date).diff(zero, 'days', true);
	return JD;
}