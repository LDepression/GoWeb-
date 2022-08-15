package controller

import (
	"GoAdvance/GoWeb/day03/bookstore/dao"
	"GoAdvance/GoWeb/day03/bookstore/model"
	"GoAdvance/GoWeb/day03/bookstore/utils"
	"net/http"
	"text/template"
)

//Logout 处理用户注销的函数
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的id值
		cookieValue := cookie.Value
		//删除数据库中与之对应的session
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		//将修改之后的cookie传给浏览器
		http.SetCookie(w, cookie)
	}
	//去首页
	GetPageBooksByPrice(w, r)
}

//Login 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, _ := dao.CheckLogin(r)
	if flag == true {
		GetPageBooksByPrice(w, r)
	} else {
		//获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		//调用userdao中验证用户名和密码的方法
		user, _ := dao.CheckUserNameAndPassWord(username, password)
		if user.ID > 0 {
			//用户密码正确
			//生成uuid作为session的id
			uuid := utils.CreateUUID()
			//创建一个session
			sess := &model.Session{
				SessionID: uuid,
				UserName:  user.UserName,
				UserID:    user.ID,
			}
			//将session保存到数据库中去
			dao.AddSession(sess)
			//创建cookie与session相关联
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			//将cookie发送给浏览器
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			//用户名或者是密码不正确
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "用户名或密码不正确")
		}
	}
}
func Register(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户已经存在了
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已经存在了")
	} else {
		//用户名不存在
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

//CheckUserName 通过发送Ajax请求验证用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户名
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户已经存在了
		w.Write([]byte("用户名已存在"))
	} else {
		//用户名不存在
		w.Write([]byte("<font style='color:green'>用户名可用</font>"))

	}

}
