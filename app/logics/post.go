package logics

import (
	"myweb/app/models"
)

func ListPost(datas []models.Post) ([]models.Post, int64, error) {
	result := db.Order("id desc").Find(&datas)
	return datas, result.RowsAffected, result.Error
}

func CreatePost(data models.Post) (int64, error) {
	result := db.Create(&data) 
	return data.ID, result.Error
}

func FindPost(id int64) (models.Post, error) {
	var model models.Post
	result := db.First(&model, id)
	return model, result.Error
}

func UpdatePost(data models.Post, id int64) (int64, error) {
	var model models.Post
	row := db.First(&model, id)
	if row.Error == nil {
		result := db.Model(&model).Updates(&data)
		return model.ID, result.Error
	}
	return 0, row.Error
}

func DeletePost(id int64) (int64, error) {
	var model models.Post
	result := db.Delete(&model, id)
	return result.RowsAffected, result.Error
}