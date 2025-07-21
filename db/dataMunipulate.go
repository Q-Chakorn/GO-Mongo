package db

import (
	"GO-Mongo/config"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// ฟังก์ชันสำหรับตรวจสอบและสร้าง database ถ้าไม่มี
func CheckAndCreateDatabase(ctx context.Context, client *mongo.Client, databaseName string) {
	// ดึงรายชื่อ database ทั้งหมด
	databases, err := client.ListDatabaseNames(ctx, struct{}{})
	if err != nil {
		log.Fatal("Error listing databases:", err)
	}

	// ตรวจสอบว่ามี database ที่ต้องการหรือไม่
	databaseExists := false
	for _, db := range databases {
		if db == databaseName {
			databaseExists = true
			fmt.Printf("Database '%s' already exists\n", databaseName)
			break
		}
	}

	// ถ้าไม่มี database ให้สร้างใหม่
	if !databaseExists {
		fmt.Printf("Database '%s' not found. Creating...\n", databaseName)
		// สร้าง collection dummy เพื่อให้ database ถูกสร้างขึ้น (MongoDB สร้าง database เมื่อมี collection แรก)
		collection := client.Database(databaseName).Collection("temp_collection")
		_, err := collection.InsertOne(ctx, map[string]interface{}{"temp": "data"})
		if err != nil {
			log.Fatal("Error creating database:", err)
		}
		// ลบ document ที่เพิ่งสร้าง (optional)
		_, err = collection.DeleteOne(ctx, map[string]interface{}{"temp": "data"})
		if err != nil {
			log.Fatal("Error cleaning up temp data:", err)
		}
		fmt.Printf("Database '%s' created successfully\n", databaseName)
	}
}

// ฟังก์ชันสำหรับแสดงข้อมูลใน collection ที่กำหนดใน config
func ShowDocument(ctx context.Context, cfg *config.LoginWithParam) {
	// ดึงข้อมูลทั้งหมดจาก collection ที่กำหนด
	showdocument, err := Collection.Find(ctx, struct{}{})
	if err != nil {
		log.Fatal(err)
	}
	defer showdocument.Close(ctx) // ปิด cursor เมื่อจบการใช้งาน

	fmt.Printf("Documents in collection %s:\n", cfg.MongoDB.Collection)
	for showdocument.Next(ctx) {
		var result map[string]interface{}
		if err := showdocument.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result) // แสดงข้อมูลที่ดึงมา
		var resultJSON map[string]interface{}
		if err := showdocument.Decode(&resultJSON); err != nil {
			log.Fatal(err)
		}
		jsonData, err := json.MarshalIndent(resultJSON, "", "  ") // แปลงข้อมูลเป็น JSON format
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData)) // แสดงข้อมูลในรูปแบบ JSON
	}
}
