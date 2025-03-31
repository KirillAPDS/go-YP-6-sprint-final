package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(s string) bool {
	for _, ch := range s {
		if ch != '.' && ch != '-' && ch != ' ' {
			return false
		}
	}
	return true
}

func Convert(s string) (string, error) {
	trim := strings.TrimSpace(s)
	if len(trim) == 0 {
		return "", errors.New("Пустая строка")
	}

	if isMorse(trim) {
		return morse.ToText(trim), nil
	}
	
	return morse.ToMorse(trim), nil
}

// В этом пакете вы реализуете функцию автоматического определения 
// кода Морзе или обычного текста из переданной строки. 
// Если передан обычный текст, функция должна переконвертировать его 
// в код Морзе и вернуть; и наоборот — если был передан код Морзе, 
// функция должна переконвертировать его в обычный текст и вернуть. 
// Для реализации этой функции придётся обратиться 
// к стандартной библиотеке, а именно — к пакету strings. 
// В этом пакете есть хорошие примеры, которые демонстрируют, 
// как можно решить эту задачу. 
// Не забудьте обработать ошибки и вернуть их.
