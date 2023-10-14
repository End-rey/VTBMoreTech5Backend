package entity

import (
    _ "gorm.io/gorm"
    "time"
)

type WorksTimeOffice struct {
    ID          int64 `gorm:"primaryKey"`
    OfficeID    int64
    WeeksDay    string
    WorkTimeStart time.Time `gorm:"type:time;not null"`
    WorkTimeEnd   time.Time `gorm:"type:time;not null"`
    LunchTimeStart time.Time `gorm:"type:time"`
	LunchTimeEnd   time.Time `gorm:"type:time"`
    IsLegal     bool
    Office      Office `gorm:"foreignkey:OfficeID"`
}