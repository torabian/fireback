package fireback

import (
	"context"

	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
)

type TaskCtx struct {
	ctx context.Context
}

// NewTaskCtx builds a TaskCtx around ctx. It exists so packages that can't
// see TaskCtx's unexported field (such as modules/fireback/clitools, which
// hosts the actual asynq worker server) can still construct one.
func NewTaskCtx(ctx context.Context) *TaskCtx {
	return &TaskCtx{ctx: ctx}
}

type TaskAction struct {
	HandlerFunc func(ctx *TaskCtx, content []byte) error
	Name        string
	Cli         func(context.Context, *cli.Command) error
	Flags       []cli.Flag
}

type TaskMessage struct {
	Name    string
	Payload []byte
}

type TaskEnqueueResult struct {
	ID string
}

// EnqueueTask and LiftAsyncqWorkerServer talk to asynq/redis, which needs a
// real OS process (goroutine scheduling, network sockets) unavailable under
// wasm. Real implementation lives in modules/fireback/clitools (tagged
// !wasm) and registers itself here via init().
var EnqueueTask func(task *TaskMessage) (*TaskEnqueueResult, error)
var LiftAsyncqWorkerServer func(tasks []*TaskAction)

func GetApplicationTasks(xapp *application.Application) *cli.Command {
	sub := []*cli.Command{}

	// for _, m := range xapp.Modules {
	// for _, t := range m.Tasks {
	// 	sub = append(sub, &cli.Command{
	// 		Name:   t.Name,
	// 		Flags:  t.Flags,
	// 		Action: t.Cli,
	// 	})
	// }
	// }

	return &cli.Command{
		Name:  "tasks",
		Usage: "Actions related to the project tasks, running them in background, list, etc.",
		Commands: []*cli.Command{

			{
				Name:     "enqueue",
				Usage:    "Enqueues a task to the queue so worker can pick it up",
				Commands: sub,
			},
			{
				Name:  "list",
				Usage: "Lists all of the tasks in the app",
				Action: func(ctx context.Context, c *cli.Command) error {
					// for _, m := range xapp.Modules {
					// 	for _, t := range m.Tasks {

					// 		fmt.Println(t.Name)
					// 	}
					// }
					return nil
				},
			},
			{
				Name:  "start",
				Usage: "Starts the background worker server",
				Action: func(ctx context.Context, c *cli.Command) error {
					taskServerLifter(xapp)
					return nil
				},
			},
		},
	}
}

func taskServerLifter(xapp *application.Application) {

	tasks := []*TaskAction{}
	// for _, m := range xapp.Modules {
	// 	for _, t := range m.Tasks {
	// 		tasks = append(tasks, t)
	// 	}
	// }

	LiftAsyncqWorkerServer(tasks)
}
