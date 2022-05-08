package respx

import "reflect"

var (
	SUCCESS_MSG = "请求成功"
)

//StateData 状态响应结构体
type StateData struct {
	Status  bool   `json:"success"` // 请求状态。
	Message string `json:"message"` // 响应消息
}

//SingleData 单数据响应结构体
type SingleData struct {
	Status  bool   `json:"success"` // 请求状态。可有可无
	Message string `json:"message"` // 响应消息
	Data    any    `json:"data"`    // 返回数据
}

//ListData 列表响应数据结构体
type ListData struct {
	Status   bool   `json:"success"`  // 【必需】请求状态
	Message  string `json:"message"`  // 【必需】响应消息
	Data     any    `json:"data"`     // 【选填】返回数据
	Total    int64  `json:"total"`    // 【选填】数据总个数
	PageSize int64  `json:"pageSize"` // 【选填】单页数量
	Current  int64  `json:"current"`  // 【选填】当前页码
	Page     int64  `json:"page"`     // 【自定】总共有多少页，根据前端的pageSize来计算
}

func New(msg string, status bool) *StateData {
	return &StateData{
		Status:  status,
		Message: msg,
	}
}

//StateResp 返回响应状态
func StateResp(alertMsg string) *StateData {
	return &StateData{
		Status:  true,
		Message: alertMsg,
	}
}

//SingleResp 返回单数据响应
func SingleResp(alertMsg string, data any) *SingleData {
	return &SingleData{
		Status:  true,
		Message: alertMsg,
		Data:    data,
	}
}

/** ListResp 返回列表数据响应，用于分页查询数据
* alertMsg string 提示信息
* data any 返回数据
* totalRecord in64 总数据量
* pageSize int64 单页数据量
* Current int64 当前页码
 */
func ListResp(alertMsg string, listData any, totalRecord, pageSize, Current int64) *ListData {
	data := listData
	if listData != nil {
		if reflect.TypeOf(listData).Kind().String() != "slice" {
			data = []any{listData}
		}
	} else {
		data = []any{}
	}
	if pageSize == 0 {
		pageSize = 1
	}
	return &ListData{
		Status:   true,
		Message:  alertMsg,
		Data:     data,
		Total:    totalRecord,
		PageSize: pageSize,
		Current:  Current,
		Page:     (totalRecord + pageSize - 1) / pageSize,
	}
}
