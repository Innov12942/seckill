package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User info
type sc_user struct {
	ID            int    `gorm:"primaryKey"`
	Username      string `gorm:"not null"`
	Password      string `gorm:"not null"`
	Realname      string `gorm:"default:'NoName'"`
	Group         int    `gorm:"not null"`
	Status        int    `gorm:"default:1"`
	Lastlogintime int    `gorm:"default:0"`
	Lastloginip   string `gorm:"default:''"`
}

type sc_admin_group struct {
	Aid   int    `gorm:"primaryKey"`
	Id    int    `gorm:"not null"`
	Rules string `gorm:"not null"`
}

type sc_menu struct {
	Id      int    `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	C       string `gorm:"not null"`
	A       string `gorm:"not null"`
	Display int    `gorm:"not null"`
}

type sc_good struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Price    int    `gorm:"not null"`
	Expire   int64  `gorm:"not null"`
	Totalnum int    `gorm:"not null"`
	Remain   int    `gorm:"not null"`
}

type sc_good_user struct {
	Gid int `gorm:"primaryKey"`
	Uid int `gorm:"not null"`
}

var db *gorm.DB

// ConnDB : initialize connection to mysql
func ConnDB() {
	databs := "phpdemo:4wj7fGAKJR2ddXDx@(localhost)/phpdemo?charset=utf8&parseTime=true&loc=Local"
	dbg, err := gorm.Open(mysql.Open(databs), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!")
	} else {
		fmt.Printf("Connection to mysql established\n")
	}
	db = dbg
	db.AutoMigrate(&sc_user{})
	db.AutoMigrate(&sc_admin_group{})
	db.AutoMigrate(&sc_menu{})
	db.AutoMigrate(&sc_good{})
	db.AutoMigrate(&sc_good_user{})

}

func insertDB() {
	//insert building info
	users := [...]sc_user{
		sc_user{3, "admin3", "3f4d2b6862a102567ceb5fbda25cb314", "zhangsan3", 0, 1, 0, ""},
		sc_user{2, "admin2", "626674251740ac921d0c5b76083c0a25", "zhangsan2", 0, 1, 0, ""},
		sc_user{4, "user1", "upw1", "userzhangsan", 0, 1, 0, ""},
	}
	for _, bd := range users {
		if db.Find(&bd).RowsAffected == 0 {
			db.Create(&bd)
		}
	}

	admingroups := [...]sc_admin_group{
		sc_admin_group{1, 0, "1"},
		sc_admin_group{2, 0, "2"},
		sc_admin_group{3, 0, "3"},
		sc_admin_group{4, 0, "4"},
		sc_admin_group{5, 0, "5"},
	}

	for _, et := range admingroups {
		if db.Find(&et).RowsAffected == 0 {
			db.Create(&et)
		}
	}

	menuss := [...]sc_menu{
		sc_menu{1, "添加商品", "Goods", "add_goods", 1},
		sc_menu{2, "删除商品", "Goods", "remove_goods", 1},
		sc_menu{3, "修改密码", "admin", "change_pw", 0},
		sc_menu{4, "添加商品do", "goods", "do_add", 0},
		sc_menu{5, "删除商品do", "goods", "do_remove", 0},
	}

	for _, met := range menuss {
		if db.Find(&met).RowsAffected == 0 {
			db.Create(&met)
		}
	}
}

func InitMysql() {
	ConnDB()
	insertDB()
}
