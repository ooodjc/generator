package utils

import (
	"generator/module"
	"log"

	"github.com/spf13/viper"
)

//var TestDb module.Db
var Config module.AppConfig

func init() {
	configs := module.AppConfig{}

	v := viper.New()
	v.SetConfigName("application") //这里就是上面我们配置的文件名称，不需要带后缀名
	v.AddConfigPath("./resources") //文件所在的目录路径
	v.SetConfigType("yml")         //这里是文件格式类型
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置文件失败：", err)
		return
	}
	maps := v.AllSettings()
	for k, val := range maps {
		v.SetDefault(k, val)
	}
	err = v.Unmarshal(&configs) //反序列化至结构体
	if err != nil {
		log.Fatal("读取配置错误：", err)
	}
	//TestDb = configs.Db
	Config = configs
}
