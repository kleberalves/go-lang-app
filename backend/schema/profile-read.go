package schema

import "time"

// Composite Key: Type + UserID
type ProfileRead struct {
	Type      int      `gorm:"primaryKey;autoIncrement:false"`
	UserID    int      `gorm:"primaryKey;autoIncrement:false;"`
	User      UserRead `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}
