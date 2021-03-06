## 什么是filter
用于判断预先定义好的一系列条件是否成立的组件跟协议；filter是由一组json格式的数组定义的

## filter格式
filter 由一个大的json格式的数组内部至少要有两个数组构成；

filter分为条件项跟执行项；最后一个数组表示执行项；最后一个数组前面的数组都为条件项

### filter举例

###### 条件项：

1、订单数量大于

2、最后一次登录时间小于当前时间

###### 执行项：（条件项1跟2同时成立时才执行）

1、设置coupon_id字段为1000
```
[
    [
        "order_number",
        ">",
        1
    ],
    [
        "last_login_time",
        "<",
        "now"
    ],
    [
        "coupon_id",
        "=",
        "1000"
    ]
]
```


###### 条件项目：

1、用户所在城市为 上海,天津,北京 其中的一个

2、当前是在10点到19点时

###### 执行项

1、设置name字段为golang
```text
[
    [
        "city" ,
        "in",
        "上海,天津,北京"
    ],
    [
        "hour",
        "between",
        "10,19"
    ],
    [
       "name","=","golang"
    ]
]

```

#### 因为最后一个数组项目为执行项，所以最后一项也可以是一个多项数组，用于设置多个值
###### 条件项目：

1、用户所在城市为 上海,天津,北京 其中的一个

2、当前是在10点到19点时

###### 执行项（条件项1跟2同时成立时才执行）

1、设置name字段为 golang

2、设置的desc字段为 description
```text
[
    [
        "city" ,
        "in",
        "上海,天津,北京"
    ],
    [
        "hour",
        "between",
        "10,19"
    ],
    [
        [
            "name","=","golang"
        ],
        [
            "desc","=","description"
        ]
    ]
]

```

#### 多条件组合
###### 条件项：

1、用户所在城市为 上海,天津,北京 其中的一个

2、当前是在10点到19点时

###### 执行项：（条件项1成立或者2成立执行）

1、设置name字段为golang

```text
[
    [
        "or",
        "=>",
        [
            [
                "city",
                "in",
                "上海,天津,北京"
            ],
            [
                "hour",
                "between",
                "10,19"
            ]
        ]
    ],
    [
        "name",
        "=",
        "golang"
    ]
]
```

###### 条件项：

1、用户所在城市为 上海,天津,北京 其中的一个

2、当前是在10点到19点时

###### 执行项：（条件项1跟2同时成立时才执行）

1、设置name字段为golang

```text
[
    [
        "and",
        "=>",
        [
            [
                "city",
                "in",
                "上海,天津,北京"
            ],
            [
                "hour",
                "between",
                "10,19"
            ]
        ]
    ],
    [
        "name",
        "=",
        "golang"
    ]
]
```


###### 条件项：

1、用户所在城市为 上海,天津,北京 其中的一个

2、当前是在10点到19点时

###### 执行项：（条件项1跟2都不成立时才执行）

1、设置name字段为golang

```text
[
    [
        "not",
        "=>",
        [
            [
                "city",
                "in",
                "上海,天津,北京"
            ],
            [
                "hour",
                "between",
                "10,19"
            ]
        ]
    ],
    [
        "name",
        "=",
        "golang"
    ]
]
```

#### 内部多条件组合
###### 条件项：分为两大组合块

1、
* 用户所在城市为 上海,天津,北京 其中的一个

* 当前时间在10点到19点之间

2、
* 当前版本大于3.4.5

* 用户IP在192.168.1.1/24端

###### 执行项：（条件项1全部成立且条件2全部成立才执行）

1、设置name字段为golang

```text
[
    [
        "or",
        "=>",
        [
            [
                [
                    "city",
                    "in",
                    "上海,天津,北京"
                ],
                [
                    "hour",
                    "between",
                    "10,19"
                ]
            ],

            [
                [
                    "version",
                    "vgt",
                    "3.4.5"
                ],
                [
                    "ip",
                    "iir",
                    "192.168.1.1/24"
                ]
            ]
        ]
    ],
    [
        "name",
        "=",
        "golang"
    ]
]
```

