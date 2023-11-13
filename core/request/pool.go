package request

import (
	"ApricusFuzz/core/response"
	"github.com/panjf2000/ants/v2"
	"log"
	"sync"
)

var TaskWaitGroup sync.WaitGroup

type Pool struct {
	Pool *ants.PoolWithFunc
}

func (pool *Pool) Wait() {
	TaskWaitGroup.Wait()
}

func (pool *Pool) Close() {
	pool.Pool.Release()
}

func FuzzRequestWorker(fuzzReqInterface interface{}) {
	defer TaskWaitGroup.Done()
	fuzzReq, ok := fuzzReqInterface.(*Request)
	if !ok {
		log.Fatalln("Failed to convert to type FuzzRequest")
	}

	resp, err := fuzzReq.Request.Send()
	if err != nil {
		// fuzzReq.Channel() <- resp
		log.Printf("%v => [ERR]%s\n", fuzzReq.Data, resp.Request.URL)
	} else {
		fuzzReq.Response.RespChannel <- resp
	}
}
func InitPool(threadNum int) *Pool {
	pool, err := ants.NewPoolWithFunc(threadNum, FuzzRequestWorker)
	if err != nil {
		log.Fatalf("Init Request Pool Error: %s\n", err.Error())
	}

	return &Pool{
		Pool: pool,
	}
}
func (pool *Pool) Start(fuzz *FuzzRequest, Response *response.Response) {
	TaskWaitGroup.Add(1)
	defer TaskWaitGroup.Done()

	reqChan := make(chan *Request)
	fuzz.Iterator.Exec(fuzz.Payloads)

	go func() {
		for req := range reqChan {
			if err := pool.Pool.Invoke(req); err != nil {
				log.Println(err.Error())
			}
		}
	}()

	for {
		data, alive := <-fuzz.Iterator.Channel()
		if alive {
			reqChan <- fuzz.DoRequest(data, Response)
			TaskWaitGroup.Add(1)
		} else {
			break
		}
	}

}
