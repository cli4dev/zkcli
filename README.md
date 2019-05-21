# zkcli
zookeeper命令行工具

#### 1. 连接到zookeeper

```sh
$ zkcli "zk://192.168.0.101"
Connected to 192.168.0.101:2181
Authenticated: id=101927630269714731, timeout=4000
Re-submitting `0` credentials after reconnect
zk:/ >

```

#### 2. `ls`列出子节点
```sh
zk:/ >ls
us            
monitor_debug              gsms_debug                 merge_debug                ServerMonitor              
sinopec_coupon             recv-coupon_debug          convoy_collect_debug       dxcsh_debug                
sso                        hydra-20_debug             
```

#### 3. `cd`进入目录
```sh
zk:/ >cd us
zk:/us >

```

#### 4. `cat`查看节点内容
```sh
zk:/us/ums1/web/t >cat conf
/us/ums1/web/t/conf:
{"address":":8081"}

zk:/us/ums1/web/t >
```
#### 5. `cd /xx`进入指定目录
```sh
zk:/us/ums1 >cd /us
zk:/us >
```

#### 6. `cd  ..`返回上级目录
```sh
zk:/us/ums1 >cd ..
zk:/us >

```
#### 7. `cd  xx`模糊匹配
```sh
zk:/us >cd u
zk:/us/ums1 >
```



#### 8. `pwd`显示当前路径
```sh
zk:/us/ums1 >pwd
/us/ums1
```

#### 9. `rm` 删除节点
```sh
zk:/us/ums1/web/t/conf >ls
static 
zk:/us/ums1/web/t/conf >rm static
remove /us/ums1/web/t/conf/static	ok
zk:/us/ums1/web/t/conf >
```
