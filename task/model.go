package task

import (
	"encoding/json"
)

type Level string

const (
	Parent Level = "1" // 父任务
	Child  Level = "2" // 子任务(依赖任务)
)

type DependencyStatus string

const (
	DependencyStatusStrong DependencyStatus = "1" // 强依赖
	DependencyStatusWeak   DependencyStatus = "2" // 弱依赖
)

type HTTPMethod string

const (
	HTTPMethodGet  HTTPMethod = "1"
	HttpMethodPost HTTPMethod = "2"
)

type Protocol string

const (
	HTTP Protocol = "1" // HTTP协议
	RPC  Protocol = "2" // RPC方式执行命令
)

type Status string

const (
	Disabled Status = "0 " // 禁用
	Failure  Status = "0"  // 失败
	Enabled  Status = "1"  // 启用
	Running  Status = "1"  // 运行中
	Finish   Status = "2"  // 完成
	Cancel   Status = "3"  // 取消
)

type Task struct {
	//id:3
	ID string `json:"id"`
	//name:api test
	Name string `json:"name"`
	//tag:
	Tag string `json:"tag"`
	//level:1
	Level Level `json:"level"`
	//dependency_status:1
	DependencyStatus DependencyStatus `json:"dependency_status"`
	//dependency_task_id:
	DependencyTaskID string `json:"dependency_task_id"`
	//spec:* * * * * *
	Spec string `json:"spec"`
	//protocol:2
	Protocol Protocol `json:"protocol"`
	//http_method:1
	HTTPMethod HTTPMethod `json:"http_method"`
	//command:echo "test"
	Command string `json:"command"`
	//host_id:1
	HostID string `json:"host_id"`
	//timeout:0
	Timeout string `json:"timeout"`
	//multi:2
	Multi string `json:"multi"` //是否允许多实例运行
	//notify_status:1
	Status Status `json:"status"` // 状态 1:正常 0:停止
	//retry_times:0
	RetryTimes string `json:"retry_times"` // 重试次数
	//retry_interval:0
	RetryInterval string `json:"retry_interval"` // 重试间隔时间
	NotifyStatus  string `json:"notify_status"`  // 任务执行结束是否通知 0: 不通知 1: 失败通知 2: 执行结束通知 3: 任务执行结果关键字匹配通知
	//notify_type:2
	NotifyType string `json:"notify_type"` // 通知类型 1: 邮件 2: slack 3: webhook
	//notify_receiver_id:
	NotifyReceiverId string `json:"notify_receiver_id"` // 通知接受者ID, setting表主键ID，多个ID逗号分隔
	//notify_keyword:
	NotifyKeyword string `json:"notify_keyword" `
	//remark:
	Remark string `json:"remark"` // 备注

}

type ListResponse struct {
	Data struct {
		Data []*Task `json:"data"`
	} `json:"data"`
}

func byte2ListResponse(in []byte) ([]*Task, error) {
	res := new(ListResponse)
	err := json.Unmarshal(in, &res)
	return res.Data.Data, err
}

type ActionResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func byte2ActionResponse(in []byte) (*ActionResponse, error) {
	res := new(ActionResponse)
	err := json.Unmarshal(in, &res)
	return res, err
}
