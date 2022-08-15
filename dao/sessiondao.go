package dao

import (
	"GoAdvance/GoWeb/day03/bookstore/model"
	"GoAdvance/GoWeb/day03/bookstore/utils"
	"net/http"
)

//AddSession 向数据库中添加session
func AddSession(sess *model.Session) error {
	//写sql语句
	sqlStr := "insert into sessions values(?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession 删除数据库中的session
func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id=?"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

//GetSessionByID 根据session的id值从数据库中查询session
func GetSessionByID(sessID string) (*model.Session, error) {
	//写sql语句
	sqlStr := "select session_id,username,user_id from sessions where session_id=?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(sessID)
	//创建session
	sess := &model.Session{}
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, nil
}
func CheckLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//说明此时cookie存在
		//获取cookie中的value值
		cookieValue := cookie.Value
		session, _ := GetSessionByID(cookieValue)
		if session.UserID > 0 {
			return true, session
		}
	}
	return false, nil
}
