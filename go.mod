module github.com/go // go mod init <module> คำสั่งที่ใช้ในการสร้างไฟล์ go.mod ในโปรเจค github.com/go เพื่อสร้างไฟล์ go.mod

go 1.23.2

// go get <package> คือคำสั่งที่ใช้ในการดึง package ที่ต้องการใช้งาน

// go list -m all คำสั่งที่ใช้ในการดู package ทั้งหมดที่ใช้งานอยู่ในโปรเจค

require github.com/google/uuid v1.6.0 // indirect go get github.com/google/uuid@v1.6.0 สำหรับดึง package ที่ต้องการใช้งาน
