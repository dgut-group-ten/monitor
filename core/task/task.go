package task

import (
	"fmt"
	"monitor/core/cache"
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

	for _, key := range keys {
		keyStr, _ := key.Str()
		logs, _ := redisPool.Cmd("LRANGE", keyStr, "0", "-1").Array()

		fmt.Printf("拿出 %s 中的数据条数:%d\n", keyStr, len(logs))

		for _, log := range logs {
			logStr, _ := log.Str()
			fmt.Println(logStr)
		}
	}

	fmt.Printf("更新用户行为数据结束\n")
}

func UpdatePVUV() {

	fmt.Printf("开始更新pvuv数据\n")
	redisPool := cache.RedisPool()

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
			_ = strings.Split(keyStr, "_")
			resources, _ := redisPool.Cmd("ZRANGE", keyStr, "0", "-1").Array()
			for _, resource := range resources {
				_, _ = resource.Str()
				_, _ = redisPool.Cmd("ZSCORE", keyStr, resource).Str()
				//fmt.Println(items, resID+" "+score)
			}
		}
	}

	fmt.Printf("更新pvuv数据结束\n")

}
