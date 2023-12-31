# Distance Calculator Go Package

[![Go Reference](https://pkg.go.dev/badge/github.com/daveroberts0321/distancecalculator/.svg)](https://pkg.go.dev/github.com/daveroberts0321/distancecalculator/)

[![Go Reference](https://pkg.go.dev/badge/github.com/daveroberts0321/distancecalculator/blob/main/LICENSE.svg)](https://pkg.go.dev/github.com/daveroberts0321/distancecalculator/blob/main/LICENSE)

The Distance Calculator Go package allows you to easily calculate distances between geographical coordinates using the Haversine formula.

## Installation

To use the Distance Calculator Go package in your Go project, you can follow these steps:

1. Open your terminal or command prompt.

2. Run the following command to install the package:

   ```shell
   go get github.com/daveroberts0321/distancecalculator

## All Coordinates are based on a string Reference point, and float64 Coordinates

```go
    type Coord struct {
        Ref  string  // string reference
        Lat  float64 // latitude of point
        Long float64 //longitude of point
    }
```
Conversions are available in Meters, Miles, and Kilometers. 



## Usage Example

```go
    package main

    import (
        "fmt"

        "github.com/daveroberts0321/distancecalculator"
    )

    func main() {
        // Haversine formula to calculate distance between two points
        fmt.Println("Starting distance calculation")

        // Define coordinates for the origin and destinations
        origin := distancecalculator.Coord{
            Ref:  "reference1",
            Lat:  40.7128,  // Latitude of the origin (New York City)
            Long: -74.0060, // Longitude of the origin (New York City)
        }

        destinations := []distancecalculator.Coord{
            {
                Ref:  "reference2",
                Lat:  34.0522,   // Latitude of destination (Los Angeles)
                Long: -118.2437, // Longitude of destination (Los Angeles)
            },
            // Add more destination coordinates as needed
        }

        // Calculate distances in meters
        distancesInMeters, err := distancecalculator.Meters(origin, destinations)
        if err != nil {
            fmt.Println("Error calculating distances:", err)
            return
        }

        // Calculate distances in miles
        distancesInMiles, err := distancecalculator.Miles(origin, destinations)
        if err != nil {
            fmt.Println("Error calculating distances:", err)
            return
        }

        // Calculate distances in kilometers
        distancesInKilometers, err := distancecalculator.Kilometers(origin, destinations)
        if err != nil {
            fmt.Println("Error calculating distances:", err)
            return
        }

        // Display the results
        for i, dest := range destinations {
            fmt.Printf("Distance from %s to %s:\n", origin.Ref, dest.Ref)
            fmt.Printf("- Meters: %.2f m\n", distancesInMeters[i])
            fmt.Printf("- Miles: %.2f miles\n", distancesInMiles[i])
            fmt.Printf("- Kilometers: %.2f km\n", distancesInKilometers[i])
        }
    }




