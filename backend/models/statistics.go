package models

import "time"

type TrackUser struct {
	ID           uint64 `gorm:"primarykey"`
	VisitedAt    time.Time
	IsFirstTime  bool
	IP           string
	ReferralLink string
}
