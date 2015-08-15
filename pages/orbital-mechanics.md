title: Orbital Mechanics
summary: A small workbook on orbital mechanics problems with runnable code.
date: 2015-08-13T00:00:00Z

<script src="/js/emit.js"></script>
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

In general, three observations of an object in orbit are required to calculate
the six orbital elements. Two other quantities often used to describe orbits are
period and true anomaly. *Period*, $P$, is the length of time required for a
satellite to complete one orbit. *True anomaly*, $\nu$, is the angular distance
of a point in an orbit past the point of periapsis, measured in degrees.

### Types Of Orbits

For a spacecraft to achieve Earth orbit, it must be launched to an elevation
above the Earth's atmosphere and accelerated to orbital velocity. The most
energy efficient orbit, that is one that requires the least amount of
propellant, is a direct low inclination orbit. To achieve such an orbit, a
spacecraft is launched in an eastward direction from a site near the Earth's
equator. The advantage being that the rotational speed of the Earth contributes
to the spacecraft's final orbital speed. At the United States' launch site in
Cape Canaveral (28.5 degrees north latitude) a due east launch results in a
"free ride" of 1,471 km/h (914 mph). Launching a spacecraft in a direction other
than east, or from a site far from the equator, results in an orbit of higher
inclination. High inclination orbits are less able to take advantage of the
initial speed provided by the Earth's rotation, thus the launch vehicle must
provide a greater part, or all, of the energy required to attain orbital
velocity. Although high inclination orbits are less energy efficient, they do
have advantages over equatorial orbits for certain applications. Below we
describe several types of orbits and the advantages of each:

**Geosynchronous orbits** (Geo) are circular orbits around the Earth having a
period of 24 hours. A geosynchronous orbit with an inclination of zero degrees
is called a *geostationary orbit*. A spacecraft in a geostationary orbit appears
to hang motionless above one position on the Earth's equator. For this reason,
they are ideal for some types of communication and meteorological satellites. A
spacecraft in an inclined geosynchronous orbit will appear to follow a regular
figure-8 pattern in the sky once every orbit. To attain geosynchronous orbit, a
spacecraft is first launched into an elliptical orbit with an apogee of 35,786
km (22,236 miles) called a *geosynchronous transfer orbit* (GTO). The orbit is
then circularized by firing the spacecraft's engine at apogee.

**Polar orbits** (PO) are orbits with an inclination of 90 degrees. Polar orbits
are useful for satellites that carry out mapping and/or surveillance operations
because as the planet rotates the spacecraft has access to virtually every point
on the planet's surface.

**Walking orbits**: An orbiting satellite is subjected to a great many
gravitational influences. First, planets are not perfectly spherical and they
have slightly uneven mass distribution. These fluctuations have an effect on a
spacecraft's trajectory. Also the sun, moon, and planets contribute a
gravitational influence on an orbiting satellite. With proper planning it is
possible to design an orbit which takes advantage of these influences to induce
a precession in the satellite's orbital plane. The resulting orbit is called a
*walking orbit*, or a precessing orbit.

**Sun synchronous orbits** (SSO) are walking orbits whose orbital plane
precesses with the same period as the planet's solar orbit period. In such an
orbit, a satellite crosses periapsis at about the same local time every orbit.
This is useful if a satellite is carrying instruments which depend on a certain
angle of solar illumination on the planet's surface. In order to maintain an
exact synchronous timing, it may be necessary to conduct occasional propulsive
maneuvers to adjust the orbit.

**Molniya orbits** are highly eccentric Earth orbits with a period of
approximately 12 hours (2 revolutions per day). The orbital inclination is
chosen so the rate of change of perigee is zero, thus both apogee and perigee
can be maintained over fixed latitudes. This condition occurs at inclinations of
63.4 and 116.6 degrees. For these orbits the argument of perigee is typically
placed in the southern hemisphere, so the satellite remains above the northern
hemisphere near apogee for approximately 11 hours per orbit. This orientation
can provide good ground coverage at high northern latitudes.

