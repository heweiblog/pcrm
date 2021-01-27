## 接口概览
| URL                                                | 方法 | 描述                                                         |
| -------------------------------------------------- | ---- | ------------------------------------------------------------ |
| http://ip:port/api/ybind/v1.0/match-clients        | GET  | [获取match-clients配置项](#获取match-clients配置项)          |
| http://ip:port/api/ybind/v1.0/match-clients        | PUT  | [更新match-clients配置项](#更新match-clients配置项)          |
| http://ip:port/api/ybind/v1.0/match-destinations   | GET  | [获取match-destinations配置项](#获取match-destinations配置项) |
| http://ip:port/api/ybind/v1.0/match-destinations   | PUT  | [更新match-destinations配置项](#更新match-destinations配置项) |
| http://ip:port/api/ybind/v1.0/match-domains        | GET  | [获取match-domains配置项](#获取match-domain配置项)           |
| http://ip:port/api/ybind/v1.0/match-domains        | PUT  | [更新match-domains配置项](#更新match-domain配置项)           |
| http://ip:port/api/ybind/v1.0match-recursive-only  | GET  | [获取match-recursive-only配置项](#获取match-recursive-only配置项) |
| http://ip:port/api/ybind/v1.0/match-recursive-only | PUT  | [更新match-recursive-only配置项](#更新match-recursive-only配置项) |

#### 概述
```
根据每个配置项的优先级不同，配置多个不同优先级的视图，用户请求时，将会匹配优先级最高的视图，回复该视图下的结果给用户。
```

#### 配置项说明：
```
match-recursive-only yes;	//RD标志位开关，默认优先级为1；
match-clients { address_match_list } ;	//源ip，默认优先级为2；
match-destinations { address_match_list };	//目的ip，默认优先级为4；
match-domains{ zone };	//域/域名，默认优先级为8；
注：1.这些配置项只能配置在view中，为可选配置项；
2.视图优先级为个配置项优先级相加；
```

#### 获取match-clients配置项

通过本接口获取match-clients配置项
- 请求URL：`http://ip:port/api/ybind/v1.0/match-clients`
- HTTP方法：GET
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| -------- | -------- | ---- |
| view*    | String   | 视图 |

- 响应参数

| 参数名称 | 参数类型 | 描述       |
| -------- | -------- | ---------- |
| data     | Array    | 返回的数据 |

**请求示例**

```
指定视图：
GET https://ip:port/api/ybind/v1.0/match-clients?view=externel
```

**返回示例**
```
# 成功返回
指定视图返回：
{
    "description": "Success",
    "rcode": 0,
    "data":["192.168.16.0/24","192.168.17.4/32"]
}

# 失败返回
{
    "description": "Bad Parameter Format",
    "rcode": 1
}
```

#### 新增/更新match-clients配置项
通过本接口更新match-clients配置项
- 请求URL：`http://ip:port/api/ybind/v1.0/match-clients`
- HTTP方法：PUT
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| -------- | -------- | ---- |
| view*    | String   | 视图 |

以JSON的方式在body中携带

| 参数名称 | 参数类型 | 描述                                                    |
| -------- | -------- | ------------------------------------------------------- |
| N/A      | Array    | 配置项值(不指定或者为空[]时表示删除match-clients配置项) |

- 响应参数

  无

**请求示例**

```
新增/更新
PUT http://ip:port/api/ybind/v1.0/match-clients?&view=dns1
["192.168.16.0/24","192.168.1.0/24"]
删除
PUT http://ip:port/api/ybind/v1.0/match-clients?&view=dns1
[]
```

**返回示例**
```
# 成功返回
{
    "description": "Success",
    "rcode": 0
}

# 失败返回
{
    "description": "Bad Parameter Format",
    "rcode": 1
}
```

#### 获取match-destinations配置项

通过本接口获取视图匹配match-destinations配置项
- 请求URL：`http://ip:port/api/ybind/v1.0/match-destinations`
- HTTP方法：GET
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| -------- | -------- | ---- |
| view*    | String   | 视图 |

- 响应参数

| 参数名称 | 参数类型 | 描述       |
| -------- | -------- | ---------- |
| data     | Array    | 返回的数据 |

**请求示例**

```
指定视图：
GET https://ip:port/api/ybind/v1.0/match-destinations?view=externel
```

**返回示例**
```
# 成功返回
指定视图返回：
{
    "description": "Success",
    "rcode": 0,
    "data":["192.168.16.0/24","192.168.17.4/32"]
}

# 失败返回
{
    "description": "Bad Parameter Format",
    "rcode": 1
}
```

