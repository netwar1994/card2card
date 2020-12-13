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
	if s.CardSvc.SearchByNumber(from) != nil && s.CardSvc.SearchByNumber(to) != nil {
		if s.CardSvc.SearchByNumber(from).Balance > amount {
			s.CardSvc.SearchByNumber(from).Balance -= amount
			s.CardSvc.SearchByNumber(to).Balance += amount
			total = amount
			ok = true
			return
		}
		total = amount
		ok = false
		return
	}

	if s.CardSvc.SearchByNumber(from) != nil {
		if s.CardSvc.SearchByNumber(from).Balance > amount {
			if s.Min > (amount / 100_00 * s.Percent){
				total = amount + s.Min
				s.CardSvc.SearchByNumber(from).Balance -= total
				ok = true
				return
			}
			total = amount + (amount / 100_00 * s.Percent)
			s.CardSvc.SearchByNumber(from).Balance -= total
			ok = true
			return
		}
		total = amount
		ok = false
		return
	}

	if s.CardSvc.SearchByNumber(to) != nil {
		if s.Min > (amount / 100_00 * s.Percent) {
			s.CardSvc.SearchByNumber(to).Balance += amount
			total = amount + s.Min
			ok = true
			return
		}
		s.CardSvc.SearchByNumber(to).Balance += amount
		total = amount + (amount / 100_00 * s.Percent)
		ok = true
		return
	}

	if s.Min > (amount / 100_00 * s.Percent) {
		total = amount + s.Min
		ok = true
		return
	}
	total = amount + (amount / 100_00 * s.Percent)
	ok = true
	return
}