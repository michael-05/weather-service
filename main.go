package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type position struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type WeatherInfo struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type SubMainInfo struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type WindInfo struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type CloudInfo struct {
	All int `json:"all"`
}

type SysInfo struct {
	ID      int    `json:"id"`
	Type    int    `json:"type"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type MainInfo struct {
	Coord      position
	Weather    []WeatherInfo
	Base       string `json:"base"`
	Main       SubMainInfo
	Visibility int `json:"visibility"`
	Wind       WindInfo
	Clouds     CloudInfo
	Dt         int `json:"dt"`
	Sys        SysInfo
	Timezone   int    `json:"timezone"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Cod        int    `json:"cod"`
}

type ResultInfo struct {
	Weather string `json: "weather"`
	Base string `json: "base"`
}
var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	router := gin.Default()
	router.GET("/get-weather", getWeatherByPosition)

	router.Run("127.0.0.1:3000")
}

func getWeatherByPosition(c *gin.Context) {
	lat := c.Query("lat")
	lon := c.Query("lon")
	url := "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=68f2c51c10fb42b906d3364af7115828"
	
	info := new(MainInfo)
	getJson(url, info)

	result := new(ResultInfo)

	result.Weather = info.Weather[0].Main

	switch  {
	case info.Main.Temp > 286:
		result.Base = "Hot"
	case info.Main.Temp <= 286:
		result.Base = "Cold" 		
	}

	c.IndentedJSON(http.StatusOK, result)
}
