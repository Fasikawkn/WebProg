package handler

// type AdminRoleHandler struct{
// 	templ *template.template
// 	roleServ menu.RoleService
// }
// func NewAdminRoleHandler(T *template.Template,Rs menu.RoleService) *AdminRoleHandler{
// 	return &AdminRoleHandler{tmpl: T, roleServ:Rs}

// }

// func (arh AdminRoleHandler) AdminRolesNew(w http.ResponseWriter,r *http.Request){
// 	if r.Methd == http.MethdPost{
// 		name := r.FormValue("name")
// 	role := entity.Role{Name: name}

// 	err := arh.roleServ.StoreRole(role)

// 	if err != nil{
// 		panic(err)
// 	}
// 	}else{
// 		arc.temple.ExecuteTemplate(w,"",nil)
// 	}

// 	arh.templ
// 	http.Redirect(w,r,"admin/roles",http.StatusSeeOther)

// }
