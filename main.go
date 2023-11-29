package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/game"
	"github.com/dwethmar/judoka/system"
	"github.com/dwethmar/judoka/system/actor"
	"github.com/dwethmar/judoka/system/actor/player"
	"github.com/dwethmar/judoka/system/camera"
	"github.com/dwethmar/judoka/system/input"
	"github.com/dwethmar/judoka/system/render"
	"github.com/dwethmar/judoka/system/terrain"
	"github.com/dwethmar/judoka/system/terrain/perlin"
	"github.com/dwethmar/judoka/system/velocity"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	WindowWidth        = 800
	WindowHeight       = 600
	PositionResolution = 10
)

func main() {
	// game
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Judoka")

	// register
	register, err := registry.New()
	if err != nil {
		log.Fatal(err)
	}

	// logger
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)
	terrainGenerator := perlin.New()
	// terrainGenerator := debug.New()

	inputSystem := input.New(input.Options{
		Logger:             logger,
		Register:           register,
		PositionResolution: PositionResolution,
	})

	// drawing systems
	terrainSystem := terrain.New(terrain.Options{
		Logger:             logger,
		Register:           register,
		PositionResolution: PositionResolution,
		Generator:          terrainGenerator,
	})

	renderSystem := render.New(render.Options{
		Logger:             logger,
		Register:           register,
		PositionResolution: PositionResolution,
	})

	// other systems
	velocitySystem := velocity.New(velocity.Options{
		Logger:   logger,
		Register: register,
	})

	// actor syb systems
	playerSubSystem := player.New(player.Options{
		Logger:   logger,
		Register: register,
	})

	actorSystem := actor.New(actor.Options{
		Logger:             logger,
		Register:           register,
		PositionResolution: PositionResolution,
		ActorSubSystems:    []actor.SubSystem{playerSubSystem},
	})

	cameraSystem := camera.New(camera.Options{
		Logger:             logger,
		Register:           register,
		PositionResolution: PositionResolution,
		Viewport:           register.Root(),
	})

	inputSystem.Init()
	terrainSystem.Init(cameraSystem.Camera())
	renderSystem.Init()
	velocitySystem.Init()
	actorSystem.Init(cameraSystem.Camera())
	cameraSystem.Init()

	// systems
	var systems []system.System = []system.System{
		inputSystem,
		terrainSystem,
		renderSystem,
		velocitySystem,
		actorSystem,
		cameraSystem,
	}

	if err := ebiten.RunGame(
		game.New(slog.Default(), systems),
	); err != nil {
		log.Fatal(err)
	}
}
