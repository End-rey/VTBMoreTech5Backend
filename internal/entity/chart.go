package entity

import (
    _ "gorm.io/gorm"
    "time"
)

type ChartsOffices struct {
    ID         int64 `gorm:"primaryKey"`
    OfficeID   int64
    WeeksDay   string
    ChartsTime *time.Time `gorm:"type:time"`
    Chart      int
    Office     Office `gorm:"foreignkey:OfficeID"`
}