package main // package declaration
// ต้องมี package ที่ชื่อ main ในไฟล์ที่จะรันโปรแกรมเสมอ

import (
	"fmt" // import statement fmt เพื่อใช้ในการแสดงผล
	"regexp"
	"sort"    // import statement sort เพื่อใช้ในการเรียงลำดับข้อมูล
	"strings" // import statement strings เพื่อใช้ในการจัดการข้อความ

	"github.com/google/uuid" // import statement uuid เพื่อใช้ในการสร้าง uuid ใหม่

	"github.com/go/hun" // import statement hun เพื่อใช้ในการเรียกฟังก์ชัน Hello จากไฟล์ hun/hun.go

	"strconv" // import statement strconv เพื่อใช้ในการแปลงชนิดข้อมูล

	"reflect" // import statement reflect เพื่อใช้ในการแสดงชนิดข้อมูล
)

// ต้องมีฟังก์ชัน main ในไฟล์ที่จะรันโปรแกรมเสมอ
func main() {
	fmt.Println("Hello, World!")   // Println คือฟังก์ชันที่ใช้ในการแสดงผลข้อความโดยขึ้นบรรทัดใหม่
	fmt.Print("สวัสดี, โลก!")      // Print คือฟังก์ชันที่ใช้ในการแสดงผลข้อความโดยไม่ขึ้นบรรทัดใหม่
	fmt.Printf("Hello, World! \n") // Printf คือฟังก์ชันที่ใช้ในการแสดงผลข้อความโดยสามารถใส่ตัวแปรเข้าไปได้ โดยไม่ขึ้นบรรทัดใหม่
	fmt.Printf("Hello %d \n", 123) // แสดงผลขึ้นบรรทัดใหม่
	// %d คือ format specifier ที่ใช้ในการแสดงผลตัวเลข
	// %s คือ format specifier ที่ใช้ในการแสดงผลข้อความ
	// %f คือ format specifier ที่ใช้ในการแสดงผลตัวเลขทศนิยม
	// %v คือ format specifier ที่ใช้ในการแสดงผลค่าของตัวแปร
	// %T คือ format specifier ที่ใช้ในการแสดงชนิดข้อมูลของตัวแปร

	// \t คือ escape character ที่ใช้ในการเว้นวรรค
	// \n คือ escape character ที่ใช้ในการขึ้นบรรทัดใหม่

	uu := "123e4567-e89b-12d3-a456-426614174000" // การประกาศตัวแปรแบบไม่กำหนดชนิดข้อมูล และเป็น short declaration
	// := == var
	fmt.Println(uu) // แสดงผลตัวแปร uuid

	var name string = "John Doe" // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล
	fmt.Println(name)            // แสดงผลตัวแปร name
	fmt.Printf(uu + name)        // แสดงผลตัวแปร uu และ name โดยไม่ขึ้นบรรทัดใหม่
	var age int = 20             // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล
	fmt.Println(age)             // แสดงผลตัวแปร age
	var weight = 70.5            // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล
	fmt.Println(weight)          // แสดงผลตัวแปร weight
	var isTrue bool              // การประกาศตัวแปรแบบไม่กำหนดค่า
	isTrue = true                // การกำหนดค่าให้กับตัวแปร
	fmt.Println(isTrue)          // แสดงผลตัวแปร isTrue

	var one, two, three int = 1, 2, 3 // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล แบบหลายตัวพร้อมกัน
	fmt.Println(one, two, three)      // แสดงผลตัวแปร one, two, three
	fmt.Println(one + two + three)    // แสดงผลผลรวมของตัวแปร one, two, three

	fmt.Println(one + two)           // การบวก
	fmt.Println(one - two)           // การลบ
	fmt.Println(one * two)           // การคูณ
	fmt.Println(one / two)           // การหารเอาส่วน
	fmt.Println(one % two)           // การหารเอาเศษ
	fmt.Println(one == two)          // การเปรียบเทียบว่าตัวแปร one มีค่าเท่ากับตัวแปร two หรือไม่
	fmt.Println("one" + " " + "two") // การเชื่อมข้อความ (Concatenation)

	fmt.Println(" ")
	fmt.Println("ตัวดำเนินการ")
	// การเปรียบเทียบ
	// == คือ การเปรียบเทียบว่าสองค่าเท่ากันหรือไม่
	// != คือ การเปรียบเทียบว่าสองค่าไม่เท่ากันหรือไม่
	// > คือ การเปรียบเทียบว่าค่าซ้ายมากกว่าค่าขวาหรือไม่
	// < คือ การเปรียบเทียบว่าค่าซ้ายน้อยกว่าค่าขวาหรือไม่
	// >= คือ การเปรียบเทียบว่าค่าซ้ายมากกว่าหรือเท่ากับค่าขวาหรือไม่
	// <= คือ การเปรียบเทียบว่าค่าซ้ายน้อยกว่าหรือเท่ากับค่าขวาหรือไม่
	fmt.Println(one == two) // การเปรียบเทียบว่าตัวแปร one มีค่าเท่ากับตัวแปร two หรือไม่
	fmt.Println(one != two) // การเปรียบเทียบว่าตัวแปร one มีค่าไม่เท่ากับตัวแปร two หรือไม่
	fmt.Println(one > two)  // การเปรียบเทียบว่าตัวแปร one มีค่ามากกว่าตัวแปร two หรือไม่
	fmt.Println(one < two)  // การเปรียบเทียบว่าตัวแปร one มีค่าน้อยกว่าตัวแปร two หรือไม่
	fmt.Println(one >= two) // การเปรียบเทียบว่าตัวแปร one มีค่ามากกว่าหรือเท่ากับตัวแปร two หรือไม่
	fmt.Println(one <= two) // การเปรียบเทียบว่าตัวแปร one มีค่าน้อยกว่าหรือเท่ากับตัวแปร two หรือไม่

	fmt.Println(" ")
	fmt.Println("or and not")
	// || คือ การเชื่อมเงื่อนไข หรือ (OR)
	// && คือ การเชื่อมเงื่อนไข และ (AND)
	// ! คือ การเชื่อมเงื่อนไข ทางตรงกันข้าม (NOT)
	fmt.Println(true || false) // การเชื่อมเงื่อนไข หรือ (OR)
	fmt.Println(true && false) // การเชื่อมเงื่อนไข และ (AND)
	fmt.Println(!true)         // การเชื่อมเงื่อนไข ทางตรงกันข้าม (NOT)

	const pi = 3.14 // การประกาศค่าคงที่
	fmt.Println(pi) // แสดงผลค่าคงที่ pi

	// การรับค่าจากผู้ใช้
	fmt.Println(" ")
	fmt.Println("Scanf")        // แสดงข้อความให้ผู้ใช้ระบุชื่อ
	var name1 string            // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล
	fmt.Scanf("%s", &name1)     // รับค่าจากผู้ใช้ %s คือ format specifier ที่ใช้ในการรับค่าข้อความ
	fmt.Println(name1, "name1") // แสดงผลค่าที่รับมาจากผู้ใช้

	fmt.Println(" ")
	fmt.Println("split")
	text := "123,456,789"                          // การประกาศตัวแปรแบบไม่กำหนดชนิดข้อมูล
	fmt.Println(strings.Split(text, ","), "split") // การแยกข้อความด้วยเครื่องหมาย ,
	fmt.Println("Fields")
	// ถ้าต้องการแยกข้อความโดยใช้ช่องว่าง (space, tab, หรือ newline) เป็นตัวคั่น:
	text2 := "123       456  789"                // การประกาศตัวแปรแบบไม่กำหนดชนิดข้อมูล
	fmt.Println(strings.Fields(text2), "Fields") // การแยกข้อความด้วยช่องว่าง
	// ใช้ Regular Expression แยกข้อความด้วยตัวเลข
	fmt.Println("Regular Expression")
	text3 := "one1two2three3four"
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString(text3, -1))

	var age1 int              // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล
	fmt.Scanf("%d", &age1)    // รับค่าจากผู้ใช้ %d คือ format specifier ที่ใช้ในการรับค่าตัวเลข
	fmt.Println(age1, "age1") // แสดงผลค่าที่รับมาจากผู้ใช้

	// if else statement
	fmt.Println(" ")
	fmt.Println("if else statement")
	count := 12      // การประกาศตัวแปรแบบไม่กำหนดชนิดข้อมูล
	if count >= 18 { // ถ้าตัวแปร age1 มีค่ามากกว่าหรือเท่ากับ 18
		fmt.Println("You are an adult") // แสดงผลข้อความ You are an adult
	} else if count >= 12 { // ถ้าตัวแปร age1 มีค่ามากกว่าหรือเท่ากับ 12
		fmt.Println("You are a teenager") // แสดงผลข้อความ You are a teenager
	} else { // ถ้าไม่ตรงเงื่อนไขข้างบน
		fmt.Println("You are a child") // แสดงผลข้อความ You are a child
	}

	// switch case statement
	fmt.Println(" ")
	fmt.Println("switch case statement")
	day := 3     // การประกาศตัวแปรแบบไม่กำหนดชนิดข้อมูล
	switch day { // การเช็คค่าของตัวแปร day
	case 1: // ถ้าตัวแปร day มีค่าเท่ากับ 1
		fmt.Println("Monday") // แสดงผลข้อความ Monday
	case 2: // ถ้าตัวแปร day มีค่าเท่ากับ 2
		fmt.Println("Tuesday") // แสดงผลข้อความ Tuesday
	case 3: // ถ้าตัวแปร day มีค่าเท่ากับ 3
		fmt.Println("Wednesday") // แสดงผลข้อความ Wednesday
	case 4: // ถ้าตัวแปร day มีค่าเท่ากับ 4
		fmt.Println("Thursday") // แสดงผลข้อความ Thursday
	case 5: // ถ้าตัวแปร day มีค่าเท่ากับ 5
		fmt.Println("Friday") // แสดงผลข้อความ Friday
	case 6: // ถ้าตัวแปร day มีค่าเท่ากับ 6
		fmt.Println("Saturday") // แสดงผลข้อความ Saturday
	case 7: // ถ้าตัวแปร day มีค่าเท่ากับ 7
		fmt.Println("Sunday") // แสดงผลข้อความ Sunday
	default: // ถ้าไม่ตรงเงื่อนไขข้างบน
		fmt.Println("Invalid day") // แสดงผลข้อความ Invalid day
	}

	// array
	fmt.Println(" ")
	fmt.Println("array")
	var numbers [5]int                   // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล และจำนวนข้อมูล
	fmt.Println(numbers)                 // แสดงผลตัวแปร numbers [0 0 0 0 0]
	numbers[0] = 1                       // การกำหนดค่าให้กับตัวแปร numbers ที่ index 0
	numbers[1] = 2                       // การกำหนดค่าให้กับตัวแปร numbers ที่ index 1
	numbers[2] = 3                       // การกำหนดค่าให้กับตัวแปร numbers ที่ index 2
	numbers[3] = 4                       // การกำหนดค่าให้กับตัวแปร numbers ที่ index 3
	numbers[4] = 5                       // การกำหนดค่าให้กับตัวแปร numbers ที่ index 4
	fmt.Println(numbers)                 // แสดงผลตัวแปร numbers
	fmt.Println(numbers[0] + numbers[1]) // การบวกค่าของตัวแปร numbers ที่ index 0 และ 1
	var numbers2 = [5]int{1, 2, 3, 4, 5} // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล และค่าข้อมูล
	fmt.Println(numbers2)                // แสดงผลตัวแปร numbers2
	// numbers2.append(6) // ไม่สามารถใช้ได้ เพราะ array ไม่สามารถเพิ่มข้อมูลได้

	// slice
	// slice คือ array แต่เป็น dynamic array ที่สามารถเพิ่มข้อมูลได้
	fmt.Println(" ")
	fmt.Println("slice")
	numbers4 := []int{1, 2, 3, 4, 5, 6}      // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล และค่าข้อมูล
	fmt.Println(numbers4)                    // แสดงผลตัวแปร numbers4
	numbers4 = append(numbers4, 7, 8, 9, 10) // การเพิ่มข้อมูลใน slice
	fmt.Println(numbers4)
	fmt.Println(numbers4[4:6])             // แสดงผลข้อมูลตั้งแต่ index 4 ถึง 5 ของตัวแปร numbers4
	fmt.Println(numbers4[:6])              // แสดงผลข้อมูลตัวแรกของตัวแปร numbers4 ถึง index 5
	fmt.Println(numbers4[4:])              // แสดงผลข้อมูลตั้งแต่ index 4 ของตัวแปร numbers4 ถึงตัวสุดท้าย
	fmt.Println(numbers4[0] + numbers4[1]) // การบวกค่าของตัวแปร numbers4 ที่ index 0 และ 1
	numbers5 := []int{}                    // การประกาศตัวแปรแบบไม่กำหนดชนิดข้อมูล
	fmt.Println(numbers5)                  // แสดงผลตัวแปร numbers5
	numbers6 := [...]int{}
	fmt.Println(numbers6)
	srcSlice := []int{1, 2, 3, 4, 5} // การประกาศตัวแปรแบบไม่กำหนดชนิดข้อมูล
	a1 := []int{1, 2}
	b1 := []int{3, 4}
	c1 := []int{5, 6}
	// รวมทั้งหมดใน slice เดียว
	// Variadic Parameters
	combined := append(append(a1, b1...), append([]int{7, 8, 9}, c1...)...)
	fmt.Println(combined, "Variadic Parameters") // Output: [1 2 3 4 7 8 9 5 6]
	combined2 := append(a1, b1...)
	fmt.Println(combined2, "Variadic Parameters") // Output: [1 2 3 4]
	fmt.Println("join")
	fmt.Println(strings.Join([]string{"one", "two", "three"}, ","), "join") // การเชื่อมข้อมูลใน slice ด้วยเครื่องหมาย ,

	destSlice := make([]int, len(srcSlice))
	copy(destSlice, srcSlice)
	fmt.Println(destSlice, "copy")

	fmt.Println(" ")
	fmt.Println("sort")
	sort2 := []int{3, 2, 1, 5, 4}
	sort.Ints(sort2)
	fmt.Println(sort2, "sort")

	// map
	// map คือ key-value pair
	fmt.Println(" ")
	fmt.Println("map")
	student := map[string]string{} // การประกาศตัวแปรแบบไม่กำหนดชนิดข้อมูล
	fmt.Println(student)           // แสดงผลตัวแปร student
	student["name"] = "John Doe"   // การกำหนดค่าให้กับตัวแปร student ที่ key name
	student["age"] = "20"          // การกำหนดค่าให้กับตัวแปร student ที่ key age
	student["weight"] = "70.5"     // การกำหนดค่าให้กับตัวแปร student ที่ key weight
	fmt.Println(student)           // แสดงผลตัวแปร student
	fmt.Println(student["name"])   // แสดงผลค่าของตัวแปร student ที่ key name

	students := map[string]string{
		"name":   "John Doe",
		"age":    "20",
		"weight": "70.5",
	} // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล และค่าข้อมูล
	fmt.Println(students) // แสดงผลตัวแปร students
	student3 := map[int]bool{
		1: true,
	}
	fmt.Println(student3)
	value, ok := student["name"]    // การเช็คว่า key name มีอยู่ในตัวแปร student หรือไม่
	fmt.Println(value, ok)          // แสดงผลค่าของตัวแปร value และ ok
	value2, ok2 := student["name2"] // การเช็คว่า key name2 มีอยู่ในตัวแปร student หรือไม่
	fmt.Println(value2, ok2)        // แสดงผลค่าของตัวแปร value2 และ ok2

	// for loop
	fmt.Println(" ")
	fmt.Println("for loop")
	for i := 0; i < 5; i++ { // การกำหนดค่าเริ่มต้น, เงื่อนไข, การเพิ่มค่า
		fmt.Println(i) // แสดงผลตัวแปร i
	}
	list := []int{10, 20, 30, 40, 50} // การประกาศตัวแปรแบบไม่กำหนดชนิดข้อมูล
	for i := 0; i < len(list); i++ {  // การกำหนดค่าเริ่มต้น, เงื่อนไข, การเพิ่มค่า
		fmt.Println(list[i]) // แสดงผลข้อมูลของตัวแปร list ที่ index i
	}
	for index, value := range list { // การวนลูปข้อมูลของตัวแปร list
		if value == 20 {
			continue // การข้ามการทำงานของลูป
		}

		if value == 40 { // การเช็คเงื่อนไข
			break // การหยุดการทำงานของลูป
		}

		fmt.Println(index, value) // แสดงผล index และ value
	}

	for _, value := range student { // _ คือการไม่ใช้ตัวแปร
		fmt.Println(value) // แสดงผล value
	}

	// map loop (range)
	languages := map[string]string{
		"en": "English",
		"th": "Thai",
		"jp": "Japanese",
	}
	for key, value := range languages { // การวนลูปข้อมูลของตัวแปร languages
		fmt.Println(key, value) // แสดงผล key และ value
	}

	// while loop
	fmt.Println(" ")
	fmt.Println("while loop")
	j := 0      // การกำหนดค่าเริ่มต้น
	for j < 5 { // การกำหนดเงื่อนไข
		fmt.Println(j) // แสดงผลตัวแปร j
		j++            // การเพิ่มค่า
	}

	// do while loop
	fmt.Println(" ")
	fmt.Println("do while loop")
	k := 0 // การกำหนดค่าเริ่มต้น
	for {  // การกำหนดเงื่อนไข
		fmt.Println(k) // แสดงผลตัวแปร k
		k++            // การเพิ่มค่า
		if k >= 5 {    // การเช็คเงื่อนไข
			break // การหยุดการทำงานของลูป
		}
	}

	// function
	fmt.Println(" ")
	fmt.Println("function")
	fmt.Println(addNum(1, 2), "1, 2")                          // การเรียกใช้ฟังก์ชัน add และส่งค่า 1 และ 2 เข้าไป
	SayHello()                                                 // การเรียกใช้ฟังก์ชัน SayHello
	fmt.Println(sumAll([]int{1, 2, 3, 4, 5}), "1, 2, 3, 4, 5") // การเรียกใช้ฟังก์ชัน sumAll และส่งค่า 1, 2, 3, 4, 5 เข้าไป
	fmt.Println(sumAll2(1, 2, 3, 4, 5), "1, 2, 3, 4, 5")       // การเรียกใช้ฟังก์ชัน sumAll2 และส่งค่า 1, 2, 3, 4, 5 เข้าไป

	// multiple return
	fmt.Println(" ")
	fmt.Println("multiple return")
	isAdd, sum := isAddAndValue(2, 2) // ต้องมีตัวแปรในการรับค่าที่ return มา 2 ตัว
	fmt.Println(isAdd, sum, "2, 2")   // การเรียกใช้ฟังก์ชัน isAddAndValue และส่งค่า 1 และ 2 เข้าไป
	a, b := isAddAndValue(2, 3)
	fmt.Println(a, b, "2, 3") // การเรียกใช้ฟังก์ชัน isAddAndValue และส่งค่า 1 และ 2 เข้าไป

	//
	//
	// struct
	fmt.Println(" ")
	fmt.Println("struct")
	type Person struct { // การประกาศ struct การประกาศชนิดข้อมูล
		Name   string  // การประกาศตัวแปร Name ที่เป็น string
		Age    int     // การประกาศตัวแปร Age ที่เป็น int
		Weight float64 // การประกาศตัวแปร Weight ที่เป็น float64
	}

	person := Person{ // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล และค่าข้อมูล
		Name:   "John Doe",
		Age:    20,
		Weight: 70.5,
	}
	val := reflect.ValueOf(person)             // การแสดงชนิดข้อมูลของตัวแปร person
	fmt.Println(val.Field(1), "vale of stuct") // แสดงชนิดข้อมูลของตัวแปร person

	fmt.Println(person, person.Age) // แสดงผลตัวแปร person

	type test1 struct {
		user string
	}

	type Test struct {
		Name      string
		Age       int
		Weight    float64
		something []int // slice ใน struct
		TestList  []test1
		test      test1
	}
	test := Test{
		Name:      "John Doe",
		Age:       20,
		Weight:    70.5,
		something: []int{1, 2, 3, 4, 5},
		TestList: []test1{
			{user: "John"},
			{user: "Doe"},
		},
		test: test1{user: "John Doe"},
	}
	fmt.Println(test, test.something, test.TestList[0].user, test.test.user) // แสดงผลตัวแปร test
	fmt.Println(test.Name)                                                   // แสดงผลค่าของตัวแปร Name ใน struct test
	test.Name = "Hun"                                                        // การเปลี่ยนค่าของตัวแปร Name ใน struct test
	fmt.Println(test.Name)                                                   // แสดงผลค่าของตัวแปร Name ใน struct test
	fmt.Println(test.TestList)
	test.TestList = append(test.TestList, test1{user: "Hun"}) // การเพิ่มข้อมูลใน TestList
	fmt.Println(test.TestList)

	// package
	fmt.Println(" ")
	fmt.Println("package")
	fmt.Println(uuid.New())          // สร้าง uuid ใหม่
	fmt.Println(uuid.New().String()) // แสดง uuid ใหม่เป็น string

	fmt.Println(" ")
	fmt.Println(" ")
	hun.Hello()  // เรียกใช้ฟังก์ชัน Hello จากไฟล์ hun/hun.go
	hun.Hello2() // เรียกใช้ฟังก์ชัน Hello2 จากไฟล์ hun/hun2.go
	// จะอ่านผ่าน package ที่เรียกใช้ ไม่ใข่ชื่อโฟลเดอร์ และชื่อ package มีได้ 1 ชื่อต่อโฟลเดอร์

	q := 123
	qq := 123.123
	fmt.Println(
		float64(q) + qq,
	)

	fmt.Println(hun.AddTwoNum(1, 2)) // เรียกใช้ฟังก์ชัน AddTwoNum จากไฟล์ hun/hun2.go

	//
	//
	// check type and convert type
	// onee := 1 // การประกาศตัวแปร onee
	const onee = 1                                          // การประกาศค่าคงที่
	fmt.Println(strconv.Itoa(onee), onee)                   // แสดงผลตัวแปร one
	fmt.Println(reflect.TypeOf(onee), "onee")               // แสดงชนิดข้อมูลของตัวแปร onee
	fmt.Println(reflect.TypeOf(strconv.Itoa(onee)), "onee") // แสดงชนิดข้อมูลของตัวแปร onee

	//
	//
	//
	fmt.Println(" ")
	fmt.Println("Method with a receiver") // Method
	student2 := Student{                  // การประกาศตัวแปรแบบกำหนดชนิดข้อมูล และค่าข้อมูล
		firstName: "John",
		lastName:  "Doe",
	}
	fullName := student2.FullName("123123") // การเรียกใช้ฟังก์ชัน FullName
	fmt.Println(fullName)                   // แสดงผล fullName

}

