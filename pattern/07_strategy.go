package main

import "fmt"

/*
Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Применимость паттерна "стратегия":
1. Когда в одном и том же месте в зависимости от текущего состояния системы (или её окружения) должны использоваться различные алгоритмы
2. Когда в наличии много похожих классов, которые отличаются только способом выполнения некоторого поведения
3. Когда необходимо изолировать бизнес-логику класса от деталей реализации алгоритмов
4. Когда необходимо заменить массивный условный оператор, который переключается между различными вариантами одного и того же алгоритма

Плюсы паттерна "стратегия":
1. Отказ от использования переключателей и/или условных операторов
2. Принцип открытости/закрытости: Добавление новых стратегий не требует изменения кода контекста
3. Улучшение тестируемости: Каждая стратегия может быть протестирована независимо.

Минусы паттерна "стратегия":
1. Увеличение числа классов: Каждая стратегия требует свой класс, что может увеличить количество классов в системе.
2. Код может стать более сложным: В программе с небольшим количеством редко меняющихся алгоритмов, применение паттерна может привести к избыточности кода и усложнению его структуры
*/

type IDBconnection interface {
	Connect()
}

type SqlConnection struct {
	connectionString string
}

func (con SqlConnection) Connect() {
	fmt.Println(("Sql " + con.connectionString))
}

type OracleConnection struct {
	connectionString string
}

func (con OracleConnection) Connect() {
	fmt.Println("Oracle " + con.connectionString)
}

type DBConnection struct {
	db IDBconnection
}

func (con DBConnection) DBConnect() {
	con.db.Connect()
}

func main() {
	sqlConnection := SqlConnection{connectionString: "Connection is connected"}
	con := DBConnection{db: sqlConnection}
	con.DBConnect()
	orcConnection := OracleConnection{connectionString: "Connection is connected"}
	con2 := DBConnection{db: orcConnection}
	con2.DBConnect()
}
