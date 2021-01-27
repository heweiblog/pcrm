| 版本 | 日期       | 更改记录 | 作者   |
| :--- | :--------- | :------- | ------ |
| 0.1  | 2020-04-14 | 初稿     | 纵杜齐 |



------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [修改/新建/更新/删除](#修改)

------------

## 接口概览
| URL                                    | 方法 | 描述                         |
| -------------------------------------- | ---- | ---------------------------- |
| http://ip:port/api/ybind/v1.0/category | GET  | [获取](#获取)                |
| http://ip:port/api/ybind/v1.0/category | PUT  | [修改/新建/更新/删除](#修改) |

## 概述
* 语法：
```
logging {
      channel <string> {
              buffered <boolean>;
              file <quoted_string> [ versions ( unlimited | <integer> ) ]
                  [ size <size> ] [ suffix ( increment | timestamp ) ];
              null;
              print-category <boolean>;
              print-severity <boolean>;
              print-time ( iso8601 | iso8601-utc | local | <boolean> );
              severity <log_severity>;
              stderr;
              syslog [ <syslog_facility> ];
      };
      ...
      category <string> { <string>; ... };
      ...
};
```
* 概念：该logging语句为名称服务器配置了多种日志记录选项。
  
   ```
它的channel短语将输出方法，格式选项和严重性级别与一个名称相关联，然后可以将其与该category短语一起使用以选择如何记录各种消息类别。
   配置日志时，首先要定义通道，然后将不同的日志类别的数据指派到指定的通道上输出。
   ```
   
* category中的string必须为以下指定的项：
	
	具体的可参照BIND日志模块说明文档中的 1.2.2、Category语句中的bind9类别
	
* category中的{ }中的string为channel的名称或者定义为null（丢弃）

## 获取

- 请求URL：http://ip:port/api/ybind/v1.0/category
- HTTP方法：GET
- 请求参数：以querystring的方式携带

| 名称 | 类型   | 默认值 | 描述                                                      |
| :--- | :----- | :----- | :-------------------------------------------------------- |
| name | String | N/A    | **说明**：category 的名称<br>**缺省**：表示所有的category |

* returnBody：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| rcode*       | Int    | N/A    | 业务执行码                                                   |
| description* | String | N/A    | `rcode`的文字描述                                            |
| data         | Dict   | N/A    | **缺省**：业务执行失败<br>**Dict**：`name`缺省时表示所有的信息 |

### 返回码

| rcode | description           | 说明                           |
| ----- | --------------------- | ------------------------------ |
| 0     | Success               | 查询成功                       |
| 1     | Bad Parameter Format  | `name`格式错误                 |
| 404   | Not Found             | 没有找到`name`指定category配置 |
| 408   | Request Timeout       | 请求超时                       |
| 500   | Internal Server Error | 程序运行错误                   |

### 示例

* 现有策略：

```
logging {
      ...
      category client { client-default; };
      category config { config-default; };

};
```

#### 获取特定策略

* 请求示例：

```
GET http://ip:port/api/ybind/v1.0/category?name=client
BODY   :
```

* 返回示例：

```
{
    "rcode": 0,
    "description": "Success",
    "data": ["client-default"]
}

```

#### 获取全量策略

* 请求示例：

```
GET http://ip:port/api/ybind/v1.0/category
BODY   : 

```

* 返回示例：

```

{
    "rcode": 0,
    "description": "Success",
    "data": {
    "client":["client-default"],
    "config":["config-default"]
     }
}

```



## 修改/新建/更新/删除


- 请求URL：http://ip:port/api/ybind/v1.0/category
- HTTP方法：PUT
- 请求参数：以querystring的方式携带


| 名称 | 类型   | 默认值 | 描述                                                     |
| :--- | :----- | :----- | :------------------------------------------------------- |
| name | String | N/A    | **说明**：category的名称<br>**缺省**：表示覆盖所有的配置 |

* body：


| 名称 | 类型 | 默认值 | 描述                                                         |
| :--- | :--- | :----- | :----------------------------------------------------------- |
| N/A* | Dict | N/A    | **说明**：根据传入的类型更新指定`name`的配置或者覆盖所有配置 <br>**缺省**：当Dict缺省时删除name下的配置 |

* returnBody：


| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

#### 返回示例


* 成功返回

```
{
    "rcode": 0,
    "description": "Success"
}
```


### 修改


* 现有策略：

```
logging {
      ...
      category client { client-default; };
      category config { config-default; };

};

```

* 请求示例：
	
	PUT下当name、BODY存在，且原先name就有的情况下，会把name的策略修改

```
PUT http://ip:port/api/ybind/v1.0/category?name=client
BODY   :["config-default"]
```

* 修改后的策略：

```
logging {
      ...
      category client { config-default; };
      category config { config-default; };
};
```

#### 新建


* 请求示例：
	
	PUT下当name存在，BODY存在的情况下，会加到策略中

```
PUT http://ip:port/api/ybind/v1.0/category?name=dnstap
BODY   :["dnstap-default"]
```

* 新建后的策略：

```
logging {
      ...
      category client { client-default; };
      category config { config-default; };
      category dnstap { dnstap-default; };
};
```

#### 更新

* 请求示例：
	
	PUT下当name不存在，可以通过BODY，将所有的策略覆盖掉

```
PUT  http://ip:port/api/ybind/v1.0/category
BODY   :{
		"dnstap":["dnstap-default"]
}
```

* 更新后的策略：

```
logging {
      ...
      category dnstap { dnstap-default; };
};
```


#### 删除指定策略

* 请求示例：
	
	PUT下当name的存在，BODY不存在的情况下，会把指定的name下的策略删掉

```
PUT http://ip:port/api/ybind/v1.0/category?name=client
BODY   :{

}
```

* 删除指定策略后：

```
logging {
      ...
      category config { config-default; };
};
```

#### 删除所有策略

* 请求示例：
	
	PUT下当name、BODY不存在的情况下，会把name所有的策略删掉

```
PUT http://ip:port/api/ybind/v1.0/category
BODY   :{

}
```

* 删除所有策略后：

```
logging {
      ...
};
```