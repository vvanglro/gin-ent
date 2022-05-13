package user

import (
	"Firstgin/global"
	"Firstgin/models/ent"
	"Firstgin/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

type login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func authenticator(c *gin.Context) (interface{}, error) {
	var loginVal login
	if err := c.ShouldBind(&loginVal); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userName := loginVal.Username
	password := loginVal.Password
	u, err := Service(c).FindByUsername(userName)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}
	if err := utils.DecryptPassword(u.Password, password); err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return u, nil

}
func authorizator(data interface{}, ctx *gin.Context) bool {
	v, ok := data.(*ent.User)
	u, err := Service(ctx).FindByUsername(v.Name)
	if err != nil {
		return false
	}
	if ok && u != nil {
		return true
	}
	return false
}

// AuthInit jwt验证new
func AuthInit() (*jwt.GinJWTMiddleware, error) {
	var identityKey = "id"
	jwtConf := global.AppSetting.Jwt
	ginJWTMiddleware := &jwt.GinJWTMiddleware{
		//可以理解成该中间件的名称，用于展示，默认值为gin jwt
		Realm: "test zone",
		//签名算法，默认值为HS256
		SigningAlgorithm: "HS256",
		//服务端密钥
		Key: []byte(jwtConf.JwtSecret),
		//token 过期时间
		Timeout: time.Duration(jwtConf.TokenExpireDuration*24) * time.Hour,
		//token 更新时间
		MaxRefresh: time.Duration(jwtConf.TokenExpireDuration*24) * time.Hour,
		//身份验证的key值
		IdentityKey: identityKey,
		//函数：根据登录信息对用户进行身份验证的回调函数
		Authenticator: authenticator,
		//登录后验证传入的 token 方法，可在此处写权限验证逻辑
		Authorizator: authorizator,
		//token检索模式，用于提取token
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		//token在请求头时的名称，默认值为Bearer
		TokenHeadName: "Bearer",
		//测试或服务器在其他时区可设置该属性，默认值为time.Now
		TimeFunc: time.Now,
		//验证失败后的函数调用，可用于自定义返回的 JSON 格式之类的
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(200, gin.H{
				"code":    code,
				"message": message,
			})
		},
		//函数：解析并设置用户身份信息
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &ent.User{
				Name: claims[identityKey].(string),
			}
		},
		//添加额外业务相关的信息
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			v, ok := data.(*ent.User)
			if ok {
				return jwt.MapClaims{identityKey: v.Name}
			}
			return jwt.MapClaims{}
		},
	}
	return jwt.New(ginJWTMiddleware)
}
