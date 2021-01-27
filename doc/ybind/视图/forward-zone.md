| 版本 | 日期       | 更改记录                                                     | 作者 |
| :--- | :--------- | :----------------------------------------------------------- | ---- |
| 0.1  | 2020-04-05 | 初稿                                                         | 程俊 |
| 0.2  | 2020-04-08 | 添加配置项的默认值提示<br>`GET`方法删除body部分<br>`POST`冲突返回409<br>`PUT`方法body不能缺省 | 程俊 |
| 0.3  | 2020-04-10 | 增加algo字段 order,weight,srtt算法                           | 韩冬 |
| 0.4  | 2020-04-10 | 修改forwaders ip为字符串类型,weight/order为数字类型          | 韩冬 |
| 0.5  | 2020-04-17 | 修改"mode"字段为"forward","algo"为"forward_algo"             | 韩冬 |
| 0.6  | 2020-04-22 | 规范化返回码，将forward拆分为model、algo                     | 程俊 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [新增](#新增)
* [修改](#修改)
* [删除](#删除)

------------

## 接口概览
| URL                                        | 方法   | 描述          |
| ------------------------------------------ | ------ | ------------- |
| http://ip:port/api/ybind/v1.0/forward-zone | GET    | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/forward-zone | POST   | [新增](#新增) |
| http://ip:port/api/ybind/v1.0/forward-zone | PUT    | [修改](#修改) |
| http://ip:port/api/ybind/v1.0/forward-zone | DELETE | [删除](#删除) |

## 概述
* 语法：
```
zone string [ class ] {
type forward;
forward ( first | only ) algo ( srtt | weight | order);
forwarders [ port integer ] [ dscp integer ] { ( ipv4_address | -
ipv6_address ) [ port integer ] [ dscp integer ] [(weight | order) interger];  ... };
};
```
* 概念：转发域。将符合条件的请求转发出去。
* 支持的配置项：

| 名称                                                         | 默认值 | 描述                                            |
| ------------------------------------------------------------ | ------ | ----------------------------------------------- |
| [forward](http://192.168.15.206:4999/web/#/3?page_id=142 "单独接口") | N/A    | 转发模式：first、only。 算法: weight,order,srtt |
| [forwarders](http://192.168.15.206:4999/web/#/3?page_id=143 "单独接口") | N/A    | 转发的目标ip地址                                |
* 注意项：
	* `type`默认为forward，不需要指定, algo 配置没有则默认为srtt算法
	* `forwarders`里面只有ip地址,wrr算法(weight | order),其他例如port等不支持

## 获取

### URL
http://ip:port/api/ybind/v1.0/forward-zone

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
	    type forward;
	    forward first algo weight;
		forwarders { 1.1.1.1 weight 1; 2.2.2.2 weight 2; 3.3.3.3 weight 3; };
	};
	
	zone "google.com" {
	    type forward;
	    forward only algo order;
		forwarders { 8.8.8.8 order 1; 114.114.114.114 order 2; };
	};
}
```

#### 获取特定策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/forward-zone?name=yamu.com&view=__default
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "forward": {
            "mode": "only",
            "algo": "srtt"
        },
        "forwarders": [
            {
                "ip": "1.1.1.1"
            }
        ],
        "type": "forward"
    }
}
```

#### 获取全量策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/forward-zone?view=__default
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "google.com": {
            "forward": {
                "mode": "first",
                "algo": "srtt"
            },
            "forwarders": [
                {
                    "ip": "2.2.2.2"
                }
            ],
            "type": "forward"
        },
        "yamu.com": {
            "forward": {
                "mode": "only",
                "algo": "srtt"
            },
            "forwarders": [
                {
                    "ip": "1.1.1.1"
                }
            ],
            "type": "forward"
        }
    }
}
```

## 新增

### URL
http://ip:port/api/ybind/v1.0/forward-zone

### 方法
`POST`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| name* | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-<br>**举例**：baidu.com |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**举例**：__default |

* body：

  **说明**：`name`下的策略字典

| 名称        | 类型   | 默认值 | 描述                                    |
| :---------- | :----- | :----- | :-------------------------------------- |
| type*       | String | N/A    | **说明**：域类型<br>**取值**：`forward` |
| forward*    | Dict   | N/A    | **说明**：mode和algo的字典              |
| forwarders* | Array  | N/A    | **说明**：地址列表                      |

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
	    type forward;
	    forward first algo weight;
		forwarders { 1.1.1.1 weight 1; 2.2.2.2 weight 2; 3.3.3.3 weight 3; };
	};
	
	zone "google.com" {
	    type forward;
	    forward only algo order;
		forwarders { 8.8.8.8 order 1; 114.114.114.114 order 2; };
	};
}
```

#### 增加特定策略
* 请求：
```
METHOD : POST
URL    : http://ip:port/api/ybind/v1.0/forward-zone?name=baidu.com&view=__default
BODY   : {
	"type": "forward",
	"forward": {
		"mode": "only",
		"algo": "srtt"
	},
	"forwarders": [
		{
			"ip": "1.1.1.1"
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
	zone "yamu.com" {
	    type forward;
	    forward first algo weight;
		forwarders { 1.1.1.1 weight 1; 2.2.2.2 weight 2; 3.3.3.3 weight 3; };
	};
	
	zone "google.com" {
	    type forward;
	    forward only algo order;
		forwarders { 8.8.8.8 order 1; 114.114.114.114 order 2; };
	};
	zone "baidu.com" {
	    type forward;
	    forward only algo srtt;
		forwarders { 1.1.1.1; };
	}
}
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/forward-zone

### 方法
`PUT`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| name  | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示所有的zone<br>**举例**：yamu.com |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**举例**：default |

* body：

  **说明**：根据传入的数据更新指定`name`的配置或者覆盖所有配置

| 名称        | 类型   | 默认值 | 描述                                    |
| :---------- | :----- | :----- | :-------------------------------------- |
| type*       | String | N/A    | **说明**：域类型<br>**取值**：`forward` |
| forward*    | Dict   | N/A    | **说明**：mode和algo的字典              |
| forwarders* | Array  | N/A    | **说明**：地址列表                      |

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
	    type forward;
	    forward first algo weight;
		forwarders { 1.1.1.1 weight 1; 2.2.2.2 weight 2; 3.3.3.3 weight 3; };
	};
	
	zone "google.com" {
	    type forward;
	    forward only algo order;
		forwarders { 8.8.8.8 order 1; 114.114.114.114 order 2; };
	};
}
```

#### 修改特定策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/forward-zone?name=google.com&view=__default
BODY   : {
		"forward": {
            "mode": "first",
            "algo": "weight"
        },
        "forwarders": [
            {
                "ip": "223.5.5.5",
				"weight": 2
            },
			{
                "ip": "223.6.6.6",
				"weight": 3
            }
        ],
        "type": "forward"
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
	    type forward;
	    forward first algo weight;
		forwarders { 1.1.1.1 weight 1; 2.2.2.2 weight 2; 3.3.3.3 weight 3; };
	};
	
	zone "google.com" {
	    type forward;
	    forward first algo weight;
		forwarders { 223.5.5.5 weight 2; 223.6.6.6 weight 3; };
	};
}
```

#### 更新全部策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/forward-zone?view=_default
BODY   : {
        "google.com": {
            "forward": {
                "mode": "first",
                "algo": "srtt"
            },
            "forwarders": [
                {
                    "ip": "2.2.2.2"
                }
            ],
            "type": "forward"
        },
        "yamu.com": {
            "forward": {
                "mode": "only",
                "algo": "srtt"
            },
            "forwarders": [
                {
                    "ip": "1.1.1.1"
                }
            ],
            "type": "forward"
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
	zone "yamu.com" {
	    type forward;
	    forward only algo srtt;
		forwarders { 1.1.1.1;};
	};
	
	zone "google.com" {
	    type forward;
	    forward first algo srtt;
		forwarders { 2.2.2.2;};
	};
}
```

## 删除

### URL
http://ip:port/api/ybind/v1.0/forward-zone

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
	    type forward;
	    forward first algo weight;
		forwarders { 1.1.1.1 weight 1; 2.2.2.2 weight 2; 3.3.3.3 weight 3; };
	};
	
	zone "google.com" {
	    type forward;
	    forward only algo order;
		forwarders { 8.8.8.8 order 1; 114.114.114.114 order 2; };
	};
}
```

#### 删除特定策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/forward-zone?name=yamu.com&view=__default
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
	    type forward;
	    forward only algo order;
		forwarders { 8.8.8.8 order 1; 114.114.114.114 order 2; };
	};
}
```

#### 删除全部策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/forward-zone?view=__default
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