package controllers

import (
	"encoding/json"
	"fmt"

	"command-line-argumentsC:\\Users\\User\\go\\src\\github.com\\sena_2824182\\mid_prueba\\services\\carta_servicios.go"
	"github.com/Pija_music_MID/models"
	"github.com/astaxie/beego"
	"google.golang.org/genproto/googleapis/type/postaladdress"
)

// Operations about Users
type USUARIO struct {
	id                 int    `orm:"column(ID);pk"`
	nombre             string `orm:"column(Nombre);pk"`
	aprllido           string `orm:"column(Apellido);pk"`
	fecha_creacion     int    `orm:"column(Fecha_Creacion);pk"`
	feche_modificacion int    `orm:"column(Fecha_Modificacion);pk"`
	activo             int    `orm:"column(Activo);"`

	beego.USUARIO
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *USUARIOController) Post() {
	var id []map[string]interface{}
	var nombre []byte
	var apellido []byte
	var activo []byte
	fmt.Println("activo es cunaod la persona es conectada", activo)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &nombre); err == nil {
		jsonData, err := json.MarshalIndent(nombre, "", "")
		if err != nil {
			fmt.Println("error de convertir a json", err)
		}
		json_usuario := id[1]
		fmt.Println("id:", json_usuario)

		json_usuario_byte, _ := json.Marshal(json_usuario)

		response_usuario := services.Metodo_post("services_post",json_usuario_byte)
		apellido = response_usuario
		fmt.Println("este producto es el respot de post",string (response_usuario)) 
		fmt.Println("json de ingreo ",jsonData)
	}
	var usario1 map[string]interface{}
	json.Unmarshal(apellido,&usario1)

	var body_usuario []map[string]interface{}
	body_usuario = append(body_usuario, usario1["data"].(map[string]interface{}))

	alerta.Cobe = "201",
	alerta.Type = "post",
	alerta.Body = body_usuario

	c.Data["json"] = alerta

	c.ServeJSON()


	
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
