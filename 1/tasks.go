package lesson2

import (
	"math"
)

// FizzBuzz возвращает FizzBuzz, Fizz, Buzz или пустую строку для обычного числа.
// Для n <= 0 возвращает invalid.
func FizzBuzz(n int) string {
	switch {
	case n <= 0:
		return "invalid"
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return ""
	}
}

// Sum возвращает сумму. Для пустого среза и nil возвращает 0.
func Sum(values []int64) int64 {
	if len(values) == 0 {
		return 0
	}
	var sum int64 = 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

// Min возвращает минимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Min(values []int64) int64 {
	if len(values) == 0 {
		return 0
	}
	var minNumber int64 = math.MaxInt64
	for _, v := range values {
		if v < minNumber {
			minNumber = v
		}
	}
	return minNumber
}

// Max возвращает максимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Max(values []int64) int64 {
	if len(values) == 0 {
		return 0
	}
	var minNumber int64 = math.MinInt64
	for _, v := range values {
		if v > minNumber {
			minNumber = v
		}
	}
	return minNumber
}

// Average возвращает среднее как float64. Для пустого среза и nil возвращает 0.
func Average(values []int64) float64 {
	if len(values) == 0 {
		return 0
	}
	var summ int64 = 0
	for _, v := range values {
		summ += v
	}
	return float64(summ) / float64(len(values))
}

// Category возвращает invalid, small, regular или large по сумме в копейках.
func Category(amount int64) string {
	switch {
	case amount < 0:
		return "invalid"
	case amount <= 9999:
		return "small"
	case amount <= 49999:
		return "regular"
	default:
		return "large"
	}
}

// RuneCount считает кодовые точки Unicode в строке через range.
func RuneCount(text string) int {
	count := 0

	for range text {
		count++
	}

	return count
}
