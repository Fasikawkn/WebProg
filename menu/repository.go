package menu

import "WebProg/entity"

// CategoryRepository specifies menu category related database operations
type CategoryRepository interface {
	Categories() ([]entity.Category, error)
	Category(id int) (entity.Category, error)
	UpdateCategory(category entity.Category) error
	DeleteCategory(id int) error
	StoreCategory(category entity.Category) error
}

type UserRepository interface {
	Users() ([]entity.User, error)
	User(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	AddUser(user entity.User) error
}

type RoleRepository interface {
	Roles() ([]entity.Role, error)
}
