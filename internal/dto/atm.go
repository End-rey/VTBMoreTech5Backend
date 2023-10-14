package dto

import "time"

type AtmDTO struct {
	ID            int64
	Address       string
	Latitude      float32
	Longitude     float32
	WorkTimeStart time.Time
	WorkTimeEnd   time.Time
	Services      []ServiceDTO
}