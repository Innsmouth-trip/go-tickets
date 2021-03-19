package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/liderman/go-aviasales-api.v1"
)

var from string    //Откуда лететь
var to string      //Куда лететь
var DepDate string //месяц вылета (формат YYYY-MM).

func hello() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==================\nДешевые авиабилеты\n==================")

	fmt.Print("Откуда собираешься лететь, друг? : ")
	from, _ = reader.ReadString('\n')
	from = strings.Replace(from, "\n", "", -1)

	fmt.Print("А куда? : ")
	to, _ = reader.ReadString('\n')
	to = strings.Replace(to, "\n", "", -1)

	fmt.Print("А когда? (формат YYYY-MM) : ")
	DepDate, _ = reader.ReadString('\n')
	DepDate = strings.Replace(DepDate, "\n", "", -1)

	from = getCities(from)
	to = getCities(to)

}

var aviaApi = aviasales.NewAviasalesApi("946fd8fc7c91540c230cee20c68b3d5e") //создаем токен апи

func main() {

	hello()
	getPricesCheap()

}

func getPricesCheap() {

	airlines := aviasales.InputPricesCheap{
		Origin:      from,
		Destination: to,
		DepartDate:  DepDate,
		Currency:    "RUB",
	}

	PricesCheap, err := aviaApi.PricesCheap(airlines)
	if err != nil {
		log.Fatal("Запрос вернул: ", err)
		// выводим ошибку чтобы понимать что вернули нил, в документации сказано,
		// что ошибки не будет и вернется пустое значение
	}

	//Сериализовываем данные в структуру airlines

	res, _ := json.MarshalIndent(PricesCheap, "", "  ")
	if err != nil {
		log.Fatal("Ошибка в маршалиндент", err)
	}

	newres := aviasales.DataFlight{}

	err = json.Unmarshal(res, &newres) //создает мапу
	if err != nil {
		fmt.Printf("произошла ошибка при декодировании json. err = %s", err)
		return
	}

	fmt.Printf("%-15v %5v %13v %5v %11v\n", "Компания", "Рейс", "Дата вылета", "Время", "Цена")
	fmt.Println("=====================================================")

	for _, value := range newres.Data {
		for _, val := range value {
			airlineName := getAirlines(val.Airline)
			t, _ := time.Parse(time.RFC3339, val.DepartureAt)
			tf := t.Add(time.Hour * 3).Format("2006-01-02 15:04:05")
			fmt.Printf("%-15v %5v %21v %8v₽\n", airlineName, val.FlightNumber, tf, val.Price)
		}
	}

}

func getAirlines(airname string) string {

	type AirlinesDataStruct []struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	airlinesName, err := aviaApi.DataAirlines()
	if err != nil {
		log.Fatal("Запрос вернул: ", err)
	}

	res, _ := json.MarshalIndent(airlinesName, "", "  ")
	if err != nil {
		log.Fatal("Ошибка в маршалиндент", err)
	}

	var newres = AirlinesDataStruct{}

	err = json.Unmarshal(res, &newres)
	if err != nil {
		fmt.Printf("произошла ошибка при декодировании json. err = %s", err)
	}

	for _, elem := range newres {
		if elem.Code == airname {
			airname = elem.Name
		}
	}

	return airname

}

func getCities(cityName string) string {

	type getCityStruct []struct {
		Code string `json:"code" bson:"code"`
		Name string `json:"name" bson:"name"`
	}

	getcityName, err := aviaApi.DataCities()
	if err != nil {
		log.Fatal("Запрос вернул: ", err)
	}

	res, _ := json.MarshalIndent(getcityName, "", "  ")
	if err != nil {
		log.Fatal("Ошибка в маршалиндент", err)
	}

	var newres = getCityStruct{}

	err = json.Unmarshal(res, &newres)
	if err != nil {
		fmt.Printf("произошла ошибка при декодировании json. err = %s", err)
	}

	for _, elem := range newres {
		if elem.Name == cityName {
			cityName = elem.Code
		}
	}

	return cityName

}