### 已支持变量
变量名 | 结果 | 描述
--- | --- | ---
success | 1 | 该变量永远返回1
rand  | 随机值 | 获取[1-100]内的随机值
ip | 当前用户IP | 需要先将用户IP注入的Context中
country | 获取注入IP所在国家 | 需要先将用户IP注入的Context中
province | 获取注入IP所在省份 | 需要先将用户IP注入的Context中
city | 获取注入IP所在城市 | 需要先将用户IP注入的Context中
timestamp | 获取当前时间时间戳（单位秒）| number
ts_simple | 获取当前时间的时间戳int64类型| number
second | 获取当前时间的秒 | number
minute | 获取当前时间的分钟 | number
hour | 获取当前时间的小时 | number
day | 获取当前时间在当年中的第几天 | number
month | 获取当前时间的月份 | number
year | 获取当前时间的年份 | number
wday | 获取当时时间为本周的第几天 | 周日为0，周六为6
date | 获取当前时间的的日期（具体到天） | string(2006-01-02)
time | 获取当前时间 | string(2006-01-02 15:04:05)
ua | 获取user_agent信息 | 需要先注入的Context中
referer| 获取refer信息 | 需要先注入的Context中
is_login | 获取当前用户是否登录 | 需要先蒋user_id注入的Context中
version | 获取版本信息 | 需要先注入的Context中
platform | 获取平台信息 | 需要先注入的Context中
channel | 获取渠道信息 | 需要先注入的Context中
uid | 获取用户ID | 需要先注入的Context中
device | 获取设备信息| 需要先注入的Context中
user_tag | 获取用户标签 | 需要先注入的Context中
get.xxx | 获取注入到context中form信息 | 需要先注入的Context中
data.xxx | 获取传入的对象的字段值 | 传入的对象可以是ptr或者struct或者是map
calc.experation | 获取计算表达式 | 计算experation 表达式例如（calc.__value1 * __value2）需要业务方实现CalcFactorGet接口获取变量的值，需要返回float64类型的值
freq.xxx | 获取xxx对应的频次 | 用户频次控制；需要业务方实现FrequencyGetter接口
ctx.xxx | 获取对应Context中的值 | 需要先注入的Context中


### 已支持运算符 

> var = value: var 称为变量; = 叫做操作符；value 叫做比较值

 运算符 | 描述 | 逻辑
 ---  | --- | ---
   =  | 比较运算符 		| 比较 var 跟value 字符串或者数字是否相等
   != | 比较运算符 		| 与 = 逻辑互斥
   <> | 比较运算符 		| 与 != 等价
   \>  | 比较运算符 		| 比较字符串或者数字是否大于另外一个值
   \>= | 比较运算符 		| 比较字符串或者数字是否大于或者等于另外一个值
   <  | 比较运算符 		| 比较字符串或者数字是否小于另外一个值
   <= | 比较运算符 		| 比较字符串或者数字是否小于或者等于另外一个值
   ~  | 正则或者字符串匹配 | 以/开头且以/结尾的字符串表示正则（比较值的正则是否配置变量）；否则会判断目标字符串是否包含当前字符串（内部会转化为小写之后再判断变量是否包含比较值）
   !~ | 正则或者字符串匹配 | 与 ~ 逻辑相反
   ~* | 正则或者字符串匹配 | 匹配任意其中一个，比较值可以是一个字符串或者正则或者是用英文逗号(,)分割的字符串跟正则表达式,以/开头且以/结尾的字符串表示正则（比较值的正则是否配置变量）；否则会判断目标字符串是否包含当前字符串（内部会转化为小写之后再判断变量是否包含比较值）
   !~* | 正则或者字符串匹配 | 与 !~* 逻辑相反 
   between | 区间变量运算符 | 变量值是否在给定的 比较值中间；需要比较值是一个用英文逗号(,)分割的字符串或者数字(var>=value[0] && var<=value[1])
   in | 判断数组中是否存在某个值 | 给定的value需要是一个用英文逗号(,)分割的字符串或者数字或者是单一的字符串；判断变量是否是给定的比较值中的其中一个
   nin | 判断数组中是否不存在某个值| 与 in 逻辑相反（not in）
   any | 数组中的任意一个值匹配即可 | 需要var跟给定的value是用英文逗号(,)分割的字符串或者数字或者是单一的字符串；value中的任意一个元素出现在var中即可
   has | 数组中必须存在某一个值 | 需要var跟给定的value是用英文逗号(,)分割的字符串或者数字或者是单一的字符串；value中的所有元素必须都出现在var元素中
   none | 数组中不能存在某个值 | 需要var跟给定的value是用英文逗号(,)分割的字符串或者数字或者是单一的字符串；value中的所有元素都不能出现在var元素中
   vgt | 比较版本号是否大于 | 版本值为 xx.xx.xx
   vgte | 比较版本号是否大于或者等于 | 版本值为 xx.xx.xx
   vlt | 比较本班好是否小于 | 版本值为 xx.xx.xx
   vlte | 比较版本号是否小于或者等于 | 版本值为 xx.xx.xx
   iir | 判断IP是否在某个IP段中 | in ip range 需要给定的value是用英文逗号(,)分割的字符串或者数字或者是单一的字符串(127.0.0.1/24,192.168.0.1/24)；（判断varIP是否在给定的valueIP段中）
   niir | 判断IP是否不在某个IP段中 | not in ip range 与iir 逻辑相反
   
   tips：操作符、运算符、以及赋值运算符如果不满足需求业务方可以按照自己的要求实现对应的接口然后注册之后即可
   ## TODO 剩余文档后续补全