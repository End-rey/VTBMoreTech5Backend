package entity

type AtmsService struct {
	ID        int64 `gorm:"primaryKey"`
	AtmID     int64
	ServiceID int64
}