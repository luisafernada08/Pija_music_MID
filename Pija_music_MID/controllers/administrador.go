package controllers

import (
	"github.com/sena_2824182/Pija_music_MID/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Administrador
type AdministradorController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create Administrador
// @Param	body		body 	models.Administrador	true		"body for user content"
// @Success 200 {int} models.Administrador.Id
// @Failure 403 body is empty
// @router / [post]
func (c *AdministradorController) Post() {
	var body_ingresa []map[string]interface{}
	var alerta models.Alert
	var temporal []byte
	var temporal_producto []byte

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err == nil {
		//fmt.Println("Body que ingresa:", body_ingresa)

		jsonData, err := json.MarshalIndent(body_ingresa, "", "")
		if err != nil {
			fmt.Println("Error la convertir a JSON", err)
		}

		json_usuario := body_ingresa[0]
		json_producto := body_ingresa[1]
		fmt.Println("Body usuario", json_usuario)
		fmt.Println("Body producto", json_producto)

		json_usuario_byte, _ := json.Marshal(json_usuario)

		response_usuario, _ := services.Metodo_post("Servicio_Post", json_usuario_byte)

		json_producto_byte, _ := json.Marshal(json_producto)

		response_producto, _ := services.Metodo_post("Servicio_post", json_producto_byte)

		temporal = response_usuario
		temporal_producto = response_producto

		//fmt.Println("Esto responde el post de producto:", string(response_producto))

		//fmt.Println("Esto responde el post del usuario:", string(response_usuario))

		fmt.Println("Body de ingreso en JSON", string(jsonData))
	}

	var temporal2 map[string]interface{}
	var temporal3 map[string]interface{}

	json.Unmarshal(temporal, &temporal2)
	json.Unmarshal(temporal_producto, &temporal3)
	var body_final []map[string]interface{}
	body_final = append(body_final, temporal2["data"].(map[string]interface{}))
	body_final = append(body_final, temporal3["data"].(map[string]interface{}))

	alerta.Code = "201"
	alerta.Type = "Post"
	alerta.Body = body_final
	c.Data["json"] = alerta
	c.ServeJSON()
}

// @Title GetAll
// @Description get all Administrador
// @Success 200 {object} models.Administrador
// @router / [get]
func (u *AdministradorController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Administrador
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *AdministradorController) Get() {
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
// @Description update the Administrador
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Administrador	true		"body for Administrador content"
// @Success 200 {object} models.Administrador
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *AdministradorController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.Administrador
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateAdministrador(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the Administrador
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *AdministradorController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs Administrador into the system
// @Param	Administrador		query 	string	true		"The Administrador for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 Administrador not exist
// @router /login [get]
func (u *AdministradorController) Login() {
	Administrador:= u.GetString("username")
	password := u.GetString("password")
	if models.Login(Administrador, password) {
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
func (u *AdministradorController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
