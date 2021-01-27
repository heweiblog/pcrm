| 版本 | 日期       | 更改记录                                                     | 作者 |
| :--- | :--------- | :----------------------------------------------------------- | ---- |
| 0.1  | 2020-04-05 | 初稿                                                         | 叶豪 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [修改](#修改)

------------

## 接口概览
| URL                                       | 方法 | 描述          |
| ----------------------------------------- | ---- | ------------- |
| http://ip:port/api/ybind/v1.0/zone-statistics | GET  | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/zone-statistics | PUT  | [修改](#修改) |

## 概述
* 语法：
```
zone-statistics  ( full | terse | none | <boolean> )
```
* 概念：服务器使用rndc stats命令时附加统计信息的文件的路径名。如果未指定，则指定缺省	值named.stats。服务器当前目录中的统计信息
* 注意项：
	* 可在`option/view/zone`中配置

## 获取

### URL
http://ip:port/api/ybind/v1.0/zone-statistics

### 方法
`GET`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| view  | String | N/A    | **说明**：view的名称，用于定位到该条view                     |
| zone  | String | N/A    | **说明**：zone的名称，用于定位到该条zone                     |

* returnBody：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| rcode*       | Int    | N/A    | 业务执行码                                                   |
| description* | String | N/A    | `rcode`的文字描述                                            |
| data         | Array  | N/A    | **缺省**：业务执行失败<br>**Array**：option中没有statistics-file配置|

### 返回码
| rcode | description           | 说明                                     |
| ----- | --------------------- | ---------------------------------------- |
| 0     | Success               | 查询成功                                 |
| 404   | Not Found             | 没有找到`option\view\zone`指定的配置 |
| 408   | Request Timeout       | 请求超时                                 |
| 500   | Internal Server Error | 程序运行错误                             |

### 示例
* 现有策略：

```
option {
	zone-statistics yes;
}
view __default {
	zone-statistics yes;
	zone "yamu.com" {
	    type master;
	    zone-statistics no;
		file "___default_yamu.com.zone";
	};
}
```

#### 获取特定策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/zone-statistics?view=__default&zone=yamu.com
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "zone-statistics": "no"
    }
}
```

#### 获取option策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/zone-statistics
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "zone-statistics": "yes"
    }
}
```

## 新增

### URL
http://ip:port/api/ybind/v1.0/zone-statistics

### 方法
`PUT`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| view  | String | N/A    | **说明**：view的名称，用于定位到该条view                     |
| zone  | String | N/A    | **说明**：zone的名称，用于定位到该条zone                     |

* body：

| 名称           | 类型   | 默认值 | 描述                                                         |
| :------------- | :----- | :----- | :----------------------------------------------------------- |
| value*          | String| boolean | N/A    可选 full | terse | none | <boolean> | yes | no      |

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
option {
	zone-statistics yes;
}
view __default {
	zone "yamu.com" {
	    type master;
	    zone-statistics no;
		file "___default_yamu.com.zone";
	};
}
```

#### 增加特定策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/zone-statistics?view=__default
BODY   : true
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
option {
	zone-statistics yes;
}
view __default {
	zone-statistics true;
	zone "yamu.com" {
	    type master;
	    zone-statistics no;
		file "___default_yamu.com.zone";
	};
}
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/zone-statistics

### 方法
`PUT`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| view  | String | N/A    | **说明**：view的名称，用于定位到该条view                     |
| zone  | String | N/A    | **说明**：zone的名称，用于定位到该条zone                     |

* body：

| 名称           | 类型   | 默认值 | 描述                                                         |
| :------------- | :----- | :----- | :----------------------------------------------------------- |
| value*          | String| boolean | N/A    可选 full | terse | none | <boolean> | yes | no      |

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
option {
	zone-statistics yes;
}
view __default {
	zone "yamu.com" {
	    type master;
	    zone-statistics no;
		file "___default_yamu.com.zone";
	};
}
```

#### 修改特定策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/zone-statistics?view=___default&zone=yamu.com
BODY   : "full"
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
option {
	zone-statistics yes;
}
view __default {
	zone "yamu.com" {
	    type master;
	    zone-statistics full;
		file "___default_yamu.com.zone";
	};
}
```

## 删除

### URL
http://ip:port/api/ybind/v1.0/zone-statistics

### 方法
PUT

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| view  | String | N/A    | **说明**：view的名称，用于定位到该条view                     |
| zone  | String | N/A    | **说明**：zone的名称，用于定位到该条zone                     |

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
option {
	zone-statistics yes;
}
view __default {
	zone-statistics no;
	zone "yamu.com" {
	    type master;
	    allow-query { "1.1.1.1"; };
		file "___default_yamu.com.zone";
	};
	
	zone "google.com" {
	    type slave;
		masters { "7.7.7.7"; };
		zone-statistics full;
	};
}
```

#### 删除特定策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/zone-statistics?view=__default
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
option {
	zone-statistics yes;
}
view __default {
	zone "yamu.com" {
	    type master;
	    allow-query { "1.1.1.1"; };
		file "___default_yamu.com.zone";
	};
	
	zone "google.com" {
	    type slave;
		masters { "7.7.7.7"; };
		zone-statistics full;
	};
}
```

#### 删除特定策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/zone-statistics
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
option {
}
view __default {
	zone-statistics no;
	zone "yamu.com" {
	    type master;
	    allow-query { "1.1.1.1"; };
		file "___default_yamu.com.zone";
	};
	
	zone "google.com" {
	    type slave;
		masters { "7.7.7.7"; };
		zone-statistics full;
	};
}
```