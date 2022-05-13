package global

import (
	"Firstgin/models/ent"
	"embed"
)

var FS embed.FS

var AppSetting *AppConfigure

// Db 全局数据库
var Db *ent.Client
