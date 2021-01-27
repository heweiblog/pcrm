## 接口概览
| URL                                           | 方法 | 描述                                                    |
| --------------------------------------------- | ---- | ------------------------------------------------------- |
| http://ip:port/api/ybind/v1.0/match-ecs-first | GET  | [获取match-ecs-first配置项](#获取match-ecs-first配置项) |
| http://ip:port/api/ybind/v1.0/match-ecs-first | PUT  | [更新match-ecs-first配置项](#更新match-ecs-first配置项) |

#### 概述
```
当存在2个优先级相同的视图时，在第一个视图配置ecs开关并打开，客户进行edns请求，将会优先匹配ecs视图的结果。
```

#### 配置项说明：
```
match-ecs-first yes;	//ecs开关
注：
1.该配置项只能配置在view中,为可选配置项；
2.当开启ecs开关时，在match-clients里IP段前配置ecs，或者在match-clients绑定的acl里面需要在地址段的前面加ecs；也可不配置match-clients配置项，对所有地址段生效；
3.多个优先级相同的视图存在时，只有ecs视图放在首位，客户端edns请求才会匹配中该视图
```

#### 获取match-ecs-first配置项

通过本接口获取match-ecs-first配置项
- 请求URL：`http://ip:port/api/ybind/v1.0/match-ecs-first`
- HTTP方法：GET
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| -------- | -------- | ---- |
| view*    | String   | 视图 |

- 响应参数

| 参数名称 | 参数类型 | 描述       |
| :------- | :------- | :--------- |
| data     | Bool     | 返回的数据 |

**请求示例**

```
指定视图：
GET https://ip:port/api/ybind/v1.0/match-ecs-first?view=externel
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

#### 新增/更新match-ecs-first配置项
通过本接口更新match-ecs-first配置项
- 请求URL：`http://ip:port/api/ybind/v1.0/match-ecs-first`
- HTTP方法：PUT
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| -------- | -------- | ---- |
| view*    | String   | 视图 |

以JSON的方式在body中携带

| 参数名称 | 数据类型 | 描述                               |
| -------- | -------- | ---------------------------------- |
| N/A      | Bool     | 配置项值(true/false)，缺省表示删除 |

- 响应参数

  无

**请求示例**

```
1.更新/新增：
PUT http://ip:port/api/ybind/v1.0/match-ecs-first?view=dns1
true
2.删除：
1)PUT http://ip:port/api/ybind/v1.0/match-ecs-first?&view=dns1
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