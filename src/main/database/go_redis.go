package database

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Pool *redis.Pool
)

func init() {
	fmt.Println("--初始化redis连接--")
	redisAddress := "ip:port"
	Pool = newPool(redisAddress)
	close()
}

func redisConn(server string) (redis.Conn, error) {
	dial, err := redis.Dial("tcp", server, redis.DialUsername(""), redis.DialPassword(""))
	if err != nil {
		return dial, err
	}
	return dial, nil
}

// 连接池连接
func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			// 设置连接地址 username password 非必填参数，根据服务器配置设置
			dial, err := redis.Dial("tcp", server, redis.DialUsername(""), redis.DialPassword(""))
			if err != nil {
				return nil, err
			}
			return dial, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

func Get(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

func Set(key string, value string) {
	conn := Pool.Get()
	defer conn.Close()
	redis.Strings(conn.Do("SET", key, value))
}

func GetKey(w http.ResponseWriter, r *http.Request) {
	Set("test", "value")
	bytes, _ := Get("test")
	fmt.Println(string(bytes))
	fmt.Fprintf(w, string(bytes))
}
