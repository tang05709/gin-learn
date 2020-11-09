package logics

import (
	"myweb/app/models"
	"myweb/app/core"
)
var db = core.Connection()

func ListCategory(datas []models.Category) ([]models.Category, error) {
	result := db.Order("id desc").Find(&datas)
	return datas, result.Error
}

func CreateCategory(data models.Category) (int64, error) {
	result := db.Create(&data) 
	return data.ID, result.Error
}

func FindCategory(id int64) (models.Category, error) {
	var model models.Category
	result := db.First(&model, id)
	return model, result.Error
}

func UpdateCategory(data models.Category, id int64) (int64, error) {
	var model models.Category
	row := db.First(&model, id)
	if row.Error == nil {
		result := db.Model(&model).Updates(&data)
		return model.ID, result.Error
	}
	return 0, row.Error
}

func DeleteCategory(id int64) (int64, error) {
	var model models.Category
	result := db.Delete(&model, id)
	return result.RowsAffected, result.Error
}