package main

import (
	"GoAdvance/GoWeb/day03/bookstore/controller"
	"net/http"
)

//indexhandle 去首页

func main() {
	//设置处理静态资源，就是css和js类型的等静态文件，说人话：就是图片之类的
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	//去登录
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	//去注册
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//获取所有图书
	//http.HandleFunc("/getBooks", controller.GetBooks)
	http.HandleFunc("/GetPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	//http.HandleFunc("/AddBook", controller.AddBook)
	http.HandleFunc("/DelBook", controller.DelBook)
	http.HandleFunc("/ToUpdateBookPage", controller.UpDateBookById)
	http.HandleFunc("/UpDateOrAddBook", controller.UpDateOrAddBook)
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	http.HandleFunc("/getCardInfo", controller.GetCardInfo)
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getOrders", controller.GetOrders)
	http.HandleFunc("/getOrderItemsInfo", controller.GetOrderItemsInfo)
	http.HandleFunc("/getMyOrder", controller.GetMyOrdersByUserID)
	http.ListenAndServe(":8080", nil)

}
