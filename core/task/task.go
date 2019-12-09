package task

import "monitor/core/util"

func Print5s() {
	util.Log.Infoln("Every 5 second!!")
}

func Print5m() {
	util.Log.Infoln("Every 5 min!!")
}
