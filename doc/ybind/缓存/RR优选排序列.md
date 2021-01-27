优先级低
#### 配置说明
```
sortlist ：sortlist { <address_match_element>; ... };
```
一个DNS请求的响应由多个资源记录组成，它们形成一个资源记录集。名称服务器在回应客户端的请求时，其回复的资源记录的顺序是不确定的(BIND9默认是采用轮询机制)。BIND中排序功能是基于配置文件sortlist在DNS服务器上实现的。sortlist排序功能是由两个元素决定，这两个元素缺一不可。
Sortlist 语句是由两个元素组成的，一个是address_match_list（客户端IP地址/IP段）;另一个是设置资源记录所对应的IP或所在的网段。具体的sortlist配置语句如下所示:

~~~
Options {

Sortlist{
{192.168.1.0/24;                         #------规定客户端的范围；
{192.168.1.0/24;                         #------规定优先返回的资源记录的范围
{192.168.2.0/24;192.168.3.0/24;};};};    #------规定次优先返回的资源记录的范围
};

}
~~~

如果只配置了sortlist，而没有配置rrset-order，那么默认的rrset-order为随机（random），即最高优先级随机IP出现在第一位。

### 获取sortlist配置项

通过本接口获取sortlist配置项
- 请求URL：`http://ip:port/ybind/cache/sortlist`
- HTTP方法：GET
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述                                          |
| :------- | :------- | --------------------------------------------- |
| view     | String   | 视图                                          |
| key      | String   | sorlist排序策略名(不传返回整个sortlist配置项) |
- 响应参数

| 参数名称 | 参数类型 | 描述       |
| :------- | :------- | ---------- |
| data     | dict     | 返回的数据 |

**请求示例**

```
GET https://ip:port/ybind/cache/sortlist?view=dns1&key=xxx;1.1.1.1,2.2.2/24
```

**返回示例**

```
# 成功返回
{
    "data": {},
    "description": "Success",
    "rcode": 0
}

# 失败返回
{
    "description": "Bad Parameter Format",
    "rcode": 1
}
```
**GET的其他方法**
```
获取视图下所有的sortlist: GET https://ip:port/ybind/cache/sortlist?view=dns1
获取options 下所有的sortlist: GET https://ip:port/ybind/cache/sortlist
获取options 下单条sortlist: GET https://ip:port/ybind/cache/sortlist?key=xxx;1.1.1.1,2.2.2/24
```

### 新增和更新sortlist配置项
通过本接口更新缓存
- 请求URL：`http://ip:port/ybind/cache/sortlist`
- HTTP方法：PUT
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述                      |
| :------- | :------- | ------------------------- |
| view     | String   | 视图（可缺省表示options）                      |
| key*     | String   | sorlist排序策略名(策略id)（可缺省，表示全量配置） |
| body     | json     | 全量配置为字典/非全量配置为列表|
- 响应参数

  无

**请求示例**

```
1. 单个配置
[
    ["1.1.1.1/24","2.2.2.2/24"],
    ["3.3.3.3/24","4.4.4.4/24"]
]

2. 全量配置
{
    "1.1.1/24;2.2.2/24":[
        ["3.3.3/24","4.4.4.4"]
        ["5.5.5/24","6.6.6.6"]
    ]
}


```

**响应示例**
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

### 删除sortlist配置项
通过本接口获取缓存
- 请求URL：`http://ip:port/ybind/cache/sortlist`
- HTTP方法：PUT
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述                                          |
| :------- | :------- | --------------------------------------------- |
| view     | String   | 视图 (不传表示options)                                         |
| key      | String   | sorlist排序策略名(不传删除整个sortlist配置项) |
- 响应参数

  无

**请求示例**
```
1.视图
    PUT https://ip:port/ybind/cache/sortlist?view=dns1
2.视图+key
    PUT https://ip:port/ybind/cache/sortlist?view=dns1&key=xxx,1.1.1.1,2.2.2/24
3.全局options
    PUT https://ip:port/ybind/cache/sortlist
4.全局+key
    PUT https://ip:port/ybind/cache/sortlist?key=xxx,1.1.1.1,2.2.2/24
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