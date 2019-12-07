# monitor
用户行为记录\PV,UV数据接口

## 安装依赖并启动

```shell script
go mod tidy
go mod download
go run main.go
```

## 如何团队项目保持同步(重要)

([附上IDEA可视化操作](https://blog.csdn.net/autfish/article/details/52513465))

第一次时需要,与团队仓库建立联系

```shell script
git remote add upstream https://github.com/dgut-group-ten/monitor.git
```

工作前后要运行这几条命令,和团队项目保持同步

```shell script
git fetch upstream
git merge upstream/master
```

## 参考文章

- [IDEA sql自动补全/sql自动提示/sql列名提示](https://www.cnblogs.com/jpfss/p/11051015.html)