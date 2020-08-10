package models

//User ...
type User struct {
	ID         int16  `gorm:"column:id"`
	Name       string `gorm:"column:username"`
	Pass       string `gorm:"column:password"`
	StoreID    string `gorm:"column:store_id"`
	AdminLevel int    `gorm:"column:admin_level"`
	WorkGroup  int    `gorm:"column:work_group"`
}

//TableName Product table
func (User) TableName() string {
	return "admin_accounts"
}
