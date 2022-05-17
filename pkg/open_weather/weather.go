package openweather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type City struct {
	Name    string  `json:"name"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
	Country string  `json:"country"`
}

//Вся погода
type MainWeather struct {
	Lat            float64  `json:"lat"`
	Lon            float64  `json:"lon"`
	Timezone       string   `json:"timezone"`
	TimezoneOffset int      `json:"timezone_offset"`
	Current        Current  `json:"current"`
	Hourly         []Hourly `json:"hourly"`
	Daily          []Daily  `json:"daily"`
}

//Текущая погода
type Current struct {
	Dt         int       `json:"dt"`
	Sunrise    int       `json:"sunrise"`
	Sunset     int       `json:"sunset"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`  //Атмосферное давление
	Humidity   int       `json:"humidity"`  //Влажность
	DewPoint   float64   `json:"dew_point"` //Атмосферная температура
	Uvi        float64   `json:"uvi"`
	Clouds     int       `json:"clouds"` //Облачность %
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"` //Скорость ветра
	WindDeg    int       `json:"wind_deg"`   //Направлоение ветра
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
	Rain       Rain      `json:"rain"` //Кол-во дождя если есть
}

//Погода по часам 48ч
type Hourly struct {
	Dt         int       `json:"dt"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        float64   `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    int       `json:"wind_deg"`
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
	Pop        float64   `json:"pop"`
	Rain       Rain      `json:"rain,omitempty"`
}

//Погода на 7 дней + 1 день такущей
type Daily struct {
	Dt        int       `json:"dt"`
	Sunrise   int       `json:"sunrise"`
	Sunset    int       `json:"sunset"`
	Moonrise  int       `json:"moonrise"`
	Moonset   int       `json:"moonset"`
	MoonPhase float64   `json:"moon_phase"`
	Temp      Temp      `json:"temp"`
	FeelsLike FeelsLike `json:"feels_like"`
	Pressure  int       `json:"pressure"`
	Humidity  int       `json:"humidity"`
	DewPoint  float64   `json:"dew_point"`
	WindSpeed float64   `json:"wind_speed"`
	WindDeg   int       `json:"wind_deg"`
	WindGust  float64   `json:"wind_gust"`
	Weather   []Weather `json:"weather"`
	Clouds    int       `json:"clouds"`
	Pop       float64   `json:"pop"`
	Rain      float64   `json:"rain,omitempty"`
	Uvi       float64   `json:"uvi"`
}

//Кол-во осадков
type Rain struct {
	OneH float64 `json:"1h"`
}

type Temp struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type FeelsLike struct {
	Day   float64 `json:"day"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func (c City) GetWeather() (*MainWeather, error) {
	api := get_api_weather()
	link := fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?lat=%f&lon=%f&units=%s&exclude=%s&lang=%s&appid=%s", c.Lat, c.Lon, "metric", "alerts", "ru", api)

	r, err := http.Get(link)
	if err != nil {
		log.Println("Error get weather:", err)
		return nil, err
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error read body:", err)
		return nil, err
	}

	var weather MainWeather

	if err := json.Unmarshal(b, &weather); err != nil {
		log.Println("Eror parse body:", err)
		return nil, err
	}

	return &weather, nil

}

func GetCity(city_name string) (*City, error) {
	api := get_api_weather()

	link := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", city_name, api)

	r, err := http.Get(link)
	if err != nil {
		log.Println("Error get weather:", err)
		return nil, err
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error read body:", err)
		return nil, err
	}

	if err := json.Valid(body); !err {
		log.Println("Json valid:", err)
		return nil, nil
	}

	var c []City
	if err := json.Unmarshal(body, &c); err != nil {
		log.Println("Eror parse body:", err)
		return nil, err
	}

	return &c[0], nil

}

func get_api_weather() string {
	f, err := os.ReadFile("./Data/open_weather_key.txt")
	if err != nil {
		log.Println("Error read file:")
		log.Fatal(err)

	}

	return string(f)
}
