package main

/**
1.并发包含以下几种主流的实现模型。
 多进程。多进程是在操作系统层面进行并发的基本模式。同时也是开销最大的模式。在
Linux平台上，很多工具链正是采用这种模式在工作。比如某个Web服务器，它会有专门
的进程负责网络端口的监听和链接管理，还会有专门的进程负责事务和运算。这种方法
的好处在于简单、进程间互不影响，坏处在于系统开销大，因为所有的进程都是由内核
管理的。
 多线程。多线程在大部分操作系统上都属于系统层面的并发模式，也是我们使用最多的
最有效的一种模式。目前，我们所见的几乎所有工具链都会使用这种模式。它比多进程
的开销小很多，但是其开销依旧比较大，且在高并发模式下，效率会有影响。
 基于回调的非阻塞/异步IO。这种架构的诞生实际上来源于多线程模式的危机。在很多
高并发服务器开发实践中，使用多线程模式会很快耗尽服务器的内存和CPU资源。而这
种模式通过事件驱动的方式使用异步IO，使服务器持续运转，且尽可能地少用线程，降
低开销，它目前在Node.js中得到了很好的实践。但是使用这种模式，编程比多线程要复
杂，因为它把流程做了分割，对于问题本身的反应不够自然。
 协程。协程（Coroutine）本质上是一种用户态线程，不需要操作系统来进行抢占式调度，
且在真正的实现中寄存于线程中，因此，系统开销极小，可以有效提高线程的任务并发
性，而避免多线程的缺点。使用协程的优点是编程简单，结构清晰；缺点是需要语言的
支持，如果不支持，则需要用户在程序中自行实现调度器。目前，原生支持协程的语言
还很少。

2.协程的最大优势在于其“轻量级”，可以轻松创建上百万个而不
会导致系统资源衰竭，而线程和进程通常最多也不能超过1万个。这也是协程也叫轻量级线程的
原因。


3.Go 语言标准库提供的所有系统调用操作（当然也包括所有同步 IO 操作），都会出让 CPU 给其他goroutine


4.go Add(1,2) ,当被调用的函数返回时，这个goroutine也自动结束，如果这个函数有返回值，那么这个返回值将被丢弃

5.并发编程的难度在于协调，而协调就要通过交流 ,在工程上，有两种最常见的并发通信模型：共享数据和消息传递。

6.go语言推荐：不要通过共享内存来通信，而应该通过通信来共享内存

7.select机制，早在unix时代 ，调用select()函数来监控一系列的文件句柄,一旦其中一个文件句柄发生了IO动作，该select()调用就会被返回。后来该机制也被用于
实现高并发的Socket服务器程序。Go语言直接在语言级别支持select关键字，用于处理异步IO
问题。

8.之前我们示范创建的都是不带缓冲的channel，这种做法对于传递单个数据的场景可以接受， 但对于需要持续传输大量数据的场景就有些不合适了。
接下来我们介绍如何给channel带上缓冲， 从而达到消息队列的效果

9.以上示例创建的都是不带缓冲的channel,这种做法对于传递单个数据的场景可以接受，但对于需要持续传输大量数据的场景就有些不合适了，
要创建一个带缓冲的channel，其实也非常容易：
c := make(chan int, 1024) 在调用make()时将缓冲区大小作为第二个参数传入即可，比如上面这个例子就创建了一个大小 为1024的int类型channel，
即使没有读取方，写入方也可以一直往channel里写入，在缓冲区被 填完之前都不会阻塞。
从带缓冲的channel中读取数据可以使用与常规非缓冲channel完全一致的方法，但我们也可 以使用range关键来实现更为简便的循环读取：
for i := range c {     fmt.Println("Received:", i) }

10.单向channel，只能用于发送或者接收数据。channel本身必然是同时支持读写的， 否则根本没法用。假如一个channel真的只能读，那么肯定只会是空的，
因为你没机会往里面写数 据。同理，如果一个channel只允许写，即使写进去了，也没有丝毫意义，因为没有机会读取里面 的数据。
所谓的单向channel概念，其实只是对channel的一种使用限制。


*/
/*func goAdd(x, y int, ch chan int) {
	ch <- 8
	fmt.Println("hello world1")
}*/

