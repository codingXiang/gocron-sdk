package util

import (
	"fmt"
	"github.com/codingXiang/gocron-sdk/model"
)

const (
	pageCondition     = "page=%d&"
	pageSizeCondition = "page_size=%d&"
)

func HandleEndpoint(endpoint string, info *model.PageInfo, conditions map[string]string) string {
	endpoint += "?"
	if info != nil {
		if info.Page != 0 {
			endpoint += fmt.Sprintf(pageCondition, info.Page)
		} else {
			endpoint += fmt.Sprintf(pageCondition, 1)
		}

		if info.Size != 0 {
			endpoint += fmt.Sprintf(pageSizeCondition, info.Size)
		} else {
			endpoint += fmt.Sprintf(pageSizeCondition, 10)
		}
	}
	if conditions != nil {
		for key, val := range conditions {
			endpoint += fmt.Sprintf("%s=%s&", key, val)
		}
	}
	return endpoint[:len(endpoint)-1]
}
