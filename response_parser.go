package geogo

import (
	"encoding/json"
	"strconv"
)

type Geo struct {
	Lat     float64
	Lng     float64
	Success bool
}

func ParseOsmJson(b []byte) *Geo {
	res := make([]map[string]interface{}, 0)
	json.Unmarshal(b, &res)
	if len(res) > 0 {
		latStr, _ := res[0]["lat"].(string)
		lngStr, _ := res[0]["lon"].(string)
		lat, _ := strconv.ParseFloat(latStr, 64)
		lng, _ := strconv.ParseFloat(lngStr, 64)
		return &Geo{lat, lng, true}
	} else {
		return &Geo{Success: false}

	}
}

func ParseYmapsJson(b []byte) *Geo {
	res := make(map[string]map[string]interface{}, 0)
	json.Unmarshal(b, &res)
	latStr, _ := res["Result"]["latitude"].(string)
	lngStr, _ := res["Result"]["longitude"].(string)
	lat, _ := strconv.ParseFloat(latStr, 64)
	lng, _ := strconv.ParseFloat(lngStr, 64)
	if lat == 0 && lng == 0 {
		return &Geo{Success: false}
	} else {
		return &Geo{lat, lng, true}
	}
}

func ParseGmapsJson(b []byte) *Geo {
	statusRes := make(map[string]interface{}, 0)
	json.Unmarshal(b, &statusRes)
	status, _ := statusRes["status"].(string)
	if status == "OK" {
		res := make(map[string][]map[string]map[string]map[string]interface{}, 0)
		json.Unmarshal(b, &res)
		lat, _ := res["results"][0]["geometry"]["location"]["lat"].(float64)
		lng, _ := res["results"][0]["geometry"]["location"]["lng"].(float64)
		return &Geo{lat, lng, true}
	} else {
		return &Geo{Success: false}
	}
}
