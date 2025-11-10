package setup

import (
	"context"
	"time"

	"github.com/CZnavody19/music-manager/src/graph"
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/zap"
)

func SetupCron(scheduler gocron.Scheduler, resolver *graph.Resolver) error {
	ctx := context.Background()
	options := gocron.WithSingletonMode(gocron.LimitModeReschedule)

	_, err := scheduler.NewJob(gocron.DurationJob(5*time.Minute), gocron.NewTask(resolver.Orchestrator.CRONjob, ctx), options)
	if err != nil {
		zap.S().Error("Error setting up cron job", err)
		return err
	}

	return nil
}
