package entity

type OfficesService struct {
	ID        int64 `gorm:"primaryKey"`
	OfficeID  int64
	ServiceID int64
}