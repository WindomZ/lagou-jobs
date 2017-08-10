package mobile

type JobDetail struct {
	Meta struct {
		Keywords    string // 关键字
		Description string // 描述
	}
	Title       string // 职称
	Salary      string // 薪资
	WorkAddress string // 工作地点
	JobNature   string // 全职/兼职
	WorkYear    string // 工作年限
	Education   string // 教育资历
	Temptation  string // 职位诱惑
	Company     struct {
		Title string // 公司名称
		Info  string // 公司信息
	}
	PositionDesc string // 职位描述
}
