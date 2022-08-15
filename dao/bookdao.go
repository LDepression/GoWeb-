package dao

import (
	"GoAdvance/GoWeb/day03/bookstore/model"
	"GoAdvance/GoWeb/day03/bookstore/utils"
	"strconv"
)

func GetBooks() ([]*model.Book, error) {
	sqlStr := `select id,title,author,price,sales,stock,img_path from books`
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil

}

//AddBook 添加图书
func AddBook(b *model.Book) error {
	sqlStr := `insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)`
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//DeleteBook 删除图书
func DeleteBook(bookID string) error {
	sqlStr := `delete from books where id=?`
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}
func GetBookById(bookId string) *model.Book {
	sqlStr := `select id,title,author,price,sales,stock,img_path from books where id=?` //这里不能写*号
	row := utils.Db.QueryRow(sqlStr, bookId)
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book
}

func UpDateBook(b *model.Book) error {
	sqlStr := `update books set title=?,author=?,price=?,sales=?,stock=? where id=?`
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil

}

//GetPageBook 获取带分页的图书信息
func GetPageBook(pageNo string) (*model.Page, error) {
	ipageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取图书的总记录数
	sqlStr := `select count(*) from books`
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//设置一个变量接收总页数
	var totalpageNo int64
	if totalRecord%pageSize == 0 {
		totalpageNo = totalRecord / pageSize
	} else {
		totalpageNo = totalRecord/pageSize + 1
	}
	//获取当前页面的图书
	sqlstr2 := `select id,title,author,price,sales,stock,img_path from books limit ?,?`
	rows, _ := utils.Db.Query(sqlstr2, pageSize*(ipageNo-1), pageSize)
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books里面去
		books = append(books, book)
	}
	//创建page
	page := &model.Page{
		Book:        books,
		PageNo:      ipageNo,
		PageSize:    pageSize,
		TotalPageNo: totalpageNo,
		TotalRecord: totalRecord,
	}
	return page, nil

}

func GetPageBookByPrice(pageNo string, minPrice string, maxPrice string) (*model.Page, error) {
	ipageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取图书的总记录数
	sqlStr := `select count(*) from books where price between ? and ?`
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//设置一个变量接收总页数
	var totalpageNo int64
	if totalRecord%pageSize == 0 {
		totalpageNo = totalRecord / pageSize
	} else {
		totalpageNo = totalRecord/pageSize + 1
	}
	//获取当前页面的图书
	sqlstr2 := `select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?`
	rows, _ := utils.Db.Query(sqlstr2, minPrice, maxPrice, pageSize*(ipageNo-1), pageSize)
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books里面去
		books = append(books, book)
	}
	//创建page
	page := &model.Page{
		Book:        books,
		PageNo:      ipageNo,
		PageSize:    pageSize,
		TotalPageNo: totalpageNo,
		TotalRecord: totalRecord,
	}
	return page, nil

}
