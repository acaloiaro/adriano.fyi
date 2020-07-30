---
layout: page
title: Where Am I?
url: /whereami
---

{{< rawhtml >}}
<link rel="stylesheet" href="https://unpkg.com/leaflet@1.6.0/dist/leaflet.css"
  integrity="sha512-xwE/Az9zrjBIphAcBb3F6JVqxf46+CDLwfLMHloNu6KEQCAWi6HcDUbeOfBIptF7tcCzusKFjFw2yuvEpDL9wQ=="
  crossorigin=""/>
 <script src="https://unpkg.com/leaflet@1.6.0/dist/leaflet.js"
   integrity="sha512-gZwIG9x3wUXg2hdXF6+rVkLF/0Vi9U8D2Ntg4Ga5I5BZpVkVxlJWbSQtXPSiUTtC0TjtGOmxa1AJPuV0CPthew=="
   crossorigin=""></script>

 <div id="mapid"></div>

<style>
#mapid { height: 480px; }
</style>

<script type="text/javascript">
var lat, lon

fetch('https://adriano.fyi/current_location')
  .then(response => response.json())
  .then(data => {
      var lat = data.lat
      var lon = data.lon
      var map = L.map('mapid').setView([lat, lon], 15);
      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="https://openstreetmap.org/copyright">OpenStreetMap contributors</a>'
      }).addTo(map);

      // show the scale bar on the lower left corner
      L.control.scale().addTo(map);

      // show a marker on the map
      L.marker({lon: lon, lat: lat}).bindPopup('Come say hi').addTo(map);

  });

</script>
{{< /rawhtml >}}
