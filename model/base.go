package model

type BaseModel struct {
	Page     int `xorm:"-"`
	PageSize int `xorm:"-"`
}

type PageInfo struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
