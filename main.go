package main

import (
	"log"
	"log/slog"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/game"
	"github.com/dwethmar/judoka/system"
	"github.com/dwethmar/judoka/system/render"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	WindowWidth  = 800
	WindowHeight = 600
)

func main() {
	// game
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Judoka")

	// registry
	registry := registry.New()
	AddTestEntity1(registry)

	// logger
	logger := slog.Default()

	// systems
	var systems []system.System = []system.System{
		render.New(logger, registry),
	}

	if err := ebiten.RunGame(
		game.New(slog.Default(), systems),
	); err != nil {
		log.Fatal(err)
	}
}

func AddTestEntity1(r *registry.Registry) {
	e, err := r.CreateEntity()
	if err != nil {
		log.Fatal(err)
	}

	transform := component.NewTransform(0, e, 100, 100)
	if err := r.AddTransform(transform); err != nil {
		log.Fatal(err)
	}

	sprite := component.NewSprite(0, e, 0, 0, assets.SkeletonDown1)
	if err := r.AddSprite(sprite); err != nil {
		log.Fatal(err)
	}
}
