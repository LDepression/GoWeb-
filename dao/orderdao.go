package dao

import (
	"GoAdvance/GoWeb/day03/bookstore/model"
	"GoAdvance/GoWeb/day03/bookstore/utils"
)

//AddOrder 向数据库中插入订单
func AddOrder(order *model.Order) error {
	//写sql语句
	sqlStr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, order.OrderID, order.CreatTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}

func GetOrders() ([]*model.Order, error) {
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreatTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}
func GetMyOrder(userID string) ([]*model.Order, error) {
	sqlStr := "select id,create_time,total_count,total_amount,state from orders where user_id=?"
	rows, err := utils.Db.Query(sqlStr, userID)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreatTime, &order.TotalCount, &order.TotalAmount, &order.State)
		orders = append(orders, order)
	}
	return orders, nil
}
