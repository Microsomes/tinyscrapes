package mosquescrappers

import (
	"sync"
)

func worker(id int, wg *sync.WaitGroup, toReturn chan []Prayer, cacheKey string) {
	c := make(chan []Prayer)
	go CrawlBCM(c, id, cacheKey)
	x := <-c
	toReturn <- x
	wg.Done()

}

func BCMYear(cachekey string, prayers chan [][]Prayer) {
	var wg sync.WaitGroup

	c := make(chan []Prayer)

	var allResults [][]Prayer

	for i := 1; i <= 12; i++ {
		wg.Add(1)
		go worker(i, &wg, c, cachekey)
		res := <-c
		allResults = append(allResults, res)
	}
	wg.Wait()
	prayers <- allResults
}
