package entity

import (
    _ "gorm.io/gorm"
    "time"
)

type WorksTime struct {
    ID          int64 `gorm:"primaryKey"`
    IDOffice    int64
    WeeksDay    string
    WorkTimeStart time.Time `gorm:"type:time;not null"`
    WorkTimeEnd   time.Time `gorm:"type:time;not null"`
    IsLegal     bool
    Office      Office `gorm:"foreignKey:IDOffice"`
}

func (wto *WorksTime) TableName() string {
    return "works_time_offices"
}