| 版本 | 日期       | 更改记录                                                     | 作者 |
| :--- | :--------- | :----------------------------------------------------------- | ---- |
| 0.1  | 2020-04-05 | 初稿                                                         | 叶豪 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [修改](#修改)

------------

## 接口概览
| URL                                       | 方法 | 描述          |
| ----------------------------------------- | ---- | ------------- |
| http://ip:port/api/ybind/v1.0/memstatistics-file | GET  | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/memstatistics-file  | PUT  | [修改](#修改) |

## 概述
* 语法：
```
memstatistics-file <quoted_string>;
```
* 概念：退出时服务器写入内存使用统计信息的文件的路径名。如果未指定，则默认 	named.memstats
* 注意项：
	* 只在`option`中配置

## 获取

### URL
http://ip:port/api/ybind/v1.0/memstatistics-file

### 方法
`GET`

### 参数
* queryString：
空

* returnBody：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| rcode*       | Int    | N/A    | 业务执行码                                                   |
| description* | String | N/A    | `rcode`的文字描述                                            |
| data         | Array  | N/A    | **缺省**：业务执行失败<br>**Array**：option中没有statistics-file配置|

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
option {
	memstatistics-file "/var/named_mem_stats.txt";
}
```
#### 获取option策略
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/memstatistics-file
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": "/var/named_mem_stats.txt"
}
```

## 修改

### URL
http://ip:port/api/ybind/v1.0/memstatistics-file

### 方法
`PUT`

* body：

| 名称 | 类型  | 默认值 | 描述                                                         |
| :--- | :---- | :----- | :----------------------------------------------------------- |
| N/A* | String | N/A    | **说明**：更新`option`的配置<br>**注意**：可以为空：[]，删除`option`的配置|

* returnBody：

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

### 示例
* 现有策略：

```
option {
	memstatistics-file "/var/named_mem_stats.txt";
}
```

#### 更新option策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/memstatistics-file
BODY   : "/var/mem_stats.txt"
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
option {
	memstatistics-file "/var/mem_stats.txt";
}
```

#### 删除option策略
* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/memstatistics-file
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
option {
}
```