package main

import "fmt"

// Целевой интерфейс
type Target interface {
	Request() string
}

// Реальная сущность, с которой работаем
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Специфический запрос"
}

// Адаптер, который имплементирует Target и делегирует вызов адаптируемому объекту
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}

	fmt.Println("Результат через адаптер:", adapter.Request()) // Результат через адаптер: Специфический запрос
}
