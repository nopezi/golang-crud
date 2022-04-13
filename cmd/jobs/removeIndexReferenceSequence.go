package cmd

import (
	"fmt"
	"log"
	"os"

	"eform-gateway/jobs"

	"eform-gateway/lib"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type JobService struct {
	logger        lib.Logger
	jobRepository jobs.JobRemoveRepository
}

var (
	jobRemoveCmd = &cobra.Command{
		Use:              "job-remove",
		Short:            "Eform Jobs Remove Index Reference Sequence",
		Long:             "Jobs Eform through Cron",
		PersistentPreRun: jobsRemovePreRun,
		RunE:             jobRemoveRun,
	}
)

func JobsRemoveCmd() *cobra.Command {
	return jobRemoveCmd
}

func jobsRemovePreRun(cmd *cobra.Command, args []string) {
	_ = godotenv.Load()
	fmt.Println("root Pre RUN")
}

func jobRemoveRun(cmd *cobra.Command, args []string) error {
	var job JobService

	url := os.Getenv("DBEHost")
	username := os.Getenv("DBEUsername")
	password := os.Getenv("DBEPassword")
	elastic, err := lib.New([]string{url}, username, password)
	if err != nil {
		log.Fatalln(err)
	}

	_ = jobs.NewJobRepository(elastic)
	if err != nil {
		log.Fatalln(err)
	}

	// err = job.jobRepository.JobsRemove(elastic)

	if err != nil {
		log.Fatalln(err)
	}
	_ = job.jobRepository.JobsRemove()

	return nil
}
