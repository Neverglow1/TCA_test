 package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

/*sql注入就是利用sql语言拼接的bug进行数据的批量获取*/
func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/map?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	// 最大打开的连接数
	db.SetMaxOpenConns(20)
	// 设置闲置的连接数
	db.SetMaxIdleConns(10)
	return
}

type PName struct {
	Name1 string
	Name2 string
	Worth int
}

func selectData(name string) {
	var p []PName
	i := 1
	sql := ""
	if i >= 0 {
		sql = fmt.Sprintf("select name1,name2,worth from schoolmp_atb where name1 = '%s' and id = %d",name,i)
	}else{
		sql = fmt.Sprintf("select name1,name2,worth from schoolmp_atb where name1 = '%s'",name)

	err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	for index, temp := range p {
		fmt.Println(index, temp)
	}
	return
}
func main() {
	_ = initDB()
	// #在select中起到的是注释的作用
	// 将所有数据查询出来
	 selectData("'xxx' or 1=1#")
	// 将大致数据统计出来(如果数据大于10个则将数据统计出来)
	selectData("'xxx' or (select count(*) from  schoolmp_atb)>100#")

	// selectData("'大门'")

	fmt.Println(123)
}
		

