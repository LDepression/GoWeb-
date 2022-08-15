package dao

import (
	"GoAdvance/GoWeb/day03/bookstore/model"
	"GoAdvance/GoWeb/day03/bookstore/utils"
)

//CheckUserNameAndPassWord 检查用户名和密码是否正确
func CheckUserNameAndPassWord(userName string, passWord string) (*model.User, error) {
	sqlStr := `select id,username,password,email from users where username=? and password =?`
	row := utils.Db.QueryRow(sqlStr, userName, passWord)
	user := &model.User{}
	row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	return user, nil
}

//CheckUserName 判断用户名不能重复
func CheckUserName(userName string) (*model.User, error) {
	sqlStr := `select id,username,password,email from users where username=?`
	row := utils.Db.QueryRow(sqlStr, userName)
	user := &model.User{}
	row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	return user, nil
}

func SaveUser(userName string, passWord string, Email string) error {
	sqlStr := `insert into users(username,password,email) values(?,?,?)`
	_, err := utils.Db.Exec(sqlStr, userName, passWord, Email)
	if err != nil {
		return err
	}
	return nil
}
