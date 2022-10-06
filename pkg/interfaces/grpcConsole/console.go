package grpcConsole

import (
	"fmt"
	"power4/pkg/domain"
	"power4/pkg/infrastructure/proto"
	"strconv"
	"strings"
)

type Console struct {
}

func printColumns() string {
	var printing string
	for i := 1; i < 8; i++ {
		printing += fmt.Sprintf("   %d", i)
	}
	printing += fmt.Sprintln()

	return printing
}

func (s *GameServer) ShowGrid(grid domain.Grid) {
	printing := fmt.Sprintln()
	printing += printColumns()
	grid.Process(func(matrix [6][7]rune) {
		for _, line := range matrix {
			for _, symbole := range line {
				printing += fmt.Sprint(" | ", string(symbole))
			}
			printing += fmt.Sprintln(" | ")
		}
	})
	printing += printColumns()

	resp := proto.StreamResponse{
		Event: &proto.StreamResponse_ResponseMessage{
			ResponseMessage: &proto.Message{
				From:    "Server",
				Message: printing,
			},
		},
	}
	s.broadcast(&resp)
}

func (s *GameServer) AskColumnTo(player domain.Player) (int, error) {
	normalizeColumn := func(col rune, _ int, err error) (int, error) {
		if err != nil {
			return -1, err
		}

		result, err := strconv.Atoi(string(col))
		if err != nil {
			return -1, fmt.Errorf("enter a number you dumbass!")
		}

		if result < 1 || result > 7 {
			return -1, fmt.Errorf("enter a number between 1 and 7, you dumbass!")
		}

		return result - 1, nil
	}

	msg := fmt.Sprintln()
	msg += fmt.Sprintf("player '%s' enter column: ", string(player.GetSymbol()))
	s.send("Server", s.GetActiveClient().id.String(), msg)

	event := <-s.GameChan
	m := event.GetResponseMessage().Message

	reader := strings.NewReader(m)
	col, err := normalizeColumn(reader.ReadRune())
	if err != nil {
		return -1, err
	}

	return col, nil
}

func (s *GameServer) Congratulate(player domain.Player) {
	msg := fmt.Sprintln()
	msg += fmt.Sprintf("You're just a looser '%s'...", string(player.GetSymbol()))
	msg += fmt.Sprintln()

	s.send("Server", s.GetActiveClient().id.String(), msg)

	msg = fmt.Sprintln()
	msg += fmt.Sprintf("congratulations player '%s' you win!", string(player.GetSymbol()))
	msg += fmt.Sprintln()

	s.send("Server", s.GetInactiveClient().id.String(), msg)
}

func (s *GameServer) Slap(player domain.Player, err error) {

	msg := fmt.Sprintln()
	msg += fmt.Sprintf("player '%s' %s", string(player.GetSymbol()), err.Error())
	msg += fmt.Sprintln()

	s.invertActive()

	s.send("Server", s.GetActiveClient().id.String(), msg)
}

func (s *GameServer) NoWinner() {

	msg := fmt.Sprintln()
	msg += fmt.Sprintln("nobody wins, losers!")
	msg += fmt.Sprintln()

	resp := proto.StreamResponse{
		Event: &proto.StreamResponse_ResponseMessage{
			ResponseMessage: &proto.Message{
				From:    "Server",
				Message: msg,
			},
		},
	}
	s.broadcast(&resp)

}
