title: Orbital Mechanics
summary: A small workbook on orbital mechanics problems with runnable code.

<!--table class="table table-bordered"-->
This is a copy of Robert A. Braeunig's inestimable  <a
href="http://www.braeunig.us/space/index.htm">Rocket & Space Technology</a>. I
learn best by following examples, so I created a workbook where I could copy the
text and follow along by solving the problems.

**Orbital mechanics**, also called flight mechanics, is the study of the motions
of artificial satellites and space vehicles moving under the influence of forces
such as gravity, atmospheric drag, thrust, etc. Orbital mechanics is a modern
offshoot of celestial mechanics which is the study of the motions of natural
celestial bodies such as the moon and planets. The root of orbital mechanics can
be traced back to the 17th century when mathematician Isaac Newton (1642-1727)
put forward his laws of motion and formulated his law of universal gravitation.
The engineering applications of orbital mechanics include ascent trajectories,
reentry and landing, rendezvous computations, and lunar and interplanetary
trajectories.

## Conic Sections

<img
		src="https://docs.google.com/drawings/d/10ZCeL3pW5HNIdKgpMf9yFGcJVJxbrKnePySc7P0_R0E/pub?w=814&amp;h=590"
		width="500px"
		class="pull-right" />

A **conic section**, or just **conic**, is a curve formed by passing a plane
through a right circular cone. As shown in Figure 4.1, the angular orientation
of the plane relative to the cone determines whether the conic section is a
circle, ellipse, parabola, or hyperbola. The circle and the ellipse arise when
the intersection of cone and plane is a bounded curve. The circle is a special
case of the ellipse in which the plane is perpendicular to the axis of the cone.
If the plane is parallel to a generator line of the cone, the conic is called a
parabola. Finally, if the intersection is an unbounded curve and the plane is
not parallel to a generator line of the cone, the figure is a hyperbola. In the
latter case the plane will intersect both halves of the cone, producing two
separate curves.

We can define all conic sections in terms of the eccentricity. The type of conic
section is also related to the semi-major axis and the energy. The table below
shows the relationships between eccentricity, semi-major axis, and energy and
the type of conic section.

<div class="pull-left">
Conic Section | Eccentricity, e | Semi-Major Axis | Energy
:------------ | :-------------: | :-------------: | :----:
Circle        | 0               | = radius        | < 0
Ellipse       | 0 < e < 1       | > 0             | < 0
Parabola      | 1               | infinity        | 0
Hyperbola     | > 1             | < 0             | > 0

</div>

<div class="clearfix"></div>

Satellite orbits can be any of the four conic sections. This page deals mostly
with elliptical orbits, though we conclude with an examination of the hyperbolic
orbit.

## Orbital Elements

To mathematically describe an orbit one must define six quantities, called
orbital elements. They are
* Semi-Major Axis, a
* Eccentricity, e
* Inclination, i
* Argument of Periapsis, $\omega$
* Time of Periapsis Passage, T
* Longitude of Ascending Node, $\Omega$

An orbiting satellite follows an oval shaped path known as an ellipse with the
body being orbited, called the primary, located at one of two points called
foci. An ellipse is defined to be a curve with the following property: for each
point on an ellipse, the sum of its distances from two fixed points, called
foci, is constant (see Figure 4.2). The longest and shortest lines that can be
drawn through the center of an ellipse are called the major axis and minor axis,
respectively. The **semi-major axis** is one-half of the major axis and
represents a satellite's mean distance from its primary. **Eccentricity** is the
distance between the foci divided by the length of the major axis and is a
number between zero and one. An eccentricity of zero indicates a circle.

**Inclination** is the angular distance between a satellite's orbital plane and
the equator of its primary (or the ecliptic plane in the case of heliocentric,
or sun centered, orbits). An inclination of zero degrees indicates an orbit
about the primary's equator in the same direction as the primary's rotation, a
direction called **prograde** (or direct). An inclination of 90 degrees
indicates a **polar** orbit. An inclination of 180 degrees indicates a
retrograde equatorial orbit. A **retrograde** orbit is one in which a satellite
moves in a direction opposite to the rotation of its primary.

