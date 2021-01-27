| 版本 | 日期       | 更改记录                                                     | 作者 |
| :--- | :--------- | :----------------------------------------------------------- | ---- |
| 0.1  | 2020-04-04 | 初稿                                                         | 程俊 |
| 0.2  | 2020-04-04 | 废除`POST`、`DELETE`方法，使用`PUT`代替                      | 程俊 |
| 0.3  | 2020-04-07 | rr中的rdata使用字符串代替<br>添加`recursion`、`zone`的链接   | 程俊 |
| 0.4  | 2020-04-08 | 添加配置项的默认值提示<br>`GET`方法删除body部分<br>`POST`冲突返回409<br>`PUT`方法body不能缺省 | 程俊 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [新增](#新增)
* [修改](#修改)
* [删除](#删除)

------------

## 接口概览
| URL                                | 方法   | 描述          |
| ---------------------------------- | ------ | ------------- |
| http://ip:port/api/ybind/v1.0/view | GET    | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/view | POST   | [新增](#新增) |
| http://ip:port/api/ybind/v1.0/view | PUT    | [修改](#修改) |
| http://ip:port/api/ybind/v1.0/view | DELETE | [删除](#删除) |

## 概述
* 语法：
```
view view_name [ class ] {
	match-clients { address_match_list } ;
	match-destinations { address_match_list } ;
	match-recursive-only yes_or_no ;
	[ view_option ; ... ]
	[ zone_statement ; ... ]
} ;
```
* 概念：视图语句是BIND 9的一个强大特性，它允许服务器响应DNS根据询问者的不同进行不同的查询。它对于实现split特别有用，无需运行多个服务器的DNS设置。
* view中目前支持的项：

| 名称                                                         | 默认值   | 描述                   |
| ------------------------------------------------------------ | -------- | ---------------------- |
| [match-clients](view-match.md)                               | { any; } | 请求源的匹配           |
| [match-destinations](view-match.md)                          | { any; } | 请求目的的匹配         |
| [match-recursive-only](view-match.md)                        | no       | 是否是递归请求的匹配   |
| [match-domains](view-match.md)                               | { "."; } | 请求域名匹配           |
| [match-ecs-first](ecs.md)                                    | no       | 是否是edns请求的匹配   |
| [recursion](recursion.md)                                    | yes      | 递归开关               |
| `zone`([auth-zone](auth-zone.md)、[hint-zone](hint-zone.md)、[static-stub-zone](static-stub-zone.md)、[forward-zone](forward-zone.md)) | N/A      | 区域                   |
| [allow-query](allow-query.md)                                | { any; } | 允许访问白名单         |
| [allow-query-on](allow-query-on.md)                          | { any; } | 允许访问目的地址白名单 |

* 注意项：
	* 没有开放`class`的配置，默认所有的都为IN
	* 初始化后的系统会有一个默认的视图，视图名为**_default**，内部没有任何配置项，匹配任意的客户端
	* option语句中给出的许多选项也可以在视图语句中使用，
然后仅在使用该视图解决查询时应用。当没有特定于视图的值时
给定，options语句中的值用作默认值。此外，区域选项可以有
视图语句中指定的默认值;这些特定于视图的默认值优先
于option语句中的值。
	* 视图是特定于类的。如果没有给定类，则假定类为IN。
	* 所有非IN视图必须包含hint区域，因为只有IN类编译了默认hint。

## 获取

### URL
http://ip:port/api/ybind/v1.0/view

### 方法
`GET`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| name | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示所有的view<br>**举例**：view-shanghai<br>**注意**： **default **表示默认视图 |

* returnBody：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| rcode*       | Int    | N/A    | 业务执行码                                                   |
| description* | String | N/A    | `rcode`的文字描述                                            |
| data         | Dict   | N/A    | **缺省**：业务执行失败<br>**Dict**：`name`缺省时表示所有的视图信息，这时是一个视图字典；否则，表示`name`指定的视图信息，是一个配置项字典 |

### 返回码
| rcode | description           | 说明                         |
| ----- | --------------------- | ---------------------------- |
| 0     | Success               | 查询成功                     |
| 1     | Bad Parameter Format  | `name`格式错误               |
| 404   | Not Found             | 没有找到`name`指定的view配置 |
| 408   | Request Timeout       | 请求超时                     |
| 500   | Internal Server Error | 程序运行错误                 |

### 示例
* 现有策略：

```
view newyorkv {
    match-clients { newyork; };
    recursion yes;
    zone "ecs.com" IN {
        type master;
        file "ecs.com.newyork.zone";
        allow-update { none; };
        allow-transfer { none; };
    };
};
view _default {
    recursion yes;
};
```

#### 获取特定策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/view?name=newyorkv
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "match-clients":[
            "newyork"
        ],
        "recursion": true,
        "zone":{
            "ecs.com":{
                "type": "master",
                "allow-update":[
                    "none"
                ],
                "allow-transfer":[
                    "none"
                ],
                "rr":{
                    "@":{
                        "soa":{
                            "ttl": 86400,
                            "rdata": "dns.ecs.com. admin.ecs.com. (1053891164 2M 1M 1W 1D)"
                        },
                        "ns":[
                            {
                                "ttl": 86400,
                                "rdata": "dns1"
                            },{
                                "rdata": "dns2"
                            }
                        ]
                    },
                    "dns1":{
                        "a":[
                            {
                                "ttl": 1090,
                                "rdata": "1.1.1.1"
                            },{
                                "rdata": "2.2.2.2"
                            }
                        ],
                        "aaaa":[
                            {
                                "rdata": "2001::1"
                            }
                        ]
                    },
                    "dns2":{
                        "aaaa":[
                            {
                                "rdata": "2001::2"
                            }
                        ]
                    }
                }
            }
        }
    }
}
```

#### 获取全量策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/view
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "newyorkv": {
            "match-clients":[
                "newyork"
            ],
            "recursion": true,
            "zone":{
                "ecs.com":{
                    "type": "master",
                    "allow-update":[
                        "none"
                    ],
                    "allow-transfer":[
                        "none"
                    ],
                    "rr":{
                        "@":{
                            "soa":{
                                "ttl": 86400,
                                "rdata": "dns.ecs.com. admin.ecs.com. (1053891164 2M 1M 1W 1D)"
                            },
                            "ns":[
                                {
                                    "ttl": 86400,
                                    "rdata": "dns1"
                                },{
                                    "rdata": "dns2"
                                }
                            ]
                        },
                        "dns1":{
                            "a":[
                                {
                                    "ttl": 1090,
                                    "rdata": "1.1.1.1"
                                },{
                                    "rdata": "2.2.2.2"
                                }
                            ],
                            "aaaa":[
                                {
                                    "rdata": "2001::1"
                                }
                            ]
                        },
                        "dns2":{
                            "aaaa":[
                                {
                                    "rdata": "2001::2"
                                }
                            ]
                        }
                    }
                }
            }
        },
        "_default": {
            "recursion": true
        }
    }
}
```

## 新增

### URL
http://ip:port/api/ybind/v1.0/view

### 方法
`POST`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| name* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-<br>**举例**：view-shanghai<br>**注意**：名称不能包含__ |

* body：

| 名称 | 类型 | 默认值 | 描述                     |
| :--- | :--- | :----- | :----------------------- |
| N/A* | Dict | N/A    | **说明**：`name`下的策略 |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description             | 说明                   |
| ----- | ----------------------- | ---------------------- |
| 0     | Success                 | 新增成功               |
| 1     | Bad Parameter Format    | `name`或`body`格式错误 |
| 4     | Miss Required Parameter | 缺少必选参数`name`     |
| 408   | Request Timeout         | 请求超时               |
| 409   | Conflict                | `name`策略已存在       |
| 500   | Internal Server Error   | 程序运行错误           |

### 示例
* 现有策略：

```
view newyorkv {
    match-clients { newyork; };
    recursion yes;
    zone "ecs.com" IN {
        type master;
        file "ecs.com.newyork.zone";
        allow-update { none; };
        allow-transfer { none; };
    };
};
view _default {
    recursion yes;
};
```

#### 增加特定策略
* 请求：
```
METHOD : POST
URL    : http://ip:port/api/ybind/v1.0/view?name=view-shanghai
BODY   : {
	"match-clients":[
		"6.6.6.6"
	]
}
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success"
}
```
* 策略：

```
view newyorkv {
    match-clients { newyork; };
    recursion yes;
    zone "ecs.com" IN {
        type master;
        file "ecs.com.newyork.zone";
        allow-update { none; };
        allow-transfer { none; };
    };
};
view _default {
    recursion yes;
};
view view-shanghai {
    match-clients { "6.6.6.6"; };
};
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/view

### 方法
`PUT`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| name | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示所有的view<br>**举例**：view-shanghai |

* body：

| 名称 | 类型 | 默认值 | 描述                                                         |
| :--- | :--- | :----- | :----------------------------------------------------------- |
| N/A* | Dict | N/A    | **说明**：根据传入的类型更新指定`name`的配置或者覆盖所有配置<br>**注意**：可以为空：{}，删除指定`name`的配置或者删除所有配置；删除所有策略时，不会删除**_default**视图，但是会删除其下面的所有配置项 |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description           | 说明                   |
| ----- | --------------------- | ---------------------- |
| 0     | Success               | 修改成功               |
| 1     | Bad Parameter Format  | `name`或`body`格式错误 |
| 408   | Request Timeout       | 请求超时               |
| 500   | Internal Server Error | 程序运行错误           |

### 示例
* 现有策略：

```
view newyorkv {
    match-clients { newyork; };
    recursion yes;
    zone "ecs.com" IN {
        type master;
        file "ecs.com.newyork.zone";
        allow-update { none; };
        allow-transfer { none; };
    };
};
view _default {
    recursion yes;
};
```

#### 修改特定策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/view?name=_default
BODY   :{
	"recursion": false,
	"allow-query":[
		"localhost"
	]
}
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success"
}
```
* 策略：

```
view newyorkv {
    match-clients { newyork; };
    recursion yes;
    zone "ecs.com" IN {
        type master;
        file "ecs.com.newyork.zone";
        allow-update { none; };
        allow-transfer { none; };
    };
};
view _default {
    recursion no;
	allow-query { localhost; };
};
```

#### 更新全部策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/view
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success"
}
```
* 策略：

```
view _default {
};
```

## 删除

### URL
http://ip:port/api/ybind/v1.0/view

### 方法
`DELETE`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| name | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示所有的view<br>**举例**：view-shanghai<br>**注意**：删除所有策略时，不会删除**default**视图，但是会删除其下面的所有配置项 |

* body：

| 名称 | 类型 | 默认值 | 描述 |
| :--- | :--- | :----- | :--- |
| N/A  | N/A  | N/A    | N/A  |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description           | 说明         |
| ----- | --------------------- | ------------ |
| 0     | Success               | 删除成功     |
| 408   | Request Timeout       | 请求超时     |
| 500   | Internal Server Error | 程序运行错误 |

### 示例
* 现有策略：

```
view newyorkv {
    match-clients { newyork; };
    recursion yes;
    zone "ecs.com" IN {
        type master;
        file "ecs.com.newyork.zone";
        allow-update { none; };
        allow-transfer { none; };
    };
};
view _default {
    recursion yes;
};
```

#### 删除特定策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/view?name=newyorkv
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success"
}
```
* 策略：

```
view _default {
    recursion yes;
};
```

#### 删除全部策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/view
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success"
}
```
* 策略：

```
view _default {
};
```