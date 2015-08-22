title: Orbital Mechanics
summary: A small workbook on orbital mechanics problems with runnable code.
date: 2015-08-13T00:00:00Z
published: true

<div class="alert alert-warning" role="alert">
This page is a work-in-progress.

11 out of 30 problems are complete.
  <div class="progress" style="margin: 5px 0px 0px 0px;">
    <div
      class="progress-bar"
      role="progressbar"
      aria-valuenow="37"
      aria-valuemin="0"
      aria-valuemax="100"
      style="width: 37%;">
      37%
    </div>
  </div>
</div>

<script src="/js/emit.js"></script>
<!--table class="table table-bordered"-->
<div class="well">
<p>
This is a copy of Robert A. Braeunig's inestimable
[Rocket & Space Technology](http://www.braeunig.us/space/index.htm). I learn
best by following examples, so I created a workbook where I could copy the text
and follow along by solving the problems.
</p>
</div>

## Introduction

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

where $G$ is an universal constant, called the *constant of gravitation*, and
has the value $6.67259x10^{-11} \frac{Nm^2}{kg^2}$.

Let's now look at the force that the Earth exerts on an object. If the object
has a mass $m$, and the Earth has a mass $M$, and the object's distance from the
center of the Earth is $r$, then the forst that the Earth exerts on the object
is $\frac{GmM}{r^2}$. If we drop the object, the Earth's gravity will cause it
to accelerate toward the center of the Earth. By Newton's second law ($F = ma$),
this acceleration $g$ must equal $\frac{\frac{GmM}{r^2}}{m}$, or

(4.3) $g = \frac{GM}{r^2}$

At the surface of the earth, this acceleration has the value $9.80665 \frac{m}{s^2}$.

Many of the upcoming computations will be somewhat simplified if we express the
product $GM$ as a constant, which for Earth has the value $3.986005x10^{14}\frac{m^3}{s^2}$.
The product GM is often represented by the Greek letter $\mu$.

For additional useful constants please see the appendix [Basic Constants](http://www.braeunig.us/space/constant.htm).

For a refresher on SI versus U.S. units see the appendix [Weights & Measures](http://www.braeunig.us/space/units.htm).

### Uniform Circular Motion

In the simple case of free fall, a particle accelerates toward the center of the
Earth while moving in a straight line. The velocity of the particle changes in
magnitude, but not in direction. In the case of uniform circular motion a
particle moves in a circle with constant speed. The velocity of the particle
changes continuously in direction, but not in magnitude. From Newton's laws we
see that since the direction of the velocity is changing, there is an
acceleration. This acceleration, called *centripetal acceleration* is directed
inward toward the center of the circle and is given by

(4.4) $a = \frac{v^2}{r}$

where $v$ is the speed of the particle and $r$ is the radius of the circle.
Every accelerating particle must have a force acting on it, defined by Newton's
second law ($F = ma$). Thus, a particle undergoing uniform circular motion is
under the influence of a force, called *centripetal force*, whose magnitude is
given by

(4.5) $F = \frac{mv^2}{r}$

The direction of $F$ at any instant must be in the direction of $a$ at the same
instant, that is radially inward.

A satellite in orbit is acted on only by the forces of gravity. The inward
acceleration which causes the satellite to move in a circular orbit is the
gravitational acceleration caused by the body around which the satellite orbits.
Hence, the satellite's centripetal acceleration is $g$, that is $g = \frac{v^2}
{r}$. From Newton's law of universal gravitation we know that $g = \frac{GM}
{r^2}$. Therefore, by setting these equations equal to one another we find that,
for a circular orbit,

$\frac{v^2}{r} = \frac{GM}{r^2}$, or 

(4.6) $v = \sqrt{\frac{GM}{r}}$

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-1">
    #### Problem 4.1
  </div>
  <div id="problem4-1" class="panel-body collapse">
Calculate the velocity of an artificial satellite orbiting the Earth in a
circular orbit at an altitude of 200 km above the Earth's surface.

Note that the radius of the earth is 6,378.14 km.

    <button
      class="btn btn-primary"
      data-toggle="collapse" href="#solution4-1">
      Show solution
    </button>
    <div id="solution4-1" class="collapse">
```javascript
var GM = 3.986005e14;
function v(r) {
  return Math.sqrt(GM/r);
}
var R_earth = 6378.14 * 1000; // meters
var altitude_satellite = 200 * 1000; // meters
var R_satellite = R_earth + altitude_satellite;
emit("velocity: %d m/s", v(R_satellite));
```
    </div>
  </div>
</div> 

### Motions of Planets and Satellites

Through a lifelong study of the motions of bodies in the solar system, Johannes
Kepler (1571-1630) was able to derive three basic laws known as *Kepler's laws
of planetary motion*. Using the data compiled by his mentor Tycho Brahe
(1546-1601), Kepler found the following regularities after years of laborious
calculations:

1. All planets move in elliptical orbits with the sun at one focus.
2. A line joining any planet to the sun sweeps out equal areas in equal times.
3. The square of the period of any planet about the sun is proportional to the
cube of the planet's mean distance from the sun.

These laws can be deduced from Newton's laws of motion and law of universal
gravitation. Indeed, Newton used Kepler's work as basic information in the
formulation of his gravitational theory.

As Kepler pointed out, all planets move in elliptical orbits, however, we can
learn much about planetary motion by considering the special case of circular
orbits. We shall neglect the forces between planets, considering only a planet's
interaction with the sun. These consideration apply equally well to the motion
of a satellite about a planet.

Let's examine the case of two bodies of masses $M$ and $m$ moving in circular
orbits under the influence of each other's gravitational attraction. The center
of mass of this system of two bodies lies along the line joining them at a point
$C$ such that $mr = MR$. The large body of mass $M$ moves in an orbit of
constant radius $R$ and the small body of mass $m$ in an orbit of constant
radius $r$, both having the same angular velocity $\omega$. For this to happen,
the gravitational force acting on each body must provide the necessary
centripetal acceleration. Since these gravitational forces are a simple
action-reaction pair, the centripetal forces must be equal but opposite in
direction. That is, $m\omega^2r$ must equal $M\omega^2R$. The specific
requirement then, is that the gravitational force acting on either body must
equal the centripetal force needed to keep it moving in its circular orbit, that
is

(4.7) $\frac{GMm}{(R+r)^2} = m \omega^2 r$

If one body has a much greater mass than the other, as is the case of the sun
and a planet or the Earth and a satellite, its distance from the center of mass
is much smaller than that other body. If we assume that $m$ is negligible
compared to $M$, then $R$ is negligible compared to $r$. Thus, equation (4.7)
becomes

(4.8) $GM = \omega^2 r^3$

If we express the angular velocity in terms of the period of revolution,
$\omega = \frac{2\pi}{P}$, we obtain

$GM = \frac{4\pi^2r^3}{p^2}$, or

(4.9) $p^2 = \frac{4\pi^2r^3}{GM}$

where $P$ is the period of revolution. This is a basic equation of planetary and
satellite motion. It also holds for elliptical orbits if we define $r$ to be the
semi-major axis ($a$) of the orbit.

A significant consequence of this equation is that it predicts Kepler's third
law of planetary motion, that is $p^2 \sim r^3$.

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-2">
    #### Problem 4.2
  </div>
  <div id="problem4-2" class="panel-body collapse">
Calculate the period of revolution for the satellite in problem 4.1.

    <button
      class="btn btn-primary"
      data-toggle="collapse" href="#solution4-2">
      Show solution
    </button>
    <div id="solution4-2" class="collapse">
```javascript
function p(r) {
  return Math.sqrt((4 * Math.pow(Math.PI, 2) * Math.pow(r, 3)) / GM);
}
emit("period: %d seconds", p(R_satellite));
```
    </div>
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

    <button
      class="btn btn-primary"
      data-toggle="collapse" href="#solution4-3">
      Show solution
    </button>
    <div id="solution4-3" class="collapse">
```javascript
function r(p) {
  return Math.cbrt((Math.pow(p, 2) * GM) / (4 * Math.pow(Math.PI, 2)));
}
emit("radius: %d meters", r(86164.1));
```
    </div>
  </div>
</div>

<div class="well">
In celestial mechanics where we are dealing with planetary or stellar sized bodies, it is often the case that the mass of the secondary body is significant in relation to the mass of the primary, as with the Moon and Earth. In this case the size of the secondary cannot be ignored. The distance R is no longer negligible compared to r and, therefore, must be carried through the derivation. Equation (4.9) becomes

$p^2 = \frac{4\pi^2r(R+r)^2}{GM}$

More commonly the equation is written in the equivalent form

$p^2 = \frac{4\pi^2a^3}{G(M+m)}$

where $a$ is the semi-major axis. The semi-major axis used in astronomy is
always the primary-to-secondary distance, or the *geocentric* semi-major axis.
For example, the Moon's mean geocentric distance from Earth ($a$) is 384,403
kilometers. On the other hand, the Moon's distance from the barycenter ($r$) is
379,732 km, with Earth's counter-orbit ($R$) taking up the difference of 4,671 km.
</div>

Kepler's second law of planetary motion must, of course, hold true for circular
orbits. In such orbits both $\omega$ and $r$ are constant so that equal areas
are swept out in equal times by the line joining a planet and the sun. For
elliptical orbits, however, both $\omega$ and $r$ will vary with time. Let's now
consider this case.

Figure 4.5 shows a particle revolving around $C$ along some arbitrary path. The
area swept out by the radius vector in a short time interval $\Delta t$ is shown
shaded. This area, neglecting the small triangular region at the end, is
one-half the base times the height or approximately $\frac{r(r\omega\Delta t)}
{2}$. This expression becomes more exact as $\Delta t$ approaches zero, i.e. the
small triangle goes to zero more rapidly than the large one. The rate at which
area is being swept out instantaneously is therefore

(4.10) $\lim_{t \to 0} \left [ \frac{r(r \omega \Delta t)}{2} \right ] = \frac
{\omega r^2} {2}$

For any given body moving under the influence of a central force, the value
$\omega r^2$ is constant.

Let's now consider two points, $P_1$ and $P_2$ in an orbit with radii $r_1$ and
$r_2$, and velocities $v_1$ and $v_2$. Since the velocity is always tangent to
the path, it can be seen that if $\gamma$ is the angle between $r$ and $v$,
then

(4.11) $v\sin \gamma = \omega r$

where $v\sin \gamma$ is the transverse component of $v$. Multiplying through by
$r$, we have

(4.12) $rv\sin\gamma = \omega r^2 = \text{Constant}$

or, for two points $P_1$ and $P_1$ on the orbital path

(4.13) $r_1v_1\sin\gamma_1 = r_2v_2\sin\gamma_2$

Note that at periapsis and apoapsis, $\gamma = \text{90 degrees}$. Thus, letting
$P_1$ and $P_2$ be these two points we get

(4.14) $R_pV_p = R_aV_a$

Let's now look at the energy of the above particle at points $P_1$ and $P_2$.
*Conservation of energy* states that the sum of the kinetic energy and the
potential energy of a particle remains constant. The kinetic energy $T$ of a
particle is given by $\frac{mv^2}{2}$ while the potential energy of gravity $V$
is calculated by the equation $-\frac{GMm}{r}$. Applying the conservation of
energy we have

$T_1 + V_1 = T_2 + V_2$, or

$\frac{mv_1^2}{2} - \frac{GMm}{r_1} = \frac{mv_2^2}{2} - \frac{GMm}{r_2}$, or

(4.15) $v_2^2 - v_1^2 = 2GM\left(\frac{1}{r_2} - \frac{1}{r_1}\right)$

From equations (4.14) and (4.15) we obtain

(4.16) $V_p = \sqrt{\frac{2GMR_a}{R_p(R_a+R_p)}}$, and

(4.17) $V_a = \sqrt{\frac{2GMR*p}{R_a(R_a+R_p)}}$

Rearranging terms we get

(4.18) $R_a = \frac{R_p}{\left ( \frac{2GM}{R_pV_p^2} - 1 \right )}$, and

(4.19) $R_p = \frac{R_a}{\left ( \frac{2GM}{R_aV_a^2} - 1 \right )}$

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-4">
    #### Problem 4.4
  </div>
  <div id="problem4-4" class="panel-body collapse">
An artificial Earth satellite is in an elliptical orbit which brings it to an altitude of 250 km at perigee and out to an altitude of 500 km at apogee. Calculate the velocity of the satellite at both perigee and apogee.

    <button
      class="btn btn-primary"
      data-toggle="collapse" href="#solution4-4">
      Show solution
    </button>
    <div id="solution4-4" class="collapse">
```javascript
function V_p(R_a, R_p) {
  return Math.sqrt((2 * GM * R_a) / (R_p * (R_a + R_p)));
};
function V_a(R_a, R_p) {
  return Math.sqrt((2 * GM * R_p) / (R_a * (R_a + R_p)));
};
R_a = 500*1000 + R_earth;
R_p = 250*1000 + R_earth;
emit("V_p: %d m/s", V_p(R_a, R_p));
emit("V_a: %d m/s", V_a(R_a, R_p));
```
    </div>
  </div>
</div>

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-5">
    #### Problem 4.5
  </div>
  <div id="problem4-5" class="panel-body collapse">
A satellite in Earth orbit passes through its perigee point at an altitude of 200 km above the Earth's surface and at a velocity of 7,850 m/s. Calculate the apogee altitude of the satellite.

    <button
      class="btn btn-primary"
      data-toggle="collapse" href="#solution4-5">
      Show solution
    </button>
    <div id="solution4-5" class="collapse">
$V_p = 7850 m/s$

$R_p = R_{earth} + 200 km$

want to find $R_a$

(4.18) $R_a = \frac{R_p}{\left ( \frac{2GM}{R_pV_p^2} - 1 \right )}$

```javascript
function R_a(R_p, V_p) {
  return R_p / (((2*GM)/(R_p*Math.pow(V_p, 2))) - 1);
}
emit(
  "Altitude @ apogee = %.1f km",
  (R_a(R_earth + 200*1000, 7850) - R_earth) / 1000);
```
    </div>
  </div>
</div>

The eccentricity $e$ of an orbit is given by

(4.20) $e = \frac{R_pV_p^2}{GM} - 1$

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-6">
    #### Problem 4.6
  </div>
  <div id="problem4-6" class="panel-body collapse">
Calculate the eccentricity of the orbit for the satellite in problem 4.5.

    <button
      class="btn btn-primary"
      data-toggle="collapse" href="#solution4-6">
      Show solution
    </button>
    <div id="solution4-6" class="collapse">
```javascript
function e(R_p, V_p) {
  return ((R_p * Math.pow(V_p, 2)) / GM) - 1;
}
emit("e = %.5f", e(R_earth + 200*1000, 7850));
```
    </div>
  </div>
</div>

If the semi-major axis $a$ and the eccentricity $e$ of an orbit are known, then
the periapsis and apoapsis distances can be calculated by

(4.21) $R_p = a(1-e)$, and

(4.22) $R_a = a(1+e)$

also note $R_p+R_a = 2a$

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-7">
    #### Problem 4.7
  </div>
  <div id="problem4-7" class="panel-body collapse">
A satellite in Earth orbit has a semi-major axis of 6,700 km and an eccentricity
of 0.01. Calculate the satellite's altitude at both perigee and apogee.

    <button
      class="btn btn-primary"
      data-toggle="collapse" href="#solution4-7">
      Show solution
    </button>
    <div id="solution4-7" class="collapse">
```javascript
a = 6700 * 1000; // meters
e = 0.01;
R_p = a*(1 - e);
R_a = a*(1 + e);
emit("Altitude @ perigee = %.1f km", (R_p - R_earth) / 1000);
emit("Altitude @ apogee = %.1f km", (R_a - R_earth) / 1000);
```
    </div>
  </div>
</div>

### Launch of a Space Vehicle

The launch of a satellite or space vehicle consists of a period of powered
flight during which the vehicle is lifted above the Earth's atmosphere and
accelerated to orbital velocity by a rocket, or launch vehicle. Powered flight
concludes at burnout of the rocket's last stage at which time the vehicle begins
its free flight. During free flight the space vehicle is assumed to be subjected
only to the gravitational pull of the Earth. If the vehicle moves far from the
Earth, its trajectory may be affected by the gravitational influence of the sun,
moon, or another planet.

<div class="pull-right text-center">
<img
    src="https://docs.google.com/drawings/d/18NEGLWMyfpBlYb2wkPMpeuDxxpTEtE_G9IH-smdPCxw/pub?w=355&amp;h=282">
  <p>Figure 4.7</p>
</div>

A space vehicle's orbit may be determined from the position and the velocity of
the vehicle at the beginning of its free flight. A vehicle's position and
velocity can be described by the variables $r$, $v$, and, $\gamma$ where $r$ is
the vehicle's distance from the center of the Earth, $v$ is its velocity, and
$\gamma$ is the angle between the position and the velocity vectors, called the
zenith angle (see Figure 4.7).
If we let $r_1$, $v_1$, and $\gamma_1$ be the initial (launch) values of $r$,
$v$, and $\gamma$, then we may consider these as given quantities. If we let
point $P_2$ represent the perigee, then equation (4.13) becomes

(4.23) $v_2 = V_p = \frac{r_1v_1\sin{\gamma_1}}{R_p}$

Substituting equation (4.23) into (4.15), we can obtain an equation for the
perigee radius $R_p$.

(4.24) $\frac{r_1^2 v_1^2 \sin^2 \gamma_1}{R_p^2} - v_1^2 = 2GM \left( \frac{1}
{R_p} - \frac{1}{r_1} \right)$

Multiplying through by $\frac{-R_p^2}{r_1^2v_1^2}$ and rearranging, we get

(4.25) $\left(\frac{R_p}{r_1}\right)^2(1-C) + \left(\frac{R_p}{r_1}\right)C -
\sin^2\gamma_1 = 0$

where $C = \frac{2GM}{r_1v_1^2}$

Note that this is a simple quadratic equation in the ratio $\frac{R_p
}{r_1}$ and that $\frac{2GM}{r_1v_1^2}$ is a nondimensional parameter of the
orbit.

Solving for $\frac{R_p}{r_1}$ gives

(4.26) $\left(\dfrac{R_p}{r_1}\right)_{1,2} = \dfrac{-C \pm \sqrt{C^2 - 4(1-C)
(-\sin^2\gamma_1)}} {2
(1-C)}$

Like any quadratic, the above equation yields two answers. The smaller of the
two answers corresponds to $R_p$, the periapsis radius. The other root
corresponds to the apoapsis radius, $R_a$.

Please note that in practice spacecraft launches are usually terminated at
either perigee or apogee, i.e. $\gamma = 90$. This condition results in the
minimum use of propellant.

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-8">
    #### Problem 4.8
  </div>
  <div id="problem4-8" class="panel-body collapse">
A satellite is launched into Earth orbit where its launch vehicle burns out at
an altitude of 250 km.  At burnout the satellite's velocity is 7,900 m/s with the
zenith angle equal to 89 degrees.  Calculate the satellite's altitude at perigee
and apogee.

    <button
      class="btn btn-primary"
      data-toggle="collapse" href="#solution4-8">
      Show solution
    </button>
    <div id="solution4-8" class="collapse">
```javascript
var r1 = R_earth + (250*1000);
var v1 = 7900;
var g1 = 89 * (Math.PI / 180);
var C = (2*GM) / (r1 * Math.pow(v1, 2));
var root = Math.pow(C, 2) - 4*(1 - C) * -Math.pow(Math.sin(g1), 2);
var R_p = r1 * ((-C + Math.sqrt(root)) / (2*(1-C)));
var R_a = r1 * ((-C - Math.sqrt(root)) / (2*(1-C)));
emit("Altitude @ perigee: %.1f km", (R_p - R_earth) / 1000);
emit("Altitude @ apogee: %.1f km", (R_a - R_earth) / 1000);
```
    </div>
  </div>
</div>

Equation (4.26) gives the values of $R_p$ and $R_a$ from which the eccentricity
of the orbit can be calculated, however, it may be simpler to calculate the
eccentricity $e$ directly from the equation

(4.27) $e = \sqrt{\left(\frac{r_1v_1^2}{GM}-1\right)^2\sin^2\gamma_1+\cos^2\gamma_1}$

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-9">
    #### Problem 4.9
  </div>
  <div id="problem4-9" class="panel-body collapse">
Calculate the eccentricity of the orbit for the satellite in problem 4.8.

    <button
      class="btn btn-primary"
      data-toggle="collapse" href="#solution4-9">
      Show solution
    </button>
    <div id="solution4-9" class="collapse">
```javascript
var r1 = R_earth + (250*1000);
var v1 = 7900;
var g1 = 89 * (Math.PI / 180);
function e(r1, v1, g1) {
  return Math.sqrt(
    Math.pow(((r1*Math.pow(v1, 2)) / GM) - 1, 2) * Math.pow(Math.sin(g1), 2)
    + Math.pow(Math.cos(g1), 2));
};
emit("e = %.7f", e(r1, v1, g1));
```
    </div>
  </div>
</div>

To pin down a satellite's orbit in space, we need to know the angle $\nu$, the true anomaly, from the periapsis point to the launch point. This angle is given by

(4.28) $\tan\nu = \dfrac{\left(\frac{r_1v_1^2}{GM}\right)\sin\gamma_1\cos\gamma_1}{\left(\frac{r_1v_1^2}{GM}\right)\sin^2\gamma_1 - 1}$

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-10">
    #### Problem 4.10
  </div>
  <div id="problem4-10" class="panel-body collapse">
Calculate the true anomaly, $\nu$ from perigee point to launch point for the satellite
in problem 4.8.

    <button
      class="btn btn-primary"
      data-toggle="collapse" href="#solution4-10">
      Show solution
    </button>
    <div id="solution4-10" class="collapse">
```javascript
var rvGM = (r1*Math.pow(v1, 2))/GM;
var v = Math.atan((rvGM*Math.sin(g1)*Math.cos(g1))/((rvGM*Math.pow(Math.sin(g1), 2)) - 1));
emit("true anomaly = %.3f°", v / (Math.PI / 180));
```
    </div>
  </div>
</div>

In most calculations, the complement of the zenith angle is used, denoted by
$\Phi$. This angle is called the flight-path angle, and is positive when the
velocity vector is directed away from the primary as shown in Figure 4.8. When
flight-path angle is used, equations (4.26) through (4.28) are rewritten as
follows:

(4.29) $\left(\dfrac{R_p}{r_1}\right)_{1,2} = \dfrac{-C \pm \sqrt{C^2 - 4(1-C)
(-\cos^2\Phi})} {2
(1-C)}$

where $C = \frac{GM}{rv^2}$

(4.30) $e = \sqrt{\left(\frac{rv^2}{GM}-1\right)^2\cos^2\Phi + \sin^2\Phi}$

(4.31) $\tan\nu = \dfrac{\left(\frac{rv^2}{GM}\right)\cos\Phi\sin\Phi}{\left
(\frac{rv^2}{GM}\right)\cos^2\Phi-1}$

The semi-major axis is, of course, equal to $\frac{R_p+R_a}{2}$, though it may be
easier to calculate it directly as follows:

(4.32) $a = \dfrac{1}{\left(\frac{2}{r} - \frac{v^2}{GM}\right)}$

<div class="panel panel-default">
  <div class="panel-heading" data-toggle="collapse" href="#problem4-11">
    #### Problem 4.11
  </div>
  <div id="problem4-11" class="panel-body collapse">
Calculate the semi-major axis of the orbit for the satellite in problem 4.8.

    <button
      class="btn btn-primary"
      data-toggle="collapse" href="#solution4-11">
      Show solution
    </button>
    <div id="solution4-11" class="collapse">
```javascript
function a(r, v) {
  return 1 / ((2 / r) - (Math.pow(v, 2) / GM));
}
emit("%d m", a(r1, v1));
```
    </div>
  </div>
</div>

If $e$ is solved for directly using equation (4.27) or (4.30), and $a$ is solved
for using equation (4.32), $R_p$ and $R_a$ can be solved for simply using
equations (4.21) and (4.22).

#### Orbit Tile, Rotation and Orientation

Above we determined the size and shape of the orbit, but to determine the
orientation of the orbit in space, we must know the latitude and longitude and
the heading of the space vehicle at burnout.

Figure 4.9 above illustrates the location of a space vehicle at engine burnout,
or orbit insertion. $\beta$ is the azimuth heading measured in degrees clockwise
from north, $\delta$ is the geocentric latitude (or declination) of the burnout
point, $\Delta\lambda$ is the angular distance between the ascending node and the
burnout point measured in the equatorial plane, and $\ell$ is the angular
distance between the ascending node and the burnout point measured in the
orbital plane. $\lambda_1$ and $\lambda_2$ are the geographical longitudes of
the ascending node and the burnout point at the instant of engine burnout.
Figure 4.10 pictures the orbital elements, where $i$ is the inclination,
$\Omega$ is the longitude at the ascending node, $\omega$ is the argument of
periapsis, and $\nu$ is the true anomaly.

If $\beta$, $\delta$, and $\lambda_2$ are given, the other values can be
calculated from the following relationships:

(4.33) $\cos i = \cos \delta \sin \beta$

(4.34) $\tan \ell = \dfrac{\tan\delta}{\cos\beta}$

(4.35) $\tan \Delta\lambda = \sin \delta \tan \beta$

(4.36) $\omega = \ell - \nu$

(4.37) $\lambda_1 = \lambda_2 - \Delta \lambda$

In equation (4.36), the value of $\nu$ is found using equation (4.28) or (4.31).
If $\nu$ is positive, periapsis is west of the burnout point (as shown in
figure 4.10). If $\nu$ is negative, periapsis is east of the burnout point.

The logitude of the ascending node, $\Omega$, is measured in celestial
longitude, while $\lambda_1$ is geographical longitude. The celestial longitude
of the ascending node is equal to the *local apparent sidereal time*, in
degrees, at longitude $\lambda_1$ at the time of engine burnout. Sidereal time
is defined as the hour angle of the vernal equinox at a specific locality and
time; it has the same value as the right ascension of any celestial body that is
crossing the local meridian at that same instant. At the moment when the vernal
equinox crosses the local meridian the local apparent sidereal time is 00:00.
Try the below sidereal time calculator: