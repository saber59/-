package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/saber59/base/db"
)

//返回待审核信息的集合
func Fchange(w http.ResponseWriter, r *http.Request) {
	var re Change
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
	sql := "select * from student_info where stu_mes_statu=0 and build_num=?"
	stmt, err := tx.Prepare(sql)
	if errworke(500, err, "sql初始化失败", w) {
		return
	}
	var temp Write
	var all []Write
	res, err := stmt.Query(re.Build_num)
	for res.Next() {
		res.Scan(&temp.Stu_no, &temp.Stu_name, &temp.Stu_phone, &temp.Dorm_no, &temp.Stu_class, &temp.Image, &temp.Build_num, nil)
		all = append(all, temp)
	}
	tx.Commit()
	var retu Re_change
	retu.Mes = all
	retu.Code = 200
	retu.Message = "scuess"
	end, err := json.Marshal(retu)
	errworke(500, err, "json转码失败", w)
	fmt.Fprint(w, string(end))
}
