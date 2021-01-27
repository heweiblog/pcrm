| 版本 | 日期       | 更改记录                                                     | 作者 |
| :--- | :--------- | :----------------------------------------------------------- | ---- |
| 0.1  | 2020-04-05 | 初稿                                                         | 程俊 |
| 0.2  | 2020-04-08 | 添加配置项的默认值提示<br>`GET`方法删除body部分<br>`POST`冲突返回409<br>`PUT`方法body不能缺省 | 程俊 |
| 0.3  | 2020-04-22 | 规范化返回码                                                 | 程俊 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [新增](#新增)
* [修改](#修改)
* [删除](#删除)

------------

## 接口概览
| URL                                            | 方法   | 描述          |
| ---------------------------------------------- | ------ | ------------- |
| http://ip:port/api/ybind/v1.0/static-stub-zone | GET    | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/static-stub-zone | POST   | [新增](#新增) |
| http://ip:port/api/ybind/v1.0/static-stub-zone | PUT    | [修改](#修改) |
| http://ip:port/api/ybind/v1.0/static-stub-zone | DELETE | [删除](#删除) |

## 概述
* 语法：
```
zone string [ class ] {
allow-query { address_match_element; ... };
allow-query-on { address_match_element; ... };
server-addresses { ( ipv4_address | ipv6_address ); ... };
server-names { string; ... };
};
```
* 概念：存根域。帮助迭代直接指定授权服务器的位置。
* 支持的配置项：

| 名称                                | 默认值   | 描述                   |
| ----------------------------------- | -------- | ---------------------- |
| [allow-query](allow-query.md)       | { any; } | 允许访问白名单         |
| [allow-query-on](allow-query-on.md) | { any; } | 允许访问目的地址白名单 |
| `server-addresses`                  | N/A      | 授权所在的ip地址       |
| `server-names`                      | N/A      | 授权所在的主机名       |

* 注意项：
	* 只支持静态存根`static-stub`，无需指定`type`
	* `server-addresses`直接指定了ip地址
	* `server-names`可以在ip地址不清楚的情况下，指定主机名

## 获取

### URL
http://ip:port/api/ybind/v1.0/static-stub-zone

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
view _default {
	zone "yamu.com" {
	    type static-stub;
	    server-addresses { "1.1.1.1"; };
	};
	
	zone "google.com" {
	    type static-stub;
	    server-names { "dns.google.com"; };
	};
}
```

#### 获取特定策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/static-stub-zone?name=yamu.com&view=__default
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
		"type": "static-stub",
        "server-addresses": [
            "1.1.1.1"
        ]
    }
}
```

#### 获取全量策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/static-stub-zone?view=__default
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "yamu.com": {
			"type": "static-stub",
            "server-addresses": [
                "1.1.1.1"
            ]
        },
        "google.com": {
			"type": "static-stub",
            "server-names": [
                "dns.google.com"
            ]
        }
    }
}
```

## 新增

### URL
http://ip:port/api/ybind/v1.0/static-stub-zone

### 方法
`POST`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| name* | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-<br>**举例**：baidu.com |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**举例**：__default |

* body：

  **说明**：`name`下的策略字典<br>**注意**：字典不能为空，即不能为{}

| 名称             | 类型   | 默认值 | 描述                                        |
| :--------------- | :----- | :----- | :------------------------------------------ |
| type*            | String | N/A    | **说明**：域类型<br>**取值**：`static-stub` |
| server-names     | Array  | N/A    | **说明**：域名列表                          |
| server-addresses | Array  | N/A    | **说明**：地址列表                          |

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
	    type static-stub;
	    server-addresses { "1.1.1.1"; };
	};
	
	zone "google.com" {
	    type static-stub;
	    server-names { "dns.google.com"; };
	};
}
```

#### 增加特定策略
* 请求：
```
METHOD : POST
URL    : http://ip:port/api/ybind/v1.0/static-stub-zone?name=baidu.com&view=__default
BODY   : {
    "type": "static-stub",
    "server-addresses": [
        "2.2.2.2"
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
	    type static-stub;
	    server-addresses { "1.1.1.1"; };
	};
	
	zone "google.com" {
	    type static-stub;
	    server-names { "dns.google.com"; };
	};
	
	zone "baidu.com" {
	    type static-stub;
	    server-addresses { "2.2.2.2"; };
	};
}
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/static-stub-zone

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

| 名称             | 类型   | 默认值 | 描述                                        |
| :--------------- | :----- | :----- | :------------------------------------------ |
| type*            | String | N/A    | **说明**：域类型<br>**取值**：`static-stub` |
| server-names     | Array  | N/A    | **说明**：域名列表                          |
| server-addresses | Array  | N/A    | **说明**：地址列表                          |

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
	    type static-stub;
	    server-addresses { "1.1.1.1"; };
	};
	
	zone "google.com" {
	    type static-stub;
	    server-names { "dns.google.com"; };
	};
}
```

#### 修改特定策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/static-stub-zone?name=google.com&view=__default
BODY   : {
		"type": "static-stub",
        "server-addresses": [
            "8.8.8.8"
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
	    type static-stub;
	    server-addresses { "1.1.1.1"; };
	};
	
	zone "google.com" {
	    type static-stub;
	    server-addresses { "8.8.8.8"; };
	};
}
```

#### 更新全部策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/static-stub-zone?view=__default
BODY   : {
        "baidu.com": {
            "server-addresses": [
                "2.2.2.20"
            ],
            "type": "static-stub"
        },
        "yamu.com": {
            "server-addresses": [
                "2.2.2.201"
            ],
            "type": "static-stub"
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
view _default {
	zone "yamu.com" {
	    type static-stub;
	    server-addresses { "2.2.2.201"; };
	};
	
	zone "baidu.com" {
	    type static-stub;
	    server-addresses { "2.2.2.20"; };
	};
}
```

## 删除

### URL
http://ip:port/api/ybind/v1.0/static-stub-zone

### 方法
`DELETE`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| name  | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示所有的zone<br>**举例**：yamu.com |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**举例**：__default |

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
	    type static-stub;
	    server-addresses { "1.1.1.1"; };
	};
	
	zone "google.com" {
	    type static-stub;
	    server-names { "dns.google.com"; };
	};
}
```

#### 删除特定策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/static-stub-zone?name=yamu.com&view=__default
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
	    type static-stub;
	    server-names { "dns.google.com"; };
	};
}
```

#### 删除全部策略

* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/static-stub-zone?view=__default
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