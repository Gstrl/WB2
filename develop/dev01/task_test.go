package printntp

import (
	"testing"
)

func TestNtpServerIsCorrect(t *testing.T) {
	ntpServer = "Hello, Gopher!"
	err := PrintNTP()
	if err == nil {
		t.Error(err)
	}
}
