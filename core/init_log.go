package core

import (
	"s3video/gloabl"
	"s3video/utils"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

var level log.Level

func InitLog() {
	if ok, _ := utils.PathExists("logs"); !ok { // 判断是否有Director文件夹
		fmt.Printf("创建保存日志的文件夹，目录为：%v\n", "logs")
		_ = os.Mkdir("logs", os.ModePerm)
	}
	switch gloabl.GVA_CONFIG.LogConf.Level { //初始化配置文件的Level
	case "debug":
		level = log.DebugLevel
	case "info":
		level = log.InfoLevel
	case "warn":
		level = log.WarnLevel
	case "error":
		level = log.ErrorLevel
	case "panic":
		level = log.PanicLevel
	case "fatal":
		level = log.FatalLevel
	default:
		level = log.InfoLevel
	}

	if gloabl.GVA_CONFIG.LogConf.Format == "text" {
		log.SetFormatter(&log.TextFormatter{})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}

	current_path, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取程序当前目录失败，%s", err)
		os.Exit(2)
	}
	last_path := filepath.Join(current_path, gloabl.GVA_CONFIG.LogConf.LinkName)
	writer, err := rotatelogs.New(
		filepath.Join(current_path, "logs", "bimaserver")+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(last_path),
		rotatelogs.WithMaxAge(time.Minute * time.Duration(gloabl.GVA_CONFIG.LogConf.SaveDays * 1440)),// 只保留指定时间内的日志，其他日志都波赚删除
		rotatelogs.WithRotationTime(time.Duration(gloabl.GVA_CONFIG.LogConf.RotateTime) * time.Second),// 日志多长时间循环新日志
	)
	log.SetOutput(os.Stdout)
	writers := []io.Writer{writer, os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)

	if err == nil {
		log.SetOutput(fileAndStdoutWriter)
	} else {
		fmt.Printf("添加日志文件失败，%s", err)
	}
	log.SetLevel(level)
	log.SetReportCaller(true)
}
