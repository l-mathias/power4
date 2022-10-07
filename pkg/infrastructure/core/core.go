package core

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"power4/pkg/domain"
	"power4/pkg/infrastructure/local"
	"sync"
	"time"
)

const (
	debug          bool = false
	roundOverScore      = 10
)

type Change interface{}

type RoundOverChange struct {
	Change
	Reason string
}

type RoundStartChange struct {
	Change
}

type AddPlayerChange struct {
	Change
	Player Player
}

type RemovePlayerChange struct {
	Change
	Player Player
}

type Action interface {
	Perform(game *Game)
}

type Message struct {
	From string
	Msg  string
}

type MessageAction struct {
	Message Message
	ID      uuid.UUID
	Created time.Time
}

type MessageChange struct {
	Change
	ID      uuid.UUID
	Message Message
	Created time.Time
}

type Player struct {
	UUID uuid.UUID
	Name string
}

type Game struct {
	Grid          domain.Grid
	Player1       domain.Player
	Player2       domain.Player
	Players       map[uuid.UUID]Player
	Mu            sync.RWMutex
	ChangeChannel chan Change
	ActionChannel chan Action
	Score         map[uuid.UUID]int
	RoundWinner   uuid.UUID
	WaitForRound  bool
}

func NewGame() *Game {
	game := Game{
		Grid:          local.NewGrid(),
		Player1:       local.NewPlayer('x'),
		Player2:       local.NewPlayer('o'),
		Players:       make(map[uuid.UUID]Player),
		ActionChannel: make(chan Action, 1),
		ChangeChannel: make(chan Change, 1),
		WaitForRound:  true,
		Score:         make(map[uuid.UUID]int),
	}
	return &game
}

func (game *Game) DisplayGame() {
	game.LogDebug(fmt.Sprintf("Displaying Game :\nNumber of player: %v\n", len(game.Players)))
	for _, p := range game.Players {
		game.LogDebug(fmt.Sprintf("Player: %v\n", p.Name))
	}
}

func (game *Game) Start() {
	go game.watchActions()
}

func (game *Game) watchActions() {
	for {
		action := <-game.ActionChannel
		if game.WaitForRound {
			continue
		}
		action.Perform(game)
	}
}

func (p *Player) ID() uuid.UUID {
	return p.UUID
}

func (game *Game) GetPlayer(id uuid.UUID) Player {
	return game.Players[id]
}

func (game *Game) PlayerExists(id uuid.UUID) bool {
	game.Mu.Lock()
	if _, ok := game.Players[id]; ok {
		return true
	}
	game.Mu.Unlock()
	return false
}

func (game *Game) AddPlayer(player *Player) {
	game.Players[player.ID()] = *player
}

func (game *Game) RemovePlayer(id uuid.UUID) {
	delete(game.Players, id)
}

func (message MessageAction) Perform(game *Game) {
	change := MessageChange{
		ID:      message.ID,
		Message: message.Message,
		Created: message.Created,
	}
	game.sendChange(change)
}

func (game *Game) sendChange(change Change) {
	select {
	case game.ChangeChannel <- change:
	default:
	}
}

func (game *Game) LogDebug(msg string) {
	if debug {
		log.Printf(msg)
	}
}

func (game *Game) QueueNewRound(roundWinner uuid.UUID, reason string) {
	game.WaitForRound = true
	game.RoundWinner = roundWinner
	game.sendChange(RoundOverChange{
		Reason: reason,
	})
}

func (game *Game) StartNewRound() {
	game.WaitForRound = false
	game.Score = map[uuid.UUID]int{}
	game.sendChange(RoundStartChange{})
}

func (game *Game) AddScore(id uuid.UUID) {
	game.Score[id]++
}
