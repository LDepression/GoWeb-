package dao

import (
	"GoAdvance/GoWeb/day03/bookstore/model"
	"GoAdvance/GoWeb/day03/bookstore/utils"
)

//AddCart 向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	//写sql语句
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	//获取购物车中所有的购物项
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		//将购物项插入到数据库中
		AddCartItem(cartItem)
	}
	return nil
}
func GetCartByUserID(userID int) *model.Cart {
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id=?"
	cart := &model.Cart{}
	row := utils.Db.QueryRow(sqlStr, userID)
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil
	}
	//获取当前购物车中所有的购物项
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	//将购物项添加到购物车中去
	cart.CartItems = cartItems
	return cart
}

//UpdateCart 更新购物车中总数量和总金额
func UpdateCart(cart *model.Cart) error {
	//写sql语句
	sqlStr := "update carts set total_count= ? , total_amount= ? where id = ? "
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}
func DeleteCartByCartID(cartID string) error {
	//先删除所有的购物项
	err := DeleteCartItemsByCartID(cartID)
	if err != nil {
		return err
	}
	//写sql语句
	sqlStr := "delete from carts where id=?"
	utils.Db.Exec(sqlStr, cartID)
	return nil
}
