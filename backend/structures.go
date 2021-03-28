package main

import "time"

type homeResponse struct {
	Message string `json:"Message"`
	TodayDate time.Time `json:"TodayDate"`
}


type applicationHealthResponse struct {
	Status string `json:"Status"`
	AverageUptime int32 `json:"AverageUptime"`
	NumFailures int32 `json:"NumFailures"`
}

type SubmittedUserInfo struct {
	Number int16 `json:"Number"` // is always a number between 0 and 10
	Age int16 `json:"Age"`
	Nationality string `json:"Nationality"`
	Gender string `json:"Gender"`
	DominantHand string `json:"DominantHand"`
}
