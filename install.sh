#!/bin/sh

# 根据thrift接口定义文件自动生成目标代码
thrift -out src -r --gen go weixinsender.thrift

# 编译安装weixinsender-server服务程序
go install weixinsender-server

