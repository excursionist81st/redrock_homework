package main

import (
	"fmt"
	"sync"
	"time"
)

func download(filename string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	time.Sleep(time.Second)
	results <- filename + "下载完成"
}
func check(results <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range results {
		fmt.Println(value)
	}
}
func closechan(wg1 *sync.WaitGroup, results chan<- string) {
	close(results)
	defer wg1.Done()
}
func main() {
	fmt.Println("开始下载3个文件")
	file := []string{"file1.zip", "file2.pdf", "file3.mp4"}
	results := make(chan string, 3)
	var wg, wg1 sync.WaitGroup
	wg1.Add(1)
	go check(results, &wg1)
	for i := 0; i < len(file); i++ {
		wg.Add(1)
		go download(file[i], &wg, results)
	}
	wg.Wait()
	wg1.Add(1)
	go closechan(&wg1, results)
	wg1.Wait()
	fmt.Printf("所有文件已下载完成")
}