**Periapsis** is the point in an orbit closest to the primary. The opposite of
periapsis, the farthest point in an orbit, is called **apoapsis**. Periapsis and
apoapsis are usually modified to apply to the body being orbited, such as
perihelion and aphelion for the Sun, perigee and apogee for Earth, perijove and
apojove for Jupiter, perilune and apolune for the Moon, etc. The argument of
periapsis is the angular distance between the ascending node and the point of
periapsis (see Figure 4.3). The time of periapsis passage is the time in which a
satellite moves through its point of periapsis.

Nodes are the points where an orbit crosses a plane, such as a satellite
crossing the Earth's equatorial plane. If the satellite crosses the plane going
from south to north, the node is the **ascending node**; if moving from north to
south, it is the **descending node**. The **longitude of the ascending node** is
the node's celestial longitude. Celestial longitude is analogous to longitude on
Earth and is measured in degrees counter-clockwise from zero with zero longitude
being in the direction of the vernal equinox.

(4.1) $F = ma$

(4.2) $F = G \left ( \frac{m_1m_2}{r^2} \right )$

where $G = 6.67259x10^{-11} \frac{Nm^2}{kg^2}$

(4.3) $g = \frac{GM}{r^2}$

where $GM = 3.986005x10^{14} \frac{m^3}{s^2}$

(4.4) $a = \frac{v^2}{r}$

(4.5) $F = \frac{mv^2}{r}$

(4.6) $v = \sqrt{\frac{GM}{r}}$

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-1">
    #### Problem 4.1
  </div>
  <div id="problem4-1" class="panel-body collapse">
Calculate the velocity of an artificial satellite orbiting the Earth in a
circular orbit at an altitude of 200 km above the Earth's surface.

The radius of the earth is 6,378.14 km.

```javascript
var GM = 3.986005e14; function v(r) { return Math.sqrt(GM/r); }
var R*earth = 6378.14 * 1000; // meters
var altitude*satellite = 200 * 1000; // meters
var R*satellite = R*earth + altitude*satellite; v(R*satellite) / 1000 + " km/s";
```
  </div>
</div> 

(4.7) $\frac{GMm}{(R+r)^2} = m \omega^2 r$

(4.8) $GM = \omega^2 r^3$

(4.9) $p^2 = \frac{4\pi^2r^3}{GM}$

## Problem 4.2

Calculate the period of revolution for the satellite in problem 4.1.

```javascript
function p(r) {
  return Math.sqrt((4 * Math.pow(Math.PI, 2) * Math.pow(r, 3)) / GM);
}
p(R_satellite) + " s";
```

## Problem 4.3

Calculate the radius of orbit for a Earth satellite in a geosynchronous orbit, where the Earth's rotational period is 86,164.1 seconds.

(4.9) $p^2 = \frac{4\pi^2r^3}{GM}$

We want $r$ in terms of $p$.

$p^2 GM = 4\pi^2r^3$

$\frac{p^2 GM}{4\pi^2} = r^3$

```javascript
function r(p) {
  return Math.cbrt((Math.pow(p, 2) * GM) / (4 * Math.pow(Math.PI, 2)));
}
r(86164.1) + " m";
```

(4.16) $V*p = \sqrt{\frac{2GMR*a}{R*p(R*a+R_p)}}$

(4.17) $V*a = \sqrt{\frac{2GMR*p}{R*a(R*a+R_p)}}$

(4.18) $R*a = \frac{R*p}{\left ( \frac{2GM}{R*pV*p^2} - 1 \right )}$

(4.19) $R*p = \frac{R*a}{\left ( \frac{2GM}{R*aV*a^2} - 1 \right )}$

