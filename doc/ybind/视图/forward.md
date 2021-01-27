| 版本 | 日期       | 更改记录                                                     | 作者 |
| :--- | :--------- | :----------------------------------------------------------- | ---- |
| 0.1  | 2020-04-05 | 初稿                                                         | 程俊 |
| 0.2  | 2020-04-08 | 添加配置项的默认值提示<br>`GET`方法删除body部分<br>`POST`冲突返回409<br>`PUT`方法body不能缺省 | 程俊 |
| 0.3  | 2020-04-17 | 修改"mode"字段为"forward","algo"字段为"forward_algo";<br>删除DELETE接口,使用PUT空字段来删除; <br>增加基于指定zone的配置接口 | 韩冬 |
| 0.4  | 2020-04-22 | 规范化返回码                                                 | 程俊 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [修改](#修改)

------------

## 接口概览
| URL                                   | 方法 | 描述          |
| ------------------------------------- | ---- | ------------- |
| http://ip:port/api/ybind/v1.0/forward | GET  | [获取](#获取) |
| http://ip:port/api/ybind/v1.0/forward | PUT  | [修改](#修改) |
## 概述
* 语法：
```
forward ( first | only ) algo (weight | order | srtt );
```
* 概念：转发模式。`first`：转发出去如果是servfail或者超时，尝试迭代；`only`：转发出去无论结果是什么不会走迭代。
	
	```
	算法. weight: 会根据ip附带的权重来进行优先级的排序
	     order: 该情况下会根据ip附带的次序来进行优先级排序，次序值较小的ip,优先被使用
		 srtt: 会根据ip附带的srtt值进行排序，srtt值最小的ip将作为优选IP，优先被作为目标服务器
	```
	
* 注意项：
	
	* 可以在`option`中配置也可以在`view`中配置也可以在zone中配置，**但是目前只支持到zone的配置，而且该zone必须是forward zone**

## 获取

### URL
http://ip:port/api/ybind/v1.0/forward

### 方法
`GET`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| view | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示option<br>**举例**：__default |
| zone | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示option<br>**举例**：yamu.com |


* returnBody：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| rcode*       | Int    | N/A    | 业务执行码                                                   |
| description* | String | N/A    | `rcode`的文字描述                                            |
| data         | String | N/A    | **缺省**：业务执行失败<br>**dict**：`view`缺省时option下的配置或者指定`view`的策略 |

### 返回码
| rcode | description           | 说明                                                         |
| ----- | --------------------- | ------------------------------------------------------------ |
| 0     | Success               | 查询成功                                                     |
| 2     | Bad Parameter Value   | `name`或`body`值错误，或者options、view、zone中不支持该配置项 |
| 404   | Not Found             | 没有找到指定的配置                                           |
| 408   | Request Timeout       | 请求超时                                                     |
| 500   | Internal Server Error | 程序运行错误                                                 |

### 示例

#### 获取view下zone策略

* 现有策略：

```
view __default {
	zone "yamu.com" {
		type forward;
		forward only algo srtt;
		forwarders {8.8.8.8 weight 2; 114.114.114.114 weight 3;};
	};
};
```


* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/forward?view=__default&zone=yamu.com
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "mode": "only",
        "algo": "srtt"
    }
}
```



## 添加/修改/删除

### URL
http://ip:port/api/ybind/v1.0/forward

### 方法
`PUT`

### 参数
* queryString：

| 名称 | 类型   | 默认值 | 描述                                                         |
| :--- | :----- | :----- | :----------------------------------------------------------- |
| view | String | N/A    | **说明**：view的名称，用于定位到该条view<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示option<br>**举例**：__default |
| zone | String | N/A    | **说明**：zone的名称，用于定位到该条zone<br>**格式**：数字、大小写字母、-、_<br>**缺省**：表示option<br>**举例**：yamu.com |

* body：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| forward      | String | N/A    | **说明**：更新指定`view`的配置或者options或者zone的转发模式配置<br>**注意**：可以为空：""，删除指定`view`的配置或者option配置或者zone的配置 |
| forward_algo | String | N/A    | **说明**：更新指定`view`的配置或者options或者zone的算法配置<br>**注意**：可以为空：""，删除指定`view`的配置或者option配置或者zone的配置 |

* returnBody：

| 名称         | 类型   | 默认值 | 描述              |
| :----------- | :----- | :----- | :---------------- |
| rcode*       | Int    | N/A    | 业务执行码        |
| description* | String | N/A    | `rcode`的文字描述 |

### 返回码
| rcode | description           | 说明                          |
| ----- | --------------------- | ----------------------------- |
| 0     | Success               | 修改成功                      |
| 1     | Bad Parameter Format  | `queryString`或`body`格式错误 |
| 2     | Bad Parameter Value   | `queryString`或`body`值错误   |
| 408   | Request Timeout       | 请求超时                      |
| 500   | Internal Server Error | 程序运行错误                  |

### 示例

#### 添加/修改 view下zone策略

* 现有策略：

```
view __default {
	zone "yamu.com" {
		forward only algo order ;
	}
}
```

* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/forward?view=__default&zone=yamu.com
BODY   : {
        "mode": "only",
        "algo": "srtt"
    }
```

* 返回：
```
{
    "rcode": 0,
    "description": "Success"
}
```

* 接口调用成功后策略：

```
view __default {
	zone "yamu.com" {
		forward only algo srtt;
	}
}
```


#### 删除 view下zone策略

* 现有策略：

```
view __default {
	zone "yamu.com" {
		forward only algo order ;
	}
}
```

* 请求：
```
METHOD : PUT
URL    : http://ip:port/api/ybind/v1.0/forward?view=__default&zone=yamu.com
BODY   : {}
```

* 返回：
```
{
    "rcode": 0,
    "description": "Success"
}
```

* 接口调用成功后策略：

```
option {
}
view __default {
	zone "yamu.com" {
	}
}
```