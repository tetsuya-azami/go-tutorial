package domain

import "time"

type Clock interface {
	Now() time.Time
}

type MyClock struct{}

func (_ *MyClock) Now() time.Time {
	return time.Now()
}

type FixedClock struct {
	FixedTime time.Time
}

func (fc *FixedClock) Now() time.Time {
	return fc.FixedTime
}
