package database

import (
	"time"
)

type Greeting struct {
	UserID    string `gorm:"primaryKey"`
	Channel   string
	GreetedAt time.Time
}

func NewGreeting(id, channel string) error {
	g := &Greeting{
		UserID:    id,
		Channel:   channel,
		GreetedAt: time.Now(),
	}
	result := db.Create(g)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Greeted(id string) (bool, error) {
	var greet Greeting
	result := db.Where("user_id= ?", id).Find(&greet)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, db.Error
}
