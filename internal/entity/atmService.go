package entity

type AtmService struct {
	ID        int64 `gorm:"primaryKey"`
	AtmID     int64
	ServiceID int64
}

func (as *AtmService) TableName() string {
    return "atms_services"
}