#### 新增/更新match-destinations配置项
通过本接口更新match-destinations配置项
- 请求URL：`http://ip:port/api/ybind/v1.0/match-clients`
- HTTP方法：PUT
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| -------- | -------- | ---- |
| view*    | String   | 视图 |

以JSON的方式在body中携带

| 参数名称 | 数据类型 | 描述                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| N/A      | Array    | 配置项值(不指定或者为空时[]表示删除mmatch-destinations配置项) |

- 响应参数

  无

**请求示例**
```
1.更新/新增：
PUT http://ip:port/api/ybind/v1.0/match-destinations?view=dns1
["127.0.0.1","192.168.1.0/24"]
2.删除：
1)PUT http://ip:port/api/ybind/v1.0/match-destinations?view=dns1
[]
```

**返回示例**
```
# 成功返回
{
    "description": "Success",
    "rcode": 0
}

# 失败返回
{
    "description": "Bad Parameter Format",
    "rcode": 1
}
```

#### 获取match-domains配置项

通过本接口获取视图匹配match-domains配置项
- 请求URL：`http://ip:port/api/ybind/v1.0/match-domains`
- HTTP方法：GET
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| -------- | -------- | ---- |
| view*    | String   | 视图 |

- 响应参数

| 参数名称 | 参数类型 | 描述       |
| -------- | -------- | ---------- |
| data     | Array    | 返回的数据 |

**请求示例**

```
指定视图：
GET https://ip:port/api/ybind/v1.0/match-domains?view=externel
```

**返回示例**

```
# 成功返回
指定视图返回：
{
    "description": "Success",
    "rcode": 0,
    "data":["test.com","www.baidu.com"]
}

# 失败返回
{
    "description": "Bad Parameter Format",
    "rcode": 1
}
```

#### 新增/更新match-domains配置项
通过本接口更新match-domain配置项
- 请求URL：`http://ip:port/api/ybind/v1.0/match-domains`
- HTTP方法：PUT
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| -------- | -------- | ---- |
| view*    | String   | 视图 |

以JSON的方式在body中携带

| 参数名称 | 数据类型 | 描述                                                 |
| -------- | -------- | ---------------------------------------------------- |
| N/A      | Array    | 配置项值(不指定或者为空时表示删除match-domain配置项) |

- 响应参数

  无

**请求示例**

```
1.更新/新增：
PUT http://ip:port/api/ybind/v1.0/match-domains?view=dns1
["test.com","www.baidu.com"]
2.删除：
1)PUT http://ip:port/api/ybind/v1.0/match-domains?view=dns1
[]
```

**返回示例**

```
# 成功返回
{
    "description": "Success",
    "rcode": 0
}

# 失败返回
{
    "description": "Bad Parameter Format",
    "rcode": 1
}
```

#### 获取match-recursive-only配置项

通过本接口获取match-recursive-only配置项
- 请求URL：`http://ip:port/api/ybind/v1.0/match-recursive-only`
- HTTP方法：GET
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| -------- | -------- | ---- |
| view*    | String   | 视图 |

- 响应参数

| 参数名称 | 参数类型 | 描述       |
| -------- | -------- | ---------- |
| data     | Bool     | 返回的数据 |

**请求示例**

```
指定视图：
GET https://ip:port/api/ybind/v1.0/match-recursive-only?view=externel
```

**返回示例**
```
# 成功返回
指定视图返回：
{
    "description": "Success",
    "rcode": 0,
    "data": true
}

# 失败返回
{
    "description": "Bad Parameter Format",
    "rcode": 1
}
```

#### 新增/更新match-recursive-only配置项
通过本接口更新match-domain配置项
- 请求URL：`http://ip:port/api/ybind/v1.0/match-recursive-only`
- HTTP方法：PUT
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| -------- | -------- | ---- |
| view*    | String   | 视图 |

以JSON的方式在body中携带

| 参数名称 | 数据类型 | 描述                              |
| -------- | -------- | --------------------------------- |
| N/A      | Bool     | 配置项值(true/false),缺省表示删除 |

- 响应参数

  无

**请求示例**
```
1.更新/新增：
PUT http://ip:port/api/ybind/v1.0/match-recursive-only?view=dns1
true
2.删除：
1)PUT http://ip:port/api/ybind/v1.0/match-recursive-only?view=dns1
```

**返回示例**
```
# 成功返回
{
    "description": "Success",
    "rcode": 0
}

# 失败返回
{
    "description": "Bad Parameter Format",
    "rcode": 1
}
```