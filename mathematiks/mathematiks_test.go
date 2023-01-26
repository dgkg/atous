package mathematiks_test

import (
	"atous/mathematiks"
	"math"
	"testing"
)

func TestMyIntDivide(t *testing.T) {
	data := []struct {
		input    mathematiks.MyInt
		input2   int
		expected int
		err      error
	}{
		{input: 10, input2: 2, expected: 5, err: nil},
		{input: 10, input2: 3, expected: 3, err: nil},
		//{input: 10, input2: 0, expected: 0, err: ErrTryDivideByZero},
	}
	for _, v := range data {
		gotV, gotErr := v.input.Divide(v.input2)
		if gotV != v.expected {
			t.Errorf("got %v, expected %v", gotV, v.expected)
		}
		if v.err == nil && gotErr != nil {
			t.Errorf("got %v, expected %v", gotErr, v.err)
		}
		if v.err != nil && gotErr == nil {
			t.Errorf("got %v, expected %v", gotErr, v.err)
		}
		if v.err != nil && gotErr != nil && gotErr.Error() != v.err.Error() {
			t.Errorf("got %v, expected %v", gotErr, v.err)
		}
	}
}

func TestMyIntAdd(t *testing.T) {
	data := []struct {
		title    string
		input    mathematiks.MyInt
		input2   int
		expected mathematiks.MyInt
		err      error
	}{
		{title: "test max", input: math.MaxInt32, input2: 1, expected: math.MinInt32, err: nil},
	}
	for _, v := range data {
		gotValue, _ := v.input.Add(v.input2)
		if gotValue != v.expected {
			t.Errorf("got %v, expected %v", gotValue, v.expected)
		}
	}
}

func TestMyIntSub(t *testing.T) {
}

func TestMyIntMultiply(t *testing.T) {
}

func FuzzMyIntAddSub(f *testing.F) {
	testparam := []int{1, math.MaxInt64, math.MinInt64}
	for _, tc := range testparam {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	var i mathematiks.MyInt = 0
	f.Fuzz(func(t *testing.T, orig int) {
		i, err := i.Add(orig)
		if err != nil {
			return
		}
		i, err = i.Sub(int(i))
		if err != nil {
			return
		}
		if i != 0 {
			t.Errorf("Before: %v, after: %v", 0, i)
		}
	})
}
