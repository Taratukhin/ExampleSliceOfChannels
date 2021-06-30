package main
import (
	"fmt"
)

func ResendToChannels(in chan int, n int) []chan int { // функция возвращает n каналов out и пересылает в них сообщение из канала in, когда оно в нем появляется
	out := make([]chan int,n)// создадим слайс
	for i := range out { 	// каждый канал тоже нужно создать
   		out[i] = make(chan int)
	}
	go func () {            // если не запустить горутину, то функция не вернет out, т.к. заблокируется ожиданием сообщения в канале in
		var msg = <-in  // сообщение во входящий канал ждем, естественно в горутине
		for i:=0; i<n; i++ {
			out[i]<-msg
		}
	}()
	return out
}

func main() {
	in := make(chan int)
	out := ResendToChannels(in,5)
	in <- 111 
	for i:=0; i<len(out);i++ { 
		fmt.Printf("Message from channel out[%d]=%d\n",i,<-out[i]) // в этом месте будет ждать сообщения именно в i-й канал поэтому номера каналов будут попорядку
	}
}