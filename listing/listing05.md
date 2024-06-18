Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

Программа выводит error, потому что, несмотря на то что возвращаемое значение функции test равно nil, 
при присвоении его переменной типа error интерфейсное значение содержит тип *customError, что делает его не равным nil при сравнении.

```