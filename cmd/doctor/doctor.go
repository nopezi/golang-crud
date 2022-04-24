package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/joho/godotenv"
	"eform-gateway/lib"
	"os"
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
	fmt.Println("doctor Pre RUN")
}

func runDoctor(cmd *cobra.Command, args []string) error {
	fmt.Println("Go Doctor")
	_ = godotenv.Load()
	
	url := os.Getenv("DBEHost")
	username := os.Getenv("DBEUsername")
	password := os.Getenv("DBEPassword")

	_, err := lib.New([]string{url}, username, password)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
