package model

type User struct {
	Base
	FirstName string `gorm:"type:varchar(50)" json:"first_name"`
	LastName  string `gorm:"type:varchar(50)" json:"last_name"`
	Email     string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string `gorm:"type:text;not null" json:"-"`
	RoleID    string `gorm:"type:int;not null"  json:"-"`

	Role *Role `gorm:"foreignKey:RoleID"`
}
