package transfer

import (
	"errors"
	"github.com/netwar1994/card2card/pkg/card"
	"strconv"
	"strings"
)

const lenOfNumbers = 16

var (
	ErrNotEnoughMoney    = errors.New("not enough money")
	ErrInvalidCardNumber = errors.New("invalid card number")
)

type Service struct {
	CardSvc *card.Service
	Percent int64
	Min     int64
}

func NewService(cardSvc *card.Service, percent int64, min int64) *Service {
	return &Service{CardSvc: cardSvc, Percent: percent, Min: min}
}

func IsValid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")

	if len(number) != lenOfNumbers {
		return false
	}
	numbers := strings.Split(number, "")
	numbersInt := make([]int, lenOfNumbers)

	for i, value := range numbers {
		var err interface{}
		numbersInt[i], err = strconv.Atoi(value)
		if err != nil {
			return false
		}
	}

	return checkLuhn(numbersInt)

}

func checkLuhn(numbers []int) bool {
	checkSum := 0
	for i := 0; i < lenOfNumbers; i += 2 {
		number := numbers[i] * 2

		if number > 9 {
			number -= 9
		}
		numbers[i] = number
	}

	for _, i := range numbers {
		checkSum += i
	}

	if checkSum%10 == 0 {
		return true
	}
	return false
}

func (s *Service) Card2Card(from, to string, amount int64) (int64, error) {
	if !IsValid(from) || !IsValid(to) {
		return amount, ErrInvalidCardNumber
	}

	commission := amount / 100_00 * s.Percent
	if s.Min > commission {
		commission = s.Min
	}

	fromCard, err := s.CardSvc.SearchByNumber(from)
	if err != nil {
		return amount, err
	}

	toCard, err := s.CardSvc.SearchByNumber(to)
	if err != nil {
		return amount, err
	}

	if fromCard != nil && fromCard.Balance < amount {
		return amount, ErrNotEnoughMoney
	}

	fromCard.Balance -= amount + commission
	toCard.Balance += amount + commission

	return amount + commission, nil
}
