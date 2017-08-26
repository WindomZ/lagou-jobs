# lagou-jobs

[![Build Status](https://travis-ci.org/WindomZ/lagou-jobs.svg?branch=master)](https://travis-ci.org/WindomZ/lagou-jobs)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/lagou-jobs)](https://goreportcard.com/report/github.com/WindomZ/lagou-jobs)

> 采集筛选lagou招聘信息，开箱即用，快速并发

## Feature
- 高并发爬虫实现
- 简化且实用规则：城市、关键字等搜索 + 薪资、关键字等筛选
- 更为实用且友好的筛选机制
- 简单易用的命令行工具
- 更多特性在[Roadmap](#roadmap)

## Install
```bash
go get -u github.com/WindomZ/lagou-jobs
```

## Usage

### Config
可以基于下列注释信息配置项目中提供的空白`config.json`文件：
```json
{
  "userAgent": "",       // User Agent
  "requestTimeout": 15,  // 请求超时时间
  "requestInterval": 0,  // 请求间隔时间
  "search": {
    "city": "",          // 搜索城市，必填
    "keywords": [        // 搜索关键字
      ""                 // 必填
    ],
    "company": {
      "excludeID": [     // 排除公司ID，可以基于上次搜索结果设置
        0                // 可为0
      ]
    },
    "position": {
      "excludeID": [     // 排除职位ID，可以基于上次搜索结果设置
        0                // 可为0
      ],
      "filter": {        // 职位信息过滤器
        "include": [     // 必须包含的关键字，为空则默认全通过
          ""             // 可为空
        ],
        "exclude": [     // 必须排除的关键字，可为空
          ""             // 可为空
        ]
      },
      "salary": {        // 薪资筛选
        "min": 0,        // 最低薪资，为0则不进行筛选
        "max": 0         // 最高薪资，为0则基于最低薪资筛选
      }
    }
  },
  "output": {            // 输出格式，选一个填写
    "files": {           // 输出文件
      "json": ""         // 输出JSON文件
    },
    "http": {            // 输出HTML(暂未实现)
      "port": 0          // 网页端口(暂未实现)
    }
  }
}
```

### Execute
项目路径下，在终端运行下列命令：
```bash
lagou-jobs config.json
```

将会得到配置中"output"的输出信息。

## Contributing
欢迎提出请求，在[issues page](https://github.com/WindomZ/lagou-jobs/issues)报告错误，提出建议和讨论。

如果你喜欢或支持，欢迎点击上面 :star:Star

## Roadmap

- [x] 爬虫业务整体框架
- [x] 高并发请求及筛选
- [x] 请求间隔调控设定
- [ ] 预防官方屏蔽机制
- [x] 归类公司、职位信息
- [x] 更友好的薪资筛选(一个或者范围内薪资)
- [ ] 学历筛选(很多都是本科，暂时不做)
- [ ] 输出JSON格式
- [ ] 输出HTML格式，网页可视化

## License

[MIT](https://github.com/WindomZ/lagou-jobs/blob/master/LICENSE)
