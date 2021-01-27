| 版本 | 日期       | 更改记录                                                     | 作者 |
| :--- | :--------- | :----------------------------------------------------------- | ---- |
| 0.1  | 2020-04-06 | 初稿                                                         | 叶豪 |

------------

* [接口概览](#接口概览)
* [概述](#概述)
* [获取](#获取)
* [修改](#修改)

------------

## 接口概览
| URL                                       | 方法 | 描述          |
| ----------------------------------------- | ---- | ------------- |
| http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1 | GET  | [获取完整数据](#获取) |
| http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=status | PUT  | [获取完整统计数据子集](#修改) |
| http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=server| GET  | [获取服务器和解析器的统计数据子集](#获取) |
| http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=zones | GET  | [获取域的统计信息子集](#获取) |
| http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=net | GET  | [获取网络状态和套接字统计信息](#获取) |
| http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=mem | GET  | [获取内存管理统计信息](#获取) |
| http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=tasks | GET  | [获取任务管理器统计信息](#获取) |
| http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=traffic | GET  | [获取流量大小统计信息](#获取) |

## 概述
* 语法：
```
statistics-channels配置项，已经写入到了配置文件里面(仅支持Ipv4地址)；
默认配置为：
statistics-channels {
    inet * port 9001
        allow {any;};
};
映射出一个获取bind统计相关信息的get接口
http://ip:port/api/ybind/v1.0/statistics?server=&part=
注：server为必选参数。
```

* 概念：声明了系统管理员用来访问名称服务器统计信息的通信通道
		

## 获取

### URL
http://ip:port/api/ybind/v1.0/statistics
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
| data         | Array  | N/A    | **缺省**：业务执行失败<br>**Array**                          |

### 返回码
| rcode | description           | 说明                                     |
| ----- | --------------------- | ---------------------------------------- |
| 0     | Success               | 查询成功                                 |
| 404   | Not Found             | 没有找到`statistics-channels`指定的配置  |
| 408   | Request Timeout       | 请求超时                                 |
| 500   | Internal Server Error | 程序运行错误                             |

### 示例


#### 获取完整数据
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "boot-time": "2020-06-01T09:45:38.733Z",
        "config-time": "2020-06-02T02:28:28.453Z",
        "current-time": "2020-06-02T08:04:56.724Z",
        "json-stats-version": "1.5",
        "memory": {
            "BlockSize": 4718592,
            "ContextSize": 46844744,
            "InUse": 8836716,
            "Lost": 0,
            "Malloced": 59489708,
            "TotalUse": 322265593,
            "contexts": [
                {
                    "blocksize": 2359296,
                    "hiwater": 0,
                    "id": "0xce9210",
                    "inuse": 2796436,
                    "lowater": 0,
                    "malloced": 4340900,
                    "maxinuse": 4710228,
                    "maxmalloced": 5267292,
                    "name": "main",
                    "pools": 39,
                    "references": 1397,
                    "total": 148046076
                },
```

#### 获取完整统计数据子集
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=status
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "boot-time": "2020-06-01T09:45:38.733Z",
        "config-time": "2020-06-02T02:28:28.453Z",
        "current-time": "2020-06-02T08:08:28.778Z",
        "json-stats-version": "1.5",
        "version": "9.16.0"
    }
}
```

#### 获取服务器和解析器的统计数据子集
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=server
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "boot-time": "2020-06-01T09:45:38.733Z",
        "config-time": "2020-06-02T02:28:28.453Z",
        "current-time": "2020-06-02T08:09:47.541Z",
        "json-stats-version": "1.5",
        "nsstats": {
            "QryAuthAns": 89,
            "QrySuccess": 89,
            "QryUDP": 89,
            "Requestv4": 89,
            "Response": 89
        },
        "opcodes": {
            "IQUERY": 0,
            "NOTIFY": 0,
            "QUERY": 89,
            "RESERVED10": 0,
            "RESERVED11": 0,
            "RESERVED12": 0,
            "RESERVED13": 0,
            "RESERVED14": 0,
            "RESERVED15": 0,
            "RESERVED3": 0,
            "RESERVED6": 0,
            "RESERVED7": 0,
            "RESERVED8": 0,
            "RESERVED9": 0,
            "STATUS": 0,
            "UPDATE": 0
        },
```

#### 获取域的统计信息子集
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=zones
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "boot-time": "2020-06-01T09:45:38.733Z",
        "config-time": "2020-06-02T02:28:28.453Z",
        "current-time": "2020-06-02T08:11:19.470Z",
        "json-stats-version": "1.5",
        "version": "9.16.0",
        "views": {
            "__default": {
                "zones": [
                    {
                        "class": "IN",
                        "name": "EMPTY.AS112.ARPA",
                        "serial": 0,
                        "type": "builtin"
                    },
                    {
                        "class": "IN",
                        "name": "HOME.ARPA",
                        "serial": 0,
                        "type": "builtin"
                    },
                    {
                        "class": "IN",
                        "name": "0.IN-ADDR.ARPA",
                        "serial": 0,
                        "type": "builtin"
                    },
```

#### 获取完整统计数据子集
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=net
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "boot-time": "2020-06-01T09:45:38.733Z",
        "config-time": "2020-06-02T02:28:28.453Z",
        "current-time": "2020-06-02T08:12:16.430Z",
        "json-stats-version": "1.5",
        "socketmgr": {
            "sockets": [
                {
                    "id": "0x7f45a9710178",
                    "local-address": "<unknown address, family 16>",
                    "references": 1,
                    "states": [
                        "bound"
                    ],
                    "type": "not-initialized"
                },
                {
                    "id": "0x7f45a97102e0",
                    "local-address": "0.0.0.0#9001",
                    "name": "statchannel",
                    "references": 1,
                    "states": [
                        "listener",
                        "bound"
                    ],
                    "type": "tcp"
                },
                {
                    "id": "0x7f45a9686718",
                    "local-address": "127.0.0.1#953",
                    "name": "control",
                    "references": 1,
                    "states": [
                        "listener",
                        "bound"
                    ],
                    "type": "tcp"
                },
```

#### 获取完整统计数据子集
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=mem
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "boot-time": "2020-06-01T09:45:38.733Z",
        "config-time": "2020-06-02T02:28:28.453Z",
        "current-time": "2020-06-02T08:13:32.868Z",
        "json-stats-version": "1.5",
        "memory": {
            "BlockSize": 4718592,
            "ContextSize": 46844744,
            "InUse": 8836716,
            "Lost": 0,
            "Malloced": 59489708,
            "TotalUse": 323024057,
            "contexts": [
                {
                    "blocksize": 2359296,
                    "hiwater": 0,
                    "id": "0xce9210",
                    "inuse": 2796436,
                    "lowater": 0,
                    "malloced": 4340900,
                    "maxinuse": 4710228,
                    "maxmalloced": 5267292,
                    "name": "main",
                    "pools": 39,
                    "references": 1397,
                    "total": 148804540
                },
```

#### 获取完整统计数据子集
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=tasks
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "boot-time": "2020-06-01T09:45:38.733Z",
        "config-time": "2020-06-02T02:28:28.453Z",
        "current-time": "2020-06-02T08:14:14.093Z",
        "json-stats-version": "1.5",
        "taskmgr": {
            "default-quantum": 25,
            "tasks": [
                {
                    "events": 0,
                    "id": "0x7f45a970a010",
                    "name": "server",
                    "quantum": 25,
                    "references": 16,
                    "state": "idle"
                },
                {
                    "events": 0,
                    "id": "0x7f45a970a0e8",
                    "name": "zmgr",
                    "quantum": 1,
                    "references": 5,
                    "state": "idle"
                },
```

#### 获取完整统计数据子集
* 请求：
```
METHOD : GET
URL    : http://ip:port/api/ybind/v1.0/statistics?server=127.0.0.1&part=traffic
BODY   :
```
* 返回：
```
{
    "rcode": 0,
    "description": "Success",
    "data": {
        "boot-time": "2020-06-01T09:45:38.733Z",
        "config-time": "2020-06-02T02:28:28.453Z",
        "current-time": "2020-06-02T08:14:44.968Z",
        "json-stats-version": "1.5",
        "traffic": {
            "dns-tcp-requests-sizes-received-ipv4": {},
            "dns-tcp-requests-sizes-received-ipv6": {},
            "dns-tcp-responses-sizes-sent-ipv4": {},
            "dns-tcp-responses-sizes-sent-ipv6": {},
            "dns-udp-requests-sizes-received-ipv4": {
                "16-31": 89
            },
            "dns-udp-requests-sizes-received-ipv6": {},
            "dns-udp-responses-sizes-sent-ipv4": {
                "32-47": 89
            },
            "dns-udp-responses-sizes-sent-ipv6": {}
        },
        "version": "9.16.0"
    }
}
```