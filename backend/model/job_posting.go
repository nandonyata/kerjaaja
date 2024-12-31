package model

import "github.com/shopspring/decimal"

type JobPosting struct {
	Base
	Title          string          `gorm:"type:varchar(50)" json:"title"`
	Description    string          `gorm:"type:text" json:"description"`
	Salary         decimal.Decimal `gorm:"type:decimal(14,2);not null" json:"salary"`
	IsNegotiable   bool            `gorm:"type:boolean;default:false" json:"is_negotiable"`
	Status         string          `gorm:"type:varchar(50)" json:"status"` // Open, Closed, Filled
	Category       string          `gorm:"type:varchar(50)" json:"type"`
	SkillsRequired []string        `gorm:"type:text[]"`
	Tags           []string        `gorm:"type:text[]"`

	// For these fields below
	// we are going to use this API to fetch the values on FE later
	// https://github.com/emsifa/api-wilayah-indonesia
	Province    string `gorm:"type:varchar(100)" json:"province"`
	City        string `gorm:"type:varchar(100)" json:"city"`
	SubDistrict string `gorm:"type:varchar(100)" json:"sub_district"`
}
