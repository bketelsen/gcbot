package database

import (
	"strings"

	"gorm.io/gorm"
)

const (
	Platinum = "Platinum"
	Gold     = "Gold"
	Silver   = "Silver"
	Bronze   = "Bronze"
)

// 763772624425320529
// 763772955033206784  763772955033206784
const (
	PlatinumRole = "763772838196281354"
	GoldRole     = "763772955033206784"
	SilverRole   = "763773133861552129"
	BronzeRole   = "763773314174156841"
	SponsorRole  = "763772624425320529"
)

type Sponsor struct {
	gorm.Model
	Name              string
	Level             string
	EmailAddress      string
	Channel           string
	Roles             string
	TicketCodes       []TicketCode
	ScheduledMessages []ScheduledMessage
}

func ListSponsors() ([]*Sponsor, error) {
	var sponsors []*Sponsor
	result := db.Find(&sponsors)
	return sponsors, result.Error
}

type TicketCode struct {
	gorm.Model
	Sponsor    Sponsor
	SponsorID  uint
	TicketCode string
}

func GetPromoRoles(promo string) []string {
	var tc TicketCode
	result := db.Joins("Sponsor").Where("ticket_code = ?", promo).First(&tc)
	if result.Error != nil {
		return []string{}
	}
	return strings.Split(tc.Sponsor.Roles, ",")
}

type ScheduledMessage struct {
	gorm.Model
	Sponsor       Sponsor
	SponsorID     uint
	TargetChannel string
	Message       string
	Cron          string
}

func GetScheduledMessages() ([]ScheduledMessage, error) {
	var sm []ScheduledMessage
	db.Find(&sm)
	return sm, db.Error
}
