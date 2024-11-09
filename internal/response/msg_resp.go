package response

type MsgCode struct {
	Code int
	Msg  string
}

//五位业务状态码

var (
	/* 成功 */
	SUCCESS = MsgCode{Code: 200, Msg: "成功"}

	/* 默认失败 */
	COMMON_FAIL = MsgCode{-43960, "失败"}

	/* 请求错误 <0 */
	TOKEN_IS_EXPIRED = MsgCode{-2, "token已过期"}

	/* 内部错误 60000 ~ 69999 */
	INTERNAL_ERROR             = MsgCode{60001, "内部错误, check log"}
	INTERNAL_FILE_UPLOAD_ERROR = MsgCode{60002, "文件上传失败"}

	/* 参数错误：10000 ~ 19999 */
	PARAM_NOT_VALID    = MsgCode{10001, "参数无效"}
	PARAM_IS_BLANK     = MsgCode{10002, "参数为空"}
	PARAM_TYPE_ERROR   = MsgCode{10003, "参数类型错误"}
	PARAM_NOT_COMPLETE = MsgCode{10004, "参数缺失"}

	/* 用户错误 20000 ~ 29999 */
	USER_NOT_LOGIN             = MsgCode{20001, "用户未登录"}
	USER_PASSWORD_DIFFERENT    = MsgCode{20002, "用户两次密码输入不一致"}
	USER_ACCOUNT_NOT_EXIST     = MsgCode{20003, "账号不存在"}
	USER_CREDENTIALS_ERROR     = MsgCode{20004, "密码错误"}
	USER_ACCOUNT_ALREADY_EXIST = MsgCode{20008, "账号已存在"}
	CAPTCHA_ERROR              = MsgCode{20500, "验证码错误"}
	INSUFFICENT_PERMISSIONS    = MsgCode{20403, "权限不足"}

	/*
	 USER_ACCOUNT_DISABLE(20005, "账号不可用"),
	 USER_ACCOUNT_LOCKED(20006, "账号被锁定"),
	 USER_ACCOUNT_NOT_EXIST(20007, "账号不存在"),
	 USER_ACCOUNT_USE_BY_OTHERS(20009, "账号下线"),
	 USER_ACCOUNT_EXPIRED(20010, "账号已过期"),
	*/
)
