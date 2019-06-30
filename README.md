# mqtt_to_mysql是mtqq消息存储到mysql数据库

- 是一个订阅mtqq所有消息，并且将消息存储到mysql数据库的脚本，使用go语言编写
- mtqq和mysql连接信息使用配置文件config.txt
- 两个zip压缩包为相关依赖，加压到同一目录即可使用

# 首先创建数据库
```
create database nulige character set utf8;
```

# 创建表
```
create table user_info(id int(11) not null AUTO_INCREMENT, username varchar(40), departname varchar(100), create_time timestamp , primary key(id));
```
