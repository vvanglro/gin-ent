package user

import (
	"Firstgin/controllers"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (c Controller) CreateUser(ctx *gin.Context) {
	var user Date
	if err := ctx.ShouldBindJSON(&user); err != nil {
		controllers.Response.Fail(ctx, 10013, "输入错误", err.Error(), controllers.HttpCode(401))
		return
	}
	resp, err := Service(ctx).Create(&user)
	if err != nil {
		controllers.Response.Fail(ctx, 10014, "用户已存在", err.Error(), controllers.HttpCode(401))
		return
	}
	controllers.Response.Success(ctx, resp)
}

func (c Controller) DeleteUser(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		controllers.Response.Fail(ctx, 10013, "参数错误", nil)
		return
	}
	err := Service(ctx).Delete(id)
	if err != nil {
		controllers.Response.Fail(ctx, 10014, "删除错误", err.Error(), controllers.HttpCode(401))
		return
	}
	controllers.Response.Success(ctx, id)
}

func (c Controller) UpdateUser(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		controllers.Response.Fail(ctx, 10013, "参数错误", nil)
		return
	}
	var user Date
	if err := ctx.ShouldBindJSON(&user); err != nil {
		controllers.Response.Fail(ctx, 10013, "输入错误", err.Error(), controllers.HttpCode(401))
		return
	}
	resp, err := Service(ctx).Update(id, &user)
	if err != nil {
		controllers.Response.Fail(ctx, 10014, "更新错误", err.Error(), controllers.HttpCode(401))
		return
	}
	controllers.Response.Success(ctx, resp)
}

func (c Controller) RetrieveUser(ctx *gin.Context) {
	var data DateListParams
	if err := ctx.ShouldBindJSON(&data); err != nil {
		controllers.Response.Fail(ctx, 10013, "输入错误", err.Error(), controllers.HttpCode(401))
		return
	}
	resp, err := Service(ctx).Retrieve(&data)
	if err != nil {
		controllers.Response.Fail(ctx, 10014, "查询错误", err.Error(), controllers.HttpCode(401))
		return
	}
	controllers.Response.Success(ctx, resp)
}
