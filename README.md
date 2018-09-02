# userManageSystem
# 简介
基于Beego开发的易用、易扩展、界面友好的轻量级功能权限管理系统。前端框架基于AdminLTE2进行资源整合，包含了多款优秀的插件，是笔者对多年后台管理系统开发经验精华的萃取。
本系统非常适合进行后台管理系统开发，统一的代码和交互给二次开发带来极大的方便，在没有前端工程师参与的情况下就可以进行快速的模块式开发，并保证用户使用的友好性和易用性。系统里整合了众多优秀的资源，在此感谢各位大神的无私奉献。
# 非原创
  框架来源于 https://github.com/beego/bee
本文博客
# 特点
1. finished CreateAccount(Register),Login and Logout of user interaction logic
2. This is only a demo

# 安装方法

本系统基于beego开发，默认使用mysql数据库，缓存redis 
1.安装golang环境（ 略）

2.安装本系统
```
go get github.com/Andy6666666666/userManageSystem
```

3.运行
在 userManageSystem 目录使用beego官方提供的命令运行
```
bee run
```
http: 在浏览器里打开 http://localhost:8080 进行访问
