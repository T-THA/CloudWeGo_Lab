package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	student "student/kitex/kitex_gen/student/studentservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8848")
	svr := student.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr))
	//svr := student.NewServer(new(StudentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
