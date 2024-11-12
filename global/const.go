package global

import "time"

// 所有常量文件读取位置
const (
	DEFAULT_CONFIG_FILE_PATH = "/config.yaml"
	REDIS_PHONE_CODE         = "Achobeta:phone.login.code:%s:string"
	REDIS_PHONE              = "Achobeta:phone.login.effective:%s:string"
	ATOKEN_EFFECTIVE_TIME    = time.Second * 300
	RTOKEN_EFFECTIVE_TIME    = time.Hour * 1
	AUTH_ENUMS_ATOKEN        = "atoken"
	AUTH_ENUMS_RTOKEN        = "rtoken"
)
