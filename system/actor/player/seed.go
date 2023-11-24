package player

import (
	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/registry"
)

func AddPlayer(r *registry.Register, parent entity.Entity) error {
	e, err := r.Create(parent)
	if err != nil {
		return err
	}

	velocity := component.NewVelocity(0, e, 0, 0)
	if err := r.Velocity.Add(velocity); err != nil {
		return err
	}

	controller := component.NewController(0, e)
	if err := r.Controller.Add(controller); err != nil {
		return err
	}

	actor := component.NewActor(0, e)
	actor.ActorType = component.ActorTypePlayer
	if err := r.Actor.Add(actor); err != nil {
		return err
	}

	layer := component.NewLayer(0, e)
	layer.Index = 1
	if err := r.Layer.Add(layer); err != nil {
		return err
	}

	return nil
}
