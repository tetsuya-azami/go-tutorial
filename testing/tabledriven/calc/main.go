package calc

import (
	"errors"
	"time"
)

const defaultFee = 1000

func Fee(admissionTime time.Time) (int, error) {
	fee := float64(defaultFee)
	hour := admissionTime.Hour()

	switch {
	case hour < 2:
		fee *= 0.9
	case 2 <= hour && hour < 5:
		return 0, errors.New("現在は入場できない時間帯です")
	case 5 <= hour && hour < 8:
		fee *= 0.9
	case 8 <= hour && hour < 22:
		fee *= 1.0
	case 22 < hour:
		fee *= 1.2
	}

	return int(fee), nil
}
