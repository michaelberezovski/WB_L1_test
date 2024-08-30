package task18

import (
	"fmt"
	"sync"
)

type Counter struct { //структура счетчика
	num int
	sync.Mutex
}

func (c *Counter) Inc() { //метод инкремент с блокировкой критической секции
	c.Lock()
	c.num++
	c.Unlock()
}

func (c *Counter) Value() int { //геттер
	return c.num
}

func main() {
	wg := sync.WaitGroup{}    //группа ожидания
	c := Counter{num: 0}      //экземпляр счетчика
	for i := 0; i < 10; i++ { //цикл с запуском 10 горутин и имитацией деятельности
		wg.Add(1)
		go func(c *Counter, i int, wg *sync.WaitGroup) { //запуск горутины с входными параметрамиа
			defer wg.Done()
			c.Inc() //при запуске горутины счетчик накидывает еденицу
			fmt.Println(i)
		}(&c, i, &wg)
	}
	wg.Wait()                                    //ждем завершения всех горутин
	fmt.Printf("Counter value is %d", c.Value()) //выводим показания счётчика
}
