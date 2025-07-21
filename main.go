package main

import (
	"GO-Mongo/config"
	"GO-Mongo/db"
	"context"
	"log"
)

var (
	ctx context.Context // ประกาศตัวแปร ctx สำหรับจัดการ context
)

func main() {
	cfg, err := config.LoadConfig("env.yaml")
	if err != nil {
		log.Fatal(err)
	}
	db.Connect(cfg) //เรียกใช้ฟังก์ชันเชื่อมต่อฐานข้อมูล MongoDB
	db.CheckAndCreateDatabase(ctx, db.Collection.Database().Client(), cfg.MongoDB.Database)
	db.ShowDocument(ctx, cfg) //เรียกใช้ฟังก์ชันแสดงข้อมูลใน collection ที่กำหนดใน config

}
