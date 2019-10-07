package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type logTime struct {
	StartTime string `json:"StartTime"`
	EndTime   string `json:"EndTime"`
}

const activityLogQueryDelta = 10 * time.Minute

func main() {
	startTime := time.Now()
	endTime := time.Now().Add(time.Duration(-activityLogQueryDelta))
	var logv logTime
	logt := logTime{
		StartTime: startTime.String(),
		EndTime:   endTime.String(),
	}

	val, err := json.Marshal(&logt)
	fmt.Println("Time", val, "Err", err)
	err = json.Unmarshal(val, &logv)
	fmt.Println("Time", logv, "Err", err)
}
