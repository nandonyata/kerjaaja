package model

type Role struct {
	Base
	Name string `gorm:"type:varchar(50);not null;unique" json:"name"`
}
