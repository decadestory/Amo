package amlog

import (
	amdb "atom_micro/am_db"
	ammodel "atom_micro/am_model"
	"time"
)

var chanapi = make(chan ammodel.LogAmInterface, 1000)
var chanbus = make(chan ammodel.LogAmBus, 1000)
var chanError = make(chan ammodel.LogAmError, 1000)

func init() {
	go logApiToDb()
	go logBusToDb()
	go logErrorToDb()
}

func LogApi(path, params string, excuteTime int) {
	lg := ammodel.LogAmInterface{
		LogPath: path, 
		Parameter: params,
		ExecuteTime: excuteTime, 
		AddTime: time.Now()
	}
	chanapi <- lg
}

func logApiToDb() {
	for {
		lg, ok := <-chanapi
		if ok {
			amdb.AddApiLog(lg)
		}
	}
}

func LogBus(srcGId, srcId, logLevel int, logType, logPath, logInfo, extInfo1, extInfo2 string) {
	lg := ammodel.LogAmBus{
		SrcGId:   srcGId,
		SrcId:    srcId,
		LogType:  logType,
		LogLevel: logLevel,
		LogPath:  logPath,
		LogInfo:  logInfo,
		ExtInfo1: extInfo1,
		ExtInfo2: extInfo2,
		AddTime:  time.Now()
	}
	chanbus <- lg
}

func logBusToDb() {
	for {
		lg, ok := <-chanbus
		if ok {
			amdb.AddBusLog(lg)
		}
	}
}

func LogError(srcGId, srcId, logLevel int, logType, logPath, logInfo, extInfo1, extInfo2 string) {
	lg := ammodel.LogAmError{
		SrcGId:   srcGId,
		SrcId:    srcId,
		LogType:  logType,
		LogLevel: logLevel,
		LogPath:  logPath,
		LogInfo:  logInfo,
		ExtInfo1: extInfo1,
		ExtInfo2: extInfo2,
		AddTime:  time.Now()
	}
	chanError <- lg
}

func logErrorToDb() {
	for {
		lg, ok := <-chanError
		if ok {
			amdb.AddErrorLog(lg)
		}
	}
}
