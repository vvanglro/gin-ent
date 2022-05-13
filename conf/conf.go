package conf

import (
	"Firstgin/global"
	"fmt"
	"github.com/labstack/gommon/color"
	"github.com/spf13/viper"
	"log"
)

func ConfigInit() {
	v := viper.New()
	v.SetConfigName("conf")
	v.SetConfigType("yaml")
	v.AddConfigPath("./conf")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %s", err)
	}

	if err := v.Unmarshal(&global.AppSetting); err != nil {
		log.Fatalf("serialize config failed: %s", err)
	}
	fmt.Println("config: " + color.Green("config file init success"))
}
