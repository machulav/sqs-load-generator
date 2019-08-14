package sender

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Config contains sender configuration
type Config struct {
	SqsURL         string
	SqsMessage     string
	Parallelism    int
	MessagesAmount int
}

// Run starts message sending
func (c *Config) Run() {
	successCount := uint32(0)
	errorCount := uint32(0)
	startTime := time.Now()

	var wg sync.WaitGroup
	for i := 1; i <= c.Parallelism; i++ {
		wg.Add(1)
		go c.runWorker(i, &wg, &successCount, &errorCount)
	}
	wg.Wait()

	fmt.Printf("Sent messages: %v\n", successCount)
	fmt.Printf("Errors: %v\n", errorCount)
	fmt.Printf("Time: %v\n", time.Since(startTime))
}

func (c *Config) runWorker(id int, wg *sync.WaitGroup, successCount *uint32, errorCount *uint32) {
	defer wg.Done()

	queue := getSqs()

	for i := 1; i <= c.MessagesAmount; i++ {
		_, err := queue.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(0),
			MessageBody:  &c.SqsMessage,
			QueueUrl:     &c.SqsURL,
		})

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			atomic.AddUint32(errorCount, 1)
			continue
		}

		atomic.AddUint32(successCount, 1)
	}
}

func getSqs() *sqs.SQS {
	sess := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
			}))

	return sqs.New(sess)
}