/*
var counter int = 0

func Count( lock *sync.Mutex)  {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()

}*/

/*func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}*/

func main() {
	/*	var ch chan int
		ch = make(chan int, 1)
		go goAdd(1, 2, ch)*/
	//	<-ch

	/*lock:=&sync.Mutex{}

	for i:=0;i<10 ;i++  {
		go Count(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	*/

	/*	chs := make([]chan int, 10)
		for i := 0; i < 10; i++ {
			chs[i] = make(chan int)
			go Count(chs[i])
		}
		for _, ch := range chs {
			<-ch
		}*/
	/*str := []byte{'1','2','3'}
	 // str  := []byte("123")
	fmt.Printf("%x", md5.Sum(str))*/

	/*select {
	case hh := <-ch:
		fmt.Println("listen chs[1]", hh)
	case ch <- 2:
		fmt.Println("listen chas[2]")*/
	/*case <-chs[2]:
	fmt.Println("listen chs[2")*/
	/*default:
		fmt.Println("no io async")
	}*/

	/*var ch chan int

	ch = make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case ch <- 1:
			fmt.Println("ch<-1")
		case ch <- 0:
			fmt.Println("ch<-0")
		}
		v := <-ch
		fmt.Println("receive value is", v)

	}*/

	// channel 缓冲机制
	/*
		ch := make(chan int, 10)

		go func() {
			for i := 0; i < 10; i++ {
				ch <- 1
			}
			close(ch)
		}()

		for i := range ch {

			fmt.Println("Received:", i)
		}*/

	/*
	     //这段代码超时导致的死锁问题
	   	ch := make(chan int, 10)

	   	go func() {
	   		for i := 0; i < 10; i++ {
	   			ch <- 1
	   		}
	   		//close(ch)
	   	}()
	   	for i := 0; i < 11; i++ {
	   		n := <-ch
	   		fmt.Println("Received:", n)
	   	}
	*/

	// 处理超时机制
	/*ch := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- 1
		}
		//close(ch)
	}()

	timeoutChannel := make(chan bool, 1)
	go func() {
		timeoutChannel <- true

	}()

	for i := 0; i < 11; i++ {

		select {
		case n := <-ch:
			fmt.Println("Received:", n)
		case <-timeoutChannel:
			fmt.Println("超时处理")
		}

	}*/

	// channel的传递
	/*
		type PipeData struct {
			value int
			handler func(int) int
			next chan int
		}
		func handle(queue chan *PipeData){
			for data:=range queue {
				data.next <- data.handler(data.value)
			}
		}
	*/

	// 单向channel
	/*
		var ch1 chan int       // 正常的channel，不是单向的
		var ch2 chan<- float64 // ch2是单向的，只能用于写入float64数据
		var ch3 <-chan int     // ch3是单向的，只能用于读取int数据

		// channel 类型装换

		var ch4 = make(chan int)
		ch5 := chan<- int(ch4) // ch5是一个单向的写入的channel
		ch6 := <-chan int(ch4) // ch6是一个单向的读取的channel*/

	// 限制函数只能使用单向channel
	/**
	func Parse(ch <-chan int) {
		for value := range ch {
			fmt.Println("Parsing value", value)
		}
	}
	*/

	// 关闭channel，直接使用GO语言内置close函数,如 close(ch)

	/* x,ok := <-ch  // 如果ok返回false,表示ch已经关闭
	if ok {
		fmt.Println("closed ch")
	}
	*/

	// 出让时间片
	/*go say("world")

	say("hello")*/


}

/*
func say(s string) {

	for i := 0; i < 2; i++ {

		runtime.Gosched()

		fmt.Println(s)

	}

}*/

/*
//多核并行化

 // runtime.GOMAXPROCS(16)  设置CUP逻辑核心数，估计要废弃了
 // runtime.NumCPU() //获取CPU核心数
type Vector []float64

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1
}

const NCPU = 16

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU,(i+1)*len(v)/NCPU,u,c)
	}

	for i:=0;i<NCPU;i++ {
		<-c
	}

}*/

