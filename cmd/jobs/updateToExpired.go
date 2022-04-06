package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/joho/godotenv"
	"eform-gateway/jobs"
)

var (
	jobUpdateCmd = &cobra.Command{
		Use:              "job-update",
		Short:            "Eform Jobs Update To Executed",
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
