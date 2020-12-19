package card

import (
	"errors"
	"strings"
)

type Service struct {
	BankName string
	Cards    []*Card
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

type Card struct {
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

func (s *Service) AddCard(issuer string, currency string, balance int64, number string) *Card {
	card := &Card{
		Issuer:   issuer,
		Balance:  balance,
		Currency: currency,
		Number:   number,
	}
	s.Cards = append(s.Cards, card)
	return card
}

var ErrCardNotFound = errors.New("card not found")

func (s *Service) SearchByNumber(number string) (*Card, error) {
	const prefix = "5106 21"

	if strings.HasPrefix(number, prefix) == true {
		for _, card := range s.Cards {
			if card.Number == number {
				return card, nil
			}
		}
	}

	return nil, ErrCardNotFound
}
