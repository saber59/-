package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/saber59/base/db"
)

func Fupup(w http.ResponseWriter, r *http.Request) {
	var link db.Mysql
	var up Write
	err, conn := link.LinkMysql()
	defer conn.Close()
	log.Print(err)
	if errworke(500, err, "数据库连接创建失败", w) {
		return
	}
	tx, err := conn.Begin()
	if errworke(500, err, "事物创建失败", w) {
		return
	}
	sql := "select * from student_info"
	stmt, err := tx.Prepare(sql)
	if errworke(500, err, "预处理失败", w) {
		return
	}
	var statu int
	res, err := stmt.Query()
	for res.Next() {
		err = res.Scan(&up.Stu_no, &up.Stu_name, &up.Stu_phone, &up.Dorm_no, &up.Stu_class, &up.Image, &up.Build_num, &statu, &up.Pinyin)
		log.Print(err)
		if statu == 0 {
			continue
		}
		upp, _ := json.Marshal(up)
		log.Print(up, err)
		S <- upp
	}
	defer res.Close()
	tx.Commit()
	var oo Ok
	oo.Code = 200
	oo.Message = "scuess"
	fina, _ := json.Marshal(oo)
	fmt.Fprint(w, string(fina))
}
