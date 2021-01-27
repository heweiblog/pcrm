| 版本 | 日期       | 更改记录 | 作者   |
| :--- | :--------- | :------- | ------ |
| 0.1  | 2020-07-21 | 初稿     | 邱士超 |

---

* [接口概述](#接口概述)
* [概述](#概述)
* [获取](#获取)
* [修改/新建/更新/删除](#修改/新建/更新/删除)

# 接口概述

| URL                                | 方法 | 描述                                        |
| ---------------------------------- | ---- | ------------------------------------------- |
| http://ip:port/api/ybind/v1.0/tsig | GET  | [获取](#获取)                               |
| http://ip:port/api/ybind/v1.0/tsig | PUT  | [修改/新建/更新/删除](#修改/新建/更新/删除) |

# 概述

* TSIG概念

  Bind 引入的一种新机制来保护DNS消息的安全性。可以使用TSIG保护查询、响应、区域传送和动态更新。

* 语法

  通过`key`定义TSIG使用的共享秘钥。

  ```
  key <string> {
        algorithm <string>;
        secret <string>;
  };
  ```

  其中：

  `algorithm`：算法定义；名称服务器支持的算法有`hmac-md5`, `hmac-sha1`, `hmac-sha224`, `hmac-sha256`, `hmac-sha384`, and `hmac-sha512`

  `secret`：算法使用的秘钥，base64编码的字符串格式。

  > **使用**：
  >
  > //本接口只是配置tsig密钥，不涉及tsig的使用，此处作为了解
  >
  > ​		在配置完名称服务器TSIG秘钥，可以使用TSIG保护查询、响应、区域传送和动态更新。配置的关键是`server`语句和`keys`语句，它告诉名称服务器对发送给特定的远端名称服务器的查询和区域传送数据进行签名。
  >
  > ```
  > server 192.249.249.1 {
  > 	keys {toystory-wormhole.movie.end.;}
  > }
  > zone "movie.edu" {
  > 	type slave;
  > 	file "bak.movie.edu"
  > }
  > ```
  >
  > *server子语句告诉本地名称服务器wormhole.movie.edu,使用秘钥toystory-wormhole.movie.end.对发往192.249.249.1(toystory.movie.end)地址的请求进行签名*
  >
  > ​		若仅关注区域传送，可以在masters 子语句中指定对所有slave区域进行传送的密钥
  >
  > ```
  > zone "movie.edu" {
  > 	type slave;
  > 	masters {192.249.249.1 key toystory-wormhole.movie.end.;}
  > 	file "bak.movie.edu"
  > }
  > ```
  >
  > 在toystory.movie.edu 上，限制区域传送，使用密钥进行签名
  >
  > ```
  > zone "movie.edu" {
  > 	type master;
  > 	file "bak.movie.edu"
  > 	allow-transfer {key toystory-wormhole.movie.end.;}
  > }
  > ```
  >
  > ​		还可以通过TSIG的`allow-update`和`update-policy`子语句来限制动态更新

# 获取

## url

http://ip:port/api/ybind/v1.0/tsig

## 方法

`GET`

## 参数

* queryString：

| 名称 | 类型   | 默认值 | 描述                                                     |
| :--- | :----- | :----- | :------------------------------------------------------- |
| name | String | N/A    | **说明**：tsig秘钥的名称<br>**缺省**：表示所有的tsig秘钥 |

* returnBody：

| 名称         | 类型   | 默认值 | 描述                                                         |
| :----------- | :----- | :----- | :----------------------------------------------------------- |
| rcode*       | Int    | N/A    | 业务执行码                                                   |
| description* | String | N/A    | `rcode`的文字描述                                            |
| data         | Dict   | N/A    | **缺省**：业务执行失败<br>**Dict**：`name`缺省时表示所有的信息 |

## 返回码

| rcode | description           | 说明                         |
| ----- | --------------------- | ---------------------------- |
| 0     | Success               | 查询成功                     |
| 1     | Bad Parameter Format  | `name`格式错误               |
| 404   | Not Found             | 没有找到`name`指定的密钥配置 |
| 408   | Request Timeout       | 请求超时                     |
| 500   | Internal Server Error | 程序运行错误                 |

## 示例

### 获取指定密钥

* 请求

  ```
  METHOD : GET
  URL    : http://ip:port/api/ybind/v1.0/tsig?name=host1-host2.
  BODY   :
  ```

* 返回

  ```
  {
      "rcode": 0,
      "description": "Success",
      "data": {
         	"algorithm":"hmac-md5",
         	"secret":"X4Ob8mBlDfzdQM2QrAfxhA=="
  	}
  }
  ```

### 获取全部密钥

* 请求

  ```
  METHOD : GET
  URL    : http://ip:port/api/ybind/v1.0/tsig
  BODY   :
  ```

* 返回

  ```
  {
      "rcode": 0,
      "description": "Success",
      "data": {
      	"host1-host2.":{
         		"algorithm":"hmac-md5",
         		"secret":"X4Ob8mBlDfzdQM2QrAfxhA=="
      	},
      	"xxxx" :{
      		"algorithm":"hmac-md5",
         		"secret":"X4Ob8mBlDfzdQM2QrAfxhA=="
      	}
  	}
  }
  ```

  

# 修改/新建/更新/删除

## url

http://ip:port/api/ybind/v1.0/tsig

## 方法

`PUT`

## 参数

* queryString

  | 名称 | 类型   | 默认值 | 描述                                           |
  | ---- | ------ | ------ | :--------------------------------------------- |
  | name | string | N/A    | 说明：tsig密钥的名称<br />缺省：覆盖所有的配置 |

* body

  | 名称 | 类型 | 默认值 | 描述                                                         |
  | ---- | ---- | ------ | ------------------------------------------------------------ |
  | N/A* | Dict | N/A    | 说明：修改或覆盖指定`name`的配置或者覆盖所有的配置<br />包含字段：<br />secret：base64编码的秘钥<br />algorithm：加密类型<br />限制条件：secret与algorithm要么同时赋值，代表更新/修改/删除配置；要么同时为空，代表删除配置 |

  

* returnBody

  | 名称         | 类型   | 默认值 | 描述              |
  | :----------- | :----- | :----- | :---------------- |
  | rcode*       | Int    | N/A    | 业务执行码        |
  | description* | String | N/A    | `rcode`的文字描述 |

## 返回码

| rcode | description           | 说明                   |
| ----- | --------------------- | ---------------------- |
| 0     | Success               | 新建/修改成功          |
| 1     | Bad Parameter Format  | `name`或`body`格式错误 |
| 408   | Request Timeout       | 请求超时               |
| 500   | Internal Server Error | 程序运行错误##         |

## 示例

### 新建密钥

* 请求

  ```
  METHOD : PUT
  URL    : http://ip:port/api/ybind/v1.0/tsig?name=host1-host2.
  BODY   :{
  		"algorithm":"hmac-md5",
         	"secret":"X4Ob8mBlDfzdQM2QrAfxhA==",
  }
  ```

* 返回

  ```
  {
      "rcode": 0,
      "description": "Success"
  }
  ```

### 更新指定密钥

 * 请求

   ```
   METHOD : PUT
   URL    : http://ip:port/api/ybind/v1.0/tsig?name=host1-host2.
   BODY   :{
   		"algorithm":"hmac-md5",
          	"secret":"X4Ob8mBlDfzdQM2QrAfxhA==",
   }
   ```

* 返回

  ```
  {
      "rcode": 0,
      "description": "Success"
  }
  ```

### 修改指定密钥

 * 请求

   ```
   METHOD : PUT
   URL    : http://ip:port/api/ybind/v1.0/tsig?name=host1-host2.
   BODY   :{
   		"algorithm":"hmac-md5",
          	"secret":"X4Ob8mBlDfzdQM2QrAfxhA==",
   }
   ```

* 返回

  ```
  {
      "rcode": 0,
      "description": "Success"
  }
  ```

### 删除指定密钥

 * 请求

   ```
   METHOD : PUT
   URL    : http://ip:port/api/ybind/v1.0/tsig?name=host1-host2.
   BODY   :
   ```

* 返回

  ```
  {
      "rcode": 0,
      "description": "Success"
  }
  ```

  