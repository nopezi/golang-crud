package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	doctorCmd = &cobra.Command{
		Use:              "doctor",
		Short:            "doctor check dependencies",
		Long:             "doctor tool for check dependencies",
		PersistentPreRun: doctorPreRun,
		RunE:             runDoctor,
	}
)

func DoctorCmd() *cobra.Command {
	return doctorCmd
}

func doctorPreRun(cmd *cobra.Command, args []string) {
	// log.LogInit()

	// if env.Get() == env.EnvDevelopment {
	// 	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	// 	log.Level("debug")
	// }
	fmt.Println("doctor Pre RUN")
}

func runDoctor(cmd *cobra.Command, args []string) error {
	// _ = godotenv.Load()
	// fx.New(bootstrap.Module).Run()
	fmt.Println("doctor RUN")
	return nil
}
