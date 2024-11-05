package postgres

import (
	"example/restapi/internal/model"

	"gorm.io/gorm"
)

type AuthorRepo struct {
	DB *gorm.DB
}

func (a *AuthorRepo) Create(author *model.Author) error {
	return a.DB.Create(author).Error
}

func (a *AuthorRepo) FindAll() ([]model.Author, error) {
	var authors []model.Author
	err := a.DB.Find(&authors).Error

	return authors, err
}

func (a *AuthorRepo) FindById(id uint) (model.Author, error) {
	var author model.Author
	err := a.DB.First(&author, "id = ?", id).Error

	return author, err

}

func (a *AuthorRepo) Update(author *model.Author) error {
	var oldAuthor model.Author
	a.DB.First(&oldAuthor, "id = ?", author.ID)
	return a.DB.UpdateColumns(author).Error
}

func (a *AuthorRepo) Delete(id uint) error {
	return a.DB.Delete(&model.Author{}, id).Error
}
