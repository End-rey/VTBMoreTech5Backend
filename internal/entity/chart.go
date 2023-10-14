package entity

import (
    _ "gorm.io/gorm"
    "time"
)

type Chart struct {
    ID         int64 `gorm:"primaryKey"`
    IDOffice   int64
    WeeksDay   string
    ChartsTime *time.Time `gorm:"type:time"`
    Chart      int
    Office     Office `gorm:"foreignKey:IDOffice"`
}

func (co *Chart) TableName() string {
    return "charts_offices"
}