package function

//查看个人信息
type Look struct {
	Stu_sno   string `json:"Stu_sno"`
	Build_num string `json:"Build_num"`
}

//返回个人信息
type Re_look struct {
	Stu_no        string `json:"Stu_sno"`
	Stu_name      string `json:"Stu_name"`
	Stu_phone     string `json:"Stu_phone"`
	Pinyin        string `json:"Pinyin"`
	Stu_class     string `json:"Stu_class"`
	Build_num     string `json:"Build_num"`
	Dorm_no       string `json:"Dorm_no"`
	Image         string `json:"Image"`
	Stu_mes_statu int64  `json:"Stu_mes_statu"`
}

//根据宿舍号显示宿舍人员信息
type LookRidStu struct {
	Build_num string `json:"Build_num"`
	Dorm_no   string `json:"Dorm_no"`
}

//根据宿舍号显示宿舍内异常人员信息
type Abnormal struct {
	Build_num string `json:"Build_num"`
	Dorm_no   string `json:"Dorm_no"`
}

//查看待审核人员信息
type Change struct {
	Build_num string `json:"Build_num"`
}

//写入人员信息
type Write struct {
	Stu_no    string `json:"Stu_sno"`
	Stu_name  string `json:"Stu_name"`
	Stu_phone string `json:"Stu_phone"`
	Pinyin    string `json:"Pinyin"`
	Stu_class string `json:"Stu_class"`
	Build_num string `json:"Build_num"`
	Dorm_no   string `json:"Dorm_no"`
	Image     string `json:"Image"`
	//Stu_mes_statu int64  `json:"Stu_mes_statu"`
}

// //返回宿舍人员信息
// type Re_lookRidStu struct {
// 	Code    int64    `json:"Code"`
// 	Message string   `json:"Message"`
// 	Record  []Re_mes `json:"Record"`
// }

// //查寝时所需要的个人信息
// type Re_mes struct {
// 	Stu_no    string `json:"Stu_sno"`
// 	Stu_name  string `json:"Stu_name"`
// 	Stu_phone string `json:"Stu_phone"`
// 	Image     string `json:"Image"`
// }

// //返回寝室内异常人员信息
// type Re_abnormal struct {
// 	Code    int64    `json:"Code"`
// 	Message string   `json:"Message"`
// 	Record  []Re_mes `json:"Record"`
// }

//返回所有待审核的信息
type Re_change struct {
	Code    int64   `json:"Code"`
	Message string  `json:"Message"`
	Mes     []Write `json:"Mes"`
}

//修改学生状态
type Gai struct {
	Stu_no        string `json:"Stu_sno"`
	Stu_mes_statu int64  `json:"Stu_mes_statu"`
}

//正常返回
type Ok struct {
	Code    int64  `json:"Code"`
	Message string `json:"Message"`
}

// //状态表更新
// type Update struct {
// 	Stu_no    string `json:"Stu_sno"`
// 	Stu_name  string `json:"Stu_name"`
// 	Build_num string `json:"Build_num"`
// 	Dorm_no   string `json:"Dorm_no"`
// }
