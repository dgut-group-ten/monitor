package db

import (
	"fmt"
	"monitor/core/models"
)

func UpdateUserOperationDB(uoList []*models.UserOperation) (err error) {

	sqlStr := "INSERT IGNORE INTO monitor_user_operation (" +
		"`uid`, `remote_addr`, `time_local`, `http_method`, " +
		"`res_type`, `res_id`, `status`, `body_bytes_sent`, " +
		"`http_referer`, `http_user_agent`) " +
		"VALUES "
	vals := []interface{}{}

	for _, uo := range uoList {
		sqlStr += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?), "
		vals = append(
			vals, uo.Uid, uo.RemoteAddr, uo.TimeLocal, uo.HttpMethod,
			uo.ResType, uo.ResId, uo.Status, uo.BodyBytesSent,
			uo.HttpReferer, uo.HttpUserAgent,
		)
	}
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-2]
	//prepare the statement
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("语句有问题")
		return err
	}
	defer stmt.Close()

	//format all vals at once
	_, err = stmt.Exec(vals...)
	if err != nil {
		fmt.Println("插入的时候出现问题")
		return err
	}

	return nil
}
