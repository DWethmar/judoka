package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/game"
	"github.com/dwethmar/judoka/system"
	"github.com/dwethmar/judoka/system/debug"
	"github.com/dwethmar/judoka/system/input"
	"github.com/dwethmar/judoka/system/render"
	"github.com/dwethmar/judoka/system/velocity"
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
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)

	// systems
	var systems []system.System = []system.System{
		input.New(logger, registry),
		// drawing systems
		render.New(logger, registry),
		debug.New(logger, registry),
		// state systems
		velocity.New(logger, registry),
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

	velocity := component.NewVelocity(0, e, 0, 0)
	if err := r.AddVelocity(velocity); err != nil {
		log.Fatal(err)
	}

	sprite := component.NewSprite(0, e, 0, 0, assets.SkeletonDown1)
	if err := r.AddSprite(sprite); err != nil {
		log.Fatal(err)
	}

	controller := component.NewController(0, e)
	if err := r.AddController(controller); err != nil {
		log.Fatal(err)
	}
}
