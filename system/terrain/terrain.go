package terrain

import (
	"fmt"
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
	initialized      bool
	logger           *slog.Logger
	registry         *registry.Registry
	defaultGenerator Generator
	level            *level.Level  // used to link chunks together
	terrainEntity    entity.Entity // used to group all all chunk entities
}

func New(
	logger *slog.Logger,
	registry *registry.Registry,
	generator Generator,
) *System {
	return &System{
		logger:           logger.WithGroup("terrain"),
		registry:         registry,
		defaultGenerator: generator,
		level:            level.New(ChunkSize),
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

// Update implements system.System.
func (s *System) Update() error {
	if err := s.init(); err != nil {
		return fmt.Errorf("failed to init: %w", err)
	}

	for _, e := range s.registry.Actor.Entities() {
		x, y := transform.Position(s.registry, e)

		chunkX := (x / system.PositionResolution) / (ChunkSize * TileSize)
		chunkY := (y / system.PositionResolution) / (ChunkSize * TileSize)

		// Create chunk if it does not exist
		if s.level.Chunk(chunkX, chunkY) == nil {
			// 1. Create chunk
			c, err := s.CreateChunk(s.level, chunkX, chunkY)
			if err != nil {
				return fmt.Errorf("failed to create chunk: %w", err)
			}

			// 2. Generate chunk
			if err := s.GenerateChunk(c, s.defaultGenerator); err != nil {
				return fmt.Errorf("failed to generate chunk: %w", err)
			}

			// 3. Draw chunk
			sprite, ok := s.registry.Sprite.First(c.Entity())
			if !ok {
				return fmt.Errorf("failed to get sprite from chunk")
			}

			if err := s.DrawChunk(c, sprite); err != nil {
				return fmt.Errorf("failed to draw chunk: %w", err)
			}
		}
	}

	return nil
}

// CreateChunk creates a chunk at the given chunk position.
func (s *System) CreateChunk(l *level.Level, chunkX, chunkY int) (*component.Chunk, error) {
	e, err := s.registry.Create(s.terrainEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to create chunk in registry: %w", err)
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

	if err := s.registry.Chunk.Add(c); err != nil {
		return nil, fmt.Errorf("failed to add chunk component to chunk store: %w", err)
	}

	// Add chunk to level
	l.SetChunk(c)
	s.logger.Info("created chunk", slog.Int("x", chunkX), slog.Int("y", chunkY))

	return c, nil
}

// GenerateChunk generates a chunk at the given chunk position.
func (s *System) GenerateChunk(c *component.Chunk, generator Generator) error {
	chunkX, chunkY := c.X, c.Y

	c.Tiles = generator.Generate(
		chunkX*ChunkSize,
		(chunkX*ChunkSize)+ChunkSize,
		chunkY*ChunkSize,
		(chunkY*ChunkSize)+ChunkSize,
	)

	// Create sprite
	img := ebiten.NewImage(ChunkSize*TileSize, ChunkSize*TileSize)
	spr := component.NewSprite(0, c.Entity(), 0, 0, img)
	if err := s.registry.Sprite.Add(spr); err != nil {
		return fmt.Errorf("failed to add sprite: %w", err)
	}

	s.logger.Info("generated chunk", slog.Int("x", chunkX), slog.Int("y", chunkY))

	return nil
}

func (s *System) DrawChunk(c *component.Chunk, g *component.Sprite) error {
	g.Image.Clear()

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
				g.Image.DrawImage(image, op)
			}
		drawdebug:
			wX := i + c.X*ChunkSize
			wY := j + c.Y*ChunkSize
			text.Draw(g.Image, fmt.Sprintf("T%d\nX%d\nY%d", tile, wX, wY), assets.GetVGAFonts(1), x+2, y+8, colornames.Yellow700)
		}
	}

	s.logger.Info("draw chunk", slog.Int("x", c.X), slog.Int("y", c.Y))

	return nil
}