## Problem 4.4

An artificial Earth satellite is in an elliptical orbit which brings it to an altitude of 250 km at perigee and out to an altitude of 500 km at apogee. Calculate the velocity of the satellite at both perigee and apogee.

```javascript
function V_p(R_a, R_p) {
  return Math.sqrt((2 * GM * R_a) / (R_p * (R_a + R_p)));
};
function V_a(R_a, R_p) {
  return Math.sqrt((2 * GM * R_p) / (R_a * (R_a + R_p)));
};
R_a = 500*1000 + R_earth;
R_p = 250*1000 + R_earth;
[V_p(R_a, R_p) + " m/s", V_a(R_a, R_p) + " m/s"];
```

## Problem 4.5

A satellite in Earth orbit passes through its perigee point at an altitude of 200 km above the Earth's surface and at a velocity of 7,850 m/s. Calculate the apogee altitude of the satellite.

$V_p = 7850 m/s$

$R_p = R_{earth} + 200 km$

want to find $R_a$

(4.18) $R_a = \frac{R_p}{\left ( \frac{2GM}{R_pV_p^2} - 1 \right )}$

```javascript
function R_a(R_p, V_p) {
  return R_p / (((2*GM)/(R_p*Math.pow(V_p, 2))) - 1);
}
(R_a(R_earth + 200*1000, 7850) - R_earth) / 1000
```

(4.20) $e = \frac{R_pV_p^2}{GM} - 1$

## Problem 4.6

Calculate the eccentricity of the orbit for the satellite in problem 4.5.

```javascript
function e(R_p, V_p) {
  return ((R_p * Math.pow(V_p, 2)) / GM) - 1;
}
e(R_earth + 200*1000, 7850);
```

(4.21) $R_p = a(1-e)$

(4.22) $R_a = a(1+e)$

also note $R_p+R_a = 2a$

## Problem 4.7

A satellite in Earth orbit has a semi-major axis of 6,700 km and an eccentricity of 0.01. Calculate the satellite's altitude at both perigee and apogee.

```javascript
a = 6700 * 1000;
e = 0.01;
R_p = a*(1 - e);
R_a = a*(1 + e);
[((R_p - R_earth) / 1000) + " km", ((R_a - R_earth) / 1000) + " km"];
```

Launch of a Space Vehicle The launch of a satellite or space vehicle consists
-----------------------------------------------------------------------------

of a period of powered flight during which the vehicle is lifted above the Earth's atmosphere and accelerated to orbital velocity by a rocket, or launch vehicle. Powered flight concludes at burnout of the rocket's last stage at which time the vehicle begins its free flight. During free flight the space vehicle is assumed to be subjected only to the gravitational pull of the Earth. If the vehicle moves far from the Earth, its trajectory may be affected by the gravitational influence of the sun, moon, or another planet.

<img
    src="https://docs.google.com/drawings/d/18NEGLWMyfpBlYb2wkPMpeuDxxpTEtE_G9IH-smdPCxw/pub?w=355&amp;h=282"
    class="pull-right">

A space vehicle's orbit may be determined from the position and the velocity of the vehicle at the beginning of its free flight. A vehicle's position and velocity can be described by the variables $r$, $v$, and, $\gamma$ where $r$ is the vehicle's distance from the center of the Earth, $v$ is its velocity, and $\gamma$ is the angle between the position and the velocity vectors, called the zenith angle (see [Figure 4.7](/fig4-7.gif)). If we let $r*1$, $v*1$, and $\gamma*1$ be the initial (launch) values of $r$, $v$, and $\gamma$, then we may consider these as given quantities. If we let point $P*2$ represent the perigee, then equation (4.13) becomes

(4.23) $v*2 = V*p = \frac{r*1v*1\sin{\gamma*1}}{R*p}$

Substituting equation (4.23) into (4.15), we can obtain an equation for the perigee radius $R_p$.

(4.24) $$
