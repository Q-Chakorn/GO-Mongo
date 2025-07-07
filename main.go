package main

import (
	"GO-Mongo/db"
)

func main() {
	db.Connect() //เรียกใช้ฟังก์ชันเชื่อมต่อฐานข้อมูล MongoDB

	//router := gin.Default() //สร้าง router(obj กำหนดเส้นทาง HTTP request) หลักสำหรับรับส่ง HTTP request
}
