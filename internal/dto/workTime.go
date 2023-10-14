package dto

import "time"

type WorkTimeDTO struct {
	WeeksDay      string
	WorkTimeStart time.Time
	WorkTimeEnd   time.Time
	IsLegal       bool
}