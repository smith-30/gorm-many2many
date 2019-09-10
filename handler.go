package main

import "github.com/gin-gonic/gin"

func CreateRole(context *gin.Context)  {
	var reqdata CreateRoleReq
	err := context.BindJSON(&reqdata)
	if err != nil {
		err.Error()
	}

	role := &Role{
		Slug:  reqdata.Slug,
		Title: reqdata.Title,
	}
	newRole := db.Create(role)
	context.JSON(200, newRole)
}

func CreateUser(context *gin.Context)  {
	var reqdata CreateUserReq
	err := context.BindJSON(&reqdata)
	if err != nil {
		err.Error()
	}

	var roles []Role
	db.Where("slug IN (?)", reqdata.Roles).Find(&roles)

	user := &User{Mobile:reqdata.Mobile, Roles:roles}
	newUser := db.Create(user)

	// Instead of line 30 & 31 for creating newUser
	// you can also use these two line of codes(34 & 35).
	//user := &User{Mobile:reqdata.Mobile}
	//newUser := db.Create(&user).Association("Roles").Append(roles)

	context.JSON(201, newUser)
}

func GetUserRole(context *gin.Context)  {
	userId := context.Param("user_id")

	var roles []Role
	var user User
	db.Where("ID=  ?", userId).First(&user)
	db.Model(&user).Association("Roles").Find(&roles)

	context.JSON(200, roles)

}

func UpdateUser(context *gin.Context)  {
	userId := context.Param("user_id")
	var reqdata UpdateUserReq
	err := context.BindJSON(&reqdata)
	if err != nil {
		err.Error()
	}

	var user User
	db.Where("ID=  ?", userId).First(&user)

	var roles []Role
	db.Where("slug IN (?)", reqdata.Roles).Find(&roles)

	user.Mobile = reqdata.Mobile
	db.Save(&user)
	db.Model(&user).Association("Roles").Replace(roles)

	context.JSON(200, user)
}