package ini

import (
	"fmt"
	"log"
	"path"
	"sunylive-gallery/config"
	"sunylive-gallery/global"
	"sunylive-gallery/logx"
	"sunylive-gallery/mysql"
	"sunylive-gallery/utils"
	"time"

	"github.com/spf13/viper"
)

func Initialize() {
	time.Local = time.FixedZone("CST", 8*3600) //InitTime 初始化时区
	initRootPath()
	initViper()
	initConfig()
	initLog()
	initGorm()

}

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	configPath := path.Join(utils.CurrentProjectRootPath(), "deploy/")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.WatchConfig() //程序在运行中实时读取配置文件
}

func initConfig() {
	global.Config = &config.Config{
		Passport: viper.GetString("login.passport"),
		Password: viper.GetString("login.password"),
		Addr:     viper.GetString("addr"),
	}
}

func initLog() {
	global.Log = logx.NewLogger()
}

func initGorm() {
	global.DB = mysql.NewGromDB()
}

func initRootPath() {
	global.RootPath = utils.CurrentProjectRootPath()
}

func initMode() {
	global.Mode = viper.GetString("model")
	if global.Mode == "" {
		global.Mode = "prod"
	}
}
