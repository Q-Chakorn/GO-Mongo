package db

import (
	"GO-Mongo/config"
	"context"
	"fmt"
	"log"
)

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
	}
}
