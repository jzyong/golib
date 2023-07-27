package math

import "math"

// ExponentialMovingAverage 指数移动平均线 ,复制于Mirror ExponentialMovingAverage
// https://en.wikipedia.org/wiki/Exponential_smoothing
type ExponentialMovingAverage struct {
	alpha             float64
	initialized       bool
	Value             float64
	Variance          float64
	StandardDeviation float64 // absolute value, see test
}

// NewExponentialMovingAverage 新建
func NewExponentialMovingAverage(n int) *ExponentialMovingAverage {
	// standard N-day EMA alpha calculation
	ema := &ExponentialMovingAverage{
		alpha:             2.0 / (float64(n) + 1),
		initialized:       false,
		Value:             0,
		Variance:          0,
		StandardDeviation: 0,
	}
	return ema
}

func (ema *ExponentialMovingAverage) Add(newValue float64) {
	// simple algorithm for EMA described here:
	// https://en.wikipedia.org/wiki/Moving_average#Exponentially_weighted_moving_variance_and_standard_deviation
	// https://en.wikipedia.org/wiki/Exponential_smoothing  公式参考
	if ema.initialized {
		delta := newValue - ema.Value
		ema.Value += ema.alpha * delta
		ema.Variance = (1 - ema.alpha) * (ema.Variance + ema.alpha*delta*delta) //方差
		ema.StandardDeviation = math.Sqrt(ema.Variance)                         //标准差
	} else {
		ema.Value = newValue
		ema.initialized = true
	}
}

func (ema *ExponentialMovingAverage) Reset() {
	ema.initialized = false
	ema.Value = 0
	ema.Variance = 0
	ema.StandardDeviation = 0
}
