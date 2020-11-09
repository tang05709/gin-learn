package models

type Category struct {
	ID         int64
	Name       string `form:"name"`
	Sort       int `form:"sort"`
	CreatedAt  int64
	UpdatedAt  int64
}
