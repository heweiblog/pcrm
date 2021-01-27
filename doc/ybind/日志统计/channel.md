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
| URL                                   | 方法 | 描述                         |
| ------------------------------------- | ---- | ---------------------------- |
| http://ip:port/api/ybind/v1.0/channel | GET  | [获取](#获取)                |
| http://ip:port/api/ybind/v1.0/channel | PUT  | [修改/新建/更新/删除](#修改) |


## 概述
* 语法：
```
logging {
      category <string> { <string>; ... };
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
};
```
* 概念：该logging语句为名称服务器配置了多种日志记录选项。
       
  
   ```
   它的channel短语将输出方法，格式选项和严重性级别与一个名称相关联，然后可以将其与该category短语一起使用以选择如何记录各种消息类别。
   配置日志时，首先要定义通道，然后将不同的日志类别的数据指派到指定的通道上输出。
   ```

## 获取

### URL
http://ip:port/api/ybind/v1.0/channel

### 方法
`GET`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                   |
| :--- | :----- | :----- | :----------------------------------------------------- |
| name | String | N/A    | **说明**：channel的名称<br>**缺省**：表示所有的channel |

* returnBody：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| rcode*       | Int    | N/A    | 业务执行码                                                   |
| description* | String | N/A    | `rcode`的文字描述                                            |
| data         | Dict   | N/A    | **缺省**：业务执行失败<br>**Dict**：`name`缺省时表示所有的信息 |

### 返回码
| rcode | description           | 说明                            |
| ----- | --------------------- | ------------------------------- |
| 0     | Success               | 查询成功                        |
| 1     | Bad Parameter Format  | `name`格式错误                  |
| 404   | Not Found             | 没有找到`name`指定的channel配置 |
| 408   | Request Timeout       | 请求超时                        |
| 500   | Internal Server Error | 程序运行错误                    |

### 示例
* 现有策略：

```
logging {
      channel client-default{
              file "/var/log/ybind/client.log";
              print-category yes;
              print-severity yes;
              print-time yes;
              severity info;
      };
     channel config-default{
              file "/var/log/ybind/config.log";
              print-category yes;
              print-severity yes;
              print-time yes;
              severity info;
      };
      category client { client-default; };
      category config { config-default; };

};
```

#### 获取特定策略
* 请求：

```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/channel?name=client-default
BODY   :
```

* 返回：

```
{
    "rcode": 0,
    "description": "Success",
    "data": {
         "file":"/var/log/ybind/client.log",
         "print-category":true,
         "print-severity": true,
         "print-time": true,
         "severity":"info"
        }
}
```

#### 获取全量策略
* 请求：

```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/channel
BODY   : 
```
* 返回：

```
{
    "rcode": 0,
    "description": "Success",
    "data": {
    "client-default":{
                      "file":"/var/log/ybind/client.log",
                       "print-category":true,
                       "print-severity": true,
                       "print-time": true,
                       "severity":"info"
                     },
     "config-default":{
                       "file":"/var/log/ybind/config.log",
                       "print-category":true,
                       "print-severity": true,
                       "print-time": true,
                       "severity":"info"
                      }
    }
}
```



## 修改/新建/更新/删除

### URL
http://ip:port/api/ybind/v1.0/channel

### 方法
`PUT`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                |
| :--- | :----- | :----- | :-------------------------------------------------- |
| name | String | N/A    | **说明**：channel的名称<br>**缺省**：覆盖所有的配置 |

* body：

| 名称 | 类型 | 默认值 | 描述                                                         |
| :--- | :--- | :----- | :----------------------------------------------------------- |
| N/A* | Dict | N/A    | **说明**：根据传入的类型修改指定`name`的配置或者覆盖所有配置 <br/> **缺省**：当Dict缺省时删除name下的配置 |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description           | 说明                   |
| ----- | --------------------- | ---------------------- |
| 0     | Success               | 新建/修改成功          |
| 1     | Bad Parameter Format  | `name`或`body`格式错误 |
| 408   | Request Timeout       | 请求超时               |
| 500   | Internal Server Error | 程序运行错误           |

### 示例
* 现有策略：

```
logging {
      channel client-default{
              file "/var/log/ybind/client.log";
              print-category yes;
              print-severity yes;
              print-time yes;
              severity info;
      };
     channel config-default{
              file "/var/log/ybind/config.log";
              print-category yes;
              print-severity yes;
              print-time yes;
              severity info;
      };
      category client { client-default; };
      category config { config-default; };
};
```

#### 修改特定策略
* 请求：

```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/channel?name=client-default
BODY   :{
	"file":"/var/log/ybind/config.log",
	"severity":"debug"
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
logging {
      channel client-default{
              file "/var/log/ybind/client.log";
              severity debug;
      };
      channel config-default{
              file "/var/log/ybind/config.log";
              print-category yes;
              print-severity yes;
              print-time yes;
              severity info;
      };
      category client { client-default; };
      category config { config-default; };
};
```

#### 新建
* 请求：

```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/channel?name=dnstap-default
BODY   :{
			"file":"/var/log/ybind/dnstap.log",
			"print-category":true,
			"print-severity": true,
			"print-time": true,
			"severity":"info"
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
logging {
 	  channel client-default{
              file "/var/log/ybind/client.log";
              print-category yes;
              print-severity yes;
              print-time yes;
              severity info;
      };
      channel config-default{
              file "/var/log/ybind/config.log";
              print-category yes;
              print-severity yes;
              print-time yes;
              severity info;
      };
      channel dnstap-default{
              file "/var/log/ybind/dnstap.log";
              print-category yes;
              print-severity yes;
              print-time yes;
              severity info;
      };
      category client { client-default; };
      category config { config-default; };

};
```

#### 更新
* 请求：

```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/channel
BODY   :{
		"dnstap":{
		"file":"/var/log/ybind/dnstap.log",
		"print-category":true,
		"print-severity": true,
		"print-time": true,
		"severity":"info"
		}

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
logging {
      channel dnstap-default{
              file "/var/log/ybind/dnstap.log";
              print-category yes;
              print-severity yes;
              print-time yes;
              severity info;
      };
      category client { client-default; };

};
```

#### 删除指定策略
* 请求：

```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/channel?name=client
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
logging {
     channel config-default{
              file "/var/log/ybind/config.log";
              print-category yes;
              print-severity yes;
              print-time yes;
              severity info;
      };
      category client { client-default; };
      category config { config-default; };
};
```