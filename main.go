package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func connection1(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
}


var GLOBAL_DB *gorm.DB

// 可自己选择配置的连接
func connection2(dsn string) {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		//DefaultStringSize: 256, // string 类型字段的默认长度 utf8
		// 需要注意, 否则使用字符串当成主键会造成索引超长
		DefaultStringSize: 171, // string 类型字段的默认长度 utf8
		DSN:               dsn,
	}), &gorm.Config{
		// 为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）
		SkipDefaultTransaction: false, // 是否跳过事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		// 关闭外键约束, 使用逻辑外键,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("开启连接池失败")
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 自动创建表
	//db.AutoMigrate(&User{})

	// 手动创建表
	//M := db.Migrator()

	// 创建表
	//M.CreateTable(&User{})

	// 判断表是否存在
	//fmt.Println(M.HasTable(&User{})) // true
	//fmt.Println(M.HasTable("t_user")) // true

	// 修改表名
	//M.RenameTable("t_user", "user_back")

	// 删除表
	//M.DropTable("user_back")
	GLOBAL_DB = db
	//fmt.Println(db)
}



func main() {
	// 想要正确的处理 time.Time ，您需要带上 parseTime 参数，
	// (更多参数) 要支持完整的 UTF-8 编码，您需要将 charset=utf8 更改为 charset=utf8mb4
	dsn := "root:123abc@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
	// 基础连接
	//connection1(dsn)
	connection2(dsn)

	CreateUser()
	//DROpUser()

}
