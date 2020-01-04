package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"WebProg/delivery/http/handler"
	"WebProg/menu/repository"
	"WebProg/menu/service"
)

func main() {
	Templates
	dbconn, err := sql.Open("postgres", "postgres://postgres:user1@localhost/restaurantdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("ui/templates/*"))

	categoryRepo := repository.NewCategoryRepositoryImpl(dbconn)
	categoryServ := service.NewCategoryServiceImpl(categoryRepo)

	userRepo := repository.NewUsersRepositoryImpl(dbconn)
	userServ := service.NewUserServiceImpl(userRepo)

	adminCatgHandler := handler.NewAdminCategoryHandler(tmpl, categoryServ)
	adminUserHandler := handler.NewAdminUserHandler(tmpl, userServ)
	menuHandler := handler.NewMenuHandler(tmpl, categoryServ)

	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", menuHandler.Index)
	http.HandleFunc("/about", menuHandler.About)
	http.HandleFunc("/contact", menuHandler.Contact)
	http.HandleFunc("/menu", menuHandler.Menu)
	http.HandleFunc("/admin", menuHandler.Admin)

	http.HandleFunc("/admin/categories", adminCatgHandler.AdminCategories)
	http.HandleFunc("/admin/categories/new", adminCatgHandler.AdminCategoriesNew)
	http.HandleFunc("/admin/categories/update", adminCatgHandler.AdminCategoriesUpdate)
	http.HandleFunc("/admin/categories/delete", adminCatgHandler.AdminCategoriesDelete)

	http.HandleFunc("/admin/users", adminUserHandler.AdminUsers)
	http.HandleFunc("/admin/users/new", adminUserHandler.AdminUsersNew)
	http.HandleFunc("/admin/users/update", adminUserHandler.AdminUsersUpdate)
	http.HandleFunc("/admin/users/delete", adminUserHandler.AdminUsersDelete)
	http.ListenAndServe(":8181", nil)

}
