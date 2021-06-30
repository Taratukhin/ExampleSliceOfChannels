package main
import (
	"fmt"
	"sync"
)

func ResendToChannels(in chan int, n int) (out []chan int) { // функция возвращает n каналов out и пересылает в них сообщение из канала in, когда оно в нем появляется
	out = make([]chan int,n)// создадим слайс
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
	var wg sync.WaitGroup // будем ждать окончания запущенных горутин
	for i:=0; i<len(out);i++ {
		go func (j int) { 	// приходится передавать номер канала, иначе для всех горутин будет использовано одно значение i 
					// (последнее значение счетчика после выхода из цикла) 
			wg.Add(1)	// на одну горутину больше
			fmt.Printf("Message from channel out[%d]=%d\n",j,<-out[j]) // с какого канала прийдет сообщение первым неизвестно
			wg.Done() 	// на одну горутину меньше
		}(i) 
	}
	in <- 111
	wg.Wait()
}