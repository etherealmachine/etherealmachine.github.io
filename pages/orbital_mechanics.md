title: Orbital Mechanics
summary: A small workbook on orbital mechanics problems with runnable code.

(4.1) $F = ma$

(4.2) $F = G \left \( \frac{m*1m*2}{r^2} \right \)$

where $G = 6.67259x10^{-11} \frac{Nm^2}{kg^2}$

(4.3) $g = \frac{GM}{r^2}$

where $GM = 3.986005x10^{14} \frac{m^3}{s^2}$

(4.4) $a = \frac{v^2}{r}$

(4.5) $F = \frac{mv^2}{r}$

(4.6) $v = \sqrt{\frac{GM}{r}}$

<div class="well">

## Problem 4.1

Calculate the velocity of an artificial satellite orbiting the Earth in a circular orbit at an altitude of 200 km above the Earth's surface.

The radius of the earth is 6,378.14 km.

```javascript
var GM = 3.986005e14; function v(r) { return Math.sqrt(GM/r); }
var R*earth = 6378.14 * 1000; // meters
var altitude*satellite = 200 * 1000; // meters
var R*satellite = R*earth + altitude*satellite; v(R*satellite) / 1000 + " km/s";
```

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

(4.18) $R*a = \frac{R*p}{\left \( \frac{2GM}{R*pV*p^2} - 1 \right \)}$

(4.19) $R*p = \frac{R*a}{\left \( \frac{2GM}{R*aV*a^2} - 1 \right \)}$

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

$R\_p = R\_{earth} + 200 km$

want to find $R_a$

(4.18) $R*a = \frac{R*p}{\left \( \frac{2GM}{R*pV*p^2} - 1 \right \)}$

```javascript
function R_a(R_p, V_p) {
  return R_p / (((2*GM)/(R_p*Math.pow(V_p, 2))) - 1);
}
(R_a(R_earth + 200*1000, 7850) - R_earth) / 1000
```

(4.20) $e = \frac{R*pV*p^2}{GM} - 1$

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

also note $R*p+R*a = 2a$

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
