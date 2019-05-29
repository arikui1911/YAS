package yas

import (
	"fmt"
	"math"
)

type FloatValue float64

func (f FloatValue) InspectType() string {
	return "float"
}

func (f FloatValue) opBinPlus(other Value) (Value, error) {
	switch o := other.(type) {
	case FloatValue:
		return FloatValue(f+o), nil
	case IntValue:
		return FloatValue(float64(f)+float64(o)), nil
	}
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to float", other, other.InspectType())
}

func (f FloatValue) opBinMinus(other Value) (Value, error) {
	switch o := other.(type) {
	case FloatValue:
		return FloatValue(f-o), nil
	case IntValue:
		return FloatValue(float64(f)-float64(o)), nil
	}
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to float", other, other.InspectType())
}

func (f FloatValue) opBinMul(other Value) (Value, error) {
	switch o := other.(type) {
	case FloatValue:
		return FloatValue(f*o), nil
	case IntValue:
		return FloatValue(float64(f)*float64(o)), nil
	}
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to float", other, other.InspectType())
}

func (f FloatValue) opBinDiv(other Value) (Value, error) {
	switch o := other.(type) {
	case FloatValue:
		return FloatValue(f/o), nil
	case IntValue:
		return FloatValue(float64(f)/float64(o)), nil
	}
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to float", other, other.InspectType())
}

func (f FloatValue) opBinMod(other Value) (Value, error) {
	switch o := other.(type) {
	case FloatValue:
		return FloatValue(math.Mod(float64(f), float64(o))), nil
	case IntValue:
		return FloatValue(math.Mod(float64(f), float64(o))), nil
	}
	return nil, fmt.Errorf("type mismatch: %v(%s) is not compatible to float", other, other.InspectType())
}