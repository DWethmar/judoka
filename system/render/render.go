package render

import (
	"fmt"
	"log/slog"

	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system"
	"github.com/dwethmar/judoka/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MaxLayers = 5
)

var _ system.System = (*System)(nil)

// System is a render system.
type System struct {
	logger             *slog.Logger
	register           *registry.Register
	positionResolution int                           // used to divide X and Y positions
	layerSorters       map[int]func([]entity.Entity) // used to sort entities in layers
}

// Options are used to configure a new render system.
type Options struct {
	Logger             *slog.Logger
	Register           *registry.Register
	PositionResolution int
	LayerSorters       map[int]func([]entity.Entity)
}

// New creates a new Render system.
func New(
	opt Options,
) *System {
	return &System{
		logger:             opt.Logger,
		register:           opt.Register,
		positionResolution: opt.PositionResolution,
	}
}

func (s *System) Init() error {
	return nil
}

// Draw implements system.System.
func (r *System) Draw(screen *ebiten.Image) error {
	layers := make([][]entity.Entity, MaxLayers)

	for _, e := range r.register.Sprite.Entities() {
		layer := 0
		if l, ok := r.register.Layer.First(e); ok {
			layer = l.Index
		}

		// check if layers is out of bounds
		if layer < 0 {
			return fmt.Errorf("layer is smaller than 0: %d", layer)
		}

		if layer >= MaxLayers {
			return fmt.Errorf("layer is bigger than max layers(%d): %d", MaxLayers, layer)
		}

		if layers[layer] == nil {
			layers[layer] = []entity.Entity{}
		}

		layers[layer] = append(layers[layer], e)
	}

	for layer, entities := range layers {
		if len(entities) == 0 {
			continue
		}

		// sort entities
		if len(entities) > 0 {
			if sorter, ok := r.layerSorters[layer]; ok {
				sorter(entities)
			}
		}

		r.DrawEntities(screen, entities)
	}

	return nil
}

// DrawEntities draws a layer.
func (r *System) DrawEntities(screen *ebiten.Image, entities []entity.Entity) {
	for _, e := range entities {
		for _, sprite := range r.register.Sprite.List(e) {
			if sprite.Image == nil {
				continue
			}

			x, y := transform.Position(r.register, e)
			r.drawSprite(screen, x, y, sprite)
		}
	}
}

func (r *System) drawSprite(screen *ebiten.Image, x, y int, sprite *component.Sprite) {
	nX := x / r.positionResolution
	nY := y / r.positionResolution

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(nX+sprite.OffsetX),
		float64(nY+sprite.OffsetY),
	)

	// filter
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(sprite.Image, op)
}

// Update implements system.System.
func (*System) Update() error {
	return nil
}
