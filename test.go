package main
import (
    "time"
   "runtime"
)
func test(c chan bool, b bool) {
x := 0
for i := 0; i < 1000000; i++ {
x += i
for m := 0; m < 50; m++ {
x += 0
}
}

if b {
println(x);
 c <- true; 
 }
}
func main() {
 t := time.Now().UnixNano()
  runtime.GOMAXPROCS(8)
c := make(chan bool)
for i := 0; i < 6; i++ {
go test(c, i == 5)
}
<- c
t2 := time.Now().UnixNano()
println((t2-t))
//下面是得到消耗的秒数
//println((t2-t)/10000000)
}
