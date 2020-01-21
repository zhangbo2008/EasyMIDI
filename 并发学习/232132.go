package main

import (

  "sync"
  "time"
)

var (
  count int
  l     = sync.Mutex{}

)

//全局变量并发写 导致计数错误
func  vari  ()(int) {
  for i := 0; i < 1000; i++ {
          go func(i int) {

            l.Lock()

            count++

             l.Unlock()
          }(i)
        }

return count


}

//map 并发写 不加锁 fatal error: concurrent map writes


func main() {
  a:=vari()
  println(a)
     //看最后是不是打印1万,加锁了,但是不是全跑完之后才打印,如何保证呢?
  //mp()
  time.Sleep(3 * time.Second)
}
