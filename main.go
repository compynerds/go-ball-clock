package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Ask for number of balls // valid between 27-127
	fmt.Println("Enter how many balls valid must be between 27-127: ")
	var ballCountInput string
	_, err := fmt.Scanln(&ballCountInput)
	if err != nil {
		panic(err)
	}
	ballCount, err := strconv.Atoi(ballCountInput)
	if err != nil {
		panic(err)
	}

	if ballCount < 27 || ballCount > 127 {
		panic("input out of range 27-127")
	}

	balls := NewBalls(ballCount)
	continueCondition := true
	halfDaysCounted := 0
	minutes := make([]Ball, 0)
	fiveMinuteQueue := make([]Ball, 0)
	fiveMinuteMax := 12
	hourQueue := make([]Ball, 0)
	hourMax := 12

	for continueCondition {
		balls, minutes = Dequeue(balls, 5) // 5 minute collection
		fifthBall, reverseRemaining := HandleBalls(minutes)
		fiveMinuteQueue = append(fiveMinuteQueue, fifthBall)
		balls = append(balls, reverseRemaining...)

		if len(fiveMinuteQueue) == fiveMinuteMax {
			hourBall, reverseFiveQueue := HandleBalls(fiveMinuteQueue)
			hourQueue = append(hourQueue, hourBall)
			balls = append(balls, reverseFiveQueue...)
			fiveMinuteQueue = make([]Ball, 0)
		}

		if len(hourQueue) == hourMax {
			halfDayBall, halfDayReverse := HandleBalls(hourQueue)
			balls = append(balls, halfDayReverse...)
			balls = append(balls, halfDayBall)
			hourQueue = make([]Ball, 0)
			halfDaysCounted++
		}

		if isInOriginalOrder(balls) {
			continueCondition = false
		}
	}
	fmt.Println(halfDaysCounted/2, " days")
}
