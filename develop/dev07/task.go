package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	res := make(chan interface{})

	// Выгружаем все из канала
	closeChan := func(channel <-chan interface{}) {
		for val := range channel {
			res <- val
		}
		// Ждем указанное в горутине время и помечаем что сделали
		wg.Done()
	}
	// Добавляем вг
	wg.Add(len(channels))
	for _, channel := range channels {
		// Запускаем параллельно каналы
		go closeChan(channel)
	}
	go func() {
		// Тут ожидаем
		wg.Wait()
		close(res)
	}()
	return res
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Millisecond),
		sig(50*time.Millisecond),
		sig(100*time.Millisecond),
		sig(1*time.Millisecond),
		sig(1*time.Millisecond),
	)

	fmt.Printf("done after %v", time.Since(start))
}
