| 版本 | 日期       | 更改记录                                                     | 作者 |
| :--- | :--------- | :----------------------------------------------------------- | ---- |
| 0.1  | 2020-04-06 | 初稿                                                         | 程俊 |
| 0.2  | 2020-04-08 | 添加配置项的默认值提示<br>`GET`方法删除body部分<br>`POST`冲突返回409<br>`PUT`方法body不能缺省 | 程俊 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [修改](#修改)

------------

## 接口概览
| URL                                        | 方法 | 描述          |
| ------------------------------------------ | ---- | ------------- |
| http://ip:port/api/ybind/v1.0/allow-update | GET  | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/allow-update | PUT  | [修改](#修改) |

## 概述
* 语法：
```
allow-update { address_match_element; ... };
```
* 概念：指定允许哪些主机向该区域提交动态DNS更新。目前只能在zone级别上设置，默认是拒绝所有主机的更新。
* 注意项：
	* 只可以在master中配置

## 获取

### URL
http://ip:port/api/ybind/v1.0/allow-update

### 方法
`GET`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**举例**：default |
| zone* | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**举例**：yamu.com |

* returnBody：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| rcode*       | Int    | N/A    | 业务执行码                                                   |
| description* | String | N/A    | `rcode`的文字描述                                            |
| data         | Array  | N/A    | **缺省**：业务执行失败<br>**Array**：指定`view`下`zone`的策略 |

### 返回码
| rcode | description           | 说明                                     |
| ----- | --------------------- | ---------------------------------------- |
| 0     | Success               | 查询成功                                 |
| 404   | Not Found             | 没有找到`view`指定的配置或者`zone`的配置 |
| 408   | Request Timeout       | 请求超时                                 |
| 500   | Internal Server Error | 程序运行错误                             |

### 示例
* 现有策略：

```
view _default {
	zone yamu.com {
		type master;
		allow-update { 8.8.8.8; };
		file "__default_yamu.com.zone";
	};
}
```

#### 获取特定策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/allow-update?view=_default&zone=yamu.com
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": [
		"8.8.8.8"
	]
}
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/allow-update

### 方法
`PUT`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| view* | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**举例**：default |
| zone* | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**举例**：yamu.com |

* body：

| 名称 | 类型  | 默认值 | 描述                                                         |
| :--- | :---- | :----- | :----------------------------------------------------------- |
| N/A* | Array | N/A    | **说明**：更新指定`view`下指定`zone`的配置<br>**注意**：可以为空：[]，删除指定`view`下指定`zone`的配置 |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description           | 说明                           |
| ----- | --------------------- | ------------------------------ |
| 0     | Success               | 修改成功                       |
| 1     | Bad Parameter Format  | `view`或`body`或`zone`格式错误 |
| 408   | Request Timeout       | 请求超时                       |
| 500   | Internal Server Error | 程序运行错误                   |

### 示例
* 现有策略：

```
view _default {
	zone yamu.com {
		type master;
		allow-update { 8.8.8.8; };
		file "__default_yamu.com.zone";
	};
}
```

#### 修改特定策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/allow-update?view=_default&zone=yamu.com
BODY   : [
	"1.1.1.1"
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
view _default {
	zone yamu.com {
		type master;
		allow-update { 1.1.1.1; };
		file "__default_yamu.com.zone";
	};
}
```

#### 删除特定策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/allow-update?view=_default&zone=yamu.com
BODY   : []
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
	zone yamu.com {
		type master;
		file "__default_yamu.com.zone";
	};
}
```