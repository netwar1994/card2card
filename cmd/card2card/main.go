package main

import (
	"fmt"
	"github.com/netwar1994/card2card/pkg/card"
	"github.com/netwar1994/card2card/pkg/transfer"
)

func main() {
	service := card.NewService("Netology Bank")
	card1 := service.AddCard("visa", "USD", 5000_00, "0001")
	card2 := service.AddCard("visa", "USD", 1000_00, "0002")

	//"Карта своего банка -> Карта своего банка (денег достаточно). Без комиссии"
	transNetToNet := transfer.NewService(service, 0, 0)
	fmt.Println(transNetToNet.Card2Card(card1.Number, card2.Number, 300_00))

	//"Карта своего банка -> Карта чужого банка (денег достаточно). Комиссия 0,5"
	transNetToOther := transfer.NewService(service, 5_0, 10_00)
	fmt.Println(transNetToOther.Card2Card(card2.Number, "0003", 100_00))
	//"Карта своего банка -> Карта чужого банка (денег недостаточно). Комиссия 0,5"
	fmt.Println(transNetToOther.Card2Card(card2.Number, "0003", 2000_00))

	//"Карта чужого банка -> Карта своего банка (денег достаточно). Комиссия 0,5"
	transOtherToNet := transfer.NewService(service, 5_0, 10_00)
	fmt.Println(transOtherToNet.Card2Card("0003", card2.Number, 50_000_00))

	//"Карта чужого банка -> Карта чужого банка (денег достаточно). Комиссия 1,5"
	transOtherToOther := transfer.NewService(service, 15_0, 30_00)
	fmt.Println(transOtherToOther.Card2Card("0003", "0004", 3000_00))

	fmt.Println("Balance of first card2card:", service.SearchByNumber("0001").Balance)
	fmt.Println("Balance of second card2card:", service.SearchByNumber("0002").Balance)

}
