package amlog

import (
	amdb "atom_micro/am_db"
	ammodel "atom_micro/am_model"
	"time"
)

var chanapi = make(chan ammodel.LogAmInterface, 1000)

func init() {
	go logApiToDb()
}

func LogApi(path, params string, excuteTime int) {
	lg := ammodel.LogAmInterface{LogPath: path, Parameter: params, ExecuteTime: excuteTime, AddTime: time.Now()}
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
