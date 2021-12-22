<!--
 * @Description: 关于go push
 * @Author: maxyang
 * @Date: 2021-04-25 00:20:22
 * @LastEditTime: 2021-12-22 15:34:58
 * @LastEditors: liutq
 * @Reference: 
-->
# Go push 

这是一个go语言开发的websocket项目。本人空余时间开发，欢迎大佬们指指点点

本项目可以支持消息推送，可以做即时通信聊天业务等，有空我会一直完善项目

## Running the Go push

1.首先你需要下载gopush
```
   go get github.com/maxyang107/gopush
```
2.进入gopush目录执行下列命令
```
   go build -o gopush ./
```
3.启动服务
```
   ./gopush
```

## 关于配置文件conf.json
   ```
   {
      "Name":           "GO PUSH SERVER",
      "Address":        "127.0.0.1",
      "Port":           8088,
      "WriteWait":      10,
      "PongWait":       60,
      "PingPeriod":     54,
      "MaxMessageSize": 1024
   }
   ```
   具体配置项说明在utils下config.go有详细说明，请前往查看配置

 ## 关于配置文件db.json
   ```
   {
      "DbHost": "127.0.0.1",
      "DbPort": 3306,
      "DbName": "gopush",
      "DbUser": "root",
      "DbPwd":  "root12345",
      "DbCharset": "utf8mb4",
      "DbParseTime": "True",
      "DbLoc": "Local",
      "DbPrefix": "",
      "DbType": "mysql"
   }
   ```  

具体配置项说明在utils下db.go有详细说明，请前往查看配置


### 前端连接websocket

1.连接地址
```
   ws://127.0.0.1:8088/ws
```

2.首次连接，客户端需要绑定客户端id，参数方式
```
   {"userId":"1234", "option":"bindUserId"}
```

3.绑定userId后，可以正常使用业务，其他业务方法调用方式同上述一致
```
   { "option":"xxxfunc","param1":"XXXXX","param2":"XXXXX"}
```

### http业务

gopush启动后会启动httpserver，访问地址：http://127.0.0.1:8088/

http服务主要提供管理websocket连接的功能，如创建分组，向分组广播信息。（应用场景如群聊，分组推送站内信等）

该方法增加了权限控制，请求参数必须携带SourceCode参数，服务端验证了SourceCode与远端请求地址Source是否符合授权，如果没有授权的请求会被拦截

调用方式：
http://127.0.0.1:8088?func=createGroup&groupName=卡尔的房间&SourceCode=02020

如此便创建好了一个房间，客户端携带groupname：卡尔的房间，即可加入该房间


### 关于http业务权限控制
权限设置接口
```
   http://127.0.0.1:8088/manager
```
请求方式POST

请求参数
```
   {
    "Source":"http://127.0.0.1:8088",//授权地址
    "SourceCode":21321,//授权码
    "SourceStatus":1,//授权状态
    "SourceTip":"优海"//授权平台
   }
```