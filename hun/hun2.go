package hun

import (
	"fmt"
	// "github.com/go/hun"

	internal_test "github.com/go/hun/internal"
)

func Hello2() {
	fmt.Println("Hello, World! 2")

	// private function จะไม่สามารถเรียกใช้จาก package อื่นได้ แต่สามารถเรียกใช้จาก package เดียวกันได้
	hello()

	internal_test.Example() // จะเรียกใช้ฟังก์ชัน Example จากไฟล์ internal/test-internal.go ได้ ถ้าอยู่ในชั้น folder เดียวกัน

	// testInternal() // เรียกใช้ฟังก์ชัน testInternal จากไฟล์ hun2.go ไม่สามารถเรียกใช้ได้เพราะเป็น private function


	internal_test.Example2() // จะเรียกใช้ฟังก์ชัน Example2 จากไฟล์ internal/test-internal.go ไม่สามารถเรียกใช้ได้เพราะเป็น private function
}

// func testInternal() {
// 	internal_test.Example()
// }

func AddTwoNum(num1, num2 int) int {
	return num1 + num2
}