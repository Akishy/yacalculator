package auth

import "time"

type User struct {
	ID             int
	Name           string
	Password       string
	IsAdmin        bool
	AmountOfAgents int
	TimeToCalc     time.Duration
}
