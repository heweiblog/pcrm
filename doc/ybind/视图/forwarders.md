| 版本 | 日期       | 更改记录                                                     | 作者 |
| :--- | :--------- | :----------------------------------------------------------- | ---- |
| 0.1  | 2020-04-06 | 初稿                                                         | 程俊 |
| 0.2  | 2020-04-08 | 添加配置项的默认值提示<br>`GET`方法删除body部分<br>`POST`冲突返回409<br>`PUT`方法body不能缺省 | 程俊 |
| 0.3  | 2020-04-20 | 删除DELETE接口,使用PUT空来删除; 增加具体到zone的GET,PUT策略  | 韩冬 |
------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [修改](#修改)
------------

## 接口概览
| URL                                      | 方法 | 描述          |
| ---------------------------------------- | ---- | ------------- |
| http://ip:port/api/ybind/v1.0/forwarders | GET  | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/forwarders | PUT  | [修改](#修改) |


## 概述
* 语法：
```
forwarders [ port integer ] [ dscp integer ] { ( ipv4_address | -
ipv6_address ) [ port integer ] [ dscp integer ] [weight | order interger]; ... };
};
```
* 概念：转发目标地址。指定了转发策略的目标地址。
* 注意项：
	* 可以在`option`中配置也可以在`view`中配置
	* 目前只支持IPV4、IPV6地址的配置

## 获取

### URL
http://ip:port/api/ybind/v1.0/forwarders

### 方法
`GET`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| view | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示option<br>**举例**：default |
| zone | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**举例**：yamu.com |
* returnBody：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| rcode*       | Int    | N/A    | 业务执行码                                                   |
| description* | String | N/A    | `rcode`的文字描述                                            |
| data         | Array  | N/A    | **缺省**：业务执行失败<br>**Array**：`view`缺省时option下的配置或者指定`view`的策略 |

### 返回码
| rcode | description           | 说明                     |
| ----- | --------------------- | ------------------------ |
| 0     | Success               | 查询成功                 |
| 404   | Not Found             | 没有找到`view`指定的配置 |
| 408   | Request Timeout       | 请求超时                 |
| 500   | Internal Server Error | 程序运行错误             |

### 示例


#### (1) 获取options策略
* 现有策略：

```
option {
}
view _default {
	forward only algo weight ;
	forwarders { 8.8.8.8 weight 2; 8.8.4.4 weight 2; 114.114.114.114 weight 4; };
}
```

* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/forwarders
BODY   : 
```

* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": []
}
```

#### (2) 获取view策略
* 现有策略：

```
option {
}
view _default {
	forward only algo weight ;
	forwarders { 8.8.8.8 weight 2; 8.8.4.4 weight 2; 114.114.114.114 weight 4; };
}
```

* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/forwarders?view=_default
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": [
		{"ip": "8.8.8.8", "weight":2},
		{"ip": "8.8.4.4", "weight":2},
		{"ip": "114.114.114.114", "weight": 4}
	]
}
```
#### (3) 获取view下zone策略
* 现有策略：

```
option {
}
view _default {
	forward only algo weight ;
	forwarders { 8.8.8.8 weight 2; 8.8.4.4 weight 2; 114.114.114.114 weight 4; };
	zone "yamu.com" IN {
		type forward;
		forward fisrt algo weight;
		forwarders {1.1.1.1 weight 1; 2.2.2.2 weight 2;};
	}
}
```

* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/forwarders?view=_default&zone=yamu.com
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": [
		{"ip": "1.1.1.1", "weight":1},
		{"ip": "2.2.2.2", "weight":2}
	]
}
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/forwarders

### 方法
`PUT`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| view | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示option<br>**举例**：default |

* body：

| 名称 | 类型  | 默认值 | 描述                                                         |
| :--- | :---- | :----- | :----------------------------------------------------------- |
| N/A* | Array | N/A    | **说明**：更新指定`view`的配置或者option的配置<br>**注意**：可以为空：""，删除指定`view`的配置或者option配置 |

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

#### (1) 添加/修改options策略

* 现有策略：

```
options {
}
view _default {
	forward first;
	forwarders { 8.8.8.8 weight 2; 8.8.4.4 weight 2; 114.114.114.114 weight 4; };
}
```

* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/forwarders
BODY   : [
	{"ip": 1.1.1.1, "weight": 2};
	{"ip": 2.2.2.2, "weight": 3}
]
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success"
}
```
* 接口调用成功后策略：

```
option {
	forwarders { 1.1.1.1 weight 2; 2.2.2.2 weight 3; };
}
view _default {
	forward first;
	forwarders { 8.8.8.8 weight 2; 8.8.4.4 weight 2; 114.114.114.114 weight 4; };
}
```


#### (2) 添加/修改view策略

* 现有策略：

```
option {
}
view _default {
	forward only algo weight;
	forwarders { 8.8.8.8 weight 2; 8.8.4.4 weight 2; 114.114.114.114 weight 4; };
}
```

* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/forwarders?view=_default
BODY   : [
	{"ip": 1.1.1.1, "weight": 2};
	{"ip": 2.2.2.2, "weight": 3}
]
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success"
}
```
* 接口调用成功后策略:

  ```
  option {
  }
  view _default {
  forward only algo weight;
  forwarders { 1.1.1.1 weight 2; 2.2.2.2 weight 3; };
  }
  ```
#### (3) 添加/修改view下zone策略

* 现有策略：

```
option {
}
view _default {
	forward only algo weight;
	forwarders { 8.8.8.8 weight 2; 8.8.4.4 weight 2; 114.114.114.114 weight 4; };
	zone "yamu.com" IN {
		type forward;
		forward first algo order;
		forwarders {1.1.1.1 order 1; 2.2.2.2 order 2;};
	}
}
```

* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/forwarders?view=_default
BODY   : [
	{"ip": "3.3.3.3", "order": 2};
	{"ip": "4.4.4.4", "order": 3}
]
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success"
}
```
* 接口调用成功后策略:

  ```
  option {
  }
  view _default {
  forward only algo weight;
  forwarders { 1.1.1.1 weight 2; 2.2.2.2 weight 3; };
  zone "yamu.com" IN {
  	type forward;
  	forward first algo order;
  	forwarders {3.3.3.3 order 2; 4.4.4.4 order 3;};
  }
  }
  ```

#### (4) 删除view策略

* 现有策略：

```
options {
}
view _default {
	forwarders { 8.8.8.8 weight 2; 8.8.4.4 weight 2; 114.114.114.114 weight 4; };
}
```

* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/forwarders
BODY   :

```
* 返回：
```
{
    "rcode": 0,
    "description": "Success"
}
```
* 接口调用成功后策略：

```
option {

}
view _default {

}
```