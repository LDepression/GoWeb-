package controller

import (
	"GoAdvance/GoWeb/day03/bookstore/dao"
	"GoAdvance/GoWeb/day03/bookstore/model"
	"GoAdvance/GoWeb/day03/bookstore/utils"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

//AddBook2Cart 添加图书到购物车中
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.CheckLogin(r)
	if flag {

		bookID := r.FormValue("bookId")
		//先通过bookID获取图书的信息
		book := dao.GetBookById(bookID)
		//获取用户id
		userID := session.UserID
		//先判断该用户有没有购物车
		cart := dao.GetCartByUserID(userID)
		if cart != nil {
			//说明该用户有购物车，此时需要判断购物车有没有这本书
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			if cartItem == nil {
				//说明没有这本书，就要添加进数据库
				//创建一个新的购物项
				fmt.Println("购物车中没有这本书")
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}

				cart.CartItems = append(cart.CartItems, cartItem)
				//添加到数据库中去
				dao.AddCartItem(cartItem)
			} else {
				//此时只需要把在数据库中数量加1就好了
				//1.获取购物车中的所有的购物项
				cts := cart.CartItems
				for _, v := range cts {
					if v.Book.ID == cartItem.Book.ID {
						v.Count = v.Count + 1
						//更新购物项中图书的数量
						dao.UpdateBookCount(v)
					}
				}
			}
			dao.UpdateCart(cart)
		} else {
			fmt.Println("该用户没有购物车。。。。")
			//该用户没有购物车
			cartID := utils.CreateUUID()
			cart := &model.Cart{
				CartID: cartID,
				UserID: userID,
			}
			//创建购物项
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cartID,
			}
			//将购物项添加到切片中
			cart.CartItems = append(cart.CartItems, cartItem)
			dao.AddCart(cart)
		}
		w.Write([]byte("您刚刚将" + book.Title + "加入到了购物车中"))
	} else {
		w.Write([]byte("请重新登录"))
	}
}

func GetCardInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.CheckLogin(r)
	//获取用户的id
	userID := session.UserID
	//获取userid对应的购物车
	cart := dao.GetCartByUserID(userID)
	if cart != nil {
		//说明该用户有购物车
		session.Cart = cart
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	} else {
		//该用户没有购物车
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	}
}

//DeleteCart 清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartID")
	dao.DeleteCartByCartID(cartID)
	GetCardInfo(w, r)
}

//DeleteCartItem 删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("CartItem")
	iCartTiemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	_, session := dao.CheckLogin(r)
	UserID := session.UserID
	//通过UserName来获取购物车
	cart := dao.GetCartByUserID(UserID)
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if v.CartItemID == iCartTiemID {
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			cart.CartItems = cartItems
			dao.DeleteCartItemByID(cartItemID)
		}
	}
	dao.UpdateCart(cart)
	GetCardInfo(w, r)
}
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemID")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	bookCount := r.FormValue("UserCount")
	ibookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	_, session := dao.CheckLogin(r)
	userID := session.UserID
	cart := dao.GetCartByUserID(userID)
	for _, v := range cart.CartItems {
		if v.CartItemID == iCartItemID {
			v.Count = ibookCount
			dao.UpdateBookCount(v)
		}
	}
	dao.UpdateCart(cart)
	GetCardInfo(w, r)
}
