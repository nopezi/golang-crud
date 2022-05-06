package cmd

import (
	"eform-gateway/lib"
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
	// fmt.Println("doctor Pre RUN")
}

func runDoctor(cmd *cobra.Command, args []string) error {
	lib.LogInfo("----------------------------------------")
	lib.LogInfo("----------- Go Doctor ------------------")
	lib.LogInfo("----------------------------------------")
	_ = godotenv.Load()

	url := os.Getenv("DBEHost")
	username := os.Getenv("DBEUsername")
	password := os.Getenv("DBEPassword")

	_, err := lib.New([]string{url}, username, password)
	if err != nil {
		fmt.Println(err)
	}
	lib.LogInfo("----------------------------------------")
	return nil
}
