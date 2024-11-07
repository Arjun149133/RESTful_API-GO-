package postgres

import (
	"example/restapi/internal/model"

	"gorm.io/gorm"
)

type AuthorRepo struct {
	DB *gorm.DB
}

func (a *AuthorRepo) Create(author *model.Author) error {
	return a.DB.Create(&model.Author{
		ID:       author.ID,
		Name:     author.Name,
		Email:    author.Email,
		Password: author.Password,
	}).Error
}

func (a *AuthorRepo) FindAuthor(email string) (*model.Author, error) {
	var author model.Author
	err := a.DB.Where("email = ?", email).First(&author).Error

	return &author, err
}

func (a *AuthorRepo) FindAll() ([]model.Author, error) {
	var authors []model.Author
	err := a.DB.Find(&authors).Error

	return authors, err
}

func (a *AuthorRepo) FindById(id string) (*model.Author, error) {
	var author model.Author
	err := a.DB.Preload("Posts").Preload("Comments").First(&author, "id = ?", id).Error

	return &author, err
}

func (a *AuthorRepo) Update(author *model.Author) error {
	var oldAuthor model.Author
	a.DB.First(&oldAuthor, "id = ?", author.ID)
	return a.DB.Model(&oldAuthor).Updates(author).Error
}

func (a *AuthorRepo) Delete(id string) error {
	return a.DB.Delete(&model.Author{}, "id = ?", id).Error
}
