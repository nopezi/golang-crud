package jobs

import (
	"database/sql"
	dio "infolelang/jobs/dio"
)

const (
	dbMaxIdleConns = 4
	dbMaxConns     = 100
	totalWorker    = 100
)

type JobsNeeded struct {
	DBdefault *sql.DB
}

func NewCronJob(DBdefault *sql.DB) *JobsNeeded {
	return &JobsNeeded{
		DBdefault: DBdefault,
	}
}

func (j JobsNeeded) ListJobs(m string) func() {
	switch m {
	// case "updateDataMasterSDM":
	// return j.updateDataMasterSDM()
	// case "updateDataDashboard":
	// 	return j.updateDataDashboard()
	// case "dioRemainder":
	// 	return j.dioRemainderBisnisReview()
	case "scheduler1":
		return j.dioRemainderBisnisReview()
	default:
		return func() {}
	}
}

func (j JobsNeeded) dioRemainderBisnisReview() func() {
	return func() {
		go dio.DioRemainder(j.DBdefault, totalWorker, dbMaxConns, dbMaxIdleConns)
	}
}
