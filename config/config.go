package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

/*
IM配置结构体
*/
type IMConfig struct {
	IMPort     		int      `json:"im_port"`     	//服务端长连接监听端口
	HttpPort   		int      `json:"http_port"`   	//服务端短连接监听端口(登录接口)
	PprofBind		string	 `json:"pprof_bind"`		//机器监控
	DBConfig   		DBConfig `json:"db_config"`   	//数据库配置
	RouterPrefix 	[]string `json:"router_prefix"` //api前缀
}

/*
数据库配置结构体
*/
type DBConfig struct {
	Host         string `json:"host"`           //连接地址
	Username     string `json:"username"`       //用户名
	Password     string `json:"password"`       //用户密码
	Name         string `json:"name"`           //数据库名
	MaxIdleConns int    `json:"max_idle_conns"` //连接池最大空闲连接数
	MaxOpenConns int    `json:"max_open_conns"` //连接池最大连接数
}

/*
读取配置文件
*/
func ReadConfig(path string) (*IMConfig, error) {
	config := new(IMConfig)
	err := config.Parse(path)
	if err != nil {
		return nil, err
	}
	return config, nil
}

/*
解析配置文件
*/
func (this *IMConfig) Parse(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &this)
	if err != nil {
		return err
	}
	return nil
}

/*
连接数据库
*/
func (this *DBConfig) Connect() (*sql.DB, error) {
	// 从配置文件中读取配置信息并初始化连接池(go中含有连接池处理机制)
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8", this.Username, this.Password, this.Host, this.Name)
	fmt.Println(url)
	db, err := sql.Open("mysql", url)
	db.SetMaxIdleConns(this.MaxIdleConns) // 最大空闲连接
	db.SetMaxOpenConns(this.MaxOpenConns) // 最大连接数
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("连接数据库成功")
	return db, nil
}
