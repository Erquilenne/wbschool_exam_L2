package main

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// EventMap хранит данные о событиях
var EventMap = &sync.Map{}

// Config представляет конфигурацию сервера.
type Config struct {
	Port int `json:"port"`
}

func main() {

	// Открыть файл config.json
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Декодировать данные из файла в структуру Config
	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/create_event", createEvent)
	http.HandleFunc("/update_event", updateEvent)
	http.HandleFunc("/delete_event", deleteEvent)
	http.HandleFunc("/events_for_day", eventsForDay)
	http.HandleFunc("/events_for_week", eventsForWeek)
	http.HandleFunc("/events_for_month", eventsForMonth)
	port := strconv.Itoa(config.Port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// createEvent обрабатывает POST-запрос для создания события
func createEvent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Println("Error parsing form data:", err)
		return
	}

	userID := r.Form.Get("user_id")
	date := r.Form.Get("date")

	if userID == "" || date == "" {
		http.Error(w, "Missing user_id or date in parameters", http.StatusBadRequest)
		log.Println("Missing user_id or date in parameters")
		return
	}

	// Проверяем, есть ли уже даты для данного user_id
	datesInterface, _ := EventMap.Load(userID)
	dates, ok := datesInterface.([]string)
	if !ok {
		dates = make([]string, 0)
	}

	dates = append(dates, date)
	EventMap.Store(userID, dates)

	json.NewEncoder(w).Encode(map[string]string{
		"result": "Event created successfully",
	})
}

// updateEvent обрабатывает POST-запрос для обновления даты события
func updateEvent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Println("Error parsing form data:", err)
		return
	}

	userID := r.Form.Get("user_id")
	oldDate := r.Form.Get("old_date")
	newDate := r.Form.Get("new_date")

	if userID == "" || oldDate == "" || newDate == "" {
		http.Error(w, "Missing user_id, old_date, or new_date in parameters", http.StatusBadRequest)
		log.Println("Missing user_id, old_date, or new_date in parameters")
		return
	}

	// Проверяем, есть ли уже даты для данного user_id
	datesInterface, _ := EventMap.Load(userID)
	dates, ok := datesInterface.([]string)
	if !ok {
		http.Error(w, "User not found or has no events", http.StatusNotFound)
		log.Printf("User %s not found or has no events", userID)
		return
	}

	// Находим и обновляем старую дату
	found := false
	for i, date := range dates {
		if date == oldDate {
			dates[i] = newDate
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Old date not found for the user", http.StatusNotFound)
		log.Printf("Old date %s not found for user %s", oldDate, userID)
		return
	}

	EventMap.Store(userID, dates)

	json.NewEncoder(w).Encode(map[string]string{
		"result": "Event date updated successfully",
	})
}

// deleteEvent обрабатывает POST-запрос для удаления даты события
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Println("Error parsing form data:", err)
		return
	}

	userID := r.Form.Get("user_id")
	dateToDelete := r.Form.Get("date")

	if userID == "" || dateToDelete == "" {
		http.Error(w, "Missing user_id or date to delete in parameters", http.StatusBadRequest)
		log.Println("Missing user_id or date to delete in parameters")
		return
	}

	// Проверяем, есть ли уже даты для данного user_id
	datesInterface, _ := EventMap.Load(userID)
	dates, ok := datesInterface.([]string)
	if !ok {
		http.Error(w, "User not found or has no events", http.StatusNotFound)
		log.Printf("User %s not found or has no events", userID)
		return
	}

	// Проверяем наличие даты для удаления и удаляем ее из списка
	found := false
	var updatedDates []string
	for _, d := range dates {
		if d != dateToDelete {
			updatedDates = append(updatedDates, d)
		} else {
			found = true
		}
	}

	if !found {
		http.Error(w, "Date to delete not found for the user", http.StatusNotFound)
		log.Printf("Date %s to delete not found for user %s", dateToDelete, userID)
		return
	}

	EventMap.Store(userID, updatedDates)

	json.NewEncoder(w).Encode(map[string]string{
		"result": "Event date deleted successfully",
	})
}

// eventsForDay обрабатывает GET-запрос для получения событий на определенный день
func eventsForDay(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	if userID == "" || date == "" {
		http.Error(w, "Missing user_id or date in parameters", http.StatusBadRequest)
		log.Println("Missing user_id or date in parameters")
		return
	}

	// Проверяем, есть ли уже даты для данного user_id
	datesInterface, _ := EventMap.Load(userID)
	dates, ok := datesInterface.([]string)
	if !ok {
		http.Error(w, "User not found or has no events", http.StatusNotFound)
		log.Printf("User %s not found or has no events", userID)
		return
	}

	// Проверяем, есть ли указанная дата в списке дат для user_id
	var eventDate string
	for _, d := range dates {
		if d == date {
			eventDate = d
			break
		}
	}

	if eventDate == "" {
		http.Error(w, "No event found for the user on the specified date", http.StatusNotFound)
		log.Printf("No event found for user %s on date %s", userID, date)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"event": eventDate,
	})
}

// eventsForWeek обрабатывает GET-запрос для получения событий на неделю, начиная с указанной даты
func eventsForWeek(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	startDate := r.URL.Query().Get("start_date")

	if userID == "" || startDate == "" {
		http.Error(w, "Missing user_id or start_date in parameters", http.StatusBadRequest)
		log.Println("Missing user_id or start_date in parameters")
		return
	}

	// Проверяем, есть ли уже даты для данного user_id
	datesInterface, _ := EventMap.Load(userID)
	dates, ok := datesInterface.([]string)
	if !ok {
		http.Error(w, "User not found or has no events", http.StatusNotFound)
		log.Printf("User %s not found or has no events", userID)
		return
	}

	// Формируем список дней в течение недели, начиная с startDate
	currentDate, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		http.Error(w, "Invalid start_date format", http.StatusBadRequest)
		log.Printf("Invalid start_date format: %s", startDate)
		return
	}

	var eventDays []string
	for day := 0; day < 7; day++ {
		formattedDate := currentDate.Format("2006-01-02")
		for _, d := range dates {
			if d == formattedDate {
				eventDays = append(eventDays, formattedDate)
				break
			}
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	json.NewEncoder(w).Encode(map[string][]string{
		"event_days": eventDays,
	})
}

// eventsForMonth обрабатывает GET-запрос для получения событий на месяц, начиная с указанной даты
func eventsForMonth(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	startDate := r.URL.Query().Get("start_date")

	if userID == "" || startDate == "" {
		http.Error(w, "Missing user_id or start_date in parameters", http.StatusBadRequest)
		log.Println("Missing user_id or start_date in parameters")
		return
	}

	// Проверяем, есть ли уже даты для данного user_id
	datesInterface, _ := EventMap.Load(userID)
	dates, ok := datesInterface.([]string)
	if !ok {
		http.Error(w, "User not found or has no events", http.StatusNotFound)
		log.Printf("User %s not found or has no events", userID)
		return
	}

	// Формируем список дней в течение месяца, начиная с startDate
	currentDate, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		http.Error(w, "Invalid start_date format", http.StatusBadRequest)
		log.Printf("Invalid start_date format: %s", startDate)
		return
	}

	year, month, _ := currentDate.Date()
	lastDayOfMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, currentDate.Location())

	var eventDays []string
	for currentDate.Before(lastDayOfMonth) {
		formattedDate := currentDate.Format("2006-01-02")
		for _, d := range dates {
			if d == formattedDate {
				eventDays = append(eventDays, formattedDate)
				break
			}
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	json.NewEncoder(w).Encode(map[string][]string{
		"event_days": eventDays,
	})
}
