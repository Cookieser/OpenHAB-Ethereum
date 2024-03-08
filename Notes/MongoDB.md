# MongoDB

> Version: v4.49. Stable from 5.2 onwards.

This is a document database (using JSON as the data model), written in C++, aimed at providing a scalable, high-performance data storage solution for WEB applications. It's a product that sits between relational and non-relational databases, supporting multi-document shard transactions (most similar to relational databases). During project iterations, MySQL requires altering the table structure, but MongoDB does not, offering great flexibility. Semi-structured: In a collection, the documents do not need to have the same fields, and there's no need to declare all fields, making it very suitable for object-oriented programming models.

## Technical Advantages

- Based on a flexible JSON document model, very suitable for agile rapid development.
  - Dynamically adding new fields, with versioning to differentiate version changes.
- High availability:
  - Self-recovery.
  - Multi-center disaster recovery capability.
  - Rolling service - minimizing service downtime.
- Horizontal scaling - sharding for capacity expansion.

## Application Scenarios

> Current specific businesses usually adopt a microservices architecture, where each service can use different databases.

- Game data: convenient for querying and updating.
- Logistics: status is constantly changing.
- Internet of Things scenarios...

## Environment Installation

Community Server v4.4

```shell
tar -xvf mongodb-linux-x86_64-ubuntu2004-4.4.28.tgz
```

## Startup

### Command Line Startup

```shell
mkdir -p /mongodb/data /mongodb/log
bin/mongod --port=27017 --dbpath=/mongodb/data --logpath=/mongodb/log/mongodb.log --bind_ip=0.0.0.0 --fork
```

Note that sudo is required for starting up!

![image.png](https://cdn.nlark.com/yuque/0/2024/png/40368069/1708737878694-1e937a46-319d-4b6d-bfac-ecb4c65017b8.png#averageHue=%23300a25&clientId=u8519a271-a58e-4&from=paste&height=78&id=u0374d33c&originHeight=97&originWidth=1716&originalType=binary&ratio=1.25&rotation=0&showTitle=false&size=29325&status=done&style=none&taskId=u67385af9-eb3b-46b2-b686-e50120a848f&title=&width=1372.8)

![image.png](https://cdn.nlark.com/yuque/0/2024/png/40368069/1708738040212-245bbb78-d0a2-4a12-8c12-a366ccd18707.png#averageHue=%23300a24&clientId=u8519a271-a58e-4&from=paste&height=423&id=ue6764f59&originHeight=529&originWidth=1705&originalType=binary&ratio=1.25&rotation=0&showTitle=false&size=122748&status=done&style=none&taskId=ub19ff17a-fbf8-4c69-ba4c-fe356bc4f29&title=&width=1364)

### Configuration File Startup

```yaml
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

## Adding Environment Variable

```bash
vim /etc/profile


export MONGODB_HOME=/usr/local/soft/mongodb
PATH=$PATH:$MONGODB_HOME/bin


source /etc/profile
```

## Shutdown

```shell
mongod --port=27017 --dbpath=/mongodb/data --shutdown
use admin
db.shutdownServer()
exit
```
