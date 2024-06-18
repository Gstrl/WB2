package transport

import (
	"WB2/develop/dev11/internal/models"
	"WB2/develop/dev11/internal/services"
	"WB2/develop/dev11/internal/services/eventsparse"
	"WB2/develop/dev11/internal/services/timeparse"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type HTTPCalendarHandler struct {
	app *services.Application
}

func NewHTTPCalendarHandler(app *services.Application) HTTPCalendarHandler {
	return HTTPCalendarHandler{app: app}
}

func (h HTTPCalendarHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		type RequestData struct {
			UserID *int    `json:"user_id"`
			Title  *string `json:"title"`
			Date   *string `json:"date"`
		}

		// Чтение тела запроса
		body, err := io.ReadAll(r.Body)
		if err != nil {
			h.mapToResponse(w, http.StatusBadRequest, "", err.Error())
			return
		}
		defer r.Body.Close()

		// Разбор JSON данных
		var data RequestData
		err = json.Unmarshal(body, &data)
		if err != nil {
			h.mapToResponse(w, http.StatusBadRequest, "", err.Error())
			return
		}

		// проверяем содержит ли url необходимые параметры
		if data.Title == nil || data.Date == nil || data.UserID == nil {
			h.mapToResponse(w, http.StatusBadRequest, "", "not all necessary parameters were passed")
			return
		}

		parsDate, err := timeparse.TimeParse(*data.Date)
		if err != nil {
			h.mapToResponse(w, http.StatusServiceUnavailable, "", "date parsing error")
			return
		}

		event := models.NewEvent(*data.UserID, *data.Title, parsDate)
		h.app.CreateEvent.Execute(event)
		h.mapToResponse(w, http.StatusOK, fmt.Sprintf("update event, user_id: %d, date: %s", *data.UserID, parsDate.String()), "")
	default:
		h.mapToResponse(w, http.StatusInternalServerError, "", "sorry, only POST methods are supported.")
	}
}

func (h HTTPCalendarHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		type RequestData struct {
			UserID  *int    `json:"user_id"`
			Title   *string `json:"title"`
			Date    *string `json:"date"`
			EventID *int    `json:"event_id"`
		}

		// Чтение тела запроса
		body, err := io.ReadAll(r.Body)
		if err != nil {
			h.mapToResponse(w, http.StatusBadRequest, "", err.Error())
			return
		}
		defer r.Body.Close()

		// Разбор JSON данных
		var data RequestData
		err = json.Unmarshal(body, &data)
		if err != nil {
			h.mapToResponse(w, http.StatusBadRequest, "", err.Error())
			return
		}

		// проверяем содержит ли url необходимые параметры
		if data.Title == nil || data.Date == nil || data.UserID == nil || data.EventID == nil {
			h.mapToResponse(w, http.StatusBadRequest, "", "not all necessary parameters were passed")
			return
		}

		parsDate, err := timeparse.TimeParse(*data.Date)
		if err != nil {
			h.mapToResponse(w, http.StatusServiceUnavailable, "", "date parsing error")
			return
		}

		event := models.NewEvent(*data.UserID, *data.Title, parsDate)
		err = h.app.UpdateEvent.Execute(*data.UserID, *data.EventID, *event)
		if err != nil {
			h.mapToResponse(w, http.StatusServiceUnavailable, "", "business logic error")
		}
		h.mapToResponse(w, http.StatusOK, fmt.Sprintf("update event, user_id: %d, event_id: %d", *data.UserID, *data.EventID), "")
	default:
		h.mapToResponse(w, http.StatusInternalServerError, "", "sorry, only POST methods are supported.")
	}
}

