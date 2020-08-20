package models

//Coupon ...
type Coupon struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `gorm:"column:name" json:"name"`
	Value       int    `gorm:"column:value" json:"value"`
	Limit       int    `gorm:"column:limit" json:"limit"`
	ExpiredDate int    `gorm:"column:expired_date" json:"expired"`
	CreatedBy   int    `gorm:"column:created_by" json:"created_by"`
	ForUser     int    `gorm:"column:for_user" json:"for_user"`
	ForPhone    string `gorm:"column:for_phone" json:"for_phone"`
	Status      int    `gorm:"column:status" json:"status"`
	CreatedAt   string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   string `gorm:"column:updated_at" json:"updated_at"`
}

//TableName Product table
func (Coupon) TableName() string {
	return "coupons"
}
