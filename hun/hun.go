package hun

import (
	"fmt"
)

// Hello ถ้าเราต้องการให้ฟังก์ชันนี้เป็น public ให้ตั้งชื่อฟังก์ชันด้วยตัวพิมพ์ใหญ่
func Hello() {
	fmt.Println("hello")
	hello()	
}

// ฟังก์ชัน hello ถ้าเราต้องการให้ฟังก์ชันนี้เป็น private ให้ตั้งชื่อฟังก์ชันด้วยตัวพิมพ์เล็ก
func hello() {
	fmt.Println("hello, World! 2 from hun.go file")
	// hun.Hello()
}