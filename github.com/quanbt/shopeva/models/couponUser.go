package models

//CouponUser ...
type CouponUser struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	CouponID  int    `gorm:"column:coupon_id" json:"coupon_id"`
	FullName  string `gorm:"column:full_name" json:"full_name"`
	Phone     string `gorm:"column:phone" json:"phone"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}

//TableName Product table
func (CouponUser) TableName() string {
	return "coupon_users"
}
