package model

import "time"

type Vote struct {
	ID           string
	CreationDate time.Time
	Category     string
}
