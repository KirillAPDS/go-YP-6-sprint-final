package server

import (
    "log"
    "net/http"
    "time"

    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
    Logger       *log.Logger
    HTTPServer   *http.Server
}

func NewServer(logger *log.Logger) *Server {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", handlers.IndexHandler)
    mux.HandleFunc("/upload", handlers.UploadHandler)

    srv := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
        ErrorLog:     logger,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  15 * time.Second,
    }

    return &Server{
        Logger:       logger,
        HTTPServer:   srv,
    }
}

// В этом пакете вы реализуете функцию для создания http-сервера.

// Алгоритм реализации: 
// 1. Создайте структуру сервера с полями для логгера (log.Logger) 
// и http-сервера (http.Server). 

// 2. Создайте функцию, в которой нужно создать http-роутер. 
// Функция принимает log.Logger и возвращает экземпляр 
// структуры вашего сервера. Зарегистрируйте ваши хендлеры в http-роутере. 
// Создайте экземпляр структуры http.Server. 
// 
// Для настойки вашего сервера используйте следующие поля: 
// Addr — используйте порт 8080. 
// Handler — передайте ваш http-роутер. 
// ErrorLog — передайте ваш логгер. 
// ReadTimeout — таймаут для чтения. 5 секунд. 
// WriteTimeout — таймаут для записи. 10 секунд. 
// IdleTimeout — таймаут ожидания следующего запроса. 15 секунд. 

// Верните ссылку на ваш сервер.