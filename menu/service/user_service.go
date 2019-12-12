package service

import (
	"WebProg/entity"
	"WebProg/menu"
)

// CategoryServiceImpl implements menu.CategoryService interface
type UserServiceImpl struct {
	userRepo menu.UserRepository
}

func (cs *UserServiceImpl) User(id int) (entity.User, error) {
	c, err := cs.userRepo.User(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// NewCategoryServiceImpl will create new CategoryService object
func NewUserServiceImpl(CatRepo menu.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: CatRepo}
}

// Categories returns list of categories
func (cs *UserServiceImpl) Users() ([]entity.User, error) {

	categories, err := cs.userRepo.Users()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// StoreCategory persists new category information
func (cs *UserServiceImpl) AddUser(category entity.User) error {

	err := cs.userRepo.AddUser(category)

	if err != nil {
		return err
	}

	return nil
}

// Category returns a category object with a given id
func (cs *UserServiceImpl) user(id int) (entity.User, error) {

	c, err := cs.userRepo.User(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateCategory updates a cateogory with new data
func (cs *UserServiceImpl) UpdateUser(category entity.User) error {

	err := cs.userRepo.UpdateUser(category)

	if err != nil {
		return err
	}

	return nil
}

// DeleteCategory delete a category by its id
func (cs *UserServiceImpl) DeleteUser(id int) error {

	err := cs.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
