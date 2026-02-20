package gamelogic

import (
	"fmt"

	"github.com/rgarcia2304/learn-pub-sub-starter/internal/routing"
	"github.com/rgarcia2304/learn-pub-sub-starter/internal/pubsub"
)

func (gs *GameState) HandlePause(ps routing.PlayingState) pubsub.AckType{
	defer fmt.Println("------------------------")
	fmt.Println()
	if ps.IsPaused {
		fmt.Println("==== Pause Detected ====")
		gs.pauseGame()
	} else {
		fmt.Println("==== Resume Detected ====")
		gs.resumeGame()
	}
	return pubsub.Ack
}
