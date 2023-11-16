package registry

import (
	"fmt"

	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
)

// AddSprite adds a sprite.
func (r *Registry) AddSprite(c *component.Sprite) error {
	c.CID = r.SpriteIDGenerator.Next()
	if err := r.Sprite.AddComponent(c); err != nil {
		return fmt.Errorf("failed to add transform: %w", err)
	}
	return nil
}

// RemoveTransform removes a transform.
func (r *Registry) RemoveSprite(c *component.Sprite) { r.Sprite.RemoveComponent(c) }

// ListSprites returns a list of all sprites.
func (r *Registry) ListSprites() []*component.Sprite {
	sprites := []*component.Sprite{}
	for _, e := range r.Sprite.ListEntities() {
		if s := r.Sprite.GetComponents(e); s != nil {
			sprites = append(sprites, s...)
		}
	}

	return sprites
}

// ListSprites returns a list of all sprites.
func (r *Registry) ListSpritesByEntity(e entity.Entity) []*component.Sprite {
	return r.Sprite.GetComponents(e)
}
