package main

import (
	"errors"
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("i am nokia")
}

type IphonePhone struct {
}

func (iphonePhone IphonePhone) call() {
	fmt.Println("i am iphone")
}

/*type Bijiao interface {
	Less() bool
}
*/
/*type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

type Rect struct {
	x, y          float64
	width, height float64
}

func (rect *Rect) Area() float64 {
	return rect.height * rect.width
}
*/
// 构造函数

/*func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height} // 实例化类型
}*/

// 匿名组合和非匿名组合
/*type Base struct {
	Name string
}

func (base *Base) Foo() {}
func (base *Base) Bar() {}

type Foo struct {
	Base

	// *Base
}

func (foo *Foo) Bar() {
	foo.Base.Bar()

}*/

// 侵入式接口(实现类需要明确声明自己实现了某个接口)

// 非入侵式接口
//一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口。
/*type File struct {
}

func (file *File) Read(buf []byte) (n int, err error) {
	return 1, errors.New("error1")
}
func (f *File) Write(buf []byte) (n int, err error) {
	return 1, errors.New("error2")
}

func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return 1, errors.New("error3")
}

func (f *File) Close() error {
	return nil
}

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}
type IWriter interface {
	Write(buf []byte) (n int, err error)
}
type ICloser interface {
	Close() error
}

var file1 IFile = new(File)
var file2 IReader = new(File)
var file3 IWriter = new(File)
var file4 ICloser = new(File)
*/
/*
Go语言的非侵入式接口，看似只是做了很小的文法调整，实则影响深远。
其一，Go语言的标准库，再也不需要绘制类库的继承树图。你一定见过不少C++、Java、C#
类库的继承树图。这里给个Java继承树图：
http://docs.oracle.com/javase/1.4.2/docs/api/overview-tree.html
在Go中，类的继承树并无意义，你只需要知道这个类实现了哪些方法，每个方法是啥含义
就足够了。
其二，实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才
合理。接口由使用方按需定义，而不用事前规划。
其三，不用为了实现一个接口而导入一个包，因为多引用一个外部的包，就意味着更多的耦
合。接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口。*/

// 接口赋值
// 1.将对象实例赋值给接口
// 2.将一个接口赋值给另外一个接口

/*type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b

}

func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAddr interface {
	Less(b Integer) bool
	Add(b Integer)
}


var a Integer =1
//var b LessAddr =a 非法赋值
var b LessAddr=&a
*/

//接口赋值并不要求两个接口必须等价。如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A
/*type InterfaceA interface {
	Music()
	Draw()
}
type InterfaceB interface {
	Draw()
}

type Person struct {
}

func (person *Person) Music() {
	fmt.Println("music func")
}

func (person *Person) Draw() {
	fmt.Println("draw func")
}

var aa InterfaceA = &Person{}
var cc InterfaceB = aa*/

// 接口查询
/*type Interface1 interface {
	Music()
}

type Person struct {
}

func (person *Person) Music() {
	fmt.Println("music")
}
*/

// 类型查询
type Stringer interface {
	String() string
}

func PrintLn(arg ...interface{}) {
	for _, v := range arg {
		switch v1 := v.(type) {
		case int:
		case string:
		case float32:
		default:
			if v1, ok := v.(Stringer); ok {
				val := v1.String()
			}else{

			}

		}
	}
}

