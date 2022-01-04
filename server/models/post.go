package models

type Comment struct {
	Id     uint
	PostId string `json:"post_id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

type Post struct {
	Id       uint
	Text     string
	Author   string
	Comments []Comment `gorm:"default=[]"`
	Likes    []uint    `gorm:"default=[]"`
}
