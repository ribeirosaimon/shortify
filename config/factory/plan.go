package factory

import (
	"time"
)

type plan struct {
	urlExpiredTime time.Time
}

func (m plan) GetPlan() plan {
	return m
}

func newPlan(time time.Time) plan {
	return plan{
		urlExpiredTime: time,
	}
}

func (m plan) TimeToExpired() time.Duration {
	return now.Sub(m.urlExpiredTime)
}
