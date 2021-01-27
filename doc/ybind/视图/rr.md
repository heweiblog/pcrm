| 版本 | 日期       | 更改记录                                                     | 作者 |
| :--- | :--------- | :----------------------------------------------------------- | ---- |
| 0.1  | 2020-04-07 | 初稿                                                         | 程俊 |
| 0.2  | 2020-04-08 | 添加配置项的默认值提示<br>`GET`方法删除body部分<br>`POST`冲突返回409<br>`PUT`方法body不能缺省 | 程俊 |
| 0.3  | 2020-04-22 | 规范化返回码                                                 | 程俊 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [修改](#修改)

------------

## 接口概览
| URL                              | 方法 | 描述          |
| -------------------------------- | ---- | ------------- |
| http://ip:port/api/ybind/v1.0/rr | GET  | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/rr | PUT  | [修改](#修改) |

## 概述
* 概念：对master域记录的增删改查。对slave域的获取。
* 注意项：
	* slave域只提供获取操作
	* soa记录不能删除

## 获取

### URL
http://ip:port/api/ybind/v1.0/rr

### 方法
`GET`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**举例**：__default |
| zone* | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**举例**：yamu.com |
| dname | String | N/A    | **说明**：domain的名称，用于定位到该条domain<br>**格式**：数字、大小写字母、-<br>**缺省**：该`view`下该`zone`的所有domain<br>**举例**：www.yamu.com. |
| type  | String | N/A    | **说明**：type的名称，用于定位到该条type<br>**格式**：a、aaaa、soa、ns、mx...<br>**缺省**：该`view`下该`zone`下该`dname`的所有domain<br>**举例**：www.yamu.com. |

* returnBody：

| 名称         | 类型       | 默认值 | 描述                                                         |
| :----------- | :--------- | :----- | :----------------------------------------------------------- |
| rcode*       | Int        | N/A    | 业务执行码                                                   |
| description* | String     | N/A    | `rcode`的文字描述                                            |
| data         | Array/Dict | N/A    | **缺省**：业务执行失败<br>**Array**：指定`view`下`zone`下`dname`下`type`的策略<br>**Dict**：指定`view`下`zone`下的策略或者指定`view`下`zone`下`dname`下的策略或者指定`view`下`zone`下`type`的策略 |

### 返回码
| rcode | description             | 说明                                                         |
| ----- | ----------------------- | ------------------------------------------------------------ |
| 0     | Success                 | 查询成功                                                     |
| 2     | Bad Parameter Value     | `name`值错误(比如"name="传来一个空值)                        |
| 4     | Miss Required Parameter | 缺少必选参数`name`或`body`                                   |
| 404   | Not Found               | 没有找到`view`指定的配置或者`zone`或者`dname`或者`type`的配置 |
| 408   | Request Timeout         | 请求超时                                                     |
| 500   | Internal Server Error   | 程序运行错误                                                 |

### 示例
* 现有策略：

```
view __default {
	zone yamu.com {
		type master;
		file "___default_yamu.com.zone";
	};
}
```
___default_yamu.com.zone:
```
@           IN  3600    SOA     dns.yamu.com.   admin.yamu.com. (1053891164 2M 1M 1W 1D)
www         IN  3001    A       10.10.10.101
www         IN  3002    A       10.10.10.102
www         IN  3003    AAAA    2001::1
mx          IN  3000    MX      10 mx.baidu.com.
dns1        IN  4000    AAAA    2001::1
dns2        IN  4000    AAAA    2001::2
```

#### 获取特定view特定zone特定dname特定type的策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/rr?view=__default&zone=yamu.com&dname=mx&type=mx
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": [
        {
            "ttl": 3000,
            "rdata": "10 mx.baidu.com."
        }
    ]
}
```

#### 获取特定view特定zone特定dname的策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/rr?view=__default&zone=yamu.com&dname=www
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "A": [
            {
                "ttl": 3001,
                "rdata": "10.10.10.101"
            },
            {
                "ttl": 3002,
                "rdata": "10.10.10.102"
            }
        ],
        "AAAA": [
            {
                "ttl": 3003,
                "rdata": "2001::1"
            }
        ]
    }
}
```

