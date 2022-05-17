package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"infolelang/jobs"

	"infolelang/lib"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	username := os.Getenv("DBUsername")
	password := os.Getenv("DBPassword")
	host := os.Getenv("DBHost")
	port := os.Getenv("DBPort")
	dbname := os.Getenv("DBName")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		lib.LogChecklist("Mysql Connection Refused", false)
	}

	lib.LogChecklist("Mysql Connection Established", true)

	result := db.Exec("truncate table reference_code_counters;")
	fmt.Println(result)

	return nil
}
