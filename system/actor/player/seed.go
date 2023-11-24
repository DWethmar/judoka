package player

import (
	"log"

	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/registry"
)

func AddPlayer(r *registry.Register, parent entity.Entity) entity.Entity {
	e, err := r.Create(parent)
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
