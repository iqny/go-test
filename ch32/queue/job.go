package queue

import (
	"context"
	"go-zh/ch32/queue/config"
)

type Job struct {
	manage *Manage
}
type JobHandle map[string]Handler

type Handler interface {
	Call(args interface{}) (traceId string, err error)
	GetJobName() string
}

func NewJob(ctx context.Context, cfg *config.Config) *Job {
	return &Job{
		manage: New(cfg),
	}
}

func (j *Job) Run() {
	j.manage.Run()
}
func (j *Job) Register(jobHandle JobHandle) {
	for name, job := range jobHandle {
		j.manage.Register(name, job.Call)
	}
}
func (j *Job) Stop() {
	//j.manage.Wait()
	//j.manage.Close()
}
