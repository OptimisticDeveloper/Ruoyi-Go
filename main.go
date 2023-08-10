package main

import (
	"fmt"
	"github.com/gin-middleware/xxl-job-executor"
	"runtime"
	"ruoyi-go/config"
	"ruoyi-go/pkg/scheduler"
	"ruoyi-go/routers"
	"ruoyi-go/utils"
	"strconv"
)

func main() {
	fmt.Println("hello ruoyi go")
	// 初始化配置文件
	config.InitAppConfig("./config.yaml")
	// 数据库初始化
	utils.MysqlInit()
	// 初始化 定时
	scheduler.InitCron()
	// xxl_job
	cron := scheduler.InitXxlJobCron()

	// 初始化路由
	r := routers.Init()

	if config.XxlJob.Enabled {
		xxl_job_executor_gin.XxlJobMux(r, cron)
	}

	//打开浏览器
	if runtime.GOOS == "windows" {
		utils.OpenWin("http://127.0.0.1:" + strconv.Itoa(config.Server.Port))
	}

	if runtime.GOOS == "darwin" {
		utils.OpenMac("http://127.0.0.1:" + strconv.Itoa(config.Server.Port))
	}

	if err := r.Run(":" + strconv.Itoa(config.Server.Port)); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}

}
