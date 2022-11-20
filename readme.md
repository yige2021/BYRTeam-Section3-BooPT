## BYR Team2022实践组考核题目3——电子书预览与发布系统（前端）
### 一、写在前面

由于我将过多的时间与精力花在尝试看懂后端代码、尝试配置环境与各种连接，以及解决各种莫名其妙的bug……（大白话：自己太菜），只实现登录注册及退出功能，其他的只画了一点界面出来……下面是我给这个项目写的文档，里面包含着我对这个项目的粗浅理解~

### 二、技术栈 

前端框架：Vue.js

前端请求库：axios
 
前端组件库：Element UI

数据库：PostgreSQL

用户认证协议：LDAP

后端：Go （Web框架：Fiber）

调试工具：Postman、Edge浏览器

### 三、目录结构

BooPT：后端
Boopt_frontend：前端

### 四、后端相关

关于对后端的理解，我在后端的源代码加了注释。

在前端发送HTTP请求时，产生了跨域问题，为此我在后端加入了Fiber CORS中间件，其他方面未做改动，应该可以吧……

### 五、开发进程记录

##### 2022年10月27日：
正式开始着手弄这个项目，先从GitHub上把后端clone了.
##### 2022年10月28日：
阅读`main.go`文件，`config.go`文件以及`database.go`文件，摸清了连接数据库的过程.
##### 2022年10月29日：
看PostgreSQL和gorm，开始尝试实现连接数据库.
##### 2022年11月3日：
成功连接到数据库并创建表；研究fiber.
##### 2022年11月5日：
到登录这一步了，学习http请求，有所进展。然而卡在了LDAP认证这一步，389端口连不上去.
##### 2022年11月6日：
发现我还没有在本地搭建LDAP服务（……）安装了OpenLDAP、LDAP Browser，建立连接之后，成功让后端通过了LDAP认证.
##### 2022年11月7日：
成功实现了登录注册功能（还没有界面），并安装了node.js、vue等.
##### 2022年11月8-9日：
CSAPP课有Bomblab，停工.
##### 2022年11月10日：
安装elementUI，开始画界面.
##### 2022年11月11日：
登录注册界面画好了，开始测试登录功能.尝试用前端向后端发送请求，使用axois请求库。然而遇到跨域问题：前端localhost:8080无法连接服务端localhost:3000.
##### 2022年11月12日：
尝试从前端解决未果，转而从后端下手：查阅fiber，调用CORS中间件，解决了跨域问题.然而还是无法成功实现登录，发送请求后报错400 bad request.
##### 2022年11月13日：
检查请求头Content-Type、负载格式等无误，但还是不行.最后发现虽然负载格式表面正确，但在浏览器里查看负载，发现和postman里成功发送的负载的格式不同.最后用param替换data，顺利实现登录,顺便把登录注册功能实现了.
##### 2022年11月14日：
把前端目录结构改善了一下，然后完善了登录注册功能，做不完了，草草画了画界面，做了注销功能，收尾.