func (h HTTPCalendarHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		type RequestData struct {
			UserID  *int `json:"user_id"`
			EventID *int `json:"event_id"`
		}

		// Чтение тела запроса
		body, err := io.ReadAll(r.Body)
		if err != nil {
			h.mapToResponse(w, http.StatusBadRequest, "", err.Error())
			return
		}
		defer r.Body.Close()

		// Разбор JSON данных
		var data RequestData
		err = json.Unmarshal(body, &data)
		if err != nil {
			h.mapToResponse(w, http.StatusBadRequest, "", err.Error())
			return
		}

		if data.EventID == nil || data.UserID == nil {
			h.mapToResponse(w, http.StatusBadRequest, "", "user_id or event_id was not passed")
			return
		}

		err = h.app.DeleteEvent.Execute(*data.UserID, *data.EventID)
		if err != nil {
			h.mapToResponse(w, http.StatusServiceUnavailable, "", err.Error())
			return
		}
		h.mapToResponse(w, http.StatusOK, fmt.Sprintf("delete event, user_id: %d, event_id: %d", *data.UserID, *data.EventID), "")
	default:
		h.mapToResponse(w, http.StatusInternalServerError, "", "sorry, only POST methods are supported.")
	}
}

func (h HTTPCalendarHandler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		params := r.URL.Query()

		userID, err := strconv.Atoi(params.Get("user_id"))
		if err != nil {
			h.mapToResponse(w, http.StatusBadRequest, "", "error reading user_id")
			return
		}

		res := h.app.EventsForDay.Execute(userID)
		strRes := eventsparse.EventsPars(res)
		h.mapToResponse(w, http.StatusOK, strRes, "")
	default:
		h.mapToResponse(w, http.StatusInternalServerError, "", "sorry, only GET methods are supported.")
	}
}

func (h HTTPCalendarHandler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		params := r.URL.Query()

		userID, err := strconv.Atoi(params.Get("user_id"))
		if err != nil {
			h.mapToResponse(w, http.StatusBadRequest, "", "error reading user_id")
			return
		}

		res := h.app.EventsForWeek.Execute(userID)
		strRes := eventsparse.EventsPars(res)
		h.mapToResponse(w, http.StatusOK, strRes, "")
	default:
		h.mapToResponse(w, http.StatusInternalServerError, "", "sorry, only GET methods are supported.")
	}
}

func (h HTTPCalendarHandler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		params := r.URL.Query()

		userID, err := strconv.Atoi(params.Get("user_id"))
		if err != nil {
			h.mapToResponse(w, http.StatusBadRequest, "", "error reading user_id")
			return
		}

		res := h.app.EventsForMonth.Execute(userID)
		strRes := eventsparse.EventsPars(res)
		h.mapToResponse(w, http.StatusOK, strRes, "")
	default:
		h.mapToResponse(w, http.StatusInternalServerError, "", "sorry, only GET methods are supported.")
	}
}

// MiddlewareLogger вывод в лог каждый обработанный запрос
func (h HTTPCalendarHandler) MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Логируем информацию о запросе
		log.Printf("Request: %s %s\n", r.Method, r.RequestURI)
		// Передаем запрос следующему обработчику
		next(w, r)
	}
}

// Парсинг ответа
func (h HTTPCalendarHandler) mapToResponse(w http.ResponseWriter, statusCode int, data string, errMessage string) {
	// Задаём JSON формат в Content-Type заголовка ответа
	w.Header().Set("Content-Type", "application/json")

	// Записываем HTTP код статуса ответа в заголовок ответа
	w.WriteHeader(statusCode)

	// Создаём JSON объект для ответа
	response := make(map[string]string)

	// В зависимости от статуса запроса, формируем JSON
	if statusCode >= 200 && statusCode < 300 {
		response["result"] = data
	} else {
		response["error"] = errMessage
	}

	// Преобразуем данные в JSON и записываем в тело ответа
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Если ошибка при кодировании данных, возвращаем HTTP 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CustomRegisterHandlers(router *http.ServeMux, h HTTPCalendarHandler) {
	router.HandleFunc("/create_event", h.MiddlewareLogger(h.CreateEvent))
	router.HandleFunc("/update_event", h.MiddlewareLogger(h.UpdateEvent))
	router.HandleFunc("/delete_event", h.MiddlewareLogger(h.DeleteEvent))
	router.HandleFunc("/events_for_day", h.MiddlewareLogger(h.GetEventsForDay))
	router.HandleFunc("/events_for_week", h.MiddlewareLogger(h.GetEventsForWeek))
	router.HandleFunc("/events_for_month", h.MiddlewareLogger(h.GetEventsForMonth))
}
