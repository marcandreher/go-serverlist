package constants

import (
	"github.com/Gigamons/common/consts"
)

type Config struct {
	Server struct {
		Host  string
		Port  int16
		Debug bool
	}
	MySQL consts.MySQLConf
	Redis consts.REDISConf
}
