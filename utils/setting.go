package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {
	// 这个地方是相对路径  相对项目的路径？
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径")
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}

func LoadServer(f *ini.File) {
	AppMode = f.Section("server").Key("AppMode").MustString("debug")
	HttpPort = f.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = f.Section("server").Key("JwtKey").MustString("q3fq234")
}

func LoadData(f *ini.File) {
	Db = f.Section("database").Key("Db").MustString("mysql")
	DbHost = f.Section("database").Key("DbHost").MustString("localhost")
	DbPort = f.Section("database").Key("DbPort").MustString("3306")
	DbUser = f.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = f.Section("database").Key("DbPassWord").MustString("admin123")
	DbName = f.Section("database").Key("DbName").MustString("ginblog")
}

func LoadQiniu(f *ini.File) {
	AccessKey = f.Section("qiniu").Key("AccessKey").String()
	SecretKey = f.Section("qiniu").Key("SecretKey").String()
	Bucket = f.Section("qiniu").Key("Bucket").String()
	QiniuServer = f.Section("qiniu").Key("QiniuServer").String()
}
