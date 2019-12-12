package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"WebProg/entity"
	"WebProg/menu"
)

type AdminUserHandler struct {
	tmpl    *template.Template
	userSRV menu.UserService
}

func NewAdminUserHandler(T *template.Template, cs menu.UserService) *AdminUserHandler {
	return &AdminUserHandler{tmpl: T, userSRV: cs}
}
func (ach *AdminUserHandler) AdminUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ach.userSRV.Users()

	if err != nil {
		panic(err)
	}

	ach.tmpl.ExecuteTemplate(w, "admin.usr.layout", users)

}

func (ach *AdminUserHandler) AdminUsersNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		ctg := entity.User{}
		ctg.Uuid = r.FormValue("uuid")
		ctg.Name = r.FormValue("fullname")
		ctg.Email = r.FormValue("email")
		ctg.Phone = r.FormValue("phone")
		ctg.Password = r.FormValue("password")

		err := ach.userSRV.AddUser(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "admin.usr.new.layout", nil)

	}
}
func (ach *AdminUserHandler) AdminUsersUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("Calling the  GET method ")

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		cat, err := ach.userSRV.User(id)

		if err != nil {
			panic(err)
		}
		fmt.Println(cat.Name)

		ach.tmpl.ExecuteTemplate(w, "admin.usr.update.layout", cat)

	} else if r.Method == http.MethodPost {

		ctg := entity.User{}
		ctg.Id, _ = strconv.Atoi(r.FormValue("id"))
		ctg.Name = r.FormValue("fullname")
		ctg.Email = r.FormValue("email")
		ctg.Phone = r.FormValue("phone")
		ctg.Password = r.FormValue("password")
		ctg.Uuid = r.FormValue("uuid")

		err := ach.userSRV.UpdateUser(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	}

}

func (ach *AdminUserHandler) AdminUsersDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = ach.userSRV.DeleteUser(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
