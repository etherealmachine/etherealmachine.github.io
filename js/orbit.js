var canvas = document.getElementById('canvas');
var w = canvas.width;
var h = canvas.height;
var scene = new THREE.Scene();
var camera = new THREE.PerspectiveCamera(75, w/h, 0.1, 1000);
var renderer = new THREE.WebGLRenderer({canvas: canvas});
var geometry = new THREE.SphereGeometry(1, 16, 16);
var material = new THREE.MeshPhongMaterial({
	color: 0x0000ff,
	specular: 0x111111,
	shininess: 30
});
var sphere = new THREE.Mesh(geometry, material);
scene.add(sphere);

var curve = new THREE.EllipseCurve(0, 0, 1.1, 1.1, 0, 2*Math.PI, false);

var path = new THREE.Path(curve.getPoints(50));

var geometry = path.createPointsGeometry(50);
var material = new THREE.LineBasicMaterial({color : 0xff0000});

var ellipse = new THREE.Line(geometry, material);
ellipse.rotateOnAxis(new THREE.Vector3(1, 0, 0), 0.5*Math.PI);
ellipse.rotateOnAxis(new THREE.Vector3(1, 0, 0), 0.1*Math.PI);
scene.add(ellipse);

scene.add(new THREE.ArrowHelper(
	new THREE.Vector3(1, 0, 0),
	new THREE.Vector3(0, 0, 0),
	2,
	0xff0000));
scene.add(new THREE.ArrowHelper(
	new THREE.Vector3(0, 1, 0),
	new THREE.Vector3(0, 0, 0),
	2,
	0x00ff00));
scene.add(new THREE.ArrowHelper(
	new THREE.Vector3(0, 0, 1),
	new THREE.Vector3(0, 0, 0),
	2,
	0x0000ff));

camera.position.z = 5;

var controls = new THREE.OrbitControls(camera);
controls.target.x = 0;
controls.target.y = 0;
controls.target.z = 0;

function render() {
	renderer.render(scene, camera);
	controls.update();
}

var light = new THREE.PointLight(0xffffff, 1, 10000);
light.position.set(10, 0, 10);
scene.add(light);

controls.addEventListener('change', render);
render();

/*
var canvas = document.getElementById('canvas');
var ctx = canvas.getContext('2d');
var w = canvas.width;
var h = canvas.height;

// Kerbin
ctx.beginPath();
ctx.lineWidth = 0;
ctx.ellipse(w/2, h/2, 6, 6, 0, 0, 2*Math.PI);
ctx.fillStyle = 'rgb(37,57,130)';
ctx.fill();

var GM = 3.5316000e12;
var eccentricity = 0;
var periapsis = 12000000;
var apoapsis = 12000000;
var semiMajorAxis = (periapsis + apoapsis) / 2;
var n = Math.sqrt(GM / Math.pow(semiMajorAxis, 3));
var M_0 = 1.7;

function propagate(dt) {
	return n*dt + M_0;
}

function radius(dt) {
	var trueAnomaly = propagate(M_0, dt);
	return (
		(semiMajorAxis*(1 - Math.pow(eccentricity, 2))) /
		(1 + eccentricity*Math.cos(trueAnomaly))
	);
}

function polar2Cartesian(r, theta) {
	return [r * Math.cos(theta), r * Math.sin(theta)];
}

var k = 100000;
var dt = 0;
var point = polar2Cartesian(radius(dt), propagate(dt));

ctx.beginPath();
ctx.lineWidth = 1;
ctx.moveTo(point[0]/k + w/2, point[1]/k + h/2);
for (; dt <= 138984; dt++) {
	point = polar2Cartesian(radius(dt), propagate(dt));
	ctx.lineTo(point[0]/k + w/2, point[1]/k + h/2);
}
ctx.stroke();

var point = polar2Cartesian(radius(0), propagate(0));
ctx.beginPath();
ctx.lineWidth = 0;
ctx.ellipse(point[0]/k + w/2, point[1]/k + h/2, 2, 2, 0, 0, 2*Math.PI);
ctx.fillStyle = 'rgb(200, 200, 200)';
ctx.fill();
*/
