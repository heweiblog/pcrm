| 版本 | 日期       | 更改记录 | 作者 |
| :--- | :--------- | :------- | ---- |
| 0.1  | 2020-04-07 | 初稿     | 程俊 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [修改](#修改)

------------

## 接口概览
| URL                                     | 方法 | 描述          |
| --------------------------------------- | ---- | ------------- |
| http://ip:port/api/ybind/v1.0/recursion | GET  | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/recursion | PUT  | [修改](#修改) |

## 概述
* 语法：
```
recursion boolean;
```
* 概念：递归开关。默认打开，请注意当关闭时表示的是一个纯授权服务器，即转发也会被关闭。
* 注意项：
	* 可以在`option`中配置也可以在`view`中配置

## 获取

### URL
http://ip:port/api/ybind/v1.0/recursion

### 方法
`GET`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| view | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示option<br>**举例**：default |

* body：

| 名称 | 类型 | 默认值 | 描述 |
| :--- | :--- | :----- | :--- |
| N/A  | N/A  | N/A    | N/A  |

* returnBody：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| rcode*       | Int    | N/A    | 业务执行码                                                   |
| description* | String | N/A    | `rcode`的文字描述                                            |
| data         | Bool   | N/A    | **缺省**：业务执行失败或者没有配置时<br>**Bool**：`view`缺省时option下的配置或者指定`view`的策略 |

### 返回码
| rcode | description           | 说明                     |
| ----- | --------------------- | ------------------------ |
| 0     | Success               | 查询成功                 |
| 404   | Not Found             | 没有找到`view`指定的配置 |
| 408   | Request Timeout       | 请求超时                 |
| 500   | Internal Server Error | 程序运行错误             |

### 示例
* 现有策略：

```
option {
}
view _default {
	recursion no;
}
```

#### 获取特定策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/recursion?view=_default
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": false
}
```

#### 获取option策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/recursion
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
}
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/recursion

### 方法
`PUT`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| view | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示option<br>**举例**：default |

* body：

| 名称 | 类型 | 默认值 | 描述                                                         |
| :--- | :--- | :----- | :----------------------------------------------------------- |
| N/A  | Bool | N/A    | **说明**：更新指定`view`的配置或者option的配置<br>**缺省**：删除指定`view`的配置或者option配置 |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description           | 说明                   |
| ----- | --------------------- | ---------------------- |
| 0     | Success               | 修改成功               |
| 1     | Bad Parameter Format  | `view`或`body`格式错误 |
| 408   | Request Timeout       | 请求超时               |
| 500   | Internal Server Error | 程序运行错误           |

### 示例
* 现有策略：

```
option {
}
view _default {
	recursion no;
}
```

#### 修改特定策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/recursion?view=_default
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
}
view _default {
	recursion yes;
}
```

#### 更新option策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/recursion
BODY   : false
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
	recursion no;
}
view _default {
	recursion no;
}
```

#### 删除option策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/recursion
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
view _default {
	recursion no;
}
```