package app

import (
	"github.com/jinzhu/gorm"

	common_utils "github.com/beautiful-code/sal/common/utils"
)

type (
	AppData struct {
		Config common_utils.AppConfig
		DB     *gorm.DB
	}
)

var Data AppData

func InitData() {
	common_utils.LoadAppConfig("config.json", &Data.Config)
	Data.DB = common_utils.NewDBConnection(&Data.Config)
}
