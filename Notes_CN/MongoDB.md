# MongoDB

> 版本：v4.49   5.2之后才会稳定

这是一个文档数据库（以JSON为数据模型），由C++语言编写，旨在为WEB应用提供可拓展的高性能数据存储解决方案。
这是一个介于关系数据库和非关系数据库之间的产品，支持多文档分片事物（最像关系型数据库）
项目迭代，MySQL需要变更表结构，但MongoDB不用,十分灵活
半结构化：在一个集合中，文档所拥有的字段并不需要是相同的，而且也不需要对所用的字段进行声明，非常适合于面向对象的编程模型

## 技术优势

- 基于灵活的JSON文档模型，非常适合做敏捷式的快速开发
  - 可动态增加新字段，加上version区别版本变化
- 高可用
  - 自恢复
  - 多中心容灾能力
  - 滚动服务 - 最小化服务终端
- 横向扩展-分片扩容

## 应用场景

> 目前的具体业务，通常为微服务架构，每个服务都可以用不同的数据库

- 游戏数据：方便查询与更新
- 物流：状态在不断变换
- 物联网场景......



## 环境安装

Community Server v4.4

```shell
tar -xvf mongodb-linux-x86_64-ubuntu2004-4.4.28.tgz
```

## 启动

### 命令启动

```shell
mkdir -p /mongodb/data /mongodb/log
bin/mongod --port=27017 --dbpath=/mongodb/data --logpath=/mongodb/log/mongodb.log --bind_ip=0.0.0.0 --fork
```

这里注意要用sudo进行启动！
![image.png](https://cdn.nlark.com/yuque/0/2024/png/40368069/1708737878694-1e937a46-319d-4b6d-bfac-ecb4c65017b8.png#averageHue=%23300a25&clientId=u8519a271-a58e-4&from=paste&height=78&id=u0374d33c&originHeight=97&originWidth=1716&originalType=binary&ratio=1.25&rotation=0&showTitle=false&size=29325&status=done&style=none&taskId=u67385af9-eb3b-46b2-b686-e50120a848f&title=&width=1372.8)
![image.png](https://cdn.nlark.com/yuque/0/2024/png/40368069/1708738040212-245bbb78-d0a2-4a12-8c12-a366ccd18707.png#averageHue=%23300a24&clientId=u8519a271-a58e-4&from=paste&height=423&id=ue6764f59&originHeight=529&originWidth=1705&originalType=binary&ratio=1.25&rotation=0&showTitle=false&size=122748&status=done&style=none&taskId=ub19ff17a-fbf8-4c69-ba4c-fe356bc4f29&title=&width=1364)

### 配置文件启动

```
systemLog:
   destination: file
   path: /mongodb/log/mongodb.log
   logAppend: true
storage:
   journal:
      enabled: true
   dbPath: /mongodb/data/
   engine: wiredTiger
processManagement:
   fork: true
net:
   port: 27017
   bindIp: 0.0.0.0
```

## 添加环境变量

```
vim /etc/profile


export MONGODB_HOME=/usr/local/soft/mongodb
PATH=$PATH:$MONGODB_HOME/bin


source/etc/profile
```

## 关闭

```shell
mongod --port=27017 --dbpath=/mongodb/data --shutdown
```

```shell
use admin
db.shutdownServer()
exit
```

