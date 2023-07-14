package main

import (
	"context"
	"errors"
	"fmt"
	student "student/kitex/kitex_gen/student"
)

var myMap = make(map[int32]student.Student)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, stu *student.Student) (resp *student.RegisterResp, err error) {
	// TODO: Your code here...
	myMap[stu.Id] = student.Student{stu.Id, stu.Name, stu.College, stu.Email}
	resp = &student.RegisterResp{true, string(stu.Id)}
	fmt.Println(myMap)
	return resp, nil
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *student.QueryReq) (resp *student.Student, err error) {
	// TODO: Your code here...
	fmt.Printf("%d %v\n", req.Id, myMap)
	if ret, EXIST := myMap[req.Id]; EXIST {
		return &ret, nil
	} else {
		err = errors.New("No such student")
		return
	}
}
