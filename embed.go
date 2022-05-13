package main

import (
	"Firstgin/global"
	"embed"
)

//go:embed static templates
var f embed.FS

// InitEmbed https://juejin.cn/post/6950112050886475790
func InitEmbed() {
	global.FS = f //初始化FS
}
