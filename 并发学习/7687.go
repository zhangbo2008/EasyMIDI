package main

import (
  "fmt"
  "sync"
  "time"
)

var (
  count int
  l     = sync.Mutex{}
  m     = make(map[int]int)
)

//全局变量并发写 导致计数错误
func vari() {
  for i := 0; i < 10000; i++ {
          go func(i int) {
            //defer l.Unlock()
            //l.Lock()
            count++
          }(i)
        }
  fmt.Println(count)   //看最后是不是打印1万,如果不加锁打印出来是不一定的.因为读写乱了.


}

//map 并发写 不加锁 fatal error: concurrent map writes
func mp() {
  for i := 0; i < 1000; i++ {
    go func() {
      defer l.Unlock()
      l.Lock()
      m[0] = 0
    }()
  }
}

func main() {
  vari()

  //mp()
  time.Sleep(3 * time.Second)
}
