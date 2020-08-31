package models

//Job ...
type Job struct {
	ID           uint   `gorm:"primary_key"`
	FullName     string `gorm:"column:full_name"`
	Phone        string `gorm:"column:phone"`
	Cast         int    `gorm:"column:cast"`
	LinkVideo    string `gorm:"column:link_video"`
	LinkFacebook string `gorm:"column:link_facebook"`
	CreatedAt    string `gorm:"column:created_at"`
	UpdatedAt    string `gorm:"column:updated_at"`
}

//TableName Product table
func (Job) TableName() string {
	return "jobs"
}
