package schema

import "time"

// Composite Key: Type + UserID
type Profile struct {
	Type      int  `gorm:"primaryKey;autoIncrement:false"`
	UserID    int  `gorm:"primaryKey;autoIncrement:false;"`
	User      User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}
