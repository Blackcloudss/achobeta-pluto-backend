package global

import (
	"tgwp/util/snowflake"
	"time"
)

// 所有常量文件读取位置
const (
	DEFAULT_CONFIG_FILE_PATH           = "/config.yaml"
	REDIS_PHONE_CODE                   = "Achobeta:phone.login.code:%s:string"
	REDIS_PHONE                        = "Achobeta:phone.login.effective:%s:string"
	ATOKEN_EFFECTIVE_TIME              = time.Second * 300
	RTOKEN_EFFECTIVE_TIME              = time.Hour * 1
	AUTH_ENUMS_ATOKEN                  = "atoken"
	AUTH_ENUMS_RTOKEN                  = "rtoken"
	DEFAULT_NODE_ID                    = 1
	TOKEN_USER_ID                      = "UserId"
	ROOT_ID                            = 1                                              // 根节点 ID
	NORMAL_ADMINISTRATOR               = 22222                                          //普通管理员
	SUPERL_ADMINISTRATOR               = 33333                                          //超级管理员
	FEISHU_APP_ID                      = "cli_a7a37dc364b2900c"                         // 飞书自建应用ID
	FEISHU_APP_SECRET                  = "cqFNuZmHaIKlMeFAr8546cHlaquXw8ep"             // 飞书自建应用Secret
	FEISHU_APP_TOKEN                   = "M5l2bHYEiaYq2esmVM1cTyamn5s"                  // 飞书自建应用Token
	FEISHU_TASK_TABLE_ID               = "tblM1AuOpuhpxBSb"                             //飞书多维表格任务表ID
	FEISHU_LIST_UPDATE_TIME            = 60 * 5                                         // 飞书任务表更新频率(秒 int格式)
	FEISHU_LIST_WILL_OVERDUE_TIME      = 60 * 60 * 24                                   // 飞书任务表即将逾期时间(秒 int格式)
	REDIS_FEISHU_UPDATA_TIME           = "Achobeta:feishu.update.time:string"           // Redis中飞书记录上次更新时间
	REDIS_FEISHU_TOTAL_TASK_CNT        = "Achobeta:feishu.total.task.cnt:%s:int"        // Redis中飞书记录用户总任务数
	REDIS_FEISHU_UNFINISHED_TASK_CNT   = "Achobeta:feishu.unfinished.task.cnt:%s:int"   // Redis中飞书记录用户未完成任务数
	REDIS_FEISHU_WILL_OVERDUE_TASK_CNT = "Achobeta:feishu.will.overdue.task.cnt:%s:int" // Redis中飞书记录用户即将逾期任务数
	REDIS_FEISHU_OVERDUE_TASK_CNT      = "Achobeta:feishu.overdue.task.cnt:%s:int"      // Redis中飞书记录用户逾期任务数
)

var Node, _ = snowflake.NewNode(DEFAULT_NODE_ID)

var NORMAL_ADMIN_URLS = []string{
	"/api/team/memberlist/delete",
	"/api/team/memberlist/put",
	"/api/team/membermsg/save",
	"/api/team/structure/collection",
}

// global 包中的定义
var SUPER_ADMIN_URLS = []string{
	"/api/team/structure/change",
	"/api/team/structure/add",
}
