//go:build !wasm

package clitools

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"github.com/robfig/cron"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	fireback.EnqueueTask = enqueueTask
	fireback.LiftAsyncqWorkerServer = liftAsyncqWorkerServer
}

func toAsyncqTask(x *fireback.TaskMessage) (*asynq.Task, error) {
	return asynq.NewTask(x.Name, x.Payload), nil
}

func enqueueTask(task *fireback.TaskMessage) (*fireback.TaskEnqueueResult, error) {

	addr := "127.0.0.1:6379"

	if config.WorkerAddress != "" {
		addr = config.WorkerAddress
	}

	asyn, err := toAsyncqTask(task)

	if err != nil {
		log.Panicln("Error casting to the task: %w", err)
	}

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: addr})
	defer client.Close()

	info, err := client.Enqueue(asyn)
	if err != nil {
		return nil, err
	}

	return &fireback.TaskEnqueueResult{
		ID: info.ID,
	}, nil
}

func liftAsyncqWorkerServer(tasks []*fireback.TaskAction) {
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
				fireback.NewTaskCtx(ctx),
				t.Payload(),
			)
		})

		// if len(task.Triggers) > 0 {
		// 	for _, trigger := range task.Triggers {
		// 		if trigger.Cron != "" {

		// 			c.AddFunc(trigger.Cron, func() { fmt.Println("Trigger: %s", trigger.Cron) })
		// 		}

		// 	}
		// }
	}

	c.Start()

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
