package models

import (
	"time"
)

// Series represents a series in database
type Series struct {
	ID            int       `gorm:"autoIncrement;"`
	Sid           int       `gorm:"not null;"`
	Title         string    `gorm:"type:varchar(150);not null;"`
	Poster        string    `gorm:"type:varchar(150);"`
	EpisodeLength int       `gorm:"not null;"`
	AddedAt       time.Time `gorm:"not null;"`
	Seasons       []Season
	UserID        string `gorm:"not null;"`
}

// SeriesInfo represents user series info
type SeriesInfo struct {
	Duration int       `json:"duration"`
	Seasons  int       `json:"seasons"`
	Episodes int       `json:"episodes"`
	Begin    time.Time `json:"beginAt"`
	End      time.Time `json:"endAt"`
}

// SeriesStat represents some information from series table
type SeriesStat struct {
	Total int `json:"total"`
}

// SeriesAddedYears represents number of series added by years
type SeriesAddedYears struct {
	Added int `json:"added"`
	Total int `json:"total"`
}
