module printNTP

go 1.22

require github.com/beevik/ntp v1.4.3

require (
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
)

//тк как модуль локальный при добавление в go.mod, replace printNTP/PrintNTP => ./develop/dev01
