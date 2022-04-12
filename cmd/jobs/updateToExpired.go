package cmd

import (
	"fmt"

	"eform-gateway/jobs"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	jobUpdateCmd = &cobra.Command{
		Use:              "job-update",
		Short:            "Eform Jobs Update To Expired",
		Long:             "Jobs Eform through Cron",
		PersistentPreRun: jobsUpdatePreRun,
		RunE:             jobsUpdateRun,
	}
)

func JobsUpdateCmd() *cobra.Command {
	return jobUpdateCmd
}

func jobsUpdatePreRun(cmd *cobra.Command, args []string) {
	// log.LogInit()

	// if env.Get() == env.EnvDevelopment {
	// 	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	// 	log.Level("debug")
	// }
	fmt.Println("root Pre RUN")
}

func jobsUpdateRun(cmd *cobra.Command, args []string) error {
	godotenv.Load()
	fmt.Println()
	jobs.JobsUpdate()
	// fx.New(bootstrap.Module).Run()
	return nil
}
