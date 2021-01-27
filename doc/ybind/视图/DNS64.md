| 版本 | 日期       | 更改记录 |  作者  |
| :--- | :--------- | :------ | ----- |
| 1.0  | 2020-05-25 | 初稿     | 张宁宁 |
| 2.0  | 2020-06-01 | 修订     | 汪凌   |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [修改](#修改)

------------

## 接口概览
| URL                                       | 方法 | 描述          |
| ----------------------------------------- | ---- | ------------- |
| http://ip:port/api/ybind/v1.0/dns64 | GET  | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/dns64 | PUT  | [修改](#修改) |

## 概述
* 语法：
```
dns64 netprefix {
break-dnssec boolean;
clients { address_match_element; ... };
exclude { address_match_element; ... };
mapped { address_match_element; ... };
recursive-only boolean;
suffix ipv6_address;
}
```
* 概念：这条指令指示named在没有找到AAAA记录时，返回由IPv4地址所映射到的AAAA查询。其意图
是用于和NAT64相配合。每个dns64定义一个DNS64前缀。可以定义多个DNS64前缀。

* 注意项：
	* 可以在`option`中配置，也可以在视图中配置多个dns64,

## 获取options所有dns64配置

### URL
http://ip:port/api/ybind/v1.0/dns64

### 方法
`GET`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| view | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示option，此时会忽略zone<br>**举例**：default |
| name | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示某个视图或者全局下的所有dns64配置，


* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :--------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |
| data         | Map    | N/A    | **失败**:返回404  |

### 返回码
| rcode | description           | 说明                                     |
| ----- | --------------------- | ---------------------------------------- |
| 0     | Success               | 查询成功                                 |
| 404   | Not Found             | 没有找到`option`指定的配置 |
| 408   | Request Timeout       | 请求超时                                 |
| 500   | Internal Server Error | 程序运行错误                             |

### 示例
* 现有策略：

```
view __default {
    dns64 2001::/32{
        clients{1.1.1.1; };
        exclude{2001::1; };
        mapped{2.2.2.2; };
        suffix ::;
        break-dnssec true;
        recursive-only false;
    };
};
options{
    dns64 2002::/32{
        clients{3.3.3.3; };
        exclude{2002::2; };
        mapped{4.4.4.4; };
        suffix ::;
        break-dnssec true;
        recursive-only false;
    };
};
```
## 获取options单个配置
* 请求:
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/dns64?name=2002::/32
```

* 返回：

```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "break-dnssec": true,
        "clients": [
            "3.3.3.3"
        ],
        "exclude": [
            "2002::2"
        ],
        "mapped": [
            "4.4.4.4"
        ],
        "recursive-only": false,
        "suffix": "::"
    }
}
```
## 获取options全量DNS64配置
* 请求:
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/dns64
```

* 返回：

```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "2002::/32": {
            "break-dnssec": true,
            "clients": [
                "3.3.3.3"
            ],
            "exclude": [
                "2002::2"
            ],
            "mapped": [
                "4.4.4.4"
            ],
            "recursive-only": false,
            "suffix": "::"
        }
    }
}
```
## 获取视图单个配置
* 请求：
```
METHOD: GET
URL   : http://ip:port/api/ybind/v1.0/dns64?view=__default&name=2001::/32
```
* 返回
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "break-dnssec": true,
        "clients": [
            "1.1.1.1"
        ],
        "exclude": [
            "2001::1"
        ],
        "mapped": [
            "2.2.2.2"
        ],
        "recursive-only": false,
        "suffix": "::"
    }
}
```

#### 获取view中全量的DNS64策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/dns64?view=__default
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "2001::/32": {
            "break-dnssec": true,
            "clients": [
                "1.1.1.1"
            ],
            "exclude": [
                "2001::1"
            ],
            "mapped": [
                "2.2.2.2"
            ],
            "recursive-only": false,
            "suffix": "::"
        }
    }
}
```

## 增加/删除/修改
### URL
http://ip:port/api/ybind/v1.0/dns64

### 方法
`PUT`
### 参数
```
* QueryString：
```
| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| view | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示option，此时会忽略zone<br>**
| *name | String | N/A    | **说明**：IPv6地址段，用于定位具体的dns64配置<br>**格式**：只允许配置IPv6地址段<br>**不可缺省**
```

*BODY：
```
| 名称 | 类型 | 默认值 | 描述                                                                  |
| :--- | :--- | :----- | :-------------------------------------------------------------------- |
| N/A* | Map | N/A    | **说明**：更新`dns64`的配置<br>**注意**：可以为空，表示删除对应得dns64配置 |

* ReturnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description           | 说明                           |
| ----- | --------------------- | ------------------------------ |
| 0     | Success               | 修改成功                       |
| 1     | Bad Parameter Format  | `body`格式错误                 |
| 408   | Request Timeout       | 请求超时                       |
| 500   | Internal Server Error | 程序运行错误                   |

### 示例:
#### <1>新增view的单个DNS64配置：
* 请求
```
METHOD : PUT
URL : http://ip:port/api/ybind/v1.0/dns64?name=2001::/32
BODY: 
{
    "break-dnssec": true,
    "clients": ["1.1.1.1"],
    "exclude": ["2001::1"],
    "mapped": ["2.2.2.2"],
    "recursive-only": false,
    "suffix": "::"
}
```
* 返回
```
{
    "rcode": 0,
    "description": "Success"
}
```
* 策略
```
view __default {
    dns64 2001::/32{
        clients{1.1.1.1; };
        exclude{2001::1; };
        mapped{2.2.2.2; };
        suffix ::;
        break-dnssec true;
        recursive-only false;
    };
};
```
* NOTE:
```
    存在即修改，不存在即新增
```
#### <2>删除view中单个DNS64配置:
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/dns64?view=__default&name=2001::/32
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
view __default{
}
```