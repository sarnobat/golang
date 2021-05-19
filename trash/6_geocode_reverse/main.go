package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/kellydunn/golang-geo"
)

type googleGeocodeResponse struct {
    Results []struct {
        AddressComponents []struct {
            LongName  string   `json:"long_name"`
            Types     []string `json:"types"`
        } `json:"address_components"`
    }
}

func main() {
    p := geo.NewPoint(49.014, 8.4043)
    geocoder := new(geo.GoogleGeocoder)
    geo.HandleWithSQL()
    data, err := geocoder.Request(fmt.Sprintf("latlng=%f,%f", p.Lat(), p.Lng()))
    if err != nil {
        log.Println(err)
    }
    var res googleGeocodeResponse
    if err := json.Unmarshal(data, &res); err != nil {
        log.Println(err)
    }
    var city string
    if len(res.Results) > 0 {
        r := res.Results[0]
    outer:
        for _, comp := range r.AddressComponents {
            // See https://developers.google.com/maps/documentation/geocoding/#Types
            // for address types
            for _, compType := range comp.Types {
                if compType == "locality" {
                    city = comp.LongName
                    break outer
                }
            }
        }
    }
    fmt.Printf("City: %s\n", city)
}
