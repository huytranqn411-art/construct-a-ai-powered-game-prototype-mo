package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// GameConfig represents the game configuration
type GameConfig struct {
	Width  int
	Height int
	FPS   int
}

// GameMonitor represents the game monitor
type GameMonitor struct {
	Config    GameConfig
	Window    *pixelgl.Window
	GameState *GameState
	AI        *AI
}

// GameState represents the game state
type GameState struct {
	Score int
}

// AI represents the AI
type AI struct {
	Name        string
	Decision    func(gs *GameState) string
	Evaluation  func(gs *GameState) float64
	LastAction  string
	LastReward  float64
}

func NewGameMonitor(config GameConfig) *GameMonitor {
	gameState := &GameState{
		Score: 0,
	}

	ai := &AI{
		Name: "Basic AI",
		Decision: func(gs *GameState) string {
			// simple decision-making logic
			if gs.Score < 10 {
				return " Move Up"
			} else {
				return " Move Down"
			}
		},
		Evaluation: func(gs *GameState) float64 {
			// simple evaluation logic
			return float64(gs.Score)
		},
	}

	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "AI-Powered Game Prototype",
		Bounds: pixel.R(0, 0, float64(config.Width), float64(config.Height)),
		VSync:  true,
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &GameMonitor{
		Config:    config,
		Window:    win,
		GameState: gameState,
		AI:        ai,
	}
}

func (gm *GameMonitor) Run() {
	rand.Seed(time.Now().UnixNano())

	for !gm.Window.Closed() {
		gm.Window.Clear(pixel.RGB(0.2, 0.2, 0.2))

		action := gm.AI.Decision(gm.GameState)
		gm.AI.LastAction = action
		gm.EvaluateAction(action)

		gm.Window.Update()
	}
}

func (gm *GameMonitor) EvaluateAction(action string) {
	// simulate game logic based on action
	score := gm.GameState.Score
	if action == " Move Up" {
		score += 1
	} else if action == " Move Down" {
		score -= 1
	}

	gm.GameState.Score = score
	gm.AI.LastReward = gm.AI.Evaluation(gm.GameState)
}

func main() {
	config := GameConfig{
		Width:  640,
		Height: 480,
		FPS:   60,
	}

	gameMonitor := NewGameMonitor(config)
	gameMonitor.Run()
}