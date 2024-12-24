package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func main() {
	// Получаем текущее время через билиотеку
	time, err := ntp.Time("time.google.com")
	//Проверка возникла ли ошибка или при успехе выводим результат
	if err != nil {
		log.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println("Текущая дата и время:", time.Format("2006-01-02 15:04:05"))
	}

}
