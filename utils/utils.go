package utils

import (
	"encoding/json"
	m "math"
	"net/http"
)

//Message - config
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond - config
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * m.Pi / 180
	lo1 = lon1 * m.Pi / 180
	la2 = lat2 * m.Pi / 180
	lo2 = lon2 * m.Pi / 180

	r = 6378.1 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + m.Cos(la1)*m.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * m.Asin(m.Sqrt(h))
}

func hsin(theta float64) float64 {
	return m.Pow(m.Sin(theta/2), 2)
}
