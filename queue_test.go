package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestDequeue(t *testing.T) {
	balls := NewBalls(10)
	old, newSet := Dequeue(balls, 5)
	assert.Equal(t, 5, len(old))
	assert.Equal(t, 5, len(newSet))
}

func TestNewBalls(t *testing.T) {
	n := 10
	balls := NewBalls(n)
	assert.Equal(t, n, len(balls))
}

func TestHandleBalls(t *testing.T) {
	n := 12
	balls := NewBalls(n)
	resultBall, revBalls := HandleBalls(balls)
	expected := []int{11,10,9,8,7,6,5,4,3,2,1}

	assert.Equal(t, n, resultBall.Number)
	for i := 0; i >= len(revBalls); i++ {
		assert.Equal(t, expected[i], revBalls[i].Number)
	}
}

func TestHandleBallsFive(t *testing.T) {
	n := 5
	balls := NewBalls(n)
	resultBall, revBalls := HandleBalls(balls)
	expected := []int{4,3,2,1}

	assert.Equal(t, n, resultBall.Number)
	for i := 0; i >= len(revBalls); i++ {
		assert.Equal(t, expected[i], revBalls[i].Number)
	}
}

func TestReverse(t *testing.T) {
	n := 12
	balls := NewBalls(n)
	result := reverse(balls)
	expected := []int{12,11,10,9,8,7,6,5,4,3,2,1}
	for i := 0; i >= len(result); i++ {
		assert.Equal(t, expected[i], result[i].Number)
	}
}
