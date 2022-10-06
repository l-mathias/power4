package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"power4/pkg/domain"
	"power4/pkg/infrastructure/core"
	"power4/pkg/infrastructure/local"
	"power4/pkg/infrastructure/proto"
	"power4/pkg/interfaces"
	"power4/pkg/interfaces/grpcConsole"
	"power4/pkg/usecase"
)

var (
	grid       domain.Grid
	player1    domain.Player
	player2    domain.Player
	game       usecase.Game
	intrface   interfaces.Interface
	serverGame *core.Game
)

func init() {
	grid = local.NewGrid()
	player1 = local.NewPlayer('x')
	player2 = local.NewPlayer('o')
	game = usecase.NewGame(grid, player1, player2)
}

func main() {

	port := flag.Int("port", 8080, "The port to listen on.")
	password := flag.String("password", "", "The pkg password.")
	flag.Parse()

	log.Printf("listening on port %d\n", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverGame = core.NewGame()

	serverGame.Start()

	s := grpc.NewServer()
	server := grpcConsole.NewGameServer(serverGame, *password)
	proto.RegisterGameServer(s, server)
	intrface = server

	go func() {
		for {
			if !serverGame.WaitForRound || serverGame.AskReset {
				log.Printf("\n\nserverGame.AskReset value: %v \n\n", serverGame.AskReset)
				serverGame.AskReset = false
				grid = local.NewGrid()
				player1 = local.NewPlayer('x')
				player2 = local.NewPlayer('o')
				game = usecase.NewGame(grid, player1, player2)
				game.Play(intrface)
			}
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
