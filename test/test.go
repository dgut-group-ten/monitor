package main

import (
	"fmt"
	"monitor/core/cache"
	"monitor/core/db"
	"monitor/core/models"
	"monitor/core/util"
	"strconv"
	"strings"
)

func main() {
	redisPool := cache.RedisPool()
	// 统计类型的种类: pv, uv
	// 资源类型: song|author|playlist|user|admin|other|media|static 等
	// 时间类型: day|hour|min

	// 关于点击量的自增(放在同一个redis的有序集合中):
	// 资源类型 资源ID 点击量
	// redis键格式: "click_" + resType
	resTypes := []string{"song", "playlist"}
	prefix := "click_"

	for _, resType := range resTypes {
		key := prefix + resType
		items, _ := redisPool.Cmd("ZRANGE", key, "0", "-1").Array()
		fmt.Println(key)
		for _, item := range items {
			resID, _ := item.Str()
			score, _ := redisPool.Cmd("ZSCORE", key, item).Str()
			fmt.Println("resID:" + resID + ", score:" + score)
		}
	}

	// 用户行为的记录审计数据格式(放在redis的列表中)
	// IP 用户ID 资源类型 资源ID 使用设备 时间
	// redis键格式: "action_" + uid
	prefix = "action_*"
	res, _ := redisPool.Cmd("scan", "0", "MATCH", prefix, "COUNT", "10000").Array()
	keys, _ := res[1].Array()

	for _, key := range keys {
		keyStr, _ := key.Str()
		fmt.Println(keyStr)
		logs, _ := redisPool.Cmd("LRANGE", keyStr, "0", "-1").Array()
		for _, log := range logs {
			logStr, _ := log.Str()
			uo := util.CutLogFetchData(logStr) //将内容装到对象中
			uo.Uid, _ = strconv.ParseInt(strings.Split(keyStr, "_")[1], 10, 64)
			err := db.UpdateUserOperationDB(uo)
			if err != nil {
				fmt.Println(err)
			}
			break
		}
	}

	// PV,UV数据格式(放在多个redis的有序集合中)
	// 统计类型 资源类型 时间类型 时间 资源ID 点击量
	// redis键格式: anyType + resType + timeType + timestamp
	for _, anyType := range []string{"pv", "uv"} {
		prefix = anyType + "_*"
		res, _ := redisPool.Cmd("scan", "0", "MATCH", prefix, "COUNT", "100000").Array()
		keys, _ := res[1].Array()

		for _, key := range keys {
			keyStr, _ := key.Str()
			items := strings.Split(keyStr, "_")
			resources, _ := redisPool.Cmd("ZRANGE", keyStr, "0", "-1").Array()
			for _, resource := range resources {
				resID, _ := resource.Str()
				score, _ := redisPool.Cmd("ZSCORE", keyStr, resource).Int64()

				//将内容装到对象中
				vc := models.VisitorCount{
					VisType:   items[0],
					ResType:   items[1],
					ResId:     resID,
					TimeType:  items[2],
					TimeLocal: items[3],
					Click:     score,
				}
				fmt.Println(vc)
				break
			}
		}
	}
}
