package task

import (
	"bytes"
	"fmt"
	sdk "github.com/codingXiang/gocron-sdk"
	"github.com/codingXiang/gocron-sdk/model"
	"github.com/codingXiang/gocron-sdk/util"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"mime/multipart"
)

const (
	GET     = "/api/task"
	Create  = "/api/task/store"
	Disable = "/api/task/disable/%s"
	Enable  = "/api/task/enable/%s"
	Remove  = "/api/task/remove/%s"
)

type Client struct {
	*sdk.Client
}

func NewClient(config *viper.Viper) *Client {
	c := new(Client)
	c.Client = sdk.NewClient(config)
	return c
}

//GetList 取得任務列表
func (p *Client) GetList(pageInfo *model.PageInfo, condition map[string]string) ([]*Response, error) {
	endpoint := util.HandleEndpoint(GET, pageInfo, condition)
	resp, err := p.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release
	return byte2ListResponse(resp.Body())
}

//Create 建立排程任務
func (p *Client) Create(task *Task) (*ActionResponse, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	if task.ID != "" {
		_ = writer.WriteField("id", task.ID)
	}
	//name:api test
	_ = writer.WriteField("name", task.Name)
	//tag:
	_ = writer.WriteField("tag", task.Tag)
	//level:1
	_ = writer.WriteField("level", string(task.Level))
	//dependency_status:1
	_ = writer.WriteField("dependency_status", string(task.DependencyStatus))
	//dependency_task_id:
	_ = writer.WriteField("dependency_task_id", task.DependencyTaskID)
	//spec:* * * * * *
	_ = writer.WriteField("spec", task.Spec)
	//	protocol:2
	_ = writer.WriteField("protocol", string(task.Protocol))
	//http_method:1
	_ = writer.WriteField("http_method", string(task.HTTPMethod))
	//command:echo "test"
	_ = writer.WriteField("command", task.Command)
	//host_id:1
	_ = writer.WriteField("host_id", task.HostID)
	//timeout:0
	_ = writer.WriteField("timeout", task.Timeout)
	//multi:2
	_ = writer.WriteField("multi", task.Multi)
	//notify_status:1
	_ = writer.WriteField("notify_status", task.NotifyStatus)
	//notify_type:2
	_ = writer.WriteField("notify_type", task.NotifyType)
	//notify_receiver_id:
	_ = writer.WriteField("notify_receiver_id", task.NotifyReceiverId)
	//notify_keyword:
	_ = writer.WriteField("notify_keyword", task.NotifyKeyword)
	//retry_times:0
	_ = writer.WriteField("retry_times", task.RetryTimes)
	//retry_interval:0
	_ = writer.WriteField("retry_interval", task.RetryInterval)
	//remark:
	_ = writer.WriteField("remark", task.Remark)

	contentType := writer.FormDataContentType()
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	resp, err := p.Post(Create, payload.Bytes(), contentType)

	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release
	return byte2ActionResponse(resp.Body())
}

//Disable 停止任務執行
func (p *Client) Disable(id string) (*ActionResponse, error) {
	endpoint := fmt.Sprintf(Disable, id)
	resp, err := p.Post(endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release
	return byte2ActionResponse(resp.Body())
}

//Enable 開始任務執行
func (p *Client) Enable(id string) (*ActionResponse, error) {
	endpoint := fmt.Sprintf(Enable, id)
	resp, err := p.Post(endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release
	return byte2ActionResponse(resp.Body())
}

//Remove 刪除任務
func (p *Client) Remove(id string) (*ActionResponse, error) {
	endpoint := fmt.Sprintf(Remove, id)
	resp, err := p.Post(endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release
	return byte2ActionResponse(resp.Body())
}
