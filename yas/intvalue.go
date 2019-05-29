package yas

import (
	"fmt"
)

type IntValue int

func (i IntValue) InspectType() string {
	return "int"
}

func (i IntValue) opBinPlus(other Value) (Value, error) {
	switch o := other.(type) {
	case IntValue:
		return IntValue(i+o), nil
	case FloatValue:
		return FloatValue(float64(i)+float64(o)), nil
	}
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to integer", other, other.InspectType())
}

func (i IntValue) opBinMinus(other Value) (Value, error) {
	switch o := other.(type) {
	case IntValue:
		return IntValue(i-o), nil
	case FloatValue:
		return FloatValue(float64(i)-float64(o)), nil
	}
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to integer", other, other.InspectType())
}

func (i IntValue) opBinMul(other Value) (Value, error) {
	switch o := other.(type) {
	case IntValue:
		return IntValue(i*o), nil
	case FloatValue:
		return FloatValue(float64(i)*float64(o)), nil
	}
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to integer", other, other.InspectType())
}

func (i IntValue) opBinDiv(other Value) (Value, error) {
	switch o := other.(type) {
	case IntValue:
		if o == 0 {
			return nil, fmt.Errorf("integer divided by zero")
		}
		return IntValue(i/o), nil
	case FloatValue:
		return FloatValue(float64(i)/float64(o)), nil
	}
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to integer", other, other.InspectType())
}

func (i IntValue) opBinMod(other Value) (Value, error) {
	switch o := other.(type) {
	case IntValue:
		if o == 0 {
			return nil, fmt.Errorf("integer divided by zero")
		}
		return IntValue(i%o), nil
	case FloatValue:
		return FloatValue(i).opBinMod(other)
	}
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to integer", other, other.InspectType())
}
