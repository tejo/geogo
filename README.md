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
}
```

output:

```
&{Lat:45.46420639999999 Lng:9.1892441 Success:true}
```
