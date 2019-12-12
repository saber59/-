package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/saber59/base/db"
)

//写入用户信息
func Fwrite(w http.ResponseWriter, r *http.Request) {
	var re Write
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
	sql := "insert into student_info values(?,?,?,?,?,?,?,?,0)"
	stmt, err := tx.Prepare(sql)
	if errworke(500, err, "SQL语句预处理出错", w) {
		return
	}
	_, err = stmt.Exec(re.Stu_no, re.Stu_name, re.Stu_phone, re.Dorm_no, re.Stu_class, re.Image, re.Build_num, re.Pinyin)
	if errworke(500, err, "插入出错", w) {
		return
	}
	tx.Commit()
	var ret Ok
	ret.Code = 200
	ret.Message = "scuess"
	like, _ := json.Marshal(ret)
	fmt.Fprint(w, string(like))
}
