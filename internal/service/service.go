
package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	hasRussianLetters := strings.Contains(input, "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")
	if hasRussianLetters {
		return morse.ToMorse(input), nil
	}
	return morse.ToText(input), nil
}
