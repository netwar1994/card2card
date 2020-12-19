package transfer

import (
	"github.com/netwar1994/card2card/pkg/card"
)

type Service struct {
	CardSvc *card.Service
	Percent int64
	Min     int64
}

func NewService(cardSvc *card.Service, percent int64, min int64) *Service {
	return &Service{CardSvc: cardSvc, Percent: percent, Min: min}
}

func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool) {
	commission := amount / 100_00 * s.Percent
	fromCard := s.CardSvc.SearchByNumber(from)
	toCard := s.CardSvc.SearchByNumber(to)

	if fromCard != nil && fromCard.Balance < amount {
		return amount, false
	}

	if s.Min > commission {
		commission = s.Min
	}

	if fromCard != nil && toCard != nil {
		fromCard.Balance -= amount
		toCard.Balance += amount
		return amount, true
	}

	if fromCard != nil {
		total = amount + commission
		fromCard.Balance -= total
		return total, true
	}

	if toCard != nil {
		toCard.Balance += amount
		return amount + commission, true
	}

	return amount + commission, true
}
