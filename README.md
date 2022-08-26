# CRIU介绍

CRIU运行在linux操作系统上的一个软件工具，
其功能是在用户空间实现Checkpoint/Restore功能。
使用这个工具，你可以冻结一个正在运行的程序，
并且checkpoint它到一系列的文件，
然后你就可以使用这些文件在任何主机重新恢复这个程序到被冻结的那个点
(白话就是实现对已运行程序的备份和恢复)。
所以criu通常被用在程序或者容器的热迁移、快照、远程调试等。


官 网：http://criu.org
源码地址：https://github.com/checkpoint-restore/criu

# 项目介绍
criu-server 可通过http请求实现对系统的冻结/恢复功能，无须再使用命令
具体参考swagger,当前系统使用criu并非criu-ns 所以多次对同一程序冻结再恢复会产生失败问题

如果需要解决命名空间问题，可修改criu-server 源码，用exec.Command("criu-ns", "dump")的方式来实现

# criu安装
https://protocode.top/pages/1535e3/

# criu-server
打包之后直接 sudo ./criu 既可