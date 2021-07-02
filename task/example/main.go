package main

import (
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/gocron-sdk/task"
	"github.com/spf13/viper"
	"log"
)

var (
	conf *viper.Viper
	t    = &task.Task{
		Name:             "sdk test",
		Tag:              "birthday",
		Level:            task.Parent,
		HostID:           "1",
		Spec:             "0 * * * * *",
		Command:          "https://www.google.com",
		Protocol:         task.HTTP,
		HTTPMethod:       task.HTTPMethodGet,
		Timeout:          "0",
		Multi:            "2",
		NotifyStatus:     "1",
		NotifyType:       "2",
		RetryTimes:       "3",
		RetryInterval:    "1",
		DependencyStatus: task.DependencyStatusStrong,
	}
)

func init() {
	c := configer.NewCore(configer.YAML, "config", "./config")
	c.SetAutomaticEnv("", ".", "_")
	if _c, err := c.ReadConfig(); err == nil {
		conf = _c
	} else {
		log.Panic(err.Error())
	}
}

func main() {
	client := task.NewClient(conf)
	resp, err := client.GetList(nil, map[string]string{
		"name": "7",
	})

	if err != nil {
		panic(err)
	}

	log.Println(resp[0].ID)
}
