package main

import (
	"log"
	"time"
)

type RemoteResult struct {
	Url    string
	Result string
}

func RemoteGet(requestUrl string, resultChan chan RemoteResult) {
	time.Sleep(100 * time.Millisecond)
	resultChan <- RemoteResult{Url: requestUrl, Result: "dddddd" + time.Now().Format(time.RFC3339)}
}
func MultiGet(urls []string) []RemoteResult {
	log.Println(time.Now())
	resultChan := make(chan RemoteResult, len(urls))
	defer close(resultChan)
	var result []RemoteResult
	//fmt.Println(result)
	for _, url := range urls {
		go RemoteGet(url, resultChan)
	}
	for i := 0; i < len(urls); i++ {
		res := <-resultChan
		result = append(result, res)
	}
	log.Println(time.Now().Format(time.RFC3339))
	return result
}

func main() {
	urls := []string{
		"http://127.0.0.1/test.php?i=13",
		"http://127.0.0.1/test.php?i=14",
		"http://127.0.0.1/test.php?i=15",
		"http://127.0.0.1/test.php?i=16",
		"http://127.0.0.1/test.php?i=17",
		"http://127.0.0.1/test.php?i=18",
		"http://127.0.0.1/test.php?i=19",
		"http://127.0.0.1/test.php?i=20"}
	content := MultiGet(urls)
	log.Println(content)
}
