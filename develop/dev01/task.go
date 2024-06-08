package printntp

import (
	"fmt"
	"github.com/beevik/ntp"
)

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
