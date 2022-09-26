package domain

import "time"

type Car struct {
	ID        int64
	ModelName string
	ExpDate   time.Time
}

func (c *Car) IsInUse() bool {
	return time.Now().Before(c.ExpDate)
}