**Hohmann transfer orbits** are interplanetary trajectories whose advantage is
that they consume the least possible amount of propellant. A Hohmann transfer
orbit to an outer planet, such as Mars, is achieved by launching a spacecraft
and accelerating it in the direction of Earth's revolution around the sun until
it breaks free of the Earth's gravity and reaches a velocity which places it in
a sun orbit with an aphelion equal to the orbit of the outer planet. Upon
reaching its destination, the spacecraft must decelerate so that the planet's
gravity can capture it into a planetary orbit

To send a spacecraft to an inner planet, such as Venus, the spacecraft is
launched an accelerated in the direction opposite of Earth's revolution around
the sun (i.e. decelerated) until it achieves a sun orbit with a perihelion equal
to the orbit of the inner planet. It should be noted that the spacecraft
continues to move in the same direction as Earth, only more slowly.

To reach a planet requires that the spacecraft be inserted into an
interplanetary trajectory at the correct time so that the spacecraft arrives at
the planet's orbit when the planet will be at the point where the spacecraft
will intercept it. This task is comparable to a quarterback "leading" his
receiver so that the football and receiver arrive at the same point at the same
time. The interval of time in which a spacecraft must be launched in order to
complete its mission is called a *launch window*.

### Newton's Laws of Motion and Universal Gravitation

*Newton's laws of motion* describe the relationship between the motion of a
particle and the forces acting on it.

The first law states that if no forces are acting, a body at rest will remain at
rest, and a body in motion will remain in motion in a straight line. Thus, if no
forces are acting, the velocity (both magnitude and direction) will remain
constant.

The second law tells us that if a force is applied there will be a change in
velocity, i.e. an acceleration, proportional to the magnitude of the force and
in the direction in which the force is applied. This law may be summarized by
the equation

(4.1) $F = ma$

where $F$ is the force, $m$ is the mass of the particle, and $a$ is the
acceleration.

The third law states that if body 1 exerts a force on body 2, then body 2 will
exert a force of equal strength, but opposite in direction, on body 1. This law
is commonly stated, "for every action there is an equal and opposite reaction".

In his *law of universal gravitation*, Newton states that two particles having
masses $m_1$ and $m_2$ and separated by a distance $r$ are attracted to each
other with equal and opposite forces directed along the line joining the
particles. The common magnitude $F$ of the two forces is

(4.2) $F = G \left ( \frac{m_1m_2}{r^2} \right )$

where $G$ is an universal constant, called the constant of gravitation, and has
the value $6.67259x10^{-11} \frac{Nm^2}{kg^2}$.

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
var GM = 3.986005e14;
function v(r) {
  return Math.sqrt(GM/r);
}
var R_earth = 6378.14 * 1000; // meters
var altitude_satellite = 200 * 1000; // meters
var R_satellite = R_earth + altitude_satellite;
emit("velocity: %.2f km/s", v(R_satellite) / 100);
```
  </div>
</div> 

(4.7) $\frac{GMm}{(R+r)^2} = m \omega^2 r$

(4.8) $GM = \omega^2 r^3$

(4.9) $p^2 = \frac{4\pi^2r^3}{GM}$

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-2">
    #### Problem 4.2
  </div>
  <div id="problem4-2" class="panel-body collapse">
Calculate the period of revolution for the satellite in problem 4.1.

```javascript
function p(r) {
  return Math.sqrt((4 * Math.pow(Math.PI, 2) * Math.pow(r, 3)) / GM);
}
emit("period: %d seconds", p(R_satellite));
```
  </div>
</div>

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-3">
    #### Problem 4.3
  </div>
  <div id="problem4-3" class="panel-body collapse">
Calculate the radius of orbit for a Earth satellite in a geosynchronous orbit, where the Earth's rotational period is 86,164.1 seconds.

Recall from (4.9) $p^2 = \frac{4\pi^2r^3}{GM}$. We want $r$ in terms of $p$.

1. $p^2 = \frac{4\pi^2r^3}{GM}$
2. $p^2 GM = 4\pi^2r^3$
3. $\frac{p^2 GM}{4\pi^2} = r^3$
4. $r = \sqrt[3]{\frac{p^2 GM}{4\pi^2}}$

```javascript
function r(p) {
  return Math.cbrt((Math.pow(p, 2) * GM) / (4 * Math.pow(Math.PI, 2)));
}
emit("radius: %.2f meters", r(86164.1));
```
  </div>
</div>

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
