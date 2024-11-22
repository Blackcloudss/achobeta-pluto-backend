package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"log"
	"strconv"
	"tgwp/global"
	"tgwp/internal/response"
	"tgwp/internal/types"
	"tgwp/log/zlog"
	"time"
)

// TenantAccessTokenResp 获取 tenant_access_token 响应
type TenantAccessTokenResp struct {
	Code              int    `json:"code"`
	Expire            int    `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

// FeiShuTaskResp 获取任务表数据响应
type FeiShuTaskResp struct {
	Code int `json:"code"`
	Data struct {
		HasMore bool `json:"has_more"`
		Items   []struct {
			Fields struct {
				TaskName   string `json:"任务名称"`
				TaskDesc   string `json:"任务描述"`
				TaskStatus string `json:"当前状态"`
				TaskOf     []struct {
					TableID string   `json:"table_id"`
					TextArr []string `json:"text_arr"`
					Type    string   `json:"type"`
				} `json:"所属需求"`
				UpdatedBy struct {
					Email  string `json:"email"`
					EnName string `json:"en_name"`
					ID     string `json:"id"`
					Name   string `json:"name"`
				} `json:"更新人"`
				UpdatedTime int64 `json:"最后更新时间"`
				WorkBy      []struct {
					Email  string `json:"email"`
					EnName string `json:"en_name"`
					ID     string `json:"id"`
					Name   string `json:"name"`
				} `json:"负责人"`
				BeganTime int64 `json:"预计开始时间"`
				EndTime   int64 `json:"预计结束时间"`
			} `json:"fields"`
			ID       string `json:"id"`
			RecordID string `json:"record_id"`
		} `json:"items"`
	} `json:"data"`
}

type UserOpenIDResp struct {
	Code int `json:"code"`
	Data struct {
		UserList []struct {
			Mobile string `json:"mobile"`
			UserID string `json:"user_id"`
		} `json:"user_list"`
	} `json:"data"`
}

func UpdateFeiShuList(ctx context.Context) (err error) {
	// 先获取 tenant_access_token
	tenant_access_token, err := GetFeiShuTenantAccessToken()
	if err != nil {
		zlog.CtxErrorf(ctx, "Unable to get tenant_access_token: ", err)
		return
	}

	// 获取任务表数据
	geq := &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": "Bearer " + tenant_access_token,
		},
	}
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables/%s/records", global.FEISHU_APP_TOKEN, global.FEISHU_TASK_TABLE_ID)
	resp, err := grequests.Get(url, geq)
	if err != nil {
		zlog.CtxErrorf(ctx, "Unable to make request: ", err)
		return
	}

	// 解析任务表数据
	var recordResp FeiShuTaskResp
	if err = json.Unmarshal([]byte(resp.String()), &recordResp); err != nil {
		zlog.CtxErrorf(ctx, "Unable to parse JSON response: ", err)
		return
	}

	// 名字列表
	nameList := make(map[string]string, 10)
	// 任务列表
	// 我的总任务数
	TotalTaskCnt := make(map[string]int, 10)
	// 我的未完成任务
	UnFinishedTaskCnt := make(map[string]int, 10)
	// 我的即将逾期任务
	WillOverdueTaskCnt := make(map[string]int, 10)
	// 我的已逾期任务
	OverdueTaskCnt := make(map[string]int, 10)
	// 总任务数

	// 解析数据
	for _, item := range recordResp.Data.Items {
		nameList[item.Fields.WorkBy[0].ID] = item.Fields.WorkBy[0].Name
		TotalTaskCnt[item.Fields.WorkBy[0].ID]++
		if item.Fields.TaskStatus != "已完成" {
			UnFinishedTaskCnt[item.Fields.WorkBy[0].ID]++
			if time.Now().UnixMilli() >= item.Fields.BeganTime {
				if time.Now().UnixMilli() >= item.Fields.EndTime+(1000*60*60*24) { // 由于结束时间应该是包含最后一天的，但是飞书传来的时间戳是当天的0点，因此需要增加一天时间
					OverdueTaskCnt[item.Fields.WorkBy[0].ID]++
				} else if time.Now().UnixMilli() >= item.Fields.EndTime+(1000*60*60*24)-(global.FEISHU_LIST_WILL_OVERDUE_TIME*1000) {
					WillOverdueTaskCnt[item.Fields.WorkBy[0].ID]++
				}
			}
		}
	}

	// 保存到 redis
	for k, v := range TotalTaskCnt {
		global.Rdb.Set(ctx, fmt.Sprintf(global.REDIS_FEISHU_TOTAL_TASK_CNT, k), v, 0)
		global.Rdb.Set(ctx, fmt.Sprintf(global.REDIS_FEISHU_UNFINISHED_TASK_CNT, k), UnFinishedTaskCnt[k], 0)
		global.Rdb.Set(ctx, fmt.Sprintf(global.REDIS_FEISHU_WILL_OVERDUE_TASK_CNT, k), WillOverdueTaskCnt[k], 0)
		global.Rdb.Set(ctx, fmt.Sprintf(global.REDIS_FEISHU_OVERDUE_TASK_CNT, k), OverdueTaskCnt[k], 0)
		//fmt.Printf("User %s(%s): Total task cnt: %d, Unfinished task cnt: %d, Will overdue task cnt: %d, Overdue task cnt: %d\n", nameList[k], k, TotalTaskCnt[k], UnFinishedTaskCnt[k], WillOverdueTaskCnt[k], OverdueTaskCnt[k])
	}

	return
}

// GetFeiShuList 获取飞书任务列表
func GetFeiShuList(ctx context.Context, openID string, forceUpdate bool) (resp types.GetFeiShuListResp, err error) {
	needUpdate := true
	// 先检查需不需要更新
	if forceUpdate {
		//强制更新
	} else {
		// 判断Redis是否存在上次更新时间记录
		var val int64
		val, err = global.Rdb.Exists(ctx, global.REDIS_FEISHU_UPDATA_TIME).Result()
		if err != nil {
			zlog.CtxErrorf(ctx, "Unable to check FEISHU_UPDATA_TIME: %v", err)
			return
		}
		// 存在则需要检验更新时间
		if val != 0 {
			var val string
			val, err = global.Rdb.Get(ctx, global.REDIS_FEISHU_UPDATA_TIME).Result()
			if err != nil {
				zlog.CtxErrorf(ctx, "Unable to get FEISHU_UPDATA_TIME: %v", err)
				return
			}
			var valInt int64
			valInt, err = strconv.ParseInt(val, 10, 64)
			if err != nil {
				zlog.CtxErrorf(ctx, "Unable to parse FEISHU_UPDATA_TIME: %v", err)
				return
			}
			//fmt.Println(time.Now().Unix()-valInt, global.FEISHU_LIST_UPDATE_TIME)
			if time.Now().Unix()-valInt < global.FEISHU_LIST_UPDATE_TIME {
				// 不需要更新的情况
				needUpdate = false
			}
		}
	}

	// 更新飞书任务列表
	if needUpdate {
		err = global.Rdb.Set(ctx, global.REDIS_FEISHU_UPDATA_TIME, time.Now().Unix(), 0).Err()
		if err != nil {
			zlog.CtxErrorf(ctx, "Unable to set FEISHU_TASK_LAST_UPDATE_TIME: %v", err)
			return
		}
		err = UpdateFeiShuList(ctx)
		if err != nil {
			return
		}
	}

	// 获得数据
	var joinErrors []error
	TotalTaskCountStr, err := global.Rdb.Get(ctx, fmt.Sprintf(global.REDIS_FEISHU_TOTAL_TASK_CNT, openID)).Result()
	if err != nil {
		joinErrors = append(joinErrors, err)
	}
	UnFinishedTaskCountStr, err := global.Rdb.Get(ctx, fmt.Sprintf(global.REDIS_FEISHU_UNFINISHED_TASK_CNT, openID)).Result()
	if err != nil {
		joinErrors = append(joinErrors, err)
	}
	WillOverdueTaskCountStr, err := global.Rdb.Get(ctx, fmt.Sprintf(global.REDIS_FEISHU_WILL_OVERDUE_TASK_CNT, openID)).Result()
	if err != nil {
		joinErrors = append(joinErrors, err)
	}
	OverdueTaskCountStr, err := global.Rdb.Get(ctx, fmt.Sprintf(global.REDIS_FEISHU_OVERDUE_TASK_CNT, openID)).Result()
	if err != nil {
		joinErrors = append(joinErrors, err)
	}
	if len(joinErrors) > 0 {
		zlog.CtxErrorf(ctx, "Unable to get redis feishu data")
		err = response.ErrResp(err, response.REDIS_ERROR)
		return
	}

	// 转换成 int
	resp.TotalTaskCount, err = strconv.Atoi(TotalTaskCountStr)
	if err != nil {
		joinErrors = append(joinErrors, err)
	}
	resp.UnFinishedTaskCount, err = strconv.Atoi(UnFinishedTaskCountStr)
	if err != nil {
		joinErrors = append(joinErrors, err)
	}
	resp.WillOverdueTaskCount, err = strconv.Atoi(WillOverdueTaskCountStr)
	if err != nil {
		joinErrors = append(joinErrors, err)
	}
	resp.OverdueTaskCount, err = strconv.Atoi(OverdueTaskCountStr)
	if err != nil {
		joinErrors = append(joinErrors, err)
	}
	if len(joinErrors) > 0 {
		zlog.CtxErrorf(ctx, "Unable to convert redis feishu data to int")
		err = response.ErrResp(err, response.INTERNAL_ERROR)
		return
	}

	return
}

func GetFeiShuTenantAccessToken() (tenant_access_token string, err error) {
	postData := map[string]string{
		"app_id":     global.FEISHU_APP_ID,
		"app_secret": global.FEISHU_APP_SECRET,
	}
	geq := &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
		JSON: postData,
	}
	resp, err := grequests.Post("https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal", geq)
	if err != nil {
		zlog.Errorf("Unable to make request: %v", err)
		return
	}

	var tenantAccessTokenResp TenantAccessTokenResp
	if err = json.Unmarshal([]byte(resp.String()), &tenantAccessTokenResp); err != nil {
		log.Fatalln("Unable to parse JSON response: ", err)
		return
	}
	tenant_access_token = tenantAccessTokenResp.TenantAccessToken
	return
}

func GetFeiShuUserOpenID(phoneNumber string) (openID string, err error) {
	// 先获取 tenant_access_token
	tenant_access_token, err := GetFeiShuTenantAccessToken()
	if err != nil {
		return
	}
	// 获取用户 open_id
	postData := map[string]interface{}{
		"mobiles": []string{phoneNumber},
	}
	geq := &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": "Bearer " + tenant_access_token,
		},
		JSON: postData,
	}
	resp, err := grequests.Post("https://open.feishu.cn/open-apis/contact/v3/users/batch_get_id", geq)
	if err != nil {
		zlog.Errorf("Unable to make request: %v", err)
		return
	}

	var userOpenIDResp UserOpenIDResp
	if err = json.Unmarshal([]byte(resp.String()), &userOpenIDResp); err != nil {
		log.Fatalln("Unable to parse JSON response: ", err)
		return
	}
	if len(userOpenIDResp.Data.UserList) == 0 {
		err = fmt.Errorf("user not found")
		return
	}
	openID = userOpenIDResp.Data.UserList[0].UserID
	return
}
