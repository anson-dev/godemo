package main

import (
    "fmt"    
    "time"
    "log"
  	"github.com/garyburd/redigo/redis"
)

func main() {

    fmt.Println("Redis:")
    conn,err := redisConn("113.108.231.10:8000","huya_gift","0");
    if err != nil {
        log.Fatal("Error: ", err)
    }

    test(conn)
}

func test(conn *RedisConn) {
 /*
    conn.Do("SET","xxx",9999)
    if xxx,err :=redis.Int(conn.Do("GET","xxx")); err == nil {
        fmt.Println("xxx:",xxx)
    }
*/
 
   
    for i := 0; i <= 5000; i++ {
        var qb string = "20170627xbuadcsohisek"

        qb = qb + fmt.Sprintf("%04d", i)

        fmt.Println("qb:",qb)
        conn.Do("lpush","qbcode",qb)
    }
    

    if qblen,err :=redis.Int(conn.Do("llen","qbcode")); err == nil {
        fmt.Println("len:",qblen)
    }

    //conn.FlushClose()
}

////////////////////////////////////////////////////////////////
type RedisConn struct {
    dbid string
    redis.Conn
}

func (r *RedisConn) FlushClose() error {
    if r.dbid != "" {
        if _, err := r.Conn.Do("SELECT", r.dbid);err != nil {
            return nil
        }
    }
    if _, err := r.Conn.Do("FLUSHDB");err != nil {
        return err
    }
    return r.Conn.Close()
}

func (r *RedisConn) Close() error {    
    return r.Conn.Close()
}

func redisConn(host,password,db string) (*RedisConn, error) {
    if host == "" {
        host =  ":6379"
    }
    //conn, err := redis.Dial( "tcp", host)
    conn , err := redis.DialTimeout("tcp", host, 0, 1*time.Second, 1*time.Second)
    if err != nil {
        return nil, err
    }

    if password != "" {
        if _, err := conn.Do("AUTH", password); err != nil {        
            conn.Close()
            return nil, err
        }
    }

    if db != "" {
        if _, err := conn.Do("SELECT", db);err != nil {    
            conn.Close()
            return nil, err
        }
    }

    return &RedisConn{dbid:db,Conn: conn}, nil
}