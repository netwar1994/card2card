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
	card1 := cardSvc.AddCard("visa", "USD", 5_000_00, "5106 2142 4434 5467")
	card2 := cardSvc.AddCard("visa", "USD", 1_000_00, "5106 2145 6743 6901")
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    error
	}{
		{name: "Недостаточно денег",
			fields: fields{
				CardSvc: cardSvc,
				Percent: 5_0,
				Min:     10_0,
			},
			args: args{
				from:   card1.Number,
				to:     card2.Number,
				amount: 6_000_00,
			},
			wantTotal: 6_000_00,
			wantOk:    ErrNotEnoughMoney},
		{name: "Достаточно денег",
			fields: fields{
				CardSvc: cardSvc,
				Percent: 5_0,
				Min:     10_00,
			},
			args: args{
				from:   card1.Number,
				to:     card2.Number,
				amount: 3_000_00,
			},
			wantTotal: 3_015_00,
			wantOk:    nil},
		{name: "Карта получателя не существует",
			fields: fields{
				CardSvc: cardSvc,
				Percent: 5_0,
				Min:     10_00,
			},
			args: args{
				from:   card1.Number,
				to:     "5106 2145 6743 5903",
				amount: 1_000_00,
			},
			wantTotal: 1_000_00,
			wantOk:    card.ErrCardNotFound},
		{name: "Неправильный номер карты",
			fields: fields{
				CardSvc: cardSvc,
				Percent: 5_0,
				Min:     10_00,
			},
			args: args{
				from:   "5106 2145 6743 5908",
				to:     card1.Number,
				amount: 6_000_00,
			},
			wantTotal: 6_000_00,
			wantOk:    ErrInvalidCardNumber},
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

func TestIsValid(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Правильный номер карты",
			args: args{
				number: "5106 2142 4434 5467",
			},
			want: true,
		},
		{name: "Неправильный номер карты",
			args: args{
				number: "5106 2142 4434 5468",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.number); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
