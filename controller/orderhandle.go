package controller

import (
	"GoAdvance/GoWeb/day03/bookstore/dao"
	"GoAdvance/GoWeb/day03/bookstore/model"
	"GoAdvance/GoWeb/day03/bookstore/utils"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

//Checkout 去结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	//获取session
	_, session := dao.CheckLogin(r)
	//获取用户的ID
	userID := session.UserID
	//获取购物车
	cart := dao.GetCartByUserID(userID)
	//生成订单号
	orderID := utils.CreateUUID()
	//生成订单的时间
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	//创建order
	order := &model.Order{
		OrderID:     orderID,
		CreatTime:   timeStr,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userID),
	}
	//将订单保存到数据库
	dao.AddOrder(order)
	//保存订单项
	//获取购物项
	cartItems := cart.CartItems
	for _, v := range cartItems {
		//创建OrderItem
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderID: orderID,
		}
		//将orderItem添加到数据库中
		dao.AddOrderItem(orderItem)
		//更新当前图书的库存与销量
		book := v.Book
		book.Sales = book.Sales + int(v.Count)
		book.Stock = book.Stock - int(v.Count)
		dao.UpDateBook(book)
	}
	//清空购物车
	dao.DeleteCartByCartID(cart.CartID)
	//将订单号设置到session中去
	session.OrderID = orderID
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, session)

}
func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, _ := dao.GetOrders()
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

//GetOrderItemsInfo 查看订单详情
func GetOrderItemsInfo(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderId")
	//通过订单号获取订单项
	orderItems, _ := dao.GetOrderItemsByID(orderID)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

//GetMyOrdersByUserID 查看我的订单
func GetMyOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	_, session := dao.CheckLogin(r)
	userID := session.UserID
	iuserName := strconv.Itoa(userID)
	orders, _ := dao.GetMyOrder(iuserName)
	session.Orders = orders
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, session)
}
