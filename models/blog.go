package models

type Blog struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Image string `json:"image"`
	UserID string `json:"userId"`
	User User `json:"user" gorm:"foreignKey:UserID"`
}
