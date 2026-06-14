# Weather Checker

Простой CLI-проект на Go, который получает текущую погоду для нескольких городов одновременно с помощью OpenWeather API.

## Возможности

- Получение погоды через OpenWeather API
- Параллельные запросы с использованием Goroutines
- Синхронизация через WaitGroup
- Парсинг JSON ответов
- Вывод температуры и погодных условий
- Использование собственного пакета внутри проекта

## Технологии

- Go
- net/http
- encoding/json
- sync.WaitGroup
- Goroutines
- OpenWeather API

## Пример вывода

```text
Moscow:      ☁️      +18.24, overcast clouds
Berlin:      ☀️      +23.15, clear sky
Tokyo:       🌧️      +27.41, light rain
Brasov:      ⛅      +20.18, few clouds
Odessa:      ☀️      +26.03, clear sky

247.832ms
```

## Структура проекта

```text
weather-checker/
│
├── main.go
├── go.mod
│
└── getall/
    └── Getall.go
```

## Как работает

1. В `main.go` создаются несколько goroutine.
2. Каждая goroutine отправляет запрос к OpenWeather API.
3. Ответ приходит в формате JSON.
4. JSON преобразуется в структуры Go через `json.Unmarshal`.
5. Результаты выводятся в консоль.
6. `WaitGroup` ожидает завершения всех запросов.

## Запуск

Склонируйте репозиторий:

```bash
git clone <repo-url>
cd weather-checker
```

Установите Go и запустите проект:

```bash
go run .
```

## Используемый API

Проект использует OpenWeather API:

https://openweathermap.org/api

## Что я изучил в этом проекте

- Struct
- Nested Struct
- Maps
- JSON Marshal / Unmarshal
- HTTP Requests
- Packages
- Goroutines
- WaitGroup
- Работа с внешними API
- Обработка ошибок

## Планы по улучшению

- [ ] Ввод города через консоль
- [ ] Вывод влажности воздуха
- [ ] Вывод скорости ветра
- [ ] Конфигурация через .env
- [ ] Улучшенная обработка ошибок
- [ ] Поддержка любого количества городов
- [ ] Юнит-тесты

## Автор

Мой первый проект на Go!;3
P.S API указанный в проекте не работает!

## ENG

# Weather Checker

A simple CLI application written in Go that fetches current weather data for multiple cities simultaneously using the OpenWeather API.

## Features

- Fetches weather data from OpenWeather API
- Concurrent requests using Goroutines
- Synchronization with WaitGroup
- JSON parsing with Go structs
- Displays temperature and weather conditions
- Uses a custom package for API interactions

## Technologies

- Go
- net/http
- encoding/json
- sync.WaitGroup
- Goroutines
- OpenWeather API

## Example Output

```text
Moscow:      ☁️      +18.24°C, overcast clouds
Berlin:      ☀️      +23.15°C, clear sky
Tokyo:       🌧️      +27.41°C, light rain
Brasov:      ⛅      +20.18°C, few clouds
Odessa:      ☀️      +26.03°C, clear sky

247.832ms
```

## Project Structure

```text
weather-checker/
│
├── main.go
├── go.mod
│
└── getall/
    └── Getall.go
```

## How It Works

1. Multiple Goroutines are started from `main.go`.
2. Each Goroutine sends a request to the OpenWeather API.
3. The API responds with JSON data.
4. JSON is parsed into Go structs using `json.Unmarshal`.
5. Weather information is printed to the console.
6. A WaitGroup waits for all requests to complete.

## Getting Started

Clone the repository:

```bash
git clone <repository-url>
cd weather-checker
```

Run the project:

```bash
go run .
```

## API

This project uses the OpenWeather API:

https://openweathermap.org/api

## What I Learned

- Structs
- Nested Structs
- Maps
- JSON Marshal / Unmarshal
- HTTP Requests
- Packages
- Goroutines
- WaitGroup
- Working with External APIs
- Error Handling

## Future Improvements

- [ ] User input for city names
- [ ] Display humidity
- [ ] Display wind speed
- [ ] Configuration via environment variables
- [ ] Better error handling
- [ ] Support for unlimited cities
- [ ] Unit tests
- [ ] Colored terminal output

## Author

**UsernameNotDefine**

My first Go project;3
P.S API in a project doesnt work, change it on yours!
