package transfer

import (
	"github.com/netwar1994/card2card/pkg/card"
)

type Service struct {
	CardSvc *card.Service
	Percent int64
	Min int64
}

func NewService(cardSvc *card.Service, percent int64, min int64) *Service {
	return &Service{CardSvc: cardSvc, Percent: percent, Min: min}
}

func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool) {
	commission := amount / 100_00 * s.Percent
	fromCard := s.CardSvc.SearchByNumber(from)
	toCard := s.CardSvc.SearchByNumber(to)

	if fromCard != nil && toCard != nil {
		if fromCard.Balance > amount {
			fromCard.Balance -= amount
			toCard.Balance += amount
			return amount, true
		}
		return amount, false
	}

	if fromCard != nil {
		if fromCard.Balance > amount {
			if s.Min > commission{
				total = amount + s.Min
				fromCard.Balance -= total
				return total, true
			}
			total = amount + commission
			fromCard.Balance -= total
			return total, true
		}
		return amount + commission, false
	}

	if toCard != nil {
		if s.Min > commission {
			toCard.Balance += amount
			return amount + s.Min, true
		}
		toCard.Balance += amount
		return amount + commission, true
	}

	if s.Min > commission {
		return amount + s.Min, true
	}

	return amount + commission, true
}

