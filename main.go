package main

import (
	"fmt"
	"grondo/db/sql"
	"grondo/services/scheduler"
	"grondo/utils"
)

func main() {
	fmt.Println(utils.Config)

	for _, job := range utils.Config.Jobs {
		fmt.Println(job.Name)
		fmt.Println(scheduler.ParseCron(job.Cron))
	}

	sql.InitDB()
	sql.TestConnection()
	sql.CloseDB()
}
