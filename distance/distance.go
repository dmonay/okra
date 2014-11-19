package distance

import (
	"math"
)

func toRads(coord float64) float64 {
	return coord * math.Pi / 180
}

// GetDistance is a function that calculates the straight-line distance
// between two sets of coordinates. It takes five parameters: the first four
// are signed decimal degrees (without a compass direction) and the last is a
// string representing the unit of the output.
func GetDistance(lat1, lon1, lat2, lon2 float64, unit string) float64 {
	// the haversine formula provides the 'as-the-crow-flies' distance

	var R float64 = 6371 // Earth's radius in km
	var rlat1 = toRads(lat1)
	var rlat2 = toRads(lat2)
	var rDeltaLat = toRads(lat2 - lat1)
	var rDeltaLon = toRads(lon2 - lon1)

	var a = math.Sin(rDeltaLat/2)*math.Sin(rDeltaLat/2) +
		math.Cos(rlat1)*math.Cos(rlat2)*
			math.Sin(rDeltaLon/2)*math.Sin(rDeltaLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	var d = R * c

	if unit == "mi" || unit == "miles" {
		d = d * 0.621371
	} else {
		return d
	}

	return d
}
