package printntp

import (
	"fmt"
	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP-библиотеки. Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

var ntpServer = "0.beevik-ntp.pool.ntp.org"

func PrintNTP() error {
	// Получаем точное время с использованием NTP
	time, err := ntp.Time(ntpServer)
	if err != nil {
		return err
	}

	// Печатаем текущее время
	fmt.Printf("Точное время: %d:%d:%d", time.Hour(), time.Minute(), time.Second())
	return nil
}
