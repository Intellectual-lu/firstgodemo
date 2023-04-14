package main

import (
	"firstGo/src/demo"
	"firstGo/src/demo/dates"
	"firstGo/src/demo/reflects"
	"firstGo/src/magazine"
	"firstGo/src/mystruct"
	"fmt"
	"github.com/Intellectual-lu/headfirstgo"
	"github.com/Intellectual-lu/headfirstgo/avg"
	"github.com/Intellectual-lu/headfirstgo/read"
	"github.com/headfirstgo/greeting"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	testPrivate()
}

// 测试封装
func testPrivate() {
	//用户可能会输出非法数据，比如月份超过12 天数超过31
	//将属性设置成非导出的，避免用户直接输入数据，提供set方法校验数据
	//now := mystruct.Date{Year: 2023, Month: 04, Day: 12}
	//fmt.Println(now)

	fmt.Println("-----------------")

	//处理非法数据
	today := mystruct.Date{}
	err := today.SetYear(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(today)

	fmt.Println("-----------------")

	year := today.Year()
	fmt.Println(year)

	fmt.Println("-----------------")

	event := mystruct.Event{}
	err = event.SetYear(2023)
	println(event.Year())
	err = event.SetTitle("中国共产主义万岁")
	if err != nil {
		log.Fatal(err)
	}
	println(event.Title())
}

// 测试方法传指针修改值
func testUpData() {
	n := Number(12)
	c := Char(12)
	fmt.Println(n, c)
	c.double()
	n.doubleByPointer()
	fmt.Println(n, c)
	n = Number(12) + 2
	fmt.Println(n)
}

type Number int
type Char int

// Double 值传递不修改原值
func (n Char) double() {
	n *= 2
}

// DoubleByPointer 指针传值可修改原值
func (n *Number) doubleByPointer() {
	*n *= 2
}

// 测试底层基础类型的定义类型

type Liters float64
type Gallons float64
type Minters float64

func testDiyField() {
	car := Liters(10.0)
	bus := Gallons(240.0)
	bike := Minters(240.0)

	//可以在具有相同基础类型的类型之间转换
	car = Liters(Gallons(10.0) * 1)

	//方法重载,调用方法
	bike.toGallons()

	fmt.Println(car.toGallons(), bus, bike.toGallons())
}

func (l Liters) toGallons() Gallons {
	return Gallons(l * 10)
}

func (l Minters) toGallons() Gallons {
	return Gallons(l * 10)
}

func testEmployee() {
	addr := magazine.Address{Street: "121 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
	emp := magazine.Employee{Name: "Peter", Salary: 600}
	emp.Address = addr
	fmt.Println(emp)
	fmt.Println(emp.State)

}

func testStruct() {
	var test mystruct.MyStruct
	fmt.Printf("%#v\n", test)
	fmt.Println(test)

	test.Name = "张三"
	fmt.Println(test.Name)

	fmt.Println("----------------------")

	//设值需要使用地址
	var a mystruct.MyStruct
	setFlag(&a)
	fmt.Println(a.Flag)
	fmt.Println("----------------------")

	pointer := returnPointer(true)
	fmt.Println(pointer.Flag)
	fmt.Println("----------------------")

	st := mystruct.MyStruct{Name: "本杰明", Number: 1.02, Flag: true}
	fmt.Println(st)
}

func setFlag(a *mystruct.MyStruct) {
	a.Flag = true
}

func returnPointer(flag bool) *mystruct.MyStruct {
	var s mystruct.MyStruct
	s.Flag = flag
	return &s
}

func calSum() {
	strings, err := reflects.GetStrings("E:\\lds\\go\\headfirstgo\\read\\strings.txt")
	if err != nil {
		log.Fatal(err)
	}
	useOne(strings)
	fmt.Println("------------------")
	useTwo(strings)
}

func useTwo(strings []string) {
	totals := make(map[string]int)
	for _, str := range strings {
		totals[str]++
	}
	fmt.Println(totals)
}

func useOne(strings []string) {
	fmt.Println(strings)
	var names []string
	var sums []int
	for _, str := range strings {
		flag := false
		for i, name := range names {
			if str == name {
				sums[i]++
				flag = true
			}
		}
		if flag == false {
			names = append(names, str)
			sums = append(sums, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, sums[i])
	}
}

func testMap() {
	m := make(map[string]string)
	m["1"] = "a"
	m["2"] = "b"
	m["3"] = "c"
	fmt.Println(m)

	//默认值为0 区分默认值和赋值0
	m2 := make(map[string]int)
	m2["1"] = 0
	fmt.Println("m2[1]:", m2["1"], ",  m2[2]:", m2["2"])
	_, ok := m2["1"]
	fmt.Print("m2[1]是否赋值:", ok)
	_, ok = m2["2"]
	fmt.Print("\tm2[2]是否赋值:", ok)
	fmt.Println()

	//遍历 (每次遍历的顺序不尽相同)
	for k, v := range m {
		fmt.Println("key:", k, "\tvalue:", v)
	}

	//删除
	delete(m2, "1")
	_, ok = m2["1"]
	fmt.Println(m2)
	fmt.Print("m2[1]是否赋值:", ok)
}

func findAverage(numbers ...float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
func findInRange(min float64, max float64, numbers ...float64) []float64 {
	var ranges []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			ranges = append(ranges, number)
		}
	}
	return ranges
}

func findMax(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func sumArg() {
	numbers := os.Args[1:]
	var num []float64
	var total float64 = 0
	for _, number := range numbers {
		float, err := strconv.ParseFloat(number, 64)
		if err != nil {
			log.Fatal(err)
		}
		num = append(num, float)
		total += float
	}
	fmt.Println("total: ", total)
	fmt.Printf("average: %.2f", total/float64(len(numbers)))
	fmt.Println()
	fmt.Printf("average: %.2f", findAverage(num...))
}

func testArg() {
	charArray5 := []string{"a", "b", "c", "d"}
	fmt.Println(charArray5, "---------")
	//获取命令行参数
	fmt.Println(os.Args[1:])
}
func testAppend() {

	//没有遵循惯例将append函数的返回值赋给传入的变量。当我们给切片charArray3的一个元素赋值的时候，
	//我们能看到charArray2中的体现。因为charArray3和charArray2碰巧都共享相同的底层数组
	//需要避免一下情况
	charArray := []string{"a", "b", "c", "d"}
	charArray2 := append(charArray, "e", "f")
	charArray3 := append(charArray2, "g", "h")
	charArray4 := append(charArray3, "i", "j")
	fmt.Println(charArray)
	fmt.Println(charArray2)
	fmt.Println(charArray3)
	fmt.Println(charArray4)
	fmt.Println("----------------")
	charArray3[0] = "z"
	fmt.Println(charArray)
	fmt.Println(charArray2)
	fmt.Println(charArray3)
	fmt.Println(charArray4)
	fmt.Println("----------------")
	//惯例是将函数的返回值赋给你传入的那个切片变量
	charArray5 := []string{"a", "b", "c", "d"}
	charArray5 = append(charArray5, "e", "f")
	charArray5 = append(charArray5, "g", "h")
	charArray5 = append(charArray5, "i", "j")
	fmt.Println(charArray5)
}

func sliceTest() {
	var a []int
	a = make([]int, 7)
	a[0] = 1
	fmt.Println(a)
	fmt.Println(len(a))

	arrays := []string{"a", "b", "c", "d"}
	fmt.Println(arrays)

	//切片
	strings := arrays[0:2]
	fmt.Println(strings)
	//修改切片的值也会修改数据的值
	strings[0] = "z"
	i := arrays[1:]
	fmt.Println(i)
	fmt.Println(arrays)
}
func readFile() {
	numbers, err := read.GetFloats("E:\\lds\\go\\headfirstgo\\read\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	avg.GetAvg(numbers)
}

func array() {
	var arr [5]string
	arr[0] = "1"
	arr[1] = "2"
	fmt.Println(arr[0])
	fmt.Println(arr[4])

	var dataArray [3]time.Time
	dataArray[0] = time.Now()
	dataArray[1] = time.Now()
	fmt.Println(dataArray[0])
	fmt.Println(dataArray[2])

	numberArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numberArray)
	fmt.Printf("%#v", numberArray)
	fmt.Println()
	for i := 0; i < len(numberArray); i++ {
		fmt.Print(numberArray[i], " ")
	}
	fmt.Println()
	for index, num := range numberArray {
		fmt.Println(index, num, " ")
	}

}

func test1() {
	greeting.Hello()
	fmt.Println("Hello----------------------------------")
	greeting.Hi()
	fmt.Println("Hi----------------------------------")
	demo.CalLiters()
	fmt.Println("CalLiters----------------------------------")
	demo.Grade()
	fmt.Println("Grade----------------------------------")
	demo.GuessNum()
	fmt.Println("GuessNum----------------------------------")
	demo.Hello()
	fmt.Println("Hello----------------------------------")
	demo.Pointer()
	fmt.Println("Pointer----------------------------------")
	demo.Sqrt()
	fmt.Println("Sqrt----------------------------------")
	input, err := keyborad.GetFloatInPut()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("you enter number is: ", input)
	fmt.Println("input----------------------------------")

	day := dates.WeekToDay(3)
	fmt.Println("WeekToDay----------------------------------", day)
	week := dates.DayToWeek(14)
	fmt.Println("DayToWeek----------------------------------", week)

	fmt.Println("demo.Days----------------------------------", dates.Days)
}
