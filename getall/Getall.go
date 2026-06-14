package getall

import "fmt"
import "encoding/json"
import "net/http"
import "time"
import "io"
import "sync"

type WeatherInfo struct {
	Coord      map[string]float64 `json:"coord"`
	Weather    []Weather          `json:"weather"`
	Base       string             `json:"base"`
	Main       map[string]float32 `json:"main"`
	Visibility int                `json:"visibility"`
	Wind       map[string]float32 `json:"wind"`
	Clouds     map[string]int     `json:"clouds"`
	Dt         int                `json:"dt"`
	Sys        Sys                `json:"sys"`
	Timezone   int                `json:"timezone"`
	Id         int                `json:"id"`
	Name       string             `json:"name"`
	Cod        int                `json:"cod"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Sys struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

var weatherEmojis = map[string]string{
	"01d": "☀️", // Clear sky, day
	"01n": "🌙",  // Clear sky, night
	"02d": "⛅",  // Few clouds, day
	"02n": "☁️", // Few clouds, night
	"03d": "☁️", // Scattered clouds
	"03n": "☁️", // Scattered clouds
	"04d": "☁️", // Broken clouds
	"04n": "☁️", // Broken clouds
	"09d": "🌧️", // Shower rain
	"09n": "🌧️", // Shower rain
	"10d": "🌦️", // Rain, day
	"10n": "🌧️", // Rain, night
	"11d": "⛈️", // Thunderstorm
	"11n": "⛈️", // Thunderstorm
	"13d": "❄️", // Snow
	"13n": "❄️", // Snow
	"50d": "🌫️", // Mist
	"50n": "🌫️", // Mist
}
var wg sync.WaitGroup

func Getall(wg *sync.WaitGroup,city string) {
	defer wg.Done()
	Apida := WeatherInfo{}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=9cf72db5170db3291218989dfd351942"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err, "First err")

	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err, "Sec err")

	}

	defer resp.Body.Close()

	responsda, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err, "third err")

	}

	err = json.Unmarshal(responsda, &Apida)
	if err != nil {
		fmt.Println(err, "Unmarshal err")

	}

	if (Apida.Main["temp"]-273.13) > 0 {
	fmt.Printf("%s:		%s		+%.2f, %s\n", Apida.Name, weatherEmojis[Apida.Weather[0].Icon], Apida.Main["temp"]-273.13, Apida.Weather[0].Description)
}else {fmt.Printf("%s:		%s		%.2f, %s\n", Apida.Name, weatherEmojis[Apida.Weather[0].Icon], Apida.Main["temp"]-273.13, Apida.Weather[0].Description)}

}
