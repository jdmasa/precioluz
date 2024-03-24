package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	// Get request
	resp, err := http.Get("https://raw.githubusercontent.com/jorgeatgu/apaga-luz/main/public/data/today_price.json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body) // response body is []byte

	var today hourlyprice
	if err := json.Unmarshal(body, &today); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// Get request
	resp2, err := http.Get("https://raw.githubusercontent.com/jorgeatgu/apaga-luz/main/public/data/tomorrow_price.json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body2, _ := io.ReadAll(resp2.Body) // response body is []byte

	var tomorrow hourlyprice
	if err := json.Unmarshal(body2, &tomorrow); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// fmt.Println(PrettyPrint(result))
	today = append(today, tomorrow...)
	hours, _, _ := time.Now().Clock()
	// Loop through the data node for the FirstName
	brightness := 0.2
	led := blinkt.NewBlinkt(brightness)

	led.SetClearOnExit(true)

	led.Setup()

	blinkt.Delay(100)
	led.Clear()
	status := false
	pixel := 0
	for _, rec := range today {
		if rec.Hour == hours {
			status = true
		}
		if pixel == 8 {
			status = false
		}
		if status {
			fmt.Println(rec.Hour)
			fmt.Println(rec.Zone)

			switch Zone := rec.Zone; Zone {
			case "valle":
				r := 0
				g := 255
				b := 0
				led.SetPixel(pixel, r, g, b)
			case "llano":
				r := 255
				g := 80
				b := 0
				led.SetPixel(pixel, r, g, b)
			case "punta":
				r := 255
				g := 0
				b := 0
				led.SetPixel(pixel, r, g, b)
			default:
				r := 0
				g := 0
				b := 0
				led.SetPixel(pixel, r, g, b)
			}
			pixel++
		}

	}
	led.Show()

}
