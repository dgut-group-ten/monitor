package task

import (
	"fmt"
	"monitor/core/cache"
	"monitor/core/db"
	"monitor/core/models"
	"monitor/core/util"
	"strconv"
	"strings"
)

func UpdateUserOperation() {

	fmt.Printf("开始更新用户行为数据\n")
	redisPool := cache.RedisPool()

	// 用户行为的记录审计数据格式(放在redis的列表中)
	// IP 用户ID 资源类型 资源ID 使用设备 时间
	// redis键格式: "action_" + uid
	prefix := "action_*"
	res, _ := redisPool.Cmd("scan", "0", "MATCH", prefix, "COUNT", "10000").Array()
	keys, _ := res[1].Array()

	uoList := []*models.UserOperation{}
	count := 0
	for _, key := range keys {
		keyStr, _ := key.Str()
		logs, _ := redisPool.Cmd("LRANGE", keyStr, "0", "-1").Array()
		//num, _ := redisPool.Cmd("DEL", keyStr).Int64()

		fmt.Printf("拿出 %s 中的数据条数:%d, 并删除 %d\n", keyStr, len(logs), 1)

		for _, log := range logs {
			logStr, _ := log.Str()
			uo := util.CutLogFetchData(logStr) //将内容装到对象中
			uo.Uid, _ = strconv.ParseInt(strings.Split(keyStr, "_")[1], 10, 64)
			uoList = append(uoList, uo)

			// 每5000条插一次数据库
			if count%5000 == 0 {
				err := db.UpdateUserOperationDB(uoList)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("成功插入%d条数据\n", len(uoList))
				uoList = []*models.UserOperation{}
			}
			count++
		}
	}
	err := db.UpdateUserOperationDB(uoList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("成功插入%d条数据\n", len(uoList))
	fmt.Printf("更新用户行为数据结束\n")
}

func UpdatePVUV() {

	fmt.Printf("开始更新pvuv数据\n")
	redisPool := cache.RedisPool()

	vcList := []*models.VisitorCount{}
	count := 0

	// PV,UV数据格式(放在多个redis的有序集合中)
	// 统计类型 资源类型 时间类型 时间 资源ID 点击量
	// redis键格式: anyType + resType + timeType + timestamp
	for _, anyType := range []string{"pv", "uv"} {
		prefix := anyType + "_*"
		res, _ := redisPool.Cmd("scan", "0", "MATCH", prefix, "COUNT", "100000").Array()
		keys, _ := res[1].Array()

		fmt.Printf("拿出 %s 中的数据条数: %d\n", prefix, len(keys))

		for _, key := range keys {
			keyStr, _ := key.Str()
			items := strings.Split(keyStr, "_")
			resources, _ := redisPool.Cmd("ZRANGE", keyStr, "0", "-1").Array()

			for _, resource := range resources {
				resID, _ := resource.Str()
				score, _ := redisPool.Cmd("ZSCORE", keyStr, resource).Int64()
				//fmt.Println(items, resID+" "+score)

				//将内容装到对象中
				vc := models.VisitorCount{
					VisType:   items[0],
					ResType:   items[1],
					ResId:     resID,
					TimeType:  items[2],
					TimeLocal: items[3],
					Click:     score,
				}
				vcList = append(vcList, &vc)

				// 每5000条插一次数据库
				if count%5000 == 0 {
					err := db.UpdateVisitorCountDB(vcList)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Printf("成功插入%d条数据\n", len(vcList))
					vcList = []*models.VisitorCount{}
				}
				count++

			}
		}
	}

	err := db.UpdateVisitorCountDB(vcList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("成功插入%d条数据\n", len(vcList))
	fmt.Printf("更新pvuv数据结束\n")

}