#### 获取特定view特定zone特定type的策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/rr?view=__default&zone=yamu.com&type=aaaa
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "dns1": {
            "AAAA": [
                {
					"ttl": 4000,
                    "rdata": "2001::1"
                }
            ]
        },
        "dns2": {
            "AAAA": [
                {
					"ttl": 4000,
                    "rdata": "2001::2"
                }
            ]
        },
		"www": {
            "AAAA": [
                {
					"ttl": 3003,
                    "rdata": "2001::1"
                }
            ]
        }
    }
}
```

#### 获取特定view特定zone的策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/rr?view=__default&zone=yamu.com
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "@": {
            "SOA": [
                "ttl": 3600,
                "rdata": "dns.yamu.com. admin.yamu.com. (1053891164 2M 1M 1W 1D)"
            ]
        },
        "www": {
            "A": [
                {
                    "ttl": 3001,
                    "rdata": "10.10.10.101"
                },
                {
                    "ttl": 3002,
                    "rdata": "10.10.10.102"
                }
            ],
            "AAAA": [
                {
                    "ttl": 3003,
                    "rdata": "2001::1"
                }
            ]
        },
        "mx": {
            "MX": [
                {
                    "ttl": 3000,
                    "rdata": "10 mx.baidu.com."
                }
            ]
        },
		"dns1": {
            "AAAA": [
                {
                    "ttl": 4000,
                    "rdata": "2001::1"
                }
            ]
        },
		"dns2": {
            "AAAA": [
                {
                    "ttl": 4000,
                    "rdata": "2001::2"
                }
            ]
        }
    }
}
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/rr

### 方法
`PUT`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**举例**：__default |
| zone* | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**举例**：yamu.com |
| dname | String | N/A    | **说明**：domain的名称，用于定位到该条domain<br>**格式**：数字、大小写字母、-<br>**缺省**：该`view`下该`zone`的所有domain<br>**举例**：www.yamu.com. |
| type  | String | N/A    | **说明**：type的名称，用于定位到该条type<br>**格式**：a、aaaa、soa、ns、mx...<br>**缺省**：该`view`下该`zone`下该`dname`的所有domain<br>**注意**：指定type时必须指定`dname`<br>**举例**：www.yamu.com. |

* body：

| 名称 | 类型       | 默认值 | 描述                                                         |
| :--- | :--------- | :----- | :----------------------------------------------------------- |
| N/A* | Array/Dict | N/A    | **说明**：更新指定`view`下`zone`或者指定`view`下`zone`下`dname`或者指定`view`下`zone`下`dname`下`type`的配置<br>**注意**：可以为空：[]/{}，删除指定`view`下`zone`或者指定`view`下`zone`下`dname`或者指定`view`下`zone`下`dname`下`type`的配置 |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description             | 说明                                            |
| ----- | ----------------------- | ----------------------------------------------- |
| 0     | Success                 | 修改成功                                        |
| 1     | Bad Parameter Format    | `view`或`body`或`zone`或`dname`或`type`格式错误 |
| 2     | Bad Parameter Value     | `name`或`body`值错误                            |
| 4     | Miss Required Parameter | 缺少必选参数`name`或`body`                      |
| 408   | Request Timeout         | 请求超时                                        |
| 500   | Internal Server Error   | 程序运行错误                                    |

### 示例
* 现有策略：

```
view __default {
	zone yamu.com {
		type master;
		file "___default_yamu.com.zone";
	};
}
```
___default_yamu.com.zone:
```
@           IN  3600    SOA     dns.yamu.com.   admin.yamu.com. (1053891164 2M 1M 1W 1D)
www         IN  3001    A       10.10.10.101
www         IN  3002    A       10.10.10.102
www         IN  3003    AAAA    2001::1
mx          IN  3000    MX      10 mx.baidu.com.
```

#### 修改特定view特定zone特定dname特定type的策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/rr?view=__default&zone=yamu.com&dname=mx&type=mx
BODY   : [
	{
		"ttl": 3000,
		"rdata": "20 mx.baidu.com."
	}
]
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
view __default {
	zone yamu.com {
		type master;
		file "___default_yamu.com.zone";
	};
}
```
___default_yamu.com.zone:
```
@           IN  3600    SOA     dns.yamu.com.   admin.yamu.com. (1053891164 2M 1M 1W 1D)
www         IN  3001    A       10.10.10.101
www         IN  3002    A       10.10.10.102
www         IN  3003    AAAA    2001::1
mx          IN  3000    MX      20 mx.baidu.com.
```

#### 修改特定view特定zone特定dname的策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/rr?view=__default&zone=yamu.com&dname=www
BODY   : {
    "AAAA": [
        {
            "ttl": 3003,
            "rdata": "2001::1"
        }
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
view __default {
	zone yamu.com {
		type master;
		file "___default_yamu.com.zone";
	};
}
```
___default_yamu.com.zone:
```
@           IN  3600    SOA     dns.yamu.com.   admin.yamu.com. (1053891164 2M 1M 1W 1D)
www         IN  3003    AAAA    2001::1
mx          IN  3000    MX      20 mx.baidu.com.
```

#### 删除特定view特定zone的策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/rr?view=__default&zone=yamu.com
BODY   : {}
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
view __default {
	zone yamu.com {
		type master;
		file "___default_yamu.com.zone";
	};
}
```
___default_yamu.com.zone:
```
@           IN  3600    SOA     dns.yamu.com.   admin.yamu.com. (1053891164 2M 1M 1W 1D)
```