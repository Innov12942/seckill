package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type mysql_good struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Price    int    `gorm:"not null"`
	Expire   int    `gorm:"not null"`
	Totalnum int    `gorm:"not null"`
}

var ctx = context.Background()

var rdb1 *redis.Client //note
var rdb2 *redis.Client //trash

func InitRedis() {
	rdb1 = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,
	})

	_, err := rdb1.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	rdb2 = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       2,
	})

	_, err = rdb2.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func InsertEntry(hashtoken string, uid int) {
	// error check
	if rdb1 == nil {
		panic("Redis not connected")
	}

	// Insert entry
	err := rdb1.Set(ctx, hashtoken, uid, 10*time.Minute).Err()
	if err != nil {
		fmt.Println("InsertEntry Set error")
		return
	}
	return
}

func FindEntry(hashtoken string) int { // 0 not found, 1 found
	// error check
	if rdb1 == nil {
		panic("Redis not connected")
	}

	val, err := rdb1.Get(ctx, hashtoken).Result()
	if err == redis.Nil {
		fmt.Println("Find failed Or other error occurs")
		return 0
	}

	uidint, _ := strconv.Atoi(val)
	return uidint
}

func CacheAll() {
	var goodses []sc_good
	db.Find(&goodses)
	var finalstr string
	finalstr = ""
	for i, vgoods := range goodses {
		vgj, _ := json.Marshal(vgoods)
		if i == 0 {
			finalstr += string(vgj)
		} else {
			finalstr += "&" + string(vgj)
		}
	}

	// error check
	if rdb2 == nil {
		panic("Redis not connected")
	}

	// Insert entry
	err := rdb2.Set(ctx, "All", finalstr, 0).Err()
	if err != nil {
		fmt.Println("InsertEntry Set error")
		return
	}
}

func GetAll() string {
	var goodses []sc_good
	db.Find(&goodses)
	var finalstr string
	finalstr = ""
	for i, vgoods := range goodses {
		vgj, _ := json.Marshal(vgoods)
		if i == 0 {
			finalstr += string(vgj)
		} else {
			finalstr += "&" + string(vgj)
		}
	}
	return finalstr
}
