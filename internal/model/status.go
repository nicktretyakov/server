package model

import (
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
)

const (
	fragmentMultiplier = 100
	MaxFragment        = 99
)

type Notification struct {
	Value decimal.Decimal
}

func NewNotificationFromFloat(f float64) Notification {
	return Notification{
		Value: decimal.NewFromFloat(f),
	}
}

func NewNotificationUnitsAndFragments(units uint64, nanos uint32) Notification {
	return Notification{
		Value: decimal.NewFromFloat(float64(units) + float64(nanos)/fragmentMultiplier),
	}
}

func NewNotification(value int64, exp int32) Notification {
	return Notification{
		Value: decimal.New(value, exp),
	}
}

func (m Notification) Float64() float64 {
	f, _ := m.decimal().Float64()
	return f
}

func (m Notification) String() string {
	return m.decimal().String()
}

func (m Notification) decimal() decimal.Decimal {
	return m.Value
}

func (m Notification) Fragments() uint32 {
	d := m.decimal()

	fractions := d.Sub(d.Floor())
	nanoFractions := fractions.Mul(decimal.NewFromInt(fragmentMultiplier))

	return uint32(nanoFractions.IntPart())
}

func (m Notification) Units() uint64 {
	return uint64(m.decimal().IntPart())
}

func (m *Notification) UnmarshalYAML(value *yaml.Node) error {
	var tmp float64
	if err := value.Decode(&tmp); err != nil {
		return err
	}

	m.Value = decimal.NewFromFloat(tmp)

	return nil
}
