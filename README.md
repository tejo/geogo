geogo
=====

go geocode package


Try to geocode provided address with google maps, open street maps and yahoo maps services  and returns the first responding api


usage
-----

install pacakge
```
go get github.com/tejo/geogo
```

use it
```
package main

import (
	"fmt"
	"github.com/tejo/geogo"
)

func main() {
	geo := geogo.NewGeocoder()
	result := geo.Geocode("duomo plaza, milan")
	fmt.Printf("%+v\n", result)
	
	//specify api services and timeout (in milliseconds)
	var geocoders = map[string]string{
		"gmaps": "http://maps.googleapis.com/maps/api/geocode/json?sensor=true&address=%s",
		"osm":   "http://nominatim.openstreetmap.org/search?format=json&q=%s",
		"ymaps": "http://gws2.maps.yahoo.com/findlocation?format=json&pf=1&locale=en_US&flags=&offset=15&gflags=&q=%s",
	}
	geo = &geogo.Geocoder{geocoders, 1000}
	result = geo.Geocode("via teodosio 65, milano")
	fmt.Printf("%+v\n", result)

}
```

output:

```
&{Lat:45.46420639999999 Lng:9.1892441 Success:true}
```
