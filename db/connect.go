package db // ประกาศ package ชื่อ db

import (
	"context" // นำเข้า context สำหรับจัดการ timeout/cancel ของ process
	"fmt"     // นำเข้า fmt สำหรับแสดงผลข้อความ
	"log"     // นำเข้า log สำหรับแสดง log ข้อผิดพลาด
	"time"    // นำเข้า time สำหรับใช้งานเกี่ยวกับเวลา

	"go.mongodb.org/mongo-driver/mongo"         // นำเข้า mongo driver สำหรับเชื่อมต่อ MongoDB
	"go.mongodb.org/mongo-driver/mongo/options" // นำเข้า options สำหรับตั้งค่าการเชื่อมต่อ MongoDB
)

var Collection *mongo.Collection // ประกาศตัวแปร global สำหรับเก็บ collection ที่จะใช้งาน

func Connect() { // ฟังก์ชันสำหรับเชื่อมต่อ MongoDB
	connectOptions := options.Client().ApplyURI("mongodb://admin:secret123@27.254.134.143:32017/?authSource=admin")
	// สร้าง connectOptions โดยกำหนด URI สำหรับเชื่อมต่อ MongoDB

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// สร้าง context ที่มี timeout 10 วินาที เพื่อป้องกันการเชื่อมต่อนานเกินไป
	defer cancel() // เมื่อฟังก์ชันจบ ให้ยกเลิก context เพื่อคืน resource

	client, err := mongo.Connect(ctx, connectOptions)
	// เชื่อมต่อ MongoDB ด้วย connectOptions และ context ที่กำหนด

	if err != nil {
		log.Fatal(err) // ถ้าเกิด error ให้แสดง log และหยุดโปรแกรม
	}

	Collection = client.Database("testdb").Collection("users")
	// กำหนดค่า Collection ให้ชี้ไปที่ collection "users" ใน database "testdb"

	fmt.Printf("Connection to MongoDB\n")
}
