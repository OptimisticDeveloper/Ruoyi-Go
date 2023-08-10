package system

import (
	"net/http"
	"ruoyi-go/app/admin/model/system"
	"ruoyi-go/app/admin/model/tools"
	"ruoyi-go/utils/R"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ListConfig(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	configName := context.DefaultQuery("configName", "")
	configKey := context.DefaultQuery("configKey", "")
	configType := context.DefaultQuery("configType", "")

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysConfig{
			ConfigName: configName,
			ConfigKey:  configKey,
			ConfigType: configType,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	var result = system.SelectConfigList(param, true)
	context.JSON(http.StatusOK, result)
}

func ExportConfig(context *gin.Context) {

	var configParam system.SysConfig
	if err := context.ShouldBind(&configParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	var param = tools.SearchTableDataParam{
		PageNum:  0,
		PageSize: 10,
		Other: system.SysConfig{
			ConfigName: configParam.ConfigName,
			ConfigKey:  configParam.ConfigKey,
			ConfigType: configParam.ConfigType,
		},
	}
	tab := system.SelectConfigList(param, false)
	list := tab.Rows.([]system.SysConfig)
	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "configId",
		"title":  "参数主键",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "configName",
		"title":  "参数名称",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "configKey",
		"title":  "参数键名",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "configValue",
		"title":  "参数键值",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "configType",
		"title":  "系统内置",
		"width":  "10",
		"is_num": "0",
	})

	//填充数据
	data := make([]map[string]interface{}, 0)
	if len(list) > 0 {
		for _, v := range list {
			configType := v.ConfigType
			var configTypeStr = ""
			if "Y" == configType {
				configTypeStr = "是"
			}
			if "N" == configType {
				configTypeStr = "否"
			}
			data = append(data, map[string]interface{}{
				"configId":    v.ConfigId,
				"configName":  v.ConfigName,
				"configKey":   v.ConfigKey,
				"configValue": v.ConfigValue,
				"configType":  configTypeStr,
			})
		}
	}
	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}

func GetConfigInfo(context *gin.Context) {
	configId := context.Param("configId")
	result := system.GetConfigInfo(configId)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": result,
	})
}

func GetConfigKey(context *gin.Context) {
	configKey := context.Param("configKey")
	var config = system.SysConfig{ConfigKey: configKey}
	var result = system.SelectConfig(config)
	context.JSON(http.StatusOK, gin.H{
		"msg":  result.ConfigValue,
		"code": http.StatusOK,
	})
}

func SaveConfig(context *gin.Context) {
	userId, _ := context.Get("userId")
	var configParam system.SysConfig
	if err := context.ShouldBind(&configParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	configParam.CreateBy = user.UserName
	configParam.CreateTime = time.Now()
	result := system.SaveConfig(configParam)
	context.JSON(http.StatusOK, result)
}

func UploadConfig(context *gin.Context) {
	userId, _ := context.Get("userId")
	var configParam system.SysConfig
	if err := context.ShouldBind(&configParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	configParam.UpdateBy = user.UserName
	configParam.UpdateTime = time.Now()
	result := system.EditConfig(configParam)
	context.JSON(http.StatusOK, result)
}

func DetectConfig(context *gin.Context) {
	userId, _ := context.Get("userId")
	println(userId)
	var configIds = context.Param("configIds")
	result := system.DelConfig(configIds)
	context.JSON(http.StatusOK, result)
}

func DeleteCacheConfig(context *gin.Context) {
	userId, _ := context.Get("userId")
	println(userId)
	var refreshCache = context.Param("refreshCache")
	result := system.DelCacheConfig(refreshCache)
	context.JSON(http.StatusOK, result)
}
