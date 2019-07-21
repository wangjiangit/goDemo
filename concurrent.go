package main

import "fmt"

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


*/
func goAdd(x, y int, ch chan int) {
	ch <- 8
	fmt.Println("hello world1")
}

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
	var ch chan int
	ch = make(chan int, 1)
	go goAdd(1, 2, ch)
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

	select {
	case hh := <-ch:
		fmt.Println("listen chs[1]", hh)
	/*case ch <- 2:
		fmt.Println("listen chas[2]")*/
	/*case <-chs[2]:
	fmt.Println("listen chs[2")*/
	/*default:
		fmt.Println("no io async")*/
	}

}
