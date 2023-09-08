package distancecalculator

import (
	"errors"
	"math"
)

// license MIT, 2023, Dave Roberts

// Coord is a struct that holds the latitude and longitude of a location
type Coord struct {
	Ref  string
	Lat  float64
	Long float64
}

// CalculateDistances calculates the distances from the origin to the destinations.
func Meters(origin Coord, destinations []Coord) ([]float64, error) {
	// Calculate distances
	distances := make([]float64, len(destinations))
	lat1 := origin.Lat * math.Pi / 180
	long1 := origin.Long * math.Pi / 180
	r := 6378100.0 // Earth radius in METERS

	for i, dest := range destinations {
		lat2 := dest.Lat * math.Pi / 180
		long2 := dest.Long * math.Pi / 180

		// Haversine formula to calculate distance between two points
		h := math.Pow(math.Sin((lat2-lat1)/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin((long2-long1)/2), 2)
		distances[i] = 2 * r * math.Asin(math.Sqrt(h))
	}

	// Error handling
	if len(distances) == 0 {
		return nil, errors.New("no distances calculated")
	}

	return distances, nil
}

// CalculateDistancesMiles calculates the distances from the origin to the destinations in miles.
func Miles(origin Coord, destinations []Coord) ([]float64, error) {
	// Calculate distances in miles
	metersDistances, err := Meters(origin, destinations)
	if err != nil {
		return nil, err // Pass along any error from CalculateDistancesMeters
	}

	milesDistances := make([]float64, len(metersDistances))

	for i, meters := range metersDistances {
		milesDistances[i] = meters * 0.000621371 // Convert meters to miles
	}

	// Error handling
	if len(milesDistances) == 0 {
		return nil, errors.New("no distances calculated")
	}

	return milesDistances, nil
}

// CalculateDistancesKilometers calculates the distances from the origin to the destinations in kilometers.
func Kilometers(origin Coord, destinations []Coord) ([]float64, error) {
	// Calculate distances in kilometers
	metersDistances, err := Meters(origin, destinations)
	if err != nil {
		return nil, err // Pass along any error from Meters
	}

	kilometersDistances := make([]float64, len(metersDistances))

	for i, meters := range metersDistances {
		kilometersDistances[i] = meters * 0.001 // Convert meters to kilometers
	}

	// Error handling
	if len(kilometersDistances) == 0 {
		return nil, errors.New("no distances calculated")
	}

	return kilometersDistances, nil
}
