package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},

		// 示例模块表
		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},

		// 自动化模块表
		// Code generated by github.com/flipped-aurora/gin-vue-admin/server Begin; DO NOT EDIT.

		cloud.VideoList{},
		cloud.FileDownload{},
		cloud.FileInput{},
		cloud.FileTrans{},
		cloud.VideoCategory{},
		cloud.VideoSpecial{},
		cloud.VideoSpecialFileTrans{},
		// Code generated by github.com/flipped-aurora/gin-vue-admin/server End; DO NOT EDIT.
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
