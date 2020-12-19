package transfer

import (
	"github.com/netwar1994/card2card/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc *card.Service
		Percent int64
		Min     int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	cardSvc := card.NewService("Our Bank")
	card1 := cardSvc.AddCard("visa", "USD", 5_000_00, "0001")
	card2 := cardSvc.AddCard("visa", "USD", 1_000_00, "0002")
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    bool
	}{
		{name: "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields: fields{
				CardSvc: cardSvc,
				Percent: 5_0,
				Min:     10_0,
			},
			args: args{
				from:   card1.Number,
				to:     card2.Number,
				amount: 1_000_00,
			},
			wantTotal: 1_000_00,
			wantOk:    true},
		{name: "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields: fields{
				CardSvc: cardSvc,
				Percent: 5_0,
				Min:     10_00,
			},
			args: args{
				from:   card1.Number,
				to:     card2.Number,
				amount: 6_000_00,
			},
			wantTotal: 6_000_00,
			wantOk:    false},
		{name: "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields: fields{
				CardSvc: cardSvc,
				Percent: 5_0,
				Min:     10_00,
			},
			args: args{
				from:   card1.Number,
				to:     "0003",
				amount: 1_000_00,
			},
			wantTotal: 1_010_00,
			wantOk:    true},
		{name: "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields: fields{
				CardSvc: cardSvc,
				Percent: 5_0,
				Min:     10_00,
			},
			args: args{
				from:   card1.Number,
				to:     "0003",
				amount: 6_000_00,
			},
			wantTotal: 6_000_00,
			wantOk:    false},
		{name: "Карта чужого банка -> Карта чужого банка",
			fields: fields{
				CardSvc: cardSvc,
				Percent: 15_0,
				Min:     30_00,
			},
			args: args{
				from:   "0004",
				to:     "0003",
				amount: 6_000_00,
			},
			wantTotal: 6_090_00,
			wantOk:    true},
		{name: "Карта чужого банка -> Карта чужого банка",
			fields: fields{
				CardSvc: cardSvc,
				Percent: 15_0,
				Min:     30_00,
			},
			args: args{
				from:   "0004",
				to:     "0003",
				amount: 1_00_00,
			},
			wantTotal: 1_30_00,
			wantOk:    true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc: tt.fields.CardSvc,
				Percent: tt.fields.Percent,
				Min:     tt.fields.Min,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
