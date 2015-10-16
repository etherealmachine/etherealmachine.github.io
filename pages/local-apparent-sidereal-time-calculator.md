title: Local Apparent Sidereal Time Calculator
summary: A local apparent sidereal time calculator written in Javascript.

<script src="/js/emit.js"></script>
<script src="/js/moment.min.js"></script>
<script src="/js/sidereal.js"></script>

```javascript
var last = local_apparent_sidereal_time(0, new Date(1994, 6, 16, 18));
emit("got: GMST=%.2f, want: GMST=%.2f", last);
```