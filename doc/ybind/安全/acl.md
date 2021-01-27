| 版本 | 日期       | 更改记录     | 作者 |
| :--- | :--------- | :----------- | ---- |
| 0.1  | 2020-04-03 | 初稿         | 程俊 |
| 0.2  | 2020-04-12 | 规范化返回码 | 程俊 |
| 0.3  | 2020-04-22 | 规范化返回码 | 程俊 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [新增](#新增)
* [修改](#修改)
* [删除](#删除)

------------

## 接口概览
| URL                               | 方法   | 描述          |
| --------------------------------- | ------ | ------------- |
| http://ip:port/api/ybind/v1.0/acl | GET    | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/acl | POST   | [新增](#新增) |
| http://ip:port/api/ybind/v1.0/acl | PUT    | [修改](#修改) |
| http://ip:port/api/ybind/v1.0/acl | DELETE | [删除](#删除) |

## 概述
* 语法：`acl string { address_match_element; ... };`
* 概念：Access Control List，访问控制列表。acl语句使用一个唯一的字符串名标记一个地址匹配列表。
* 地址匹配列表包含以下几个部分：
	* IP地址(IPV4/IPV6)
	* IP地址段(使用/做分割)
	* key的ID(已经使用key语法定义的)
	* acl的名字
* acl语句有一些预定义的项：

| 名称        | 描述                                                         |
| ----------- | ------------------------------------------------------------ |
| `any`       | Matches all hosts.                                           |
| `none`      | Matches no hosts.                                            |
| `localhost` | Matches the IPv4 and IPv6 addresses of all network interfaces on the system.<br>When addresses are added or removed, the localhost ACL element is updated to reflect the changes. |
| `localnets` | Matches any host on an IPv4 or IPv6 network for which the system has an interface.<br>When addresses are added or removed,the localnets ACL element is updated to reflect the changes.<br>Some systems do not provide a way to determine the prefix lengths of local IPv6 addresses.<br>In such a case, localnets only matches the local IPv6 addresses, just like localhost. |

* 注意项：
	* `!`如果加在一个语句前面表示取反

## 获取

### URL
http://ip:port/api/ybind/v1.0/acl

### 方法
`GET`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| name | String | N/A    | **说明**：acl的名称，用于定位到该条acl<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示所有的acl<br>**举例**：acl-shanghai |

* returnBody：

| 名称         | 类型       | 默认值 | 描述                                                         |
| :----------- | :--------- | :----- | :----------------------------------------------------------- |
| rcode*       | Int        | N/A    | 业务执行码                                                   |
| description* | String     | N/A    | `rcode`的文字描述                                            |
| data         | Array/Dict | N/A    | **缺省**：业务执行失败<br>**Array**：根据检索条件`name`返回的此名称下的配置<br>**Dict**：`name`缺省时所有名称下的配置 |

### 返回码
| rcode | description           | 说明                                  |
| ----- | --------------------- | ------------------------------------- |
| 0     | Success               | 查询成功                              |
| 2     | Bad Parameter Value   | `name`值错误(比如"name="传来一个空值) |
| 404   | Not Found             | 没有找到`name`指定的acl配置           |
| 408   | Request Timeout       | 请求超时                              |
| 500   | Internal Server Error | 程序运行错误                          |

### 示例
* 现有策略：

```
acl "acl-shanghai" {
    1.1.1.1;
    localhost;
};

acl "acl-suzhou" {
    2.2/16;
};
```

#### 获取特定策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/acl?name=acl-shanghai
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": [
    "1.1.1.1",
    "localhost"
    ]
}
```

#### 获取全量策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/acl
BODY   : 
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "acl-shanghai":[
            "1.1.1.1",
            "localhost"
        ],
        "acl-suzhou":[
            "2.2/16"
        ]
    }
}
```

## 新增

### URL
http://ip:port/api/ybind/v1.0/acl

### 方法
`POST`

### 参数
* queryString：

| 名称  | 类型   | 默认值 | 描述                                                         |
| :---- | :----- | :----- | :----------------------------------------------------------- |
| name* | String | N/A    | **说明**：acl的名称，用于定位到该条acl<br>**格式**：数字、大小写字母、-、_<br>**举例**：acl-shanghai |

* body：

| 名称 | 类型  | 默认值 | 描述                                                         |
| :--- | :---- | :----- | :----------------------------------------------------------- |
| N/A* | Array | N/A    | **说明**：`name`下的策略数组<br>**注意**：数组不能为空，即不能为[] |

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
| 409   | Conflict                | `name`已存在               |
| 500   | Internal Server Error   | 程序运行错误               |

### 示例
* 现有策略：

```
acl "acl-shanghai" {
    1.1.1.1;
    localhost;
};

acl "acl-suzhou" {
    2.2/16;
};
```

#### 增加特定策略
* 请求：
```
METHOD : POST
URL    : http://ip:port/api/ybind/v1.0/acl?name=acl-beijing
BODY   : [
	"3.3.3.3"
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
acl "acl-shanghai" {
    1.1.1.1;
    localhost;
};

acl "acl-suzhou" {
    2.2/16;
};

acl "acl-beijing" {
    3.3.3.3;
};
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/acl

### 方法
`PUT`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| name | String | N/A    | **说明**：acl的名称，用于定位到该条acl<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示所有的acl<br>**举例**：acl-shanghai |

* body：

| 名称 | 类型       | 默认值 | 描述                                                         |
| :--- | :--------- | :----- | :----------------------------------------------------------- |
| N/A  | Array/Dict | N/A    | **说明**：根据传入的类型更新指定`name`的配置或者覆盖所有配置<br>**Array**：更新`name`的配置<br>**Dict**：覆盖所有的配置 |

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
| 404   | Not Found               | 该`name`下的策略没有找到   |
| 408   | Request Timeout         | 请求超时                   |
| 500   | Internal Server Error   | 程序运行错误               |

### 示例
* 现有策略：

```
acl "acl-shanghai" {
    1.1.1.1;
    localhost;
};

acl "acl-suzhou" {
    2.2/16;
};
```

#### 修改特定策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/acl?name=acl-shanghai
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
acl "acl-shanghai" {
    1.1.1.1;
};

acl "acl-suzhou" {
    2.2/16;
};
```

#### 更新全部策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/acl
BODY   : {
    "acl-shanghai":[
        "1.1.1.1",
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
acl "acl-shanghai" {
    1.1.1.1;
};
```

## 删除

### URL
http://ip:port/api/ybind/v1.0/acl

### 方法
`DELETE`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| name | String | N/A    | **说明**：acl的名称，用于定位到该条acl<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示所有的acl<br>**举例**：acl-shanghai |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description           | 说明                 |
| ----- | --------------------- | -------------------- |
| 0     | Success               | 删除成功             |
| 2     | Bad Parameter Value   | `name`或`body`值错误 |
| 408   | Request Timeout       | 请求超时             |
| 500   | Internal Server Error | 程序运行错误         |

### 示例
* 现有策略：

```
acl "acl-shanghai" {
    1.1.1.1;
    localhost;
};

acl "acl-suzhou" {
    2.2/16;
};
```

#### 删除特定策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/acl?name=acl-shanghai
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
acl "acl-suzhou" {
    2.2/16;
};
```

#### 删除全部策略
* 请求：
```
METHOD : DELETE
URL    : http://ip:port/api/ybind/v1.0/acl
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

```