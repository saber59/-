package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/saber59/base/db"
)

//查看个人信息
func FLook(w http.ResponseWriter, r *http.Request) {
	var re Look
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
	sql := "select * from student_info where stu_sno=?"
	stmt, err := tx.Prepare(sql)
	if errworke(500, err, "SQL语句预处理出错", w) {
		return
	}
	res, err := stmt.Query(re.Stu_sno)
	defer res.Close()
	var rel Re_look
	for res.Next() {
		err = res.Scan(&rel.Stu_no, &rel.Stu_name, &rel.Stu_phone, &rel.Dorm_no, &rel.Stu_class, &rel.Image, &rel.Build_num, &rel.Pinyin, &rel.Stu_mes_statu)
		if err != nil {
			log.Print(err)

		}
	}
	temp, err := json.Marshal(rel)
	fmt.Print(rel)
	errworke(500, err, "json编码失败", w)
	//fmt.Fprint(w, temp)
	fmt.Fprintln(w, string(temp))
	//w.Write(temp)
}
