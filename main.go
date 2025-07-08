package main

import (
	"GO-Mongo/config"
	"GO-Mongo/db"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("env.yaml")
	if err != nil {
		log.Fatal(err)
	}
	db.Connect(cfg) //เรียกใช้ฟังก์ชันเชื่อมต่อฐานข้อมูล MongoDB

	//router := gin.Default() //สร้าง router(obj กำหนดเส้นทาง HTTP request) หลักสำหรับรับส่ง HTTP request
}
