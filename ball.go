package main

//Ball -
type Ball struct {
	Number int
}

//NewBalls
func NewBalls(max int) []Ball {
	balls := make([]Ball, 0)
	for i := 1; i <= max; i++ {
		balls = append(balls, Ball{Number: i})
	}
	return balls
}

//HandleMinuteBalls
func HandleBalls(balls []Ball) (Ball, []Ball) {
	lastBall := balls[len(balls)-1]
	return lastBall, reverse(balls[:len(balls)-1]) // this will need a zero slice check
}

func reverse(balls []Ball) []Ball {
	revBalls := make([]Ball, 0)
	for i := len(balls) - 1; i >= 0; i-- {
		revBalls = append(revBalls, balls[i])
	}
	return revBalls
}

func isInOriginalOrder(balls []Ball) bool {
	for key, ball := range balls {
		if !(ball.Number == key+1) {
			return false
		}
	}
	return true
}