func main() {
	/*	var c Interface1 = &Person{}

		if p,ok:=c.(*Person) ;ok{
			p.Music()
		}else{
			fmt.Println("no")
		}*/
	/*var v2 string
	  var v3 [10]int
	  var v4 []int
	  var v5 struct {
	  	f int
	  }
	  var v6 *int
	  var v7 map[string]int
	  var v8 func(a int) int

	  var (
	  	v9  int
	  	v10 struct {
	  			name string
	  		}
	  )*/

	/*	var str1 string
		str1="helloworld中"

		fmt.Println("hello world %u",len(str1))*/
	/*var str string
	  str = "Hello world"
	  ch := str[0]
	  fmt.Printf(" %s is %d\n", str, len(str))
	  fmt.Printf(" c is %c", ch)*/

	/*mySlice1:=make([]int,5)
	  mySlice2:=make([]int,3)*/

	/*	mySlice3:=make([]int,4,20)
		mySlice2:=[]int{1,2,3,4,5}
		mySlice3=append(mySlice3,mySlice2...)
		for i,v:= range mySlice3{
			fmt.Println(i,v)
		}

		fmt.Println("Cap",cap(mySlice3))
		fmt.Println("len",len(mySlice3))


	*/
	/*
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := []int{5, 6, 7}

		copy(slice1, slice2)

		for i,v:= range slice1 {
			fmt.Println(i, v)
		}*/

	// map 字典

	/*	type PersonInfo struct {
			ID      string
			Name    string
			Address string
		}
	*/
	//var myMap map[string]PersonInfo
	/*myMap := map[string]PersonInfo{
		"1234": PersonInfo{"1", "Jack", "Room 101"},
		"1235": PersonInfo{"1", "Jack", "Room 101"},
	}*/
	//myMap=make(map[string] PersonInfo)
	// 赋值
	//myMap["1234"]=PersonInfo{"1", "Jack", "Room 101"}

	//delete(myMap, "1234")

	/*value,ok :=myMap["1234"]
	  if ok{
	  	fmt.Println(ok,value.Address)
	  }*/

	// 控制语句

	//var i int = 2
	/*if i < 5 {
	  	fmt.Println("Y")
	  }else{
	  	fmt.Println("N")
	  }*/

	/*switch i {
	  case 1:
	  	fmt.Println("1")
	  case 2:
	  	fallthrough
	  case 6:
	  	//fallthrough
	  	//fmt.Println("6")
	  default:
	  	fmt.Println("default")
	  }*/

	/*	switch  {
		case i>1 && i<12 :
			fmt.Println("1<i<12")
		case i>56:
			fmt.Println(">56")

		}
	*/

	//sum := 10
	/*
		for i := 0; i < sum; i++ {
			fmt.Println(sum)
		}
		fmt.Println(sum)*/

	/*
	   	for {
	   		sum++
	   		fmt.Println(sum)
	   		if sum > 10 {
	   			goto JumpTag
	   			break
	   		}
	   	}

	   JumpTag:
	   	fmt.Println("hello")
	*/

	/*	sum, _ := Add(3, 4)
		fmt.Println(sum)*/

	/*myArr:=[]int{1,2,3,4}
	  myFunc(1,2,3,4)
	  myFunc(myArr...)*/
	/*var v1 int = 1
	  var v2 int64 = 234
	  var v3 string = "hello"
	  var v4 float32 = 1.234
	  MyPrintf(v1,v2,v3,v4)*/

	// 匿名函数
	/*
		sum := func(x, y int) int {
			return x + y
		}(3, 4)

		fmt.Println(sum)*/

	/*var j int = 5

	  a := func() (func()) {
	  	var i int = 10
	  	return func() { fmt.Printf("i, j: %d, %d\n", i, j) }
	  }()

	  a()

	  j *= 2

	  a()*/

	/*	b := func1(4, 3)
		c,d:=b(3,"hllo")

		fmt.Println(c,d)*/

	// 接口实现

	/*var phone Phone

	  phone = new(NokiaPhone)
	  phone.call()

	  phone = new(IphonePhone)
	  phone.call()*/

	/*defer func() {
			fmt.Println("clean the end1")
		}()
		defer func() {
			fmt.Println("clean the end2")
		}()
	  	 fmt.Println(error1())*/

	// panic 和recover 使用
	/*fmt.Println("c")
	  defer func() { // 必须要先声明defer，否则不能捕获到panic异常
	  	fmt.Println("d")
	  	if err := recover(); err != nil {
	  		fmt.Println(err) // 这里的err其实就是panic传入的内容
	  	}
	  	fmt.Println("e")
	  }()
	  f() //开始调用f
	  fmt.Println("f") //这里开始下面代码不会再执行*/
	/*var a Integer = 9
	  fmt.Println(a.Less(5))*/

	/*var a = [3]int{1, 2, 3}
	  var b = &a
	  b[1]++
	  fmt.Println(a, *b)*/

	// 初始化类型
	//var rect =Rect{x:100,y:200,height:3,width:4}
	//var rect =new(Rect)
	/*var rect = &Rect{x: 100, y: 200, height: 3, width: 4}
	  fmt.Println(rect.Area())*/

	//匿名组合（继承）
}

//因此需要先牢记这样的规则：小写字母开头的函数只在本包内可见，大写字母开头的函数才 能被其他包使用。 这个规则也适用于类型和变量的可见性。
func Add(a int, b int) (sum int, err error) {

	return a + b, nil

}

func myFunc(args ...int) {
	for _, arg := range args {

		fmt.Println(arg)
	}
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func func1(a int, b int) func(int, string) (int, string) {

	return func(a int, b string) (int, string) {
		return a, b
	}

}

func error1() string {
	return errors.New("error1 demo").Error()
}

func f() {
	fmt.Println("a")
	panic("异常信息")
	fmt.Println("b") //这里开始下面代码不会再执行
}
