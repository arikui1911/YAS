package yas

import (
	"fmt"
)

type StringValue string

func (s StringValue) InspectType() string {
	return "string"
}

func (s StringValue) opBinPlus(other Value) (Value, error) {
	switch o := other.(type) {
	case StringValue:
		return StringValue(s+o), nil
	}
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to string", other, other.InspectType())
}

func (s StringValue) opBinMinus(other Value) (Value, error) {
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to number", s, s.InspectType())
}

func (s StringValue) opBinMul(other Value) (Value, error) {
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to number", s, s.InspectType())
}

func (s StringValue) opBinDiv(other Value) (Value, error) {
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to number", s, s.InspectType())
}

func (s StringValue) opBinMod(other Value) (Value, error) {
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to number", s, s.InspectType())
}
