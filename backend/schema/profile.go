package schema

import "time"

type Profile struct {
	Type      int `gorm:"primaryKey;autoIncrement:false"`
	UserID    int `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time
}
