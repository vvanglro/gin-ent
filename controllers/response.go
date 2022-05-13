package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//go可选参数默认值
//https://www.austsxk.com/2020/12/24/Go%E5%8F%AF%E9%80%89%E5%8F%82%E6%95%B0%E7%9A%84%E4%BD%BF%E7%94%A8%E6%8A%80%E5%B7%A7/

var Response = response{}

const successCode = http.StatusOK
const failCode = http.StatusBadRequest

type response struct {
}

type respConfig struct {
	httpCode int
	data     ReturnMsg
}

type RespOptions func(config *respConfig)

// HttpCode 添加httpCode
func HttpCode(httpCode int) RespOptions {
	return func(resp *respConfig) {
		resp.httpCode = httpCode
	}
}

func defaultRespConfig(resp *respConfig, httpCode int) *respConfig {
	resp.httpCode = httpCode
	return resp
}

// ReturnMsg 定义返回的结构体
type ReturnMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 正确返回的json格式
func (ctr *response) Success(c *gin.Context, data interface{}, options ...RespOptions) {
	res := ReturnMsg{
		200, "success", data,
	}
	resp := &respConfig{
		data: res,
	}
	// 默认值的设定
	resp = defaultRespConfig(resp, successCode)
	// 遍历可选参数，然后分别调用匿名函数，将连接对象指针传入，进行修改
	for _, op := range options {
		// 遍历调用函数，进行数据修改
		op(resp)
	}
	c.JSON(resp.httpCode, resp.data)
}

// Fail 错误返回的json格式
func (ctr *response) Fail(c *gin.Context, code int, msg string, data interface{}, options ...RespOptions) {
	res := ReturnMsg{
		code, msg, data,
	}
	resp := &respConfig{
		data: res,
	}
	// 默认值的设定
	resp = defaultRespConfig(resp, failCode)
	// 遍历可选参数，然后分别调用匿名函数，将连接对象指针传入，进行修改
	for _, op := range options {
		// 遍历调用函数，进行数据修改
		op(resp)
	}
	c.JSON(resp.httpCode, resp.data)
	c.Abort()
}
