package cron

import (
	"context"
	"errors"
	"time"

	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"github.com/google/wire"
	"github.com/robfig/cron"
)

var (
	ErrNoSuchJob                 = errors.New("no such job")
	ErrConvertJobInterfaceFalied = errors.New("convert job interface failed")
	ErrFailedAddFunc             = errors.New("failed to add func")
)

var CronProvider = wire.NewSet(NewCron)

type CronConfig struct {
	CronName string           `mapstructure:"cron_name"`
	Jobs     map[JobName]Spec `mapstructure:"jobs"`
}

type (
	JobName = string
	Spec    = string
)

// type JobConfig struct {
// 	JobName string `mapstructure:"job_name"`
// 	Spec    string `mapstructure:"spec"`
// }

type Cron struct {
	CronName     string
	JobProviders map[string]JobInterface
	TimeLocation *time.Location
}

type JobInterface interface {
	ExecuteJob(ctx context.Context) error
	Name() string
}

func (c *Cron) Register(jobName string, jobInterface JobInterface) {
	c.JobProviders[jobName] = jobInterface
}

// default time location: "Asia/Shanghai".
func NewCron(cronName string) *Cron {
	// default time location
	timeLoc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	cron := &Cron{
		CronName:     cronName,
		JobProviders: make(map[string]JobInterface),
		TimeLocation: timeLoc,
	}

	return cron
}

func (c *Cron) StartJob(ctx context.Context, spec string, jobName string) (*cron.Cron, error) {
	// 1.获取"provider"
	var job JobInterface
	job, exist := c.JobProviders[jobName]
	if !exist {
		logger.Error(ErrConvertJobInterfaceFalied.Error())
		return nil, ErrConvertJobInterfaceFalied
	}

	// 2.执行"provider"
	cr := cron.NewWithLocation(c.TimeLocation)
	if err := cr.AddFunc(spec, func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(err)
			}
		}()
		cr.Stop()
		defer cr.Start()
		// logger.Infof("Start job, current job name is [%s]\n\n\n", jobName)
		if err := job.ExecuteJob(ctx); err != nil {
			logger.Error(err.Error())
		}
	}); err != nil {
		logger.Error(ErrFailedAddFunc.Error())

		return nil, err
	}
	cr.Start()

	return cr, nil
}
