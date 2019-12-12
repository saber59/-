package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/saber59/base/db"
)

//提交修改信息
func Fgai(w http.ResponseWriter, r *http.Request) {
	var re Gai
	rec, err := ioutil.ReadAll(r.Body)
	if errworke(406, err, "请求读取出错", w) {
		return
	}
	err = json.Unmarshal(rec, &re)
	if errworke(406, err, "json解析出错", w) {
		return
	}

	var link db.Mysql

	err, conn := link.LinkMysql()
	if errworke(500, err, "MySQL连接异常", w) {
		return
	}
	defer conn.Close()
	tx, err := conn.Begin()
	if errworke(500, err, "事物开启失败", w) {
		return
	}
	sql := "update student_info set stu_mes_statu = ? where stu_sno=?"
	stmt, err := tx.Prepare(sql)
	if errworke(500, err, "SQL语句预处理出错", w) {
		return
	}
	_, err = stmt.Exec(re.Stu_mes_statu, re.Stu_no)
	if errworke(500, err, "修改出错", w) {
		return
	}
	if re.Stu_mes_statu == 1 {
		var up Write
		sql = "select * from student_info where stu_sno=?"
		stmt, err := tx.Prepare(sql)
		log.Print(err)
		if errworke(500, err, "SQL语句预处理出错2", w) {
			return
		}
		res, _ := stmt.Query(re.Stu_no)
		res.Next()
		err = res.Scan(&up.Stu_no, &up.Stu_name, &up.Stu_phone, &up.Dorm_no, &up.Stu_class, &up.Image, &up.Build_num, nil)
		res.Close()
		upp, _ := json.Marshal(up)
		log.Print(up, err)
		S <- upp
	} else if re.Stu_mes_statu == -1 {
		sql = "delete from student_info where stu_sno=?"
		stmt, err := tx.Prepare(sql)
		log.Print(err)
		if errworke(500, err, "SQL语句预处理出错2", w) {
			return
		}
		_, err = stmt.Exec(re.Stu_no)

	}
	tx.Commit()
	var ret Ok
	ret.Code = 200
	ret.Message = "scuess"
	like, _ := json.Marshal(ret)
	fmt.Fprint(w, string(like))
}
