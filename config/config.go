package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/BurntSushi/toml"
)

type Mysql struct {
	Host      string
	Port      int
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool `toml:"parse_time"`
	Loc       string
}

type Server struct {
	IP   string
	PORT int
}

type Config struct {
	DB     Mysql  `toml:"mysql"`
	Server Server `toml:"server"`
}

var Info Config

func init() {
	fmt.Println("初始化config")
	if _, err := toml.DecodeFile("/Users/maomao/Documents/go_learn/src/my_todolist/config/config.toml", &Info); err != nil {
		panic(err)
	}

	strings.Trim(Info.DB.Host, " ")
	strings.Trim(Info.Server.IP, " ")
}

func DBConnectString() string {
	arg := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		Info.DB.Username,
		Info.DB.Password,
		Info.DB.Host,
		Info.DB.Port,
		Info.DB.Database,
		Info.DB.Charset,
		Info.DB.ParseTime,
		Info.DB.Loc,
	)

	log.Println(arg)

	return arg
}
