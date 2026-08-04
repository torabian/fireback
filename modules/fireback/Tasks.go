package fireback

import (
	"context"

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
	Triggers    []*Module3Trigger
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
