package app

import (
	"fmt"

	"github.com/suchimauz/job-scheduler/internal/config"
	"github.com/suchimauz/job-scheduler/pkg/logger"
)

func Run() {
	_, err := config.Init()
	if err != nil {
		logger.Error(fmt.Sprintf("ENV: %s", err.Error()))

		return
	}
}
