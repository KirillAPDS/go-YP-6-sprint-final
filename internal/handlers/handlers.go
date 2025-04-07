package handlers

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"

    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

// UploadHandler
// обрабатываем загрузку файла
// конвертируем содержимое
// сохраняем результат
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    // Парсим форму 
    // func (r *Request) ParseForm() error

    err := r.ParseForm()

    if err != nil {
        http.Error(w, "Ошибка парсинга формы: " + err.Error(), http.StatusInternalServerError)
        return
    }
    // Получаем файл из формы и отложенно закрываем
    // func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
    // FormFile возвращает первый файл для предоставленного ключа формы.
    file, _, err := r.FormFile("myFile")
    if err != nil {
        http.Error(w, "Ошибка получения файла: " + err.Error(), http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Читаем данные из файла
    data, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, "Ошибка чтения файла: " + err.Error(), http.StatusInternalServerError)
        return
    }

    // Конвертируем данные
    result, err := service.Convert(string(data))
    if err != nil {
        http.Error(w, "Ошибка конвертации: " + err.Error(), http.StatusInternalServerError)
        return
    }

    // Создаем локальный файл
    // func Ext(path string) string
    // Ext возвращает расширение имени файла, используемое path. 
    // Расширение — это суффикс, начинающийся с последней точки 
    // в последнем элементе path; оно пустое, если точки нет.
    ext := filepath.Ext(".txt")
    filename := fmt.Sprintf("%d%s", time.Now().UTC().String(), ext)
    //filename := filepath.Join("%d%s", time.Now().UTC().String(), ext)

    newFile, err := os.Create(filename)
    if err != nil {
        http.Error(w, "Ошибка создания файла: " + err.Error(), http.StatusInternalServerError)
        return
    }
    defer newFile.Close()

    // Записываем результат конвертации в файл.
    _, err = newFile.Write([]byte(result))
    if err != nil {
        http.Error(w, "Ошибка записи в файл: " + err.Error(), http.StatusInternalServerError)
        return
    }

    // Возвращаем результат
    w.Write([]byte("Результат конвертации:\n" + result))
}

// функция для раздачи mainHandler 
// /			GET раздача index.html //serve static files
// /upload		POST тут 
// 

// В этом пакете вы реализуете два хендлера. 
// Для корневого эндпоинта / нужно реализовать хендлер, 
// который возвращает HTML из файла index.html. 

// Второй хендлер для эндпоинта /upload 
// должен выполнять следующие действия: 

// - Парсить html-форму из файла index.html. 
// - Получить файл из формы (не забудьте его закрыть). 
// - Прочитать данные из файла. 
// - Передать эти данные в функцию автоопределения 
// из пакета service, которую вы создали, 
// чтобы получить переконвертируемую строку. 
// - Создать локальный файл. Эта операция обычно небезопасна 
// и так делать не рекомендуется, но в рамках нашего задания 
// хотелось бы более наглядного результата, 
// поэтому мы решились на этот шаг, ради видимого результата. 
// А вообще, обычно используют временные файлы. 
// - Записать в локальный файл результат конвертации строки. 
// Для генерации имени файла вы можете использовать время 
// с помощью time.Now().UTC().String(). 
// Чтобы получить расширения файла, 
// используйте filepath.Ext(). 
// - Вернуть результат конвертации строки. 
// Там, где это необходимо, обработайте возможные ошибки. 
// Статус при возникновении ошибок http.StatusInternalServerError. 