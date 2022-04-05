package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/joho/godotenv"
)

var (
	jobUpdateCmd = &cobra.Command{
		Use:              "jobs-updatetoexecuted",
		Short:            "Eform Jobs Update To Executed",
		Long:             "Jobs Eform through Cron",
		PersistentPreRun: jobsUpdatePreRun,
		RunE:             jobsUpdateRun,
	}
)
// - [ ] Cronjob update expired date by timestime, Not including this service api, registered on crontab linux
// - search index where documen if expired_date = now , create to transactionExpireds and delete index from transactions
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
	// fx.New(bootstrap.Module).Run()
	return nil
}
