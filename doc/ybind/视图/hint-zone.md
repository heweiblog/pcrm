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
| URL                                     | 方法   | 描述          |
| --------------------------------------- | ------ | ------------- |
| http://ip:port/api/ybind/v1.0/hint-zone | GET    | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/hint-zone | POST   | [新增](#新增) |
| http://ip:port/api/ybind/v1.0/hint-zone | PUT    | [修改](#修改) |
| http://ip:port/api/ybind/v1.0/hint-zone | DELETE | [删除](#删除) |

## 概述
* 语法：
```
zone string [ class ] {
rr rr_dict;
}
```
* 概念：提示域。用来指示迭代从哪里开始。
* 支持的配置项：

| 名称 | 默认值 | 描述           |
| ---- | ------ | -------------- |
| `rr` | N/A    | 存放记录的字典 |

* 注意项：
	* `type`已经默认为hint，所以无需指定`type`
	* `rr`中只有ns、a、aaaa记录
	* `rr`中的记录无需指定ttl

## 获取

### URL
http://ip:port/api/ybind/v1.0/hint-zone

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
	zone "." {
	    type hint;
		file "___default_root.zone";
	};
	zone "yamu.com" {
	    type hint;
		file "___default_yamu.com.zone";
	};
}
```

#### 获取特定策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/hint-zone?name=.&view=__default
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
		"type": "hint",
        "rr": {
			"@":{
				"NS":[
					{
						"ttl": 86400,
						"rdata": "a.root-servers.net."
					}
				]
			},
			"a.root-servers.net.":{
				"A":[
					{
						"ttl": 1090,
						"rdata": "198.41.0.4"
					}
				],
				"AAAA":[
					{
						"rdata": "2001:503:ba3e::2:30"
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
URL    : http://ip:port/api/ybind/v1.0/hint-zone?view=__default
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "yamu.com": {
			"type": "hint",
			"rr": {
            	"@":{
            	    "NS":[
            	        {
            	            "ttl": 86400,
            	            "rdata": "dns1"
            	        },{
            	            "rdata": "dns2"
            	        }
            	    ]
            	},
            	"dns1":{
            	    "A":[
            	        {
            	            "ttl": 1090,
            	            "rdata": "1.1.1.1"
            	        },{
            	            "rdata": "2.2.2.2"
            	        }
            	    ],
            	    "AAAA":[
            	        {
            	            "rdata": "2001::1"
            	        }
            	    ]
            	},
            	"dns2":{
            	    "AAAA":[
            	        {
            	            "rdata": "2001::2"
            	        }
            	    ]
            	}
			}
        },
        ".":{
			"type": "hint",
			"rr": {
				"@":{
					"NS":[
						{
							"ttl": 86400,
							"rdata": "a.root-servers.net."
						}
					]
				},
				"a.root-servers.net.":{
					"A":[
						{
							"ttl": 1090,
							"rdata": "198.41.0.4"
						}
					],
					"AAAA":[
						{
							"rdata": "2001:503:ba3e::2:30"
						}
					]
				}
			}
        }
    }
}
```

## 新增

### URL
http://ip:port/api/ybind/v1.0/hint-zone

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

| 名称  | 类型   | 默认值 | 描述                                                    |
| :---- | :----- | :----- | :------------------------------------------------------ |
| type* | String | N/A    | **说明**：域类型<br>**取值**：`hint`                    |
| rr*   | Dict   | N/A    | **说明**：记录<br>**注意**：只能有NS、A、AAAA类型的记录 |

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
	zone "." {
	    type hint;
		file "___default_root.zone";
	};
	zone "yamu.com" {
	    type hint;
		file "___default_yamu.com.zone";
	};
}
```

#### 增加特定策略
* 请求：
```
METHOD : POST
URL    : http://ip:port/api/ybind/v1.0/hint-zone?name=baidu.com&view=__default
BODY   : {
	"type": "hint",
	"rr": {
    	"@":{
    	    "NS":[
    	        {
    	            "ttl": 86400,
    	            "rdata": "a.root-servers.net."
    	        }
    	    ]
    	},
    	"a.root-servers.net.":{
    	    "A":[
    	        {
    	            "ttl": 1090,
    	            "rdata": "198.41.0.4"
    	        }
    	    ],
    	    "AAAA":[
    	        {
    	            "rdata": "2001:503:ba3e::2:30"
    	        }
    	    ]
    	}
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
	zone "." {
	    type hint;
		file "___default_root.zone";
	};
	zone "yamu.com" {
	    type hint;
		file "___default_yamu.com.zone";
	};
	zone "baidu.com" {
	    type hint;
		file "___default_baidu.com.zone";
	};
}
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/hint-zone

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

| 名称  | 类型   | 默认值 | 描述                                                    |
| :---- | :----- | :----- | :------------------------------------------------------ |
| type* | String | N/A    | **说明**：域类型<br>**取值**：`hint`                    |
| rr*   | Dict   | N/A    | **说明**：记录<br>**注意**：只能有NS、A、AAAA类型的记录 |

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
	zone "." {
	    type hint;
		file "___default_root.zone";
	};
	zone "yamu.com" {
	    type hint;
		file "___default_yamu.com.zone";
	};
}
```

#### 修改特定策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/hint-zone?name=yamu.com&view=__default
BODY   : {
	"type": "hint",
	"rr": {
    	"@":{
    	    "NS":[
    	        {
    	            "ttl": 86400,
    	            "rdata": "a.root-servers.net."
    	        }
    	    ]
    	},
    	"a.root-servers.net.":{
    	    "A":[
    	        {
    	            "ttl": 1090,
    	            "rdata": "198.41.0.4"
    	        }
    	    ],
    	    "AAAA":[
    	        {
    	            "rdata": "2001:503:ba3e::2:30"
    	        }
    	    ]
    	}
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
	zone "." {
	    type hint;
		file "___default_root.zone";
	};
	zone "yamu.com" {
	    type hint;
		file "___default_yamu.com.zone";
	};
}
```

#### 更新全部策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/hint-zone?view=__default
BODY   : {
        "yamu.com": {
            "rr": {
                "@": {
                    "NS": [
                        {
                            "ttl": 86400,
                            "rdata": "a.root-servers.net."
                        }
                    ]
                },
                "a.root-servers.net.": {
                    "A": [
                        {
                            "ttl": 1090,
                            "rdata": "198.41.0.4"
                        }
                    ],
                    "AAAA": [
                        {
                            "rdata": "2001:503:ba3e::2:30"
                        }
                    ]
                }
            },
            "type": "hint"
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
	    type hint;
		file "___default_yamu.com.zone";
	};
}
```

## 删除

### URL
http://ip:port/api/ybind/v1.0/hint-zone

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
	zone "." {
	    type hint;
		file "___default_root.zone";
	};
	zone "yamu.com" {
	    type hint;
		file "___default_yamu.com.zone";
	};
}
```

#### 删除特定策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/hint-zone?name=yamu.com&view=__default
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
	zone "yamu.com" {
	    type hint;
		file "___default_yamu.com.zone";
	};
}
```

#### 删除全部策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/hint-zone?view=__default
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