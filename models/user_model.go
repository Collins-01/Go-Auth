package models

type UserModel struct {
	ID       string `gorm"column:id" json:"id"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"password" json:"password"`
	FullName string `gorm:"column:fullname" json:"fullname"`
}
