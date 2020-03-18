# delRedisScript
High Performance Delete Redis Script

## 设计目标
配套addRedisScript脚本，增加数据   
 [addRedisScript](https://github.com/iooikaak/addRedisScript)

### Features
- 高性能，bechmark测试
- 支持开10W协程并发删除redis
- QPS可以达到100W

#	依赖
## 硬件平台
- X86 服务器

## 软件平台
- Ubuntu 18.04 LTS
- Go1.14

## bechmark测试
- 测试机器
hp latebook pro 2核 8G内存 

### 开200个协程，插入10W条数据
![image](https://github.com/iooikaak/delRedisScript/blob/master/pic/11.png)
```
vagrant@ubuntu-bionic:~/Dev/code/go/src/delRedisScript$ go run main.go -readRedisHost "127.0.0.1:6379" -readRedisPassword "" -writeRedisHost "127.0.0.1:6379" -writeRedisPassword "" -userIDMinNum 1 -userIDMaxNum 100000 -goRoutineCount 200
2020/03/18 09:46:54 Del Redis Key Script Has Done!!!
2020/03/18 09:46:54 Start Unix Time:1584524811060564195, End Unix Time:1584524814592211008, Run time:3531646813
```
QPS：100000/3.531646813= 28315
### 开1W个协程，插入1000W条数据
![image](https://github.com/iooikaak/delRedisScript/blob/master/pic/12.png)
```
vagrant@ubuntu-bionic:~/Dev/code/go/src/delRedisScript$ go run main.go -readRedisHost "127.0.0.1:6379" -readRedisPassword "" -writeRedisHost "127.0.0.1:6379" -writeRedisPassword "" -userIDMinNum 1 -userIDMaxNum 10000000 -goRoutineCount 10000
2020/03/18 09:50:10 Del Redis Key Script Has Done!!!
2020/03/18 09:50:10 Start Unix Time:1584524989419585648, End Unix Time:1584525010091981372, Run time:20672395724
```
QPS: 10000000/20.672395724=483736
### 开5W个协程，插入1000W条数据
![image](https://github.com/iooikaak/delRedisScript/blob/master/pic/13.png)
```
vagrant@ubuntu-bionic:~/Dev/code/go/src/delRedisScript$ go run main.go -readRedisHost "127.0.0.1:6379" -readRedisPassword "" -writeRedisHost "127.0.0.1:6379" -writeRedisPassword "" -userIDMinNum 1 -userIDMaxNum 10000000 -goRoutineCount 50000
2020/03/18 09:52:23 Del Redis Key Script Has Done!!!
2020/03/18 09:52:23 Start Unix Time:1584525133518306707, End Unix Time:1584525143059049486, Run time:9540742779
```
QPS: 10000000/9.540742779=924374
### 开10W个协程，插入1000W条数据
![image](https://github.com/iooikaak/delRedisScript/blob/master/pic/14.png)
```
vagrant@ubuntu-bionic:~/Dev/code/go/src/addRedisScript$ go run main.go -readRedisHost "127.0.0.1:6379" -readRedisPassword "" -writeRedisHost "127.0.0.1:6379" -writeRedisPassword "" -userIDMinNum 1 -userIDMaxNum 10000000 -goRoutineCount 100000
2020/03/18 09:19:56 Add Redis Key Script Has Done!!!
2020/03/18 09:19:56 Start Unix Time:1584523186976773055, End Unix Time:1584523196774139671, Run time:9797366616 
```
QPS:10000000/9.797366616=1020682
