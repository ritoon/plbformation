package main

import (
	"math"
	"testing"
)

func TestMyIntDivide(t *testing.T) {
	data := []struct {
		title  string
		value  myInt
		param  int
		should myInt
		err    error
	}{
		{title: "A", value: 10, param: 2, should: 5, err: nil},
		{"B", 10, 0, 0, ErrDivisionByZero},
	}
	for _, v := range data {
		result, err := v.value.Divide(v.param)
		if result != v.should || err != v.err {
			t.Errorf("for %s got result %d and error %v should got result %d and error %v", v.title, result, err, v.should, v.err)
		}
	}
}

func TestMyIntAdd(t *testing.T) {
	data := []struct {
		title  string
		value  myInt
		param  int
		should myInt
		err    error
	}{
		{title: "A", value: 10, param: 5, should: 15, err: nil},
		{"B", math.MaxInt32, 1, 0, ErrOverflow},
	}
	for _, v := range data {
		result, err := v.value.Add(v.param)
		if result != v.should || err != v.err {
			t.Errorf("for %s got result %d and error %v should got result %d and error %v", v.title, result, err, v.should, v.err)
		}
	}
}

func TestMyIntSub(t *testing.T) {
	data := []struct {
		title  string
		value  myInt
		param  int
		should myInt
		err    error
	}{
		{title: "A", value: 10, param: 5, should: 5, err: nil},
		{"B", math.MinInt32, 1, 0, ErrOverflow},
	}
	for _, v := range data {
		result, err := v.value.Sub(v.param)
		if result != v.should || err != v.err {
			t.Errorf("for %s got result %d and error %v should got result %d and error %v", v.title, result, err, v.should, v.err)
		}
	}
}

func TestMyIntMultiply(t *testing.T) {
	data := []struct {
		title  string
		value  myInt
		param  int
		should myInt
		err    error
	}{
		{title: "A", value: 10, param: 5, should: 50, err: nil},
		{"B", math.MaxInt32, 2, 0, ErrOverflow},
	}
	for _, v := range data {
		result, err := v.value.Multiply(v.param)
		if result != v.should || err != v.err {
			t.Errorf("for %s got result %d and error %v should got result %d and error %v", v.title, result, err, v.should, v.err)
		}
	}
}
