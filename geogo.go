package geogo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Result struct {
	res    []byte
	kind   string
	status int
}

var geocoders = map[string]string{
	"gmaps": "http://maps.googleapis.com/maps/api/geocode/json?sensor=true&address=%s",
	"osm":   "http://nominatim.openstreetmap.org/search?format=json&q=%s",
	"ymaps": "http://gws2.maps.yahoo.com/findlocation?format=json&pf=1&locale=en_US&flags=&offset=15&gflags=&q=%s",
}

type Geocoder struct {
	ApiUrls map[string]string
	Timeout time.Duration
}

func NewGeocoder() *Geocoder {
	return &Geocoder{geocoders, 1000}
}

func (g *Geocoder) Lookup(kind, query string) *Result {
	result, status := makeHttpRequest(fmt.Sprintf(g.ApiUrls[kind], url.QueryEscape(query)))
	return &Result{result, kind, status}
}

func (g *Geocoder) MultiLookup(query string) (result *Result) {
	c := make(chan *Result)
	for kind := range g.ApiUrls {
		go func(kind string) {
			if result := g.Lookup(kind, query); validResult(result) {
				c <- result
			}
		}(kind)
	}
	timeout := time.After(g.Timeout * time.Millisecond)
	select {
	case result := <-c:
		return result
	case <-timeout:
		log.Println("timed out")
		return &Result{status: 400}
	}
	return
}

func makeHttpRequest(addr string) ([]byte, int) {
	r, err := http.Get(addr)
	if err != nil {
		log.Println(err)
    return make([]byte,0), 500
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
    return make([]byte,0), 500
	}

	return body, r.StatusCode
}

func validResult(r *Result) bool {
	return r.status == 200
}

func (g *Geocoder) Geocode(q string) *Geo {
	result := g.MultiLookup(q)
	if result.status != 200 {
		return &Geo{Success: false}
	}
	switch result.kind {
	case "gmaps":
		return ParseGmapsJson(result.res)
	case "osm":
		return ParseOsmJson(result.res)
	case "ymaps":
		return ParseYmapsJson(result.res)
	default:
		return &Geo{Success: false}
	}
}
