package lagou

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestJobs_crawlJobDetail(t *testing.T) {
	spider, err := New("../tests/config.json")
	if err != nil {
		t.Fatal(err)
	} else if err := spider.ready(); err != nil {
		t.Fatal(err)
	}

	job, err := spider.crawlJobDetail(3100439)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Golang/Java后台开发工程师（深圳,Golang/Java后台开发工程师"+
		"（深圳招聘,腾讯科技(深圳)有限公司Golang/Java后台开发工程师（深圳招聘", job.Meta.Keywords)
	assert.Equal(t, "java,go,高并发 科兴科学园", job.Meta.Description)

	assert.Equal(t, "Golang/Java后台开发工程师（深圳", job.Title)
	assert.Equal(t, "15k-30k", job.Salary)
	assert.Equal(t, "深圳", job.WorkAddress)
	assert.Equal(t, "全职", job.JobNature)
	assert.Equal(t, "3-5年", job.WorkYear)
	assert.Equal(t, "本科及以上", job.Education)

	assert.Equal(t, "腾讯", job.Company.Title)
	assert.Equal(t, "移动互联网,游戏/上市公司/2000人以上", job.Company.Info)
}

func TestJobs_filterJobDetail(t *testing.T) {
	spider, err := New("../tests/config.json")
	if err != nil {
		t.Fatal(err)
	} else if err := spider.ready(); err != nil {
		t.Fatal(err)
	}

	assert.True(t, spider.filterJobDetail(3100439))
}
