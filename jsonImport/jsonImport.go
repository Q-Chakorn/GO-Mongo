package main

import (
	"GO-Mongo/config"
	"GO-Mongo/db"
	"context"       // นำเข้า context สำหรับจัดการ timeout/cancel ของ process
	"encoding/json" // นำเข้า json สำหรับการจัดการข้อมูล JSON
	"fmt"           // นำเข้า fmt สำหรับแสดงผลข้อความ
	"io/ioutil"     // นำเข้า ioutil สำหรับการอ่านไฟล์
	"log"           // นำเข้า log สำหรับแสดง log ข้อผิดพลาด
	"os"            // นำเข้า os สำหรับการจัดการไฟล์และระบบปฏิบัติการ

	"go.mongodb.org/mongo-driver/mongo" // นำเข้า mongo driver สำหรับเชื่อมต่อ MongoDB
	// นำเข้า time สำหรับใช้งานเกี่ยวกับเวลา
)

// ประกาศ package ชื่อ main

func main() {
	cfg, err := config.LoadConfig("env.yaml") // โหลดการตั้งค่าจากไฟล์ env.yaml
	if err != nil {
		log.Fatal("Failed to load configuration:", err) // ถ้าโหลดการตั้งค่าไม่สำเร็จ ให้แสดง log และหยุดโปรแกรม
	}
	db.Connect(cfg) // เชื่อมต่อกับ MongoDB ด้วยการตั้งค่าที่โหลดมา
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err) // ถ้าเชื่อมต่อ MongoDB ไม่สำเร็จ ให้แสดง log และหยุดโปรแกรม
	} else {
		fmt.Println("Connected to MongoDB successfully") // แสดงข้อความเมื่อเชื่อมต่อ MongoDB สำเร็จ
	}
	ctx := context.Background()                                                          // สร้าง context สำหรับการเรียกใช้งาน
	ImportJSONToMongo(ctx, db.Collection, "jsonImport/kanto/pokemon_kanto_dataset.json") // นำเข้าข้อมูล JSON ไปยัง MongoDB                                                           // แสดงข้อมูลใน collection ที่กำหนดใน config
}

func ImportJSONToMongo(ctx context.Context, collection *mongo.Collection, jsonFilePath string) {
	file, err := os.Open(jsonFilePath) // เปิดไฟล์ JSON ที่ต้องการนำเข้า
	if err != nil {
		log.Fatal("Error opening JSON file:", err) // ถ้าเกิด error ในการเปิดไฟล์ ให้แสดง log และหยุดโปรแกรม
	}
	defer file.Close() // ปิดไฟล์เมื่อเสร็จสิ้น

	byteValue, err := ioutil.ReadAll(file) // อ่านข้อมูลจากไฟล์ JSON
	if err != nil {
		log.Fatal("Error reading JSON file:", err) // ถ้าเกิด error ในการอ่านไฟล์ ให้แสดง log และหยุดโปรแกรม
	}

	var documents []interface{}                 // ประกาศตัวแปรสำหรับเก็บข้อมูล JSON ที่อ่านมา
	err = json.Unmarshal(byteValue, &documents) // แปลงข้อมูล JSON เป็น slice ของ interface{}
	if err != nil {
		log.Fatal("Error unmarshalling JSON data:", err) // ถ้าเกิด error ในการแปลงข้อมูล JSON ให้แสดง log และหยุดโปรแกรม
	}

	result, err := collection.InsertMany(ctx, documents) // แทรกข้อมูล JSON ลงใน collection ที่กำหนด
	if err != nil {
		log.Fatal("Error inserting documents into MongoDB:", err) // ถ้าเกิด error ในการแทรกข้อมูล ให้แสดง log และหยุดโปรแกรม
	}
	fmt.Println("Inserted documents:", result.InsertedIDs) // แสดง ID ของเอกสารที่ถูกแทรก
}
