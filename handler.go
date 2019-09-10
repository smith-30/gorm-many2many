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