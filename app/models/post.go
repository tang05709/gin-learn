package models

type Post struct {
	ID         int64 
	CategoryId int64 `form:"category_id"`
	Title      string `form:"title"`
	Image      string `form:"image"`
	Content    string `form:"content"`
	Sort       int `form:"sort"`
	Status     int `form:"status"`
	CreatedAt  int64
	UpdatedAt  int64
}
