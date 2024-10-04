package utils

import (
	"fmt"

	"github.com/robfig/cron"
)

type TestJob struct {
}

func (this TestJob) Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (this Test2Job) Run() {
	fmt.Println("testJob2...")
}

func Cron() {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	err := c.AddFunc(spec, func() {
		i++
		fmt.Println("cron job", i)
	})
	if err != nil {
		fmt.Println(err)
	}
	//AddJob方法
	c.AddJob(spec, TestJob{})
	c.AddJob(spec, Test2Job{})

	//启动计划任务
	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select {}
}
