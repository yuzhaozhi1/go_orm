package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AddUser(){
	// 创建一条
	// result := GLOBAL_DB.Create(&Student{
	// 	Name: "张三",
	// 	Age: 18,
	// })
	//

	// 只创建select 中指定的字段 5	李四	20	2021-08-23 21:15:56.072	2021-08-23 21:15:56.072
	// result := GLOBAL_DB.Select("Name","Age").Create(&Student{Name: "李四", Age: 20, Hobby: "玩"})

	// 指定创建时忽略的字段 6	李四		2021-08-23 21:17:44.665	2021-08-23 21:17:44.665		玩
	// result := GLOBAL_DB.Omit("Age").Create(&Student{Name: "李四", Age: 20, Hobby: "玩"})

	// 一次创建多条数据
	result := GLOBAL_DB.Create(&[]Student{
		{Name: "李五", Age: 21, Hobby: "玩"},
		{Name: "李六", Age: 22, Hobby: "玩"},
		{Name: "李七", Age: 23, Hobby: "玩"},
		{Name: "李八", Age: 24, Hobby: "玩"},
		{Name: "李九", Age: 25, Hobby: "玩"},
	})


	if result.Error != nil{
		fmt.Println(result.Error)
		return
	}
	fmt.Println(result.RowsAffected, )
}

func FindObj()  {
	var result = make(map[string]interface{})
	var studentObj Student


	// 查询第一条数据, 放到 map 中 （主键升序）
	// dbRes := GLOBAL_DB.Model(&Student{}).First(&result)

	// 查询最后一条数据, 放到 map 中 （主键降序）
	// dbRes := GLOBAL_DB.Model(&Student{}).Last(&result)

	// 获取一条记录，没有指定排序字段
	dbRes := GLOBAL_DB.Model(&Student{}).Take(&result)
	fmt.Println(errors.Is(dbRes.Error, gorm.ErrRecordNotFound))
	fmt.Println(result)

	// 查询第一条数据, 放到 机构体中 （主键升序）
	// Res := GLOBAL_DB.Model(&Student{}).First(&studentObj)
	// fmt.Println(Res.Error, Res.RowsAffected, Res.Statement)
	// fmt.Println(studentObj.ID)

	// 根据主键检索
	// GORM 允许通过内联条件指定主键来检索对象，但只支持整形数值，因为 string 可能导致 SQL 注入
	res := GLOBAL_DB.Model(&Student{}).First(&studentObj, 10)
	fmt.Println(res.Error, studentObj.ID,studentObj.Name)  // <nil> 10 李八










}