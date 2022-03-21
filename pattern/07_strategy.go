package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

//это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый
//из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

// Горячая замена алгоритмов на лету.
// Изолирует код и данные алгоритмов от остальных классов.
// Уход от наследования к делегированию.
// Реализует принцип открытости/закрытости.

// Усложняет программу за счёт дополнительных классов.
// Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

//интерфейс стратегии
type evictionAlgo interface {
	evict(c *cache)
}

//стратегия FIFO
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

//стратегия lru
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

//стратегия lfu
type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}

//кеш
type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

//пользовательский код
func main() {
	LFU := &lfu{}
	Cache := initCache(LFU)

	Cache.add("a", "1")
	Cache.add("b", "2")

	Cache.add("c", "3")

	LRU := &lru{}
	Cache.setEvictionAlgo(LRU)

	Cache.add("d", "4")

	FIFO := &fifo{}
	Cache.setEvictionAlgo(FIFO)

	Cache.add("e", "5")

}
