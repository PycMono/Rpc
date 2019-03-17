package main

import (
	"fmt"
	"moqikaka.com/RpcClient/src/RpcClient"
	_ "moqikaka.com/RpcClient/src/Slot/impl"
	"sync"
	"time"
)

var(
	wg  sync.WaitGroup
)

func init()  {
	wg.Add(1)
}



func main()  {
	//Slot.CallFunction()
	go func() {
		time.Sleep(time.Second*1)
		for i:=0;i<12000 ; i++ {
			time.Sleep(time.Millisecond*2)

			go RpcClient.StartClient(wg)
		}
	}()

	fmt.Println("准备连接")
	wg.Wait()
}
 