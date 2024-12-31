package model

// UserActivity represents a log of activities a user has taken
type UserActivity struct {
	Base
	UserID       string `gorm:"type:uuid" json:"user_id"`
	JobPostingID string `gorm:"type:uuid" json:"job_posting_id"`

	ActivityType string `gorm:"type:varchar(255);not null"` // Type of activity (e.g., "Applied", "Viewed")
	ActionStatus string `gorm:"type:varchar(50)"`           // Action status (e.g., "Success", "Failed")

	User       *User       `gorm:"foreignKey:UserID"`
	JobPosting *JobPosting `gorm:"foreignKey:JobPostingID"`
}
