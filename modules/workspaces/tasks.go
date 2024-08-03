package workspaces

import (
	"context"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"github.com/robfig/cron"
	"github.com/urfave/cli"
)

type TaskCtx struct {
	ctx context.Context
}

type TaskAction struct {
	HandlerFunc func(ctx *TaskCtx, content []byte) error
	Name        string
	Cli         func(c *cli.Context) error
	Flags       []cli.Flag
	Cron        string
}

type TaskMessage struct {
	Name    string
	Payload []byte
}

func (x *TaskMessage) ToAsyncqTask() (*asynq.Task, error) {

	return asynq.NewTask(x.Name, x.Payload), nil
}

type TaskEnqueueResult struct {
	ID string
}

func EnqueueTask(task *TaskMessage) (*TaskEnqueueResult, error) {

	addr := "127.0.0.1:6379"

	if config.WorkerAddress != "" {
		addr = config.WorkerAddress
	}

	asyn, err := task.ToAsyncqTask()

	if err != nil {
		log.Panicln("Error casting to the task: %w", err)
	}

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: addr})
	defer client.Close()

	info, err := client.Enqueue(asyn)
	if err != nil {
		return nil, err
	}

	return &TaskEnqueueResult{
		ID: info.ID,
	}, nil
}

func liftAsyncqWorkerServer(tasks []*TaskAction) {
	addr := "127.0.0.1:6379"

	if config.WorkerAddress != "" {
		addr = config.WorkerAddress
	}

	concurrency := 10
	if config.WorkerConcurrency != 0 {
		concurrency = int(config.WorkerConcurrency)
	}

	// Only asyncq for now. Implement the rabbit mq etc here
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: addr},
		asynq.Config{
			Concurrency: concurrency,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()

	c := cron.New()

	for _, task := range tasks {
		task := task
		mux.HandleFunc(task.Name, func(ctx context.Context, t *asynq.Task) error {
			return task.HandlerFunc(
				&TaskCtx{
					ctx: ctx,
				},
				t.Payload(),
			)
		})

		if task.Cron != "" {
			c.AddFunc(task.Cron, func() { fmt.Println("Trigger: %s", task.Cron) })
		}
	}

	c.Start()

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
