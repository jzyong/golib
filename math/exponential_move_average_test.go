package math

import (
	"fmt"
	"math"
	"testing"
)

// 初始化
func TestEMAInitial(t *testing.T) {
	ema := NewExponentialMovingAverage(10)
	ema.Add(3)
	fmt.Printf("%f", ema.Value)
	if ema.Value != 3 {
		t.Errorf("expire 3 get %f", ema.Value)
	}
}

func TestEMAVariance(t *testing.T) {
	ema := NewExponentialMovingAverage(10)
	ema.Add(5)
	ema.Add(6)
	ema.Add(7)
	fmt.Printf("Value=%f Variance=%f \r\n", ema.Value, ema.Variance)
	if math.Abs(ema.Variance-0.6134) > 0.0001 {
		t.Errorf("expire 0.6134 get %f", ema.Variance)
	}
}

func TestEMAValueAndVariance(t *testing.T) {
	ema := NewExponentialMovingAverage(10)
	ema.Add(5)
	ema.Add(6)
	fmt.Printf("Value=%f Variance=%f \r\n", ema.Value, ema.Variance)
	if math.Abs(ema.Value-5.1818) > 0.0001 {
		t.Errorf("expire 5.1818 get %f", ema.Value)
	}
	if math.Abs(ema.Variance-0.1487) > 0.0001 {
		t.Errorf("expire 0.1487 get %f", ema.Variance)
	}
}

func TestEMAStandardDeviation(t *testing.T) {
	ema := NewExponentialMovingAverage(10)
	ema.Add(5)
	ema.Add(600)
	ema.Add(70)
	fmt.Printf("Value=%f Variance=%f  StandardDeviation=%f \r\n", ema.Value, ema.Variance, ema.StandardDeviation)
	if math.Abs(ema.StandardDeviation-208.2470) > 0.0001 {
		t.Errorf("expire 208.2470 get %f", ema.Value)
	}
}

func TestEMAReset(t *testing.T) {
	ema := NewExponentialMovingAverage(10)
	ema.Add(500)
	ema.Add(600)
	ema.Reset()
	ema.Add(5)
	ema.Add(6)
	fmt.Printf("Value=%f Variance=%f  StandardDeviation=%f \r\n", ema.Value, ema.Variance, ema.StandardDeviation)
	if math.Abs(ema.Value-5.1818) > 0.0001 {
		t.Errorf("expire 5.1818 get %f", ema.Value)
	}
	if math.Abs(ema.Variance-0.1487) > 0.0001 {
		t.Errorf("expire 0.1487 get %f", ema.Variance)
	}
}
