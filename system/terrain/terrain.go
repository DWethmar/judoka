package terrain

import (
	"fmt"
	"image"
	"log/slog"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/level"
	"github.com/dwethmar/judoka/system"
	"github.com/dwethmar/judoka/transform"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

const (
	ChunkSize = 16
	TileSize  = 32
)

var _ system.System = &System{}

type System struct {
	initialized   bool
	logger        *slog.Logger
	registry      *registry.Registry
	level         *level.Level  // used to link chunks together
	terrainEntity entity.Entity // used to group all all chunk entities
}

func New(logger *slog.Logger, registry *registry.Registry) *System {
	return &System{
		logger:   logger.WithGroup("terrain"),
		registry: registry,
		level:    level.New(ChunkSize),
	}
}

func (s *System) init() error {
	if s.initialized {
		return nil
	}

	terrainEntity, err := s.registry.Create(s.registry.Root())
	if err != nil {
		return fmt.Errorf("failed to create terrain entity: %w", err)
	}

	s.terrainEntity = terrainEntity
	s.initialized = true

	return nil
}

// Draw implements system.System.
func (s *System) Draw(screen *ebiten.Image) error {
	return nil
}

// GenerateChunk generates a chunk at the given chunk position.
func (s *System) GenerateChunk(chunkX, chunkY int) error {
	if s.level.GetChunk(chunkX, chunkY) != nil {
		return nil
	}

	e, err := s.registry.Create(s.terrainEntity)
	if err != nil {
		return fmt.Errorf("failed to create chunk in registry: %w", err)
	}

	// set position
	t, ok := s.registry.Transform.First(e)
	if ok {
		t.X = (chunkX * ChunkSize * TileSize) * system.PositionResolution
		t.Y = (chunkY * ChunkSize * TileSize) * system.PositionResolution
	}

	// Create chunk
	c := component.NewChunk(0, e)
	c.X = chunkX
	c.Y = chunkY
	c.Tiles = Generate(image.Rect(chunkX*ChunkSize, chunkY*ChunkSize, chunkX*ChunkSize+ChunkSize, chunkY*ChunkSize+ChunkSize))
	if err := s.registry.Chunk.Add(c); err != nil {
		return fmt.Errorf("failed to add chunk: %w", err)
	}

	// Add chunk to level
	s.level.SetChunk(chunkX, chunkY, c.Tiles)

	// Create sprite
	img := ebiten.NewImage(ChunkSize*TileSize, ChunkSize*TileSize)
	spr := component.NewSprite(0, e, 0, 0, img)
	if err := s.registry.Sprite.Add(spr); err != nil {
		return fmt.Errorf("failed to add sprite: %w", err)
	}

	if err := s.DrawChunk(c, img); err != nil {
		return fmt.Errorf("failed to draw chunk: %w", err)
	}

	s.logger.Info("generated chunk", slog.Int("x", chunkX), slog.Int("y", chunkY))

	return nil
}

func (s *System) DrawChunk(c *component.Chunk, screen *ebiten.Image) error {
	for i := 0; i < ChunkSize; i++ { // x
		x := i * TileSize
		for j := 0; j < ChunkSize; j++ { // y
			y := j * TileSize
			tile := c.Tiles.Get(i, j)

			image := Shapes(i+(ChunkSize*c.X), j+(ChunkSize*c.Y), s.level)
			if image == nil {
				// skip drawing
				goto drawdebug
			}

			{
				w := image.Bounds().Max.X - image.Bounds().Min.X
				h := image.Bounds().Max.Y - image.Bounds().Min.Y

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(float64(TileSize)/float64(w), float64(TileSize)/float64(h))
				op.GeoM.Translate(float64(x), float64(y))
				screen.DrawImage(image, op)
			}
		drawdebug:
			wX := i + c.X*ChunkSize
			wY := j + c.Y*ChunkSize
			text.Draw(screen, fmt.Sprintf("T%d\nX%d\nY%d", tile, wX, wY), assets.GetVGAFonts(1), x+2, y+8, colornames.Yellow700)
		}
	}

	return nil
}

// Update implements system.System.
func (s *System) Update() error {
	if err := s.init(); err != nil {
		return fmt.Errorf("failed to init: %w", err)
	}

	for _, e := range s.registry.Actor.Entities() {
		x, y := transform.Position(s.registry, e)

		chunkX := (x / system.PositionResolution) / (ChunkSize * TileSize)
		chunkY := (y / system.PositionResolution) / (ChunkSize * TileSize)

		if err := s.GenerateChunk(chunkX, chunkY); err != nil {
			return fmt.Errorf("failed to generate chunk: %w", err)
		}
	}

	return nil
}
