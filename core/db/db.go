package db

import (
	"fmt"
	"monitor/core/models"
)

// 更新uo
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

// 更新vc
func UpdateVisitorCountDB(vcList []*models.VisitorCount) (err error) {

	sqlStr := "INSERT IGNORE monitor_visitor_count (vis_type, res_type, res_id, time_type, time_local, click) VALUES "
	vals := []interface{}{}

	for _, vc := range vcList {
		sqlStr += "(?, ?, ?, ?, ?, ?), "
		vals = append(vals, vc.VisType, vc.ResType, vc.ResId, vc.TimeType, vc.TimeLocal, vc.Click)
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

// 计算某个用户的操作的条数
func CountUserOperationDB(uid int64) (count int64, err error) {
	stmt, err := DBConn().Prepare("SELECT COUNT(`id`) FROM `monitor_user_operation` WHERE uid=?")
	if err != nil {
		fmt.Println("语句有问题")
		return count, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(uid).Scan(&count)
	if err != nil {
		fmt.Println("拿数据的时候出现问题")
		return count, err
	}

	return count, nil
}
