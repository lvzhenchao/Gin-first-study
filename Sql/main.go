package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main()  {
	//1、链接数据库
	connstr := "root:BspKCZLRZWeHeaTR@tcp(192.168.33.10:3306)/ginsql"

	db, err := sql.Open("mysql", connstr)
	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("数据库连接成功...")
	}

	//2、创建数据库
	//_, err = db.Exec("create table person(" +
	//	"id int auto_increment primary key," +
	//	"name varchar(12) not null," +
	//	"age int default 1"+
	//");")
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//}

	//插入数据到数据表
	_ ,err =db.Exec("insert into person (name, age)" +
	"value (?,?);", "lzc", 8)
	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("数据插入成功...")
	}

	//查询数据
	rows, err := db.Query("select id,name,age from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//按行
scan:
	if rows.Next() {
		person := new(Person)
		err := rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
		goto scan
	}



}

type Person struct {
	Id int
	Name string
	Age int
}
