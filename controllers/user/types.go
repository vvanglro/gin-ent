package user

import (
	"Firstgin/models/ent"
	"Firstgin/utils"
)

type Date struct {
	Username string `json:"username" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Resp struct {
	*ent.User
}

type DateListParams struct {
	utils.PageOptions
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type RespList struct {
	Data  []*ent.User `json:"data"`
	Total int         `json:"total"`
}
