package hsklib

// 基础包
import (
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	. "github.com/henrylee2cn/pholcus/app/spider"           //必需
	"github.com/henrylee2cn/pholcus/common/goquery"         //DOM解析
	"github.com/henrylee2cn/pholcus/logs"                   //信息输出
	// . "github.com/henrylee2cn/pholcus/app/spider/common"          //选用

	// net包
	// "net/http" //设置http.Header
	// "net/url"

	// 编码包
	// "encoding/xml"
	// "encoding/json"

	// 字符串处理包
	//"regexp"
	"strconv"
	//"strings"
	// 其他包
	// "fmt"
	//"math"
	// "time"
)

func init() {
	//BaiduSearch.Register()
	HSKSearch.Register()
}

var HSKSearch = &Spider{
	Name:        "hsk语料库",
	Description: "hsk语料库搜索结果",
	// Pausetime: 300,
	Keyin:        KEYIN,
	Limit:        LIMIT,
	EnableCookie: false,
	// 禁止输出默认字段 Url/ParentUrl/DownloadTime
	NotDefaultField: true,
	// 命名空间相对于数据库名，不依赖具体数据内容，可选
	Namespace: nil,
	// 子命名空间相对于表名，可依赖具体数据内容，可选
	SubNamespace: nil,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.Aid(map[string]interface{}{"loop": [2]int{0, 1}, "Rule": "生成请求"}, "生成请求")
		},

		Trunk: map[string]*Rule{

			"生成请求": {
				AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
					var duplicatable bool

					for loop := aid["loop"].([2]int); loop[0] < loop[1]; loop[0]++ {
						if loop[0] == 0 {
							duplicatable = true
						} else {
							duplicatable = false
						}
						duplicatable = true
						logs.Log.Critical("loop0 is %v,loop1 is %v,duplicatable is %v\n", loop[0], loop[1], duplicatable)
						ctx.AddQueue(&request.Request{
							Url:        "http://yuliaoku.hanyu123.cn/index.php/hsk/index/hsk?name=" + ctx.GetKeyin() + "&c=&level=0&page=" + strconv.Itoa(loop[0]),
							Rule:       aid["Rule"].(string),
							Reloadable: duplicatable,
						})
					}
					return nil
				},
				ParseFunc: func(ctx *Context) {

					query := ctx.GetDom()
					lastpage := query.Find("ul.pagination").Find("li:nth-last-child(2)").Text()
					total, err := strconv.Atoi(lastpage)
					if err != nil {
						logs.Log.Critical("%v\n", err)
						total = 1
					}
					logs.Log.Critical("total is %v\n", total)
					if total > ctx.GetLimit() {
						total = ctx.GetLimit()
					} else if total == 0 {
						logs.Log.Critical("没有抓取到内容！")
						return
					}
					// 调用指定规则下辅助函数
					ctx.Aid(map[string]interface{}{"loop": [2]int{1, total}, "Rule": "搜索结果"})
					// 用指定规则解析响应流
					ctx.Parse("搜索结果")
				},
			},

			"搜索结果": {
				//注意：有无字段语义和是否输出数据必须保持一致
				ItemFields: []string{
					"序号",
					"国家",
					"语料",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					query.Find("div.row").Find("table").Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {

						if i != 0 {
							a := s.Find("td")
							// 结果存入Response中转
							ctx.Output(map[int]interface{}{
								0: a.First().Text(),
								1: a.First().Next().Text(),
								2: a.First().Next().Next().Text(),
							})
						}
					})
				},
			},
		},
	},
}
