 
### Программа, которая выдает список дешевых рейсов используя go-aviasales-api.

### Warning!

Программа может не работать по причине изменения структуры Airline в файле data-airlines.go (из пакета https://pkg.go.dev/github.com/liderman/go-aviasales-api)

### Причина изменения:
Функция DataAirlines по умолчанию возвращает слудущую стуктуру:

```go
type Airline struct {
	Name string `json:"name" bson:"name"`
	Alias    string `json:"alias" bson:"alias"`
	Iata     string `json:"iata" bson:"iata"`
	Icao     string `json:"icao" bson:"icao"`
	Callsign string `json:"callsign" bson:"callsign"`
	Country  string `json:"country" bson:"country"`
	IsActive bool   `json:"is_active" bson:"is_active"`
}
```
Хотя [Json](http://api.travelpayouts.com/data/airlines.json), к которому она обращается имеет только 3 поля, а именно: name, code, name_translations.<br />
Поэтому, чтобы не получать пустые или ненужные данные, я закомментировал нунжные строки и добавил Code. 

```go
type Airline struct {
	Name string `json:"name" bson:"name"`
	// Alias    string `json:"alias" bson:"alias"`
	// Iata     string `json:"iata" bson:"iata"`
	// Icao     string `json:"icao" bson:"icao"`
	// Callsign string `json:"callsign" bson:"callsign"`
	// Country  string `json:"country" bson:"country"`
	// IsActive bool   `json:"is_active" bson:"is_active"`
	Code string `json:"code" bson:"code"`
}
```


### ROADMAP 

1) Русский ввод городов 
2) Получение иных данный, а не только дешевых авиабилетов.
3) Вебморда с помощью rest api 
