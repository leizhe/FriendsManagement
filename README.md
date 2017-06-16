

## 开发框架
* Go、[BeeGo](https://beego.me/)、SwaggerUI
* [PostgreSql](https://www.postgresql.org/)


## 项目背景
For any application with a need to build its own social network, "Friends Management" is a
common requirement which ussually starts off simple but can grow in complexity depending
on the application's use case.
Usually, applications would start with features like "Friend", "Unfriend", "Block", "Receive
Updates" etc.

## 如何运行
* 安装并配置[Go语言环境](http://www.runoob.com/go/go-environment.html)
* 安装[BeeGo](https://beego.me/)
```java
  go get github.com/astaxie/beego
``` 
* 安装[PostgreSql](https://www.postgresql.org/)
* 在PostgreSql建立数据库，然后在数据库中运行文件夹dbScript下的DB_Init.sql脚本初始化数据库
* 修改main.go文件的连接字符串
```java
  "postgres://[username]:[password]@[IP]:[PORT]/[DbName]?sslmode=disable"
``` 

* 在项目文件夹FriendsManagement下打开CMD，使用bee run命令运行

* 运行后在浏览器中输入地址：http://ip:port/swagger/

    如：[http://localhost:8080/swagger/](http://localhost:8080/swagger/)

## API说明

#### 1 /Firend/AddFriends   [`POST`] 

*Create a friend connection between two email addresses.*

*JSON request:*
```json
  {
    "friends":[
        "andy@example.com",
        "john@example.com"
     ]
  }
```
*JSON response on success:*
```json
  {
    "success": true
  }
```  

#### 2 /Firend/GetFriends    [`POST`] 

*Retrieve the friends list for an email address.*

*JSON request:*
```json
  {
    "email": "andy@example.com"
  }
```
*JSON response on success:*
```json
  {
     "success": true,
     "friends" : [
         "john@example.com"
     ],
     "count" : 1
  }
```  

#### 3 /Firend/GetCommonFriends    [`POST`] 

*Retrieve the common friends list between two email addresses.*

*JSON request:*
```json
  {
     "friends":[
         "andy@example.com",
         "john@example.com"
     ]
  }
```
*JSON response on success:*
```json
  {
     "success": true,
     "friends":[
         "common@example.com"
     ],
     "count" : 1
  }
```  

#### 4 /Subscribe/AddSubscribe    [`POST`] 

*Subscribe to updates from an email address.*

That "subscribing to updates" is NOT equivalent to "adding a friend connection".

*JSON request:*
```json
  {
      "requestor": "lisa@example.com",
      "target": "john@example.com"
  }
```
*JSON response on success:*
```json
  {
      "success": true
  }
```  

#### 5 /Subscribe/BlockSubscribe    [`POST`] 

*Block updates from an email address.*

*Suppose "andy@example.com" blocks "john@example.com":*
* if they are connected as friends, then "andy" will no longer receive notifications from "john"
* if they are not connected as friends, then no new friends connection can be added

*JSON request:*
```json
  {
      "requestor": "andy@example.com",
      "target": "john@example.com"
  }
```
*JSON response on success:*
```json
  {
      "success": true
  }
```  

#### 6 /Subscribe/RetrieveSubscribe   [`POST`] 

*Retrieve all email addresses that can receive updates from an email address.*

*Eligibility for receiving updates from i.e. "john@example.com":*
* has not blocked updates from "john@example.com", and
* at least one of the following:
    * has a friend connection with "john@example.com"
    * has subscribed to updates from "john@example.com"
    * has been @mentioned in the update

*JSON request:*
```json
  {
     "sender": "john@example.com",
     "text": "Hello World! kate@example.com"
  }
```
*JSON response on success:*
```json
  {
     "success": true,
     "recipients" : [
        "lisa@example.com",
        "kate@example.com"
     ]
  }
```  
