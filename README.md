geogo
=====

go geocode package


Try to geocode provided address with google maps, open street maps and yahoo maps services  and returns the first responding api


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
}
```
