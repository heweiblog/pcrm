#### 配置项说明
1. rrset-order配置项定义记录响应的返回顺序，参数有qtype（请求类型）,domain（域名）,order(返回策略)。例如：

~~~
rrset-order {
class IN type A name "host.example.com" order random;
order cyclic;
};
~~~

2. order策略说明：

| Item   | Description                                                  |
| :----- | :----------------------------------------------------------- |
| fixed  | 记录按照它们在zone中定义的顺序返回。此选项仅在编译时配置为“-enable-fixed-rrset”时可用 |
| random | 记录以随机顺序返回。                                         |
| cyclic | 记录以循环循环的顺序返回，每个查询循环一条记录。<br/>如果绑定在编译时配置为“-enable-fixed-rrset”，则RRset的初始顺序将与区域文件中指定的顺序匹配;否则，初始排序是不确定的。 |
| none   | 记录以从数据库检索到的任何顺序返回。这个顺序是不确定的，但是只要数据库没有被修改，它就是一致的。当没有指定顺序时，这是默认值。 |

3. 如果出现多个rrset-order语句，则不合并它们，只生效最后一个应用。

​     全局控制策略  rrset-order{order cyclic}
​	

4. 特定记录策略  rrset-order { class IN type A name "host.example.com" order random}

### 获取rrset-order配置项

- 请求URL：`http://ip:port/ybind/cache/rrset-order`
- HTTP方法：GET
- 请求参数：以query string的方式携带

| 参数名称 | 数据类型 | 描述                                                 |
| :------- | :------- | ---------------------------------------------------- |
| view     | String   | 视图                                                 |
| key      | String   | rrset-order排序策略名(不传获取整个rrset-order配置项) |

- 响应参数

| 参数名称 | 参数类型 | 描述       |
| :------- | :------- | ---------- |
| data     | dict     | 返回的数据 |

**请求示例**

```
GET https://ip:port/ybind/cache/rrset-order?view=dns1
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
### 新增和更新rrset-order配置项
- 请求URL：`http://ip:port/ybind/cache/rrset-order`
- HTTP方法：POST
- 请求参数：以JSON的方式在body中携带

| 参数名称 | 必选参数 | 数据类型 | 描述                                        |
| :------- | :------- | :------- | ------------------------------------------- |
| view     | N        | String   | 视图                                        |
| key      | Y        | String   | rrset-order排序策略id(策略记录唯一标识符)   |
| qtype    | N        | String   | 请求类型 (A，AAAA,NS)                       |
| domain   | N        | String   | 域名                                        |
| order    | Y        | String   | rr策略 （cyclic,random,fiexd,none）默认none |

- 响应参数

  无

**请求示例**

```
1.视图
{
    "view":"default",
    "type":"A",
	"dmain":"www.baidu.com",
	"order":"cyclic",
	"key":"www.baidu,comA"
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

### 删除rrset-order配置项

- 请求URL：`http://ip:port/ybind/cache/rrset-order`
- HTTP方法：DELETE
- 请求参数：以query string的方式携带

| 参数名称 | 必选参数 | 数据类型 | 描述                                                 |
| :------- | :------- | :------- | ---------------------------------------------------- |
| view     | N        | String   | 视图                                                 |
| key      | N        | String   | rrset-order排序策略名(不传删除整个rrset-order配置项) |
- 响应参数

  无

**请求示例**
```
1.视图
    DELETE https://ip:port/ybind/cache/rrset-order?view=dns1
2.视图+key
    DELETE https://ip:port/ybind/cache/rrset-order?view=dns1&key=www.baidu,comA
3.全局options
    DELETE https://ip:port/ybind/cache/rrset-order
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