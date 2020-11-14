package models

import (
	"github.com/jinzhu/gorm"
)

type AccessModel interface {
	BeforeSave() error
	Prepare()
	Validate(action string) error
	Save(db *gorm.DB) error
	FindByID(db *gorm.DB, uid uint32) error
	UpdatePromotion(db *gorm.DB, uid uint32, promotion string) error
	Update(db *gorm.DB, uid uint32) error
	Delete(db *gorm.DB, uid uint32) (int64, error)
}

type Factory interface {
	Create() AccessModel
}

type UserFactory struct{}

func (receiver UserFactory) Create() AccessModel {
	return &User{}
}

type CompanyFactory struct {
	Company Company
}

func (receiver CompanyFactory) Create() AccessModel {
	return &Company{}
}

func (u *User) Set(i User) {

}

func (u *User) Get() *User {
	return u
}
