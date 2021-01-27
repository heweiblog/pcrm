#### 概述
在收到递归、转发应答结果后，将满足过滤条件的A/AAAA记录去掉，然后将应答内容重新组装之后发送给用户。

#### 配置项说明：
```
	deny-answer-addresses {acl;};  //过滤ip地址列表
	deny-answer-addresses {acl;} from {“domain”;};   //过滤掉针对域/域名的应答A记录中包含在IP列表中的地址；
	deny-answer-addresses {acl;} except-from { “domain”;};    //过滤ip地址列表，但对域/域名的应答排除
注：
	1.该项可同时配置在options、view、zone中
	2.在options/view/zone中，只能配置一条配置项
	3.deny-answer-addresses {acl;}为必选参数，from/except-from { “domain”;}为可选参数。
```
#### 3.1.1 获取deny-answer-addresses配置项
通过本接口获取deny-answer-addresses配置项
- 请求URL：`http://ip:port/ybind/deny-answer-filter`
- HTTP方法：GET
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| :------- | :------- | ---- |
| view     | String   | 视图 |
| zone     | String   | 区域 |


- 响应参数

| 参数名称 | 参数类型 | 描述       |
| :------- | :------- | ---------- |
| data     | dict     | 返回的数据 |

**请求示例**

```
GET http://ip:port/ybind/deny-answer-filter?view=any&zone=yamu.com
```

**返回示例**
```
# 成功返回
实例一
{
    "data": {"iplist": ["1.1.1.1/24","2.2.2.2/24"],},
    "description": "Success",
    "rcode": 0
}
实例二
{
    "data": {
		"iplist": ["1.1.1.1","2.2.2.2"],
		"type":"from",
		"domain":["www.yamu.com"]},
    "description": "Success",
    "rcode": 0
}


# 失败返回
{
    "description": "Bad Parameter Format",
    "rcode": 1
}
```

#### 3.1.2 设置deny-answer-addresses配置项
通过本接口设置deny-answer-addresses配置项
- 请求URL：`http://ip:port/ybind/deny-answer-filter`
- HTTP方法：PUT
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述 |
| :------- | :------- | ---- |
| view     | String   | 视图 |
| zone     | String   | 区域 |

以JSON的方式在body中携带

| 参数名称 | 数据类型 | 描述                                   |
| :------- | :------- | -------------------------------------- |
| iplist*  | List     | IP列表                                 |
| type     | String   | 针对域名的过滤/排除 (from/except-from) |
| domain   | List     | 域名                                   |

注：type， domain 是一组参数，要么都传，要么都不传。


- 响应参数

  无

**请求示例**
```
1.全局 ：不加view/zone参数，表示策略配置在options中
PUT https://ip:port/ybind/deny-answer-filter
{
"iplist": ["1.1.1.1","2.2.2.2"] 
}

2.zone：不加view参数，表示在options下的yamu.com区域配置策略
PUT https://ip:port/ybind/deny-answer-filter?zone=yamu.com
{
"iplist": ["1.1.1.1","2.2.2.2"],
"type":"from",
"domain":["www.yamu.com"]
}

3.视图：表示在any视图下配置策略
PUT https://ip:port/ybind/deny-answer-filter?view=any
{
"iplist": ["1.1.1.1","2.2.2.2"],
"type":"except-from",
"domain":["yamu.com"]
}

4.视图+zone：表示在any视图中的yamu.com区域下配置的策略
PUT https://ip:port/ybind/deny-answer-filter?view=any&zone=yamu.com {
"iplist": ["1.1.1.1","2.2.2.2"],
"type":"from",
"domain":["www.yamu.com","ns.yamu.com"]
}
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