package main

import "fmt"

func AddUser(){
	result := GLOBAL_DB.Create(&Student{
		Name: "张三",
		Age: 18,
	})
	if result.Error != nil{
		fmt.Println(result.Error)
		return
	}
	fmt.Println(result.RowsAffected, )
}