type Student struct {
	firstName string
	lastName  string
}

func (s Student) FullName(ss string) string {
	return s.firstName + " " + s.lastName + " " + ss
}

// go run main.go คือคำสั่งที่ใช้ในการรันโปรแกรม file main.go
// go build main.go คือคำสั่งที่ใช้ในการคอมไพล์โปรแกรม file main.go เป็นไฟล์ exe
// go mod init ชื่อโปรเจค คือคำสั่งที่ใช้ในการสร้างไฟล์ go.mod ที่ใช้ในการจัดการ package ของโปรเจค

// การประกาศฟังก์ชัน
func addNum(a int, b int) int { // การประกาศฟังก์ชัน add ที่รับค่า a เป็น int และ b เป็น int และ return ค่าเป็น int
	return a + b // การ return ค่าผลรวมของ a และ b
}

func sumAll(a []int) int {
	sum := 0
	for _, value := range a {
		sum += value
	}

	return sum
}

func sumAll2(a ...int) int { // การรับค่าแบบไม่จำกัดจำนวน
	sum := 0
	fmt.Println(a, "a") // แสดงผลตัวแปร a ที่รับมา [1 2 3 4 5]
	for _, value := range a {
		sum += value
	}

	return sum
}

func isAddAndValue(a int, b int) (bool, int) { // การ return ค่าเป็นหลายค่า
	if (a+b)%2 == 0 {
		return false, a + b
	}

	return true, a + b
}

func SayHello() { // การประกาศฟังก์ชัน SayHello ที่ไม่มีการรับค่าและ return ค่า
	fmt.Println("Hello") // แสดงผลข้อความ Hello
}
