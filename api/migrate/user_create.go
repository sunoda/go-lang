package migrate

import (
    orm "api/database"
    "fmt"
)

type User struct {
	ID       int64  `json:"id"`       // 列名为 `id`
	Username string `json:"username"` // 列名为 `username`
	Password string `json:"password"` // 列名为 `password`
}

func UserTable()  {
	if orm.Eloquent.HasTable("User") != true {
		fmt.Printf("沒有這張資料表")
		orm.Eloquent.AutoMigrate(&User{})
		fmt.Printf("建立完成")
	}
}