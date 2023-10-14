package entity

type OfficeService struct {
	ID        int64 `gorm:"primaryKey"`
	OfficeID  int64
	ServiceID int64
}

func (os *OfficeService) TableName() string {
    return "offices_services"
}