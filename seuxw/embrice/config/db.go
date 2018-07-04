package config

import (
	"seuxw/embrice/entity"
	"gopkg.in/ini.v1"
)

var configPath = "/data/config/seuxw.cfg"

func ReadDBConfig() (entity.Database, error){

	var config entity.Database
	
	conf, err := ini.Load(configPath)   //加载配置文件
	if err != nil {
	  return config, err
	}
	
	conf.BlockMode = false
	err = conf.MapTo(&config)   //解析成结构体
	if err != nil {
	  return config, err
	}
	return config, nil
}