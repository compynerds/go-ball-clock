package main


//Dequeue -
func Dequeue(balls []Ball, n int) (old, newSet []Ball) {
	newSet = make([]Ball, 0)
	for i := 0; i < n; i++ {
		newSet = append(newSet, balls[i])
	}
	return balls[n:], newSet
}
