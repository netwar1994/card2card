package main

import (
	"fmt"
	"github.com/netwar1994/card2card/pkg/card"
	"github.com/netwar1994/card2card/pkg/transfer"
)

func main() {
	service := card.NewService("Netology Bank")
	card1 := service.AddCard("visa", "USD", 5000_00, "5106 2142 4434 5467")
	card2 := service.AddCard("visa", "USD", 1000_00, "5106 2145 6743 6901")

	//"Коммисия 0.5%, мин сумма 10 руб"
	trans := transfer.NewService(service, 5_0, 10_00)

	//"Недостаточно денег"
	fmt.Println(trans.Card2Card(card1.Number, card2.Number, 6_000_00))

	//"Достаточно денег"
	fmt.Println(trans.Card2Card(card1.Number, card2.Number, 3_000_00))

	//"Карта получателя не существует"
	fmt.Println(trans.Card2Card(card2.Number, "5106 2145 6743 5903", 100_00))

	//"Карта отправителя не существует"
	fmt.Println(trans.Card2Card("5106 2145 6743 5903", card2.Number, 2000_00))

	//"Неправильный номер карты"
	fmt.Println(trans.Card2Card("5106 2145 6743 5908", card2.Number, 2000_00))

	//fmt.Println("Balance of first card2card:", service.SearchByNumber("5106 2100 0000 0001").Balance)
	//fmt.Println("Balance of second card2card:", service.SearchByNumber("5106 2100 0000 0002").Balance)

}
