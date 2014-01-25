package geogo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
	"time"
)

func TestGeocoderFactory(t *testing.T) {
	geocoder := NewGeocoder()
	if len(geocoder.apiUrls) == 0 {
		t.Error("error initilaizing geocoder")
	}
}

/* func TestGeocoder(t *testing.T) { */

/* 	var mygeocoders = map[string]string{ */
/* 		"gmaps": "http://maps.googleapis.com/maps/api/geocode/json?sensor=true&address=%s", */
/* 		"osm":   "http://nominatim.openstreetmap.org/search?format=json&q=%s", */
/* 		"ymaps": "http://gws2.maps.yahoo.com/findlocation?format=json&pf=1&locale=en_US&flags=&offset=15&gflags=&q=%s", */
/* 	} */
/* 	geo := &Geocoder{mygeocoders, 1000} */
/* 	result := geo.Geocode("via teodosio 65, milano") */
/* 	fmt.Printf("%+v\n", result) */
/* } */

func TestGeocodeRequest(t *testing.T) {
	data, _ := NewMockResponse("fixtures/success/gmaps.json")
	gmaps := NewTestServer(data, 20)
	defer gmaps.Close()
	data, _ = NewMockResponse("fixtures/success/ymaps.json")
	ymaps := NewTestServer(data, 30)
	defer ymaps.Close()
	data, _ = NewMockResponse("fixtures/success/osm.json")
	osm := NewTestServer(data, 10)
	defer osm.Close()
	var geocoders = map[string]string{
		"gmaps": gmaps.URL + "/?%s",
		"osm":   osm.URL + "/?%s",
		"ymaps": ymaps.URL + "/?%s",
	}
	geocoder := &Geocoder{geocoders, 1000}
	geo := geocoder.Geocode("via ferretta 1/e, voghera")
	if geo.Lat != 44.9762554 || geo.Lng != 9.0313928 {
		t.Error("error geocooding")
	}
}

func TestParseGmapsJson(t *testing.T) {
	data, _ := NewMockResponse("fixtures/success/gmaps.json")
	geo := ParseGmapsJson(data)
	if geo.Lat != 44.98451 || geo.Lng != 9.01891 {
		t.Error("error geocoding data")
	}
	data, _ = NewMockResponse("fixtures/unsuccess/gmaps.json")
	geo = ParseGmapsJson(data)
	if geo.Success != false {
		t.Error("error geocoding data")
	}
}

func TestParseYmapsJson(t *testing.T) {
	data, _ := NewMockResponse("fixtures/success/ymaps.json")
	geo := ParseYmapsJson(data)
	if geo.Lat != 44.987764 || geo.Lng != 9.00233 {
		t.Error("error geocoding data")
	}
	data, _ = NewMockResponse("fixtures/unsuccess/ymaps.json")
	geo = ParseYmapsJson(data)
	if geo.Success != false {
		t.Error("error geocoding data")
	}
}

func TestParseOsmJson(t *testing.T) {
	data, _ := NewMockResponse("fixtures/success/osm.json")
	geo := ParseOsmJson(data)
	if geo.Lat != 44.9762554 || geo.Lng != 9.0313928 {
		t.Error("error geocoding data")
	}
	data, _ = NewMockResponse("fixtures/unsuccess/osm.json")
	geo = ParseOsmJson(data)
	if geo.Success != false {
		t.Error("error geocoding data")
	}
}

func NewTestServer(data []byte, ms time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(ms * time.Millisecond)
		fmt.Fprintln(w, string(data))
	}))
}

func NewMockResponse(s string) ([]byte, error) {
	dataPath := path.Join(s)
	_, readErr := os.Stat(dataPath)
	if readErr != nil && os.IsNotExist(readErr) {
		return nil, readErr
	}
	handler, handlerErr := os.Open(dataPath)
	if handlerErr != nil {
		return nil, handlerErr
	}

	data, readErr := ioutil.ReadAll(handler)

	if readErr != nil {
		return nil, readErr
	}

	return data, nil
}
