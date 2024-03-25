package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"time"

	blinkt "github.com/alexellis/blinkt_go"
)

type hourlyprice []struct {
	Day   string  `json:"day"`
	Hour  int     `json:"hour"`
	Price float64 `json:"price"`
	Zone  string  `json:"zone"`
}

func main() {

	// Get today price request
	todresp, err := http.Get("https://raw.githubusercontent.com/jorgeatgu/apaga-luz/main/public/data/today_price.json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer todresp.Body.Close()
	todbody, _ := io.ReadAll(todresp.Body) // response body is []byte

	var today hourlyprice
	if err := json.Unmarshal(todbody, &today); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// Create Slice with hourly prices and sort ascending
	var m []float64

	for _, rec := range today {
		m = append(m, rec.Price)
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i] < m[j]
	})

	// Get tomorrow price request
	tomresp, err := http.Get("https://raw.githubusercontent.com/jorgeatgu/apaga-luz/main/public/data/tomorrow_price.json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer tomresp.Body.Close()
	tombody, _ := io.ReadAll(tomresp.Body) // response body is []byte

	var tomorrow hourlyprice
	if err := json.Unmarshal(tombody, &tomorrow); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// Append tomorrow details to today
	today = append(today, tomorrow...)

	// Get current hour
	hours, _, _ := time.Now().Clock()

	// Prepare Leds
	brightness := 0.2
	led := blinkt.NewBlinkt(brightness)

	led.Setup()

	blinkt.Delay(100)
	led.Clear()
	led.Show()
	status := false
	pixel := 7
	// Iterage through all hours and light leds from current hour until next 7.
	for _, rec := range today {
		if rec.Hour == hours {
			status = true
		}

		if status {
			result := "Pixel " + strconv.Itoa(pixel) + " - Hour " + strconv.Itoa(rec.Hour)

			if rec.Price < m[8] {
				r := 0
				g := 255
				b := 0
				led.SetPixel(pixel, r, g, b)
				result += " : Green"
			} else if rec.Price < m[16] {
				r := 255
				g := 80
				b := 0
				led.SetPixel(pixel, r, g, b)
				result += " : Orange"
			} else {
				r := 255
				g := 0
				b := 0
				led.SetPixel(pixel, r, g, b)
				result += " : Red"
			}
			fmt.Println(result)
			pixel--
		}
		if pixel == -1 {
			led.Show()
			blinkt.Delay(100)
			break
		}

	}

}
