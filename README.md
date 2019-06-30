# mqtt_to_mysql是mqtt消息存储到mysql数据库

- 是一个订阅mqtt所有消息，并且将消息存储到mysql数据库的脚本，使用go语言编写
- mqtt和mysql连接信息使用配置文件config.txt
- 两个zip压缩包为相关依赖，加压到同一目录即可使用

# 首先创建数据库
```
create database nulige character set utf8;
```

# 创建表
```
create table user_info(id int(11) not null AUTO_INCREMENT, username varchar(40), departname varchar(100), create_time timestamp , primary key(id));
```

# 版本
v1.0
!(下载地址)[https://github.com/haitianjingyu/mqtt_to_mysql/blob/master/version/emqttv1.0.zip]

# 修改相应配置文件


mysql和emqq相应地址，端口，用户名，密码等。

```
[mysql]
mysqlurl=root:root@tcp(192.168.99.108:3306)/nulige?charset=utf8&parseTime=true&loc=Local
maxopenconns=10
maxidleconns=10

[emqq]
emqqurl=tcp://192.168.99.111:1883
emqquser=mysql
emqqpasswd=mysql
```
