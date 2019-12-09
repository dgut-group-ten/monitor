package db

import (
	"fmt"
	"monitor/core/models"
)

func UpdateUserOperationDB(uo *models.UserOperation) (err error) {

	stmt, err := DBConn().Prepare(
		"INSERT IGNORE INTO monitor_user_operation (" +
			"`uid`, `remote_addr`, `time_local`, `http_method`, " +
			"`res_type`, `res_id`, `status`, `body_bytes_sent`, " +
			"`http_referer`, `http_user_agent`) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		fmt.Println("准备语句有问题:")
		return err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(
		uo.Uid, uo.RemoteAddr, uo.TimeLocal, uo.HttpMethod,
		uo.ResType, uo.ResId, uo.Status, uo.BodyBytesSent,
		uo.HttpReferer, uo.HttpUserAgent,
	)
	if err != nil {
		fmt.Println("执行语句有问题:")
		return err
	}

	_, err = ret.RowsAffected()
	if nil != err {
		return err
	}
	return nil
}
