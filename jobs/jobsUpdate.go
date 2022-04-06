package jobs

import (
	"fmt"
)

// - [ ] Cronjob update expired date by timestime, Not including this service api, registered on crontab linux
// - search index where documen if expired_date = now , create to transactionExpireds and delete index from transactions

func JobsUpdate() {
	fmt.Println("jobs Update")
}