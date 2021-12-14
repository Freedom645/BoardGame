package player_type

import "errors"

type PlayerType string

const (
	Spectator = PlayerType("spectator")
	Player    = PlayerType("player")
)

func Of(pt string) (PlayerType, error) {
	switch pt {
	case string(Spectator):
		return Spectator, nil
	case string(Player):
		return Player, nil
	}

	return "", errors.New("unknown player type")
}
