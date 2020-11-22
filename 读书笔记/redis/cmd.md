redis-cli -h 127.0.0.1 -p 6379

redis-cli info clients

brew services start redis


redis 中的string的底层是一个类似 c++ vector的结构，包含 
capacity len 和 byte[]

每个redis 对象都有对象头结构
RedisObject
{
    type  
    encoding
    lru
    reference  
    *prt
}


~~~
set key value
get key

mset key1 v1 k2 v2
mget key1 key2

expire key 5


setex key 5 value    设置key 并且5s后过期，等价于 set key value , expire key 5

setnx key value 如果不存在那么就设置 key= value

set age 30
inc age  // i++
~~~