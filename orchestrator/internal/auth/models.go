package auth

import "time"

type User struct {
	ID             int
	Name           string
	AmountOfAgents int
	TimeToCalc     time.Time
}
