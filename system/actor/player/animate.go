package player

import (
	"fmt"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/direction"
	"github.com/dwethmar/judoka/images"
	"github.com/hajimehoshi/ebiten/v2"
)

func (m *System) animate(actor *component.Actor) error {
	e := actor.Entity()

	// check if sprite exists
	sprite, ok := m.register.Sprite.First(actor.Entity())
	if !ok {
		// add sprite
		offsetX, offsetY := images.ActorOffsets(assets.SkeletonKill1)
		sprite = component.NewSprite(0, actor.Entity(), offsetX, offsetY, assets.SkeletonKill1)
		if err := m.register.Sprite.Add(sprite); err != nil {
			return fmt.Errorf("failed to add sprite: %w", err)
		}
	}

	transform, ok := m.register.Transform.First(e)
	if !ok {
		return nil
	}

	controller, ok := m.register.Controller.First(e)
	if !ok {
		return nil
	}

	isMoving := controller.X != 0 || controller.Y != 0

	if isMoving {
		d := direction.Get(transform.X, transform.Y, transform.X+controller.X, transform.Y+controller.Y)
		if d != actor.Facing {
			actor.AnimationFrame = 0
		}

		var frames []*ebiten.Image
		switch d {
		case direction.Top, direction.TopLeft, direction.TopRight:
			frames = assets.SkeletonMoveUpFrames
		case direction.Bottom, direction.BottomLeft, direction.BottomRight:
			frames = assets.SkeletonMoveDownFrames
		case direction.Left:
			frames = assets.SkeletonMoveLeftFrames
		case direction.Right:
			frames = assets.SkeletonMoveRightFrames
		default:
			frames = assets.SkeletonMoveDownFrames
		}
		actor.AnimationFrame = (actor.AnimationFrame + 1) % len(frames)
		actor.Facing = d
		sprite.Image = frames[actor.AnimationFrame]
	} else {
		actor.AnimationFrame = 0
		var image *ebiten.Image
		switch actor.Facing {
		case direction.Top, direction.TopLeft, direction.TopRight:
			image = assets.SkeletonUp1
		case direction.Bottom, direction.BottomLeft, direction.BottomRight:
			image = assets.SkeletonDown1
		case direction.Left:
			image = assets.SkeletonLeft1
		case direction.Right:
			image = assets.SkeletonRight1
		default:
			image = assets.SkeletonDown1
		}
		sprite.Image = image
	}

	// calculate offset
	offsetX, offsetY := images.ActorOffsets(sprite.Image)
	sprite.OffsetX = offsetX
	sprite.OffsetY = offsetY

	return nil
}
