package main

import(
	"github.com/rgarcia2304/learn-pub-sub-starter/internal/gamelogic"
	"fmt"
	"github.com/rgarcia2304/learn-pub-sub-starter/internal/routing"
)

func handlerPause(gs *gamelogic.GameState) func(routing.PlayingState){
	
	return func(ps routing.PlayingState){
		defer fmt.Print(" >")
		gs.HandlePause(ps)
	}
}	
