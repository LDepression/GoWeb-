package dao

import (
	"GoAdvance/GoWeb/day03/bookstore/model"
	"fmt"
	"testing"
	"time"
)

//func TestMain(m *testing.M) {
//	fmt.Println("测试bookdao中的函数")
//	m.Run()
//}
func testUser(t *testing.T) {
	fmt.Println("测试userdao中的函数")
	t.Run("测试登录的账号和密码", testLogin)
	t.Run("测试注册是的账号", testRegister)
	//t.Run("保存用户信息", testSave)

}
func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassWord("admin2", "123456")
	fmt.Println("获取用户信息:", user)
}
func testRegister(t *testing.T) {
	user, _ := CheckUserName("admin2")
	fmt.Println("获取用户信息:", user)
}
func testSave(t *testing.T) {
	SaveUser("admin2", "123456", "admin2@qq.com")
}
func testBook(t *testing.T) {
	fmt.Println("测试bookdao中的相关函数")
	//t.Run("测试获取所有的图书", testGetBooks)
	//t.Run("测试添加一本书", testAddBook)
	//t.Run("按照id删除一本书", testDeleteBook)
	//t.Run("按照id去获取一本书的相关信息", testGetBookById)
	//t.Run("按照id去修改一本书的相关信息", testUpDateBook)
	//t.Run("测试带分页的图书", testGetPageBook)
	t.Run("测试带分页和价格范围的图书", testGetPageBookByPrice)
}
func testGetBooks(t *testing.T) {
	fmt.Println("测试bookdao中的相关函数")
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%d本书是%v\n", k+1, v)
	}
}
func testAddBook(t *testing.T) {
	fmt.Println("测试addBook")
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   99.99,
		Sales:   200,
		Stock:   200,
		ImgPath: "/views/static/img/default.jpg",
	}
	AddBook(book)
}
func testDeleteBook(t *testing.T) {
	DeleteBook("35")
}
func testGetBookById(t *testing.T) {
	b := GetBookById("1")
	fmt.Println("该图书的信息是:", b)
}
func testUpDateBook(t *testing.T) {
	book := &model.Book{
		ID:      46,
		Title:   "三国演义",
		Author:  "施耐庵",
		Price:   99.99,
		Sales:   200,
		Stock:   200,
		ImgPath: "/views/static/img/default.jpg",
	}
	UpDateBook(book)
}
func testGetPageBook(t *testing.T) {
	page, _ := GetPageBook("1")
	fmt.Println("当前页数是", page.PageNo)
	fmt.Println("总页数是", page.TotalPageNo)
	fmt.Println("总记录数是", page.TotalRecord)
	for _, v := range page.Book {
		fmt.Println("当前页的信息是", v)
	}
}
func testGetPageBookByPrice(t *testing.T) {
	page, _ := GetPageBookByPrice("1", "10", "30")
	fmt.Println("当前页数是", page.PageNo)
	fmt.Println("总页数是", page.TotalPageNo)
	fmt.Println("总记录数是", page.TotalRecord)
	for _, v := range page.Book {
		fmt.Println("当前页的信息是", v)
	}
}
func testSession(t *testing.T) {
	fmt.Println("测试添加session的函数")
	//t.Run("测试添加session", testAddSession)
	//t.Run("测试删除session", testDeleteSession)
	t.Run("测试获取session", testGetSessionByID)
}
func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "13038594219",
		UserName:  "lyc",
		UserID:    81,
	}
	AddSession(sess)
}
func testDeleteSession(t *testing.T) {
	DeleteSession("13038594219")
}
func testGetSessionByID(t *testing.T) {
	sess, _ := GetSessionByID("cd28be39-dd4f-4ebb-6c53-d6baf259e29a")
	fmt.Println("session的相关信息是", sess.UserName, sess.UserID)
}
func TestCart(t *testing.T) {
	fmt.Println("测试购物车的方法")
	//t.Run("测试addCart", testAddCart)
	//t.Run("测试以bookid获取购物项的方法", testGetCartItemByBookID)
	//t.Run("测试以cartid获取购物项的方法", testGetCartItemsByCartID)
	//t.Run("测试通过userid来获取购物车", testGetCartByUserID)
	//t.Run("测试清空购物车", testDeleteCartByCartID)
	//t.Run("测试删除某一购物项", testDeleteCartItemByID)
}
func testAddCart(t *testing.T) {
	//设置要买的第一本书
	book := &model.Book{
		ID:    1,
		Price: 27.2,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}
	//创建一个购物项切片
	var cartItems []*model.CartItem
	//创建两个购物项
	cartItem := &model.CartItem{
		Book:   book,
		Count:  10,
		CartID: "66668888",
	}
	cartItems = append(cartItems, cartItem)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "66668888",
	}
	cartItems = append(cartItems, cartItem2)
	//创建购物车
	cart := &model.Cart{
		CartID:    "66668888",
		UserID:    2,
		CartItems: cartItems,
	}
	AddCart(cart)
}

func testGetCartItemByBookID(t *testing.T) {
	cartItem, _ := GetCartItemByBookIDAndCartID("1", "66668888")
	fmt.Println("bookID为1的购物项的信息是:", cartItem)
}
func testGetCartItemsByCartID(t *testing.T) {
	cartItem, _ := GetCartItemsByCartID("66668888")
	for k, v := range cartItem {
		fmt.Printf("第%v个购物项的信息是%v\n", k+1, v)
	}
}
func testGetCartByUserID(t *testing.T) {
	cart := GetCartByUserID(2)
	fmt.Println("user_id为2的信息是", cart)
}
func testDeleteCartByCartID(t *testing.T) {
	DeleteCartByCartID("1aba85aa-e574-4abd-7a8b-29565a4ba7cc")
}

func testDeleteCartItemByID(t *testing.T) {
	DeleteCartItemByID("31")
}

func TestOrder(t *testing.T) {
	fmt.Println("测试订单的相关函数")
	//t.Run("测试添加订单的函数", testAddOrder)
	//t.Run("测试获取全部订单的函数", testGetOrders)
	//t.Run("测试通过订单号获取全部订单项的函数", testGetOrderItemsByID)
	t.Run("测试通过用户ID获取订单", testGetMyOrder)

}
func testAddOrder(t *testing.T) {
	//创建一个订单
	orderID := "13811118888"
	order := &model.Order{
		OrderID:     orderID,
		CreatTime:   time.Now().String(),
		TotalCount:  2,
		TotalAmount: 400,
		State:       0,
		UserID:      1,
	}
	orderItem1 := &model.OrderItem{
		Count:   1,
		Amount:  300,
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   300,
		ImgPath: "/static/img.default.jpg",
		OrderID: orderID,
	}
	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "西游记",
		Author:  "吴承恩",
		Price:   100,
		ImgPath: "/static/img.default.jpg",
		OrderID: orderID,
	}
	AddOrder(order)
	AddOrderItem(orderItem1)
	AddOrderItem(orderItem2)
}
func testGetOrders(t *testing.T) {
	orders, _ := GetOrders()
	for k, v := range orders {
		fmt.Printf("第%d个订单的信息是%v\n", k+1, v)
	}
}

func testGetOrderItemsByID(t *testing.T) {
	orderItems, _ := GetOrderItemsByID("eb17b3bf-a144-41e6-43c8-c5ffabd22139")
	for k, v := range orderItems {
		fmt.Printf("第%d个订单项是的相关信息是%v\n", k+1, v)
	}
}

func testGetMyOrder(t *testing.T) {
	orders, _ := GetMyOrder("1")
	for k, v := range orders {
		fmt.Printf("第%d个订单是%v\n", k+1, v)
	}

}
