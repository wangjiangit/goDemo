package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

// 网络编程

/**
GO 语言标准库里提供net包，支持基于IP层，TCP/UDP层，更高层次(HTTP,FTP,SMTP)的网络操作,其中IP层
的称为Raw Socket
第一：SOCKET编程
。以前我们使用Socket编程时， 会按照如下步骤展开。
(1) 建立Socket：使用socket()函数。
(2) 绑定Socket：使用bind()函数。
(3) 监听：使用listen()函数。或者连接：使用connect()函数。
(4) 接受连接：使用accept()函数。
(5) 接收：使用receive()函数。或者发送：使用send()函数。
Go语言标准库对此过程进行了抽象和封装。无论我们期望使用什么协议建立什么形式的连 接，都只需要调用net.Dial()即可。


实际上，Dial()函数是对DialTCP()、DialUDP()、DialIP()和DialUnix()的封装。我 们也可以直接调用这些函数，它们的功能是一致的。这些函数的原型如下：
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err error)
func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err error)
func DialIP(netProto string, laddr, raddr *IPAddr) (*IPConn, error)
func DialUnix(net string, laddr, raddr *UnixAddr) (c *UnixConn, err error)

net.ResolveTCPAddr() // 用户解析地址和端口号
net.DialTCP() //用于建立连接
net.ParseIP() //验证IP地址有效性
func Ipv4Mask(a,b,c,d byte) IPMask  // 创建子网掩码
func (ip IP) DefaultMask() IPMask // 获取默认子网掩码
根据域名查找IP的代码如下：
func ResolveIPAddr(net, addr string) (*IPAddr, error)
func LookupHost(name string) (cname string, addrs []string, err error)；

第二：HTTP编程
GO语言标准库内建提供了net/http包，涵盖了HTTP客户端和服务端的具体实现。
  *HTTP客户端
	net/http包的Client类型提供如下方法。
	func (c *Client) Get(url string) (r *Response ,err error)
	func (c *Client) Post(url,string,bodyType string, body io.Reader)(r *Response,err error)
    func (c *Client) PostForm(url string,data url.Values)(r *Response,err error)
	func (c *Client) Head(url string)(r *Response,err error)
	func (c *Client) Do(req *Request)(resp *Response,err error)




*/
func main() {

	//-------------------------SOCKET 编程(START)---------------------------------------

	//ICMP 示例程序
	/*if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := net.Dial("ip4:icmp", service)
	checkError(err)

	var msg [512]byte

	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum
	msg[3] = 0  // checksum
	msg[4] = 0  // identifier[0]
	msg[5] = 13 // identifier[1]
	msg[6] = 0  // sequence[0]
	msg[7] = 37 // sequence[0]
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:len])
	checkError(err)

	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("Got response")

	if msg[5] == 13 {
		fmt.Println("Identifier matches")
	}

	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}

	os.Exit(0)*/

	// TCP示例程序
	/*if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)*/
	//-------------------------SOCKET 编程(END)---------------------------------------

	//-------------------------HTTP 编程(START)------------------------------------------
	/*resp, err := http.Get("http://5izan.com")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout,resp.Body)*/

	// http.Post("http://example.com/upload", "image/jpeg", &imageDataBuf)
	/*
		resp, err := http.Post("http://example.com/upload", "image/jpeg", &imageDataBuf)
		if err != nil { // 处理错误           return }

			if resp.StatusCode != http.StatusOK {
				// 处理错误
				return
			}

			if resp.StatusCode != http.StatusOK {
				// 处理错误
				return
			}*/

/*	resp, err := http.PostForm("http://example.com/posts", url.Values{"title": {"article title"}, "content": {"article body"}})
	if err != nil {
		// 处理错误
		return
	}*/


  // resp,err :=http.Head("http://example.com/")


  //在多数情况下，http.Get()和http.PostForm() 就可以满足需求，但是如果我们发起的 HTTP请求需要更多的定制信息

 /*  req,err :=http.NewRequest("GET","http://example.com",nil)
  req.Header.Add("User-Agent","Gobook custome user-agent")

  client :=&http.Client{ }
  resp, err := client.Do(req)*/

 // HTTP 服务端技术
/*	http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path)) })
	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil))
	//或者是：
	ss := &http.Server{
		Addr:           ":10443",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(ss.ListenAndServeTLS("cert.pem", "key.pem"))*/









	//-------------------------HTTP 编程(END)--------------------------------------------

}
func checkSum(msg []byte) uint16 {
	sum := 0
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
