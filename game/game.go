package game

import (
	"fmt"
	"log/slog"

	"github.com/dwethmar/judoka/system"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	logger  *slog.Logger
	systems []system.System
}

func New(logger *slog.Logger, systems []system.System) *Game {
	return &Game{
		logger:  logger,
		systems: systems,
	}
}

func (g *Game) Update() error {
	for _, s := range g.systems {
		if err := s.Update(); err != nil {
			return fmt.Errorf("error updating system: %w", err)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, s := range g.systems {
		if err := s.Draw(screen); err != nil {
			g.logger.Error(err.Error())
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
