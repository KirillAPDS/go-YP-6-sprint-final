package main

import (
    "log"
    "os"

    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
    srv := server.NewServer(logger)
    logger.Println("Сервер запущен")
    if err := srv.HTTPServer.ListenAndServe(); err != nil {
        logger.Fatalf("Ошибка сервера: %v", err)
    }
}

// Здесь, в функции main() нужно создать логгер, 
// далее создать сервер с помощью вашей функции 
// из пакета server, и запустить его. 
// Если при запуске сервера возникают ошибки, 
// выведите её с помощью логгера на уровне Fatal.

// здесь берем хендлеры (с сервисным слоем)
// и передаем в сервак
// запускаем