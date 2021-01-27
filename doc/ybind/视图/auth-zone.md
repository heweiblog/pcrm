| 版本 | 日期       | 更改记录                                                     | 作者 |
| :--- | :--------- | :----------------------------------------------------------- | ---- |
| 0.1  | 2020-04-05 | 初稿                                                         | 程俊 |
| 0.2  | 2020-04-07 | 增加`allow-transfer`、`allow-update`、<br>`allow-notify`、`masters`、`rr`的链接 | 程俊 |
| 0.3  | 2020-04-08 | 添加配置项的默认值提示<br>`GET`方法删除body部分<br>`POST`冲突返回409<br>`PUT`方法body不能缺省 | 程俊 |
| 0.4  | 2020-04-22 | 规范化返回码                                                 | 程俊 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [新增](#新增)
* [修改](#修改)
* [删除](#删除)

------------

## 接口概览
| URL                                     | 方法   | 描述          |
| --------------------------------------- | ------ | ------------- |
| http://ip:port/api/ybind/v1.0/auth-zone | GET    | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/auth-zone | POST   | [新增](#新增) |
| http://ip:port/api/ybind/v1.0/auth-zone | PUT    | [修改](#修改) |
| http://ip:port/api/ybind/v1.0/auth-zone | DELETE | [删除](#删除) |

## 概述
* 语法：
```
zone string [ class ] {
type ( master | primary );
allow-query { address_match_element; ... };
allow-query-on { address_match_element; ... };
allow-transfer { address_match_element; ... };
allow-update { address_match_element; ... };
rr rr_dict;
}
zone string [ class ] {
type ( slave | secondary );
allow-notify { address_match_element; ... };
allow-query { address_match_element; ... };
allow-query-on { address_match_element; ... };
allow-transfer { address_match_element; ... };
masters [ port integer ] [ dscp integer ] { ( masters | ipv4_address [ -
port integer ] | ipv6_address [ port integer ] ) [ key string ]; ... -
};
};
```
* 概念：授权域。能够提供对属于此域的权威应答或者委派。
* 支持的配置项：

| 名称                                | 默认值    | 描述                                                         |
| ----------------------------------- | --------- | ------------------------------------------------------------ |
| `type`                              | N/A       | 域类型，这里支持：master、slave                              |
| [allow-query](allow-query.md)       | { any; }  | 允许访问白名单                                               |
| [allow-query-on](allow-query-on.md) | { any; }  | 允许访问目的地址白名单                                       |
| [allow-transfer](allow-transfer.md) | { any; }  | 允许通过axfr/ixfr传输源地址白名单                            |
| [allow-update](allow-update.md)     | { none; } | 允许动态更新白名单                                           |
| [allow-notify](allow-notify.md)     | { none; } | 允许接受notify的白名单，默认是允许masters中配置的主服务器，这里只是对其扩展 |
| [masters](masters.md)               | N/A       | 指定的master服务器地址                                       |
| [rr](rr.md)                         | N/A       | 存放记录的字典                                               |

* 注意项：
	* master域下面有记录项`rr`，而slave没有
	* master域下面有`allow-update`，而slave没有
	* slave域下面有`allow-notify`，而master没有
	* slave域下面有`masters`，而master没有

## 获取

### URL
http://ip:port/api/ybind/v1.0/auth-zone

### 方法
`GET`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| name  | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示所有的zone<br>**举例**：yamu.com |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**举例**：__default |

* returnBody：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| rcode*       | Int    | N/A    | 业务执行码                                                   |
| description* | String | N/A    | `rcode`的文字描述                                            |
| data         | Dict   | N/A    | **缺省**：业务执行失败<br>**Dict**：`name`缺省时所有名称下的配置或者所有的策略 |

### 返回码
| rcode | description             | 说明                         |
| ----- | ----------------------- | ---------------------------- |
| 0     | Success                 | 查询成功                     |
| 2     | Bad Parameter Value     | `name`或`body`值错误         |
| 4     | Miss Required Parameter | 缺少必选参数`name`或`body`   |
| 404   | Not Found               | 没有找到`name`指定的zone配置 |
| 408   | Request Timeout         | 请求超时                     |
| 500   | Internal Server Error   | 程序运行错误                 |

### 示例
* 现有策略：

```
view __default {
	zone "yamu.com" {
	    type master;
	    allow-query { "1.1.1.1"; };
		file "___default_yamu.com.zone";
	};
	
	zone "google.com" {
	    type slave;
		masters { "8.8.8.8"; };
	};
}
```

#### 获取特定策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/auth-zone?name=yamu.com&view=__default
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "type": "master",
        "allow-query":[
            "1.1.1.1"
        ],
        "rr":{
                "@": {
                    "NS": [
                        {
                            "ttl": 86400,
                            "rdata": "dns1"
                        },
                        {
                            "rdata": "dns2"
                        }
                    ],
                    "SOA": [
                        {
                            "ttl": 86400,
                            "rdata": "dns sa 2012137261 300 300 2592000 7200"
                        }
                    ]
                },
                "dns1": {
                    "A": [
                        {
                            "ttl": 1090,
                            "rdata": "1.1.1.1"
                        },
                        {
                            "rdata": "2.2.2.2"
                        }
                    ],
                    "AAAA": [
                        {
                            "rdata": "2001::1"
                        }
                    ]
                },
                "dns2": {
                    "AAAA": [
                        {
                            "rdata": "2001::2"
                        }
                    ]
                }
            }
    }
}
```

#### 获取全量策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/auth-zone?view=__default
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "yamu.com": {
            "type": "master",
            "allow-query":[
                "1.1.1.1"
            ],
            "rr":{
                "@": {
                    "NS": [
                        {
                            "ttl": 86400,
                            "rdata": "dns1"
                        },
                        {
                            "rdata": "dns2"
                        }
                    ],
                    "SOA": [
                        {
                            "ttl": 86400,
                            "rdata": "dns sa 2012137261 300 300 2592000 7200"
                        }
                    ]
                },
                "dns1": {
                    "A": [
                        {
                            "ttl": 1090,
                            "rdata": "1.1.1.1"
                        },
                        {
                            "rdata": "2.2.2.2"
                        }
                    ],
                    "AAAA": [
                        {
                            "rdata": "2001::1"
                        }
                    ]
                },
                "dns2": {
                    "AAAA": [
                        {
                            "rdata": "2001::2"
                        }
                    ]
                }
            }
        },
        "google.com":{
            "type": "slave",
            "masters":[
                "8.8.8.8"
            ]
        }
    }
}
```

## 新增

### URL
http://ip:port/api/ybind/v1.0/auth-zone

### 方法
`POST`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| name* | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-<br>**举例**：baidu.com |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br><br>**举例**：__default |

* body：

  **说明**：`name`下的策略字典<br>**注意**：字典不能为空，即不能为{}

| 名称           | 类型   | 默认值 | 描述                                                         |
| :------------- | :----- | :----- | :----------------------------------------------------------- |
| type*          | String | N/A    | **说明**：域类型<br>**取值**：`master`、`slave`              |
| rr             | Dict   | N/A    | **说明**：记录<br>**注意**：master必须有，slave不能有        |
| allow-update   | Array  | N/A    | **说明**：允许动态更新白名单<br>**注意**：master可以有，slave不能有 |
| allow-notify   | Array  | N/A    | **说明**：允许接受notify的白名单<br>**注意**：master不能有，slave可以有 |
| masters        | Array  | N/A    | **说明**：master服务器地址<br>**注意**：master不能有，slave必须有 |
| allow-transfer | Array  | N/A    | **说明**：允许通过axfr/ixfr传输源地址白名单                  |
| allow-query-on | Array  | N/A    | **说明**：允许访问目的地址白名单                             |
| allow-query    | Array  | N/A    | **说明**：允许访问白名单                                     |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description             | 说明                       |
| ----- | ----------------------- | -------------------------- |
| 0     | Success                 | 新增成功                   |
| 1     | Bad Parameter Format    | `name`或`body`格式错误     |
| 2     | Bad Parameter Value     | `name`或`body`值错误       |
| 4     | Miss Required Parameter | 缺少必选参数`name`或`body` |
| 408   | Request Timeout         | 请求超时                   |
| 409   | Conflict                | `name`策略已存在           |
| 500   | Internal Server Error   | 程序运行错误               |

### 示例
* 现有策略：

```
view __default {
	zone "yamu.com" {
	    type master;
	    allow-query { "1.1.1.1"; };
		file "___default_yamu.com.zone";
	};
	
	zone "google.com" {
	    type slave;
		masters { "8.8.8.8"; };
	};
}
```

#### 增加特定策略
* 请求：
```
METHOD : POST
URL    : http://ip:port/api/ybind/v1.0/auth-zone?name=baidu.com&view=__default
BODY   : {
    "type": "slave",
    "masters":[
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
view __default {
	zone "yamu.com" {
	    type master;
	    allow-query { "1.1.1.1"; };
		file "___default_yamu.com.zone";
	};
	
	zone "google.com" {
	    type slave;
		masters { "8.8.8.8"; };
	};
	
	zone "baidu.com" {
	    type slave;
		masters { "6.6.6.6"; };
	};
}
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/auth-zone

### 方法
`PUT`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| name  | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示所有的zone<br>**举例**：yamu.com |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**举例**：__default |

* body：

  **说明**：根据传入的类型更新指定`name`的配置或者覆盖所有配置<br>**Dict**：指定`name`的策略或者所有的配置

| 名称           | 类型   | 默认值 | 描述                                                         |
| :------------- | :----- | :----- | :----------------------------------------------------------- |
| type*          | String | N/A    | **说明**：域类型<br>**取值**：`master`、`slave`              |
| rr             | Dict   | N/A    | **说明**：记录<br>**注意**：master必须有，slave不能有        |
| allow-update   | Array  | N/A    | **说明**：允许动态更新白名单<br>**注意**：master可以有，slave不能有 |
| allow-notify   | Array  | N/A    | **说明**：允许接受notify的白名单<br>**注意**：master不能有，slave可以有 |
| masters        | Array  | N/A    | **说明**：master服务器地址<br>**注意**：master不能有，slave必须有 |
| allow-transfer | Array  | N/A    | **说明**：允许通过axfr/ixfr传输源地址白名单                  |
| allow-query-on | Array  | N/A    | **说明**：允许访问目的地址白名单                             |
| allow-query    | Array  | N/A    | **说明**：允许访问白名单                                     |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description             | 说明                       |
| ----- | ----------------------- | -------------------------- |
| 0     | Success                 | 修改成功                   |
| 1     | Bad Parameter Format    | `name`或`body`格式错误     |
| 2     | Bad Parameter Value     | `name`或`body`值错误       |
| 4     | Miss Required Parameter | 缺少必选参数`name`或`body` |
| 404   | Not Found               | 该`name`下的策略没有找到   |
| 408   | Request Timeout         | 请求超时                   |
| 500   | Internal Server Error   | 程序运行错误               |

### 示例
* 现有策略：

```
view __default {
	zone "yamu.com" {
	    type master;
	    allow-query { "1.1.1.1"; };
		file "___default_yamu.com.zone";
	};
	
	zone "google.com" {
	    type slave;
		masters { "8.8.8.8"; };
	};
}
```

#### 修改特定策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/auth-zone?name=google.com&view=__default
BODY   : {
    "type": "slave",
    "masters":[
        "7.7.7.7"
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
	zone "yamu.com" {
	    type master;
	    allow-query { "1.1.1.1"; };
		file "___default_yamu.com.zone";
	};
	
	zone "google.com" {
	    type slave;
		masters { "7.7.7.7"; };
	};
}
```

#### 更新全部策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/auth-zone?view=__default
BODY   : {
    "xxx.com": {
        "masters": [
            "6.6.6.0"
        ],
        "type": "slave"
    },
    "yyy.com": {
        "allow-query": [
            "1.1.1.1"
        ],
        "rr": {
            "@": {
                "NS": [
                    {
                        "ttl": 86400,
                        "rdata": "dns1"
                    },
                    {
                        "rdata": "dns2"
                    }
                ],
                "SOA": [
                    {
                        "ttl": 86400,
                        "rdata": "dns sa 2012137261 300 300 2592000 7200"
                    }
                ]
            },
            "dns1": {
                "A": [
                    {
                        "ttl": 1090,
                        "rdata": "1.1.1.1"
                    },
                    {
                        "rdata": "2.2.2.2"
                    }
                ],
                "AAAA": [
                    {
                        "rdata": "2001::1"
                    }
                ]
            },
            "dns2": {
                "AAAA": [
                    {
                        "rdata": "2001::2"
                    }
                ]
            }
        },
        "type": "master"
    }
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
	zone "yyy.com" {
	    type master;
	    allow-query { "1.1.1.1"; };
		file "___default_yamu.com.zone";
	};
	
	zone "xxx.com" {
	    type slave;
		masters { "6.6.6.0"; };
	};
}
```

## 删除

### URL
http://ip:port/api/ybind/v1.0/auth-zone

### 方法
`DELETE`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| name  | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示所有的zone<br>**举例**：yamu.com |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**举例**：default |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description             | 说明                       |
| ----- | ----------------------- | -------------------------- |
| 0     | Success                 | 删除成功                   |
| 2     | Bad Parameter Value     | `name`或`body`值错误       |
| 4     | Miss Required Parameter | 缺少必选参数`name`或`body` |
| 408   | Request Timeout         | 请求超时                   |
| 500   | Internal Server Error   | 程序运行错误               |

### 示例
* 现有策略：

```
view __default {
	zone "yamu.com" {
	    type master;
	    allow-query { "1.1.1.1"; };
		file "___default_yamu.com.zone";
	};
	
	zone "google.com" {
	    type slave;
		masters { "7.7.7.7"; };
	};
}
```

#### 删除特定策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/auth-zone?name=yamu.com&view=__default
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
view __default {
	zone "google.com" {
	    type slave;
		masters { "7.7.7.7"; };
	};
}
```

#### 删除全部策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/auth-zone?view=__default
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
view __default {
}
```