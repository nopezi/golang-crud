package cmd
import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/joho/godotenv"

)

var (
	jobRemoveCmd = &cobra.Command{
		Use:              "jobs-resetrefnumber",
		Short:            "Eform Jobs Remove Index Reference Sequence",
		Long:             "Jobs Eform through Cron",
		PersistentPreRun: jobsRemovePreRun,
		RunE:             jobRemoveRun,
	}
)
// crontjob remove index reference_sequence
func JobsRemoveCmd() *cobra.Command {
	return jobRemoveCmd
}

func jobsRemovePreRun(cmd *cobra.Command, args []string) {
	// log.LogInit()

	// if env.Get() == env.EnvDevelopment {
	// 	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	// 	log.Level("debug")
	// }
	fmt.Println("root Pre RUN")
}

func jobRemoveRun(cmd *cobra.Command, args []string) error {
	godotenv.Load()
	fmt.Println()
	// fx.New(bootstrap.Module).Run()
	return nil
}
