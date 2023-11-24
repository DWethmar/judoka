package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/game"
	"github.com/dwethmar/judoka/system"
	"github.com/dwethmar/judoka/system/actor"
	"github.com/dwethmar/judoka/system/actor/player"
	"github.com/dwethmar/judoka/system/input"
	"github.com/dwethmar/judoka/system/render"
	"github.com/dwethmar/judoka/system/terrain"
	"github.com/dwethmar/judoka/system/terrain/debug"
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

	// registry
	registry := registry.New()
	AddPlayer(registry)
	// AddTestEntity2(registry, p)

	// logger
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)
	// terrainGenerator := perlin.New()
	terrainGenerator := debug.New()

	// systems
	var systems []system.System = []system.System{
		input.New(input.Options{
			Logger:             logger,
			Registry:           registry,
			PositionResolution: PositionResolution,
		}),
		// drawing systems
		terrain.New(terrain.Options{
			Logger:             logger,
			Registry:           registry,
			PositionResolution: PositionResolution,
			Generator:          terrainGenerator,
		}),
		render.New(render.Options{
			Logger:             logger,
			Registry:           registry,
			PositionResolution: PositionResolution,
		}),
		// other systems
		velocity.New(velocity.Options{
			Logger:   logger,
			Registry: registry,
		}),
		actor.New(actor.Options{
			Logger:             logger,
			Registry:           registry,
			PositionResolution: PositionResolution,
			Managers: map[component.ActorType]actor.Manager{
				component.ActorTypePlayer: player.New(
					player.Options{
						Logger:   logger,
						Registry: registry,
					},
				),
			},
		}),
	}

	if err := ebiten.RunGame(
		game.New(slog.Default(), systems),
	); err != nil {
		log.Fatal(err)
	}
}

func AddPlayer(r *registry.Registry) entity.Entity {
	e, err := r.Create(r.Root())
	if err != nil {
		log.Fatal(err)
	}

	velocity := component.NewVelocity(0, e, 0, 0)
	if err := r.Velocity.Add(velocity); err != nil {
		log.Fatal(err)
	}

	controller := component.NewController(0, e)
	if err := r.Controller.Add(controller); err != nil {
		log.Fatal(err)
	}

	actor := component.NewActor(0, e)
	actor.ActorType = component.ActorTypePlayer
	if err := r.Actor.Add(actor); err != nil {
		log.Fatal(err)
	}

	return e
}

func AddTestEntity2(r *registry.Registry, p entity.Entity) entity.Entity {
	e, err := r.Create(p)
	if err != nil {
		log.Fatal(err)
	}

	velocity := component.NewVelocity(0, e, 0, 0)
	if err := r.Velocity.Add(velocity); err != nil {
		log.Fatal(err)
	}

	sprite := component.NewSprite(0, e, 0, 0, assets.SkeletonKill6)
	if err := r.Sprite.Add(sprite); err != nil {
		log.Fatal(err)
	}

	return e
}
