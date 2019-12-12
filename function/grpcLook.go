package function

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/saber59/base/db"
	"github.com/saber59/base/getInfo"
	"google.golang.org/grpc"
)

type Server struct{}

//服务端
func GrpcLook() {
	addy, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, err := net.ListenTCP("tcp", addy)
	inerrworke(500, err, "rpc启动失败")
	server := grpc.NewServer()
	getInfo.RegisterGetInfoServer(server, &Server{})

	server.Serve(conn)
}

func (s *Server) Info(ctx context.Context, req *getInfo.Rec) (*getInfo.Send, error) {

	var sno string
	err := json.Unmarshal([]byte(req.R), &sno)
	log.Print(err)
	var link db.Mysql

	err, conn := link.LinkMysql()
	inerrworke(500, err, "MySQL连接异常")
	defer conn.Close()
	tx, err := conn.Begin()
	inerrworke(500, err, "事物开启失败")
	sql := "select stu_name,dorm_no,build_num from student_info where stu_sno=?"
	stmt, err := tx.Prepare(sql)
	inerrworke(500, err, "SQL语句预处理出错")
	res, err := stmt.Query(sno)
	//var re_sno string
	var re_name string
	var re_dormno string
	var re_buildnum string
	for res.Next() {
		err := res.Scan(&re_name, &re_dormno, &re_buildnum)
		if err != nil {
			log.Print(err)
			return nil, err
		}

	}

	var sss getInfo.Send
	sss.StuName = re_name
	sss.DormNo = re_dormno
	sss.BuildNum = re_buildnum
	return &sss, nil //返回
}
