package system

import (
	"ruoyi-go/app/admin/model/tools"
	"ruoyi-go/utils"
	"ruoyi-go/utils/R"
	"strconv"
	"time"
)

type SysConfig struct {
	ConfigId    int       `json:"configId" gorm:"column:config_id;primaryKey"` //表示主键
	ConfigName  string    `json:"configName" gorm:"config_name"`
	ConfigKey   string    `json:"configKey" gorm:"config_key"`
	ConfigValue string    `json:"configValue" gorm:"config_value"`
	ConfigType  string    `json:"configType" gorm:"config_type"`
	CreateBy    string    `json:"createBy" gorm:"create_by"`
	CreateTime  time.Time `json:"createTime" gorm:"column:create_time;type:datetime"`
	UpdateBy    string    `json:"updateBy" gorm:"update_by"`
	UpdateTime  time.Time `json:"updateTime" gorm:"column:update_time;type:datetime"`
	Remark      string    `json:"remark" gorm:"remark"`
}

// TableName 指定数据库表名称
func (SysConfig) TableName() string {
	return "sys_config"
}

func SelectConfigList(params tools.SearchTableDataParam, isPage bool) tools.TableDataInfo {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	sysRoles := params.Other.(SysConfig)
	offset := (pageNum - 1) * pageSize
	var total int64
	var rows []SysConfig

	var db = utils.MysqlDb.Model(&rows)

	var configId = sysRoles.ConfigId
	if configId != 0 {
		db.Where("config_id = ?", configId)
	}

	var configKey = sysRoles.ConfigKey
	if configKey != "" {
		db.Where("config_key = ?", configKey)
	}

	var configName = sysRoles.ConfigName
	if configName != "" {
		db.Where("config_name like concat('%', ?, '%')", configName)
	}

	var configType = sysRoles.ConfigType
	if configType != "" {
		db.Where("config_type = ?", configType)
	}

	var beginTime = params.Params.BeginTime
	var endTime = params.Params.EndTime

	if beginTime != "" {
		//Loc, _ := time.LoadLocation("Asia/Shanghai")
		//startTime1, _ := time.ParseInLocation(constants.DateFormat, beginTime, Loc)
		//endTime = endTime + " 23:59:59"
		//endTime1, _ := time.ParseInLocation(constants.TimeFormat, endTime, Loc)
		startTime1, endTime1 := utils.GetBeginAndEndTime(beginTime, endTime)
		db.Where("create_time >= ?", startTime1)
		db.Where("create_time <= ?", endTime1)
	}

	if err := db.Count(&total).Error; err != nil {
		return tools.Fail()
	}
	if isPage {
		if err := db.Limit(pageSize).Offset(offset).Find(&rows).Error; err != nil {
			return tools.Fail()
		}
	} else {
		if err := db.Find(&rows).Error; err != nil {
			return tools.Fail()
		}
	}

	if rows == nil {
		return tools.Fail()
	} else {
		return tools.Success(rows, total)
	}
}

func SelectConfig(config SysConfig) SysConfig {
	var result SysConfig
	var dbg = utils.MysqlDb.Model(&result)
	var configId = config.ConfigId
	if configId != 0 {
		dbg.Where("config_id = ?", configId)
	}
	var configKey = config.ConfigKey
	if configKey != "" {
		dbg.Where("config_key = ?", configKey)
	}
	err := dbg.First(&result).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return result
}

func GetConfigInfo(configId interface{}) SysConfig {
	/*方式一*/
	//var sql = "select config_id, config_name, config_key, config_value, config_type, create_by, create_time, update_by, update_time, remark from sys_config where config_id = ?"
	var config SysConfig
	//utils.MysqlDb.Raw(sql, configId).Find(&config)
	/*方式二*/
	err := utils.MysqlDb.Where("config_id = ?", configId).First(&config).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return config
}

func SaveConfig(config SysConfig) R.Result {
	/*参数键名是否存在*/
	var keyCount = checkConfigKeyUnique(config.ConfigKey)
	if keyCount > 0 {
		return R.ReturnFailMsg("新增参数'" + config.ConfigName + "'失败，参数键名已存在")
	}
	err := utils.MysqlDb.Model(&SysConfig{}).Create(&config).Error
	if err == nil {
		return R.ReturnSuccess("操作成功")
	} else {
		return R.ReturnFailMsg(err.Error())
	}
}

func checkConfigKeyUnique(configKey string) int64 {
	var keyCount int64
	err := utils.MysqlDb.Model(&SysConfig{}).Where("config_key = ?", configKey).Count(&keyCount).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return keyCount
}

func SelectCaptchaEnabled() bool {
	var configValue = SelectConfigByKey("sys.account.captchaEnabled")
	boolValue, err := strconv.ParseBool(configValue)
	if err != nil {
		return true
	}
	return boolValue
}

func SelectConfigByKey(configKey string) string {
	var config SysConfig
	err := utils.MysqlDb.Where("config_key = ?", configKey).First(&config).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return config.ConfigValue
}

func EditConfig(config SysConfig) R.Result {
	err := utils.MysqlDb.Updates(&config).Error
	if err == nil {
		return R.ReturnSuccess("操作成功")
	} else {
		return R.ReturnFailMsg(err.Error())
	}
}

func DelConfig(configIds string) R.Result {
	var ids = utils.Split(configIds)
	for i := 0; i < len(ids); i++ {
		id := ids[i]
		var config = GetConfigInfo(id)
		configType := config.ConfigType
		if "Y" == configType {
			panic(R.ReturnFailMsg("内置参数" + config.ConfigKey + "不能删除 "))
		}
		DelConfigById(config.ConfigId)
	}
	return R.ReturnSuccess("操作成功")
}

func DelConfigById(configId int) {
	err := utils.MysqlDb.Where("config_id = ?", configId).Delete(&SysConfig{}).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
}

/*
加载缓存
重复初始化
*/
func loadingConfigCache() {
	//var param = tools.SearchTableDataParam{}
	//SelectConfigList(param, false)
	/*重新赋值进去*/

}

func DelCacheConfig(refreshCache string) R.Result {
	/*删除所有缓存*/
	/*重复初始化*/
	loadingConfigCache()
	return R.ReturnSuccess("操作成功")
}
