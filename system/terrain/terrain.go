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
	"github.com/dwethmar/judoka/tilebitmasking"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ChunkSize = 16
	TileSize  = 32
)

// Options are used to configure a new terrain system.
type Options struct {
	Logger             *slog.Logger
	Register           *registry.Register
	PositionResolution int
	Generator          Generator
}

var _ system.System = &System{}

// System is a terrain system.
// It is responsible for generating and drawing and linking chunks.
type System struct {
	initialized        bool
	logger             *slog.Logger
	register           *registry.Register
	camera             *system.Camera
	PositionResolution int // used to divide X and Y positions
	defaultGenerator   Generator
	level              *level.Level  // used to link chunks together
	terrainEntity      entity.Entity // used to group all all chunk entities
	debug              bool          // used to draw debug info
}

func New(
	opt Options,
) *System {
	return &System{
		logger:             opt.Logger.WithGroup("terrain"),
		register:           opt.Register,
		PositionResolution: opt.PositionResolution,
		defaultGenerator:   opt.Generator,
		level:              level.New(ChunkSize),
		debug:              false,
	}
}

// init initializes the system.
func (s *System) Init(camera *system.Camera) error {
	if s.initialized {
		return nil
	}

	s.camera = camera

	terrainEntity, err := s.register.Create(s.register.Root())
	if err != nil {
		return fmt.Errorf("failed to create terrain entity: %w", err)
	}

	s.terrainEntity = terrainEntity
	s.initialized = true

	return nil
}

// Draw implements system.System.
func (s *System) Draw(screen *ebiten.Image) error {
	minTileX := (s.camera.Bounds.Min.X / TileSize) - 1
	minTileY := (s.camera.Bounds.Min.Y / TileSize) - 1

	maxTileX := (s.camera.Bounds.Max.X / TileSize) + 1
	maxTileY := (s.camera.Bounds.Max.Y / TileSize) + 1

	for i := minTileX; i < maxTileX; i++ { // x
		chunkX := i / ChunkSize

		for j := minTileY; j < maxTileY; j++ { // y
			chunkY := j / ChunkSize

			chunk := s.level.Chunk(chunkX, chunkY)
			if chunk == nil {
				continue
			}

			neighbors := Neighbors(i, j, s.level)
			bitmask := tilebitmasking.Calculate(neighbors)

			if bitmask != tilebitmasking.AllEdges {
				s.DrawTile(screen, assets.Water1, i, j)
			}

			image := TileImage(neighbors.Center, bitmask)
			if image == nil {
				continue
			}

			s.DrawTile(screen, image, i, j)

			// text.Draw(screen, fmt.Sprintf("X%d\nY%d", i, j), assets.GetVGAFonts(1), int(dx)+1, int(dy)+7, colornames.Black)
			// text.Draw(screen, fmt.Sprintf("X%d\nY%d", i, j), assets.GetVGAFonts(1), int(dx)+2, int(dy)+8, colornames.Yellow500)
		}
	}

	return nil
}

// DrawTile draws a tile at the given tile index.
func (s *System) DrawTile(screen *ebiten.Image, image *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear

	w := image.Bounds().Dx()
	h := image.Bounds().Dy()

	op.GeoM.Scale(float64(TileSize)/float64(w), float64(TileSize)/float64(h))

	// Translate to tile position
	dx := float64(x * TileSize)
	dy := float64(y * TileSize)

	// Translate to camera position
	dx -= float64(s.camera.Bounds.Min.X)
	dy -= float64(s.camera.Bounds.Min.Y)

	op.GeoM.Translate(dx, dy)
	screen.DrawImage(image, op)
}

// Update implements system.System.
func (s *System) Update() error {
	minChunkX := (s.camera.Bounds.Min.X / (ChunkSize * TileSize)) - 1
	minChunkY := (s.camera.Bounds.Min.Y / (ChunkSize * TileSize)) - 1

	maxChunkX := (s.camera.Bounds.Max.X / (ChunkSize * TileSize)) + 1
	maxChunkY := (s.camera.Bounds.Max.Y / (ChunkSize * TileSize)) + 1

	for i := minChunkX; i < maxChunkX; i++ { // x
		for j := minChunkY; j < maxChunkY; j++ { // y
			// Create chunk if it does not exist
			if s.level.Chunk(i, j) == nil {
				// 1. Create chunk
				c, err := s.CreateChunk(s.level, i, j)
				if err != nil {
					return fmt.Errorf("failed to create chunk: %w", err)
				}

				// 2. Generate chunk
				if err := s.GenerateChunk(c, s.defaultGenerator); err != nil {
					return fmt.Errorf("failed to generate chunk: %w", err)
				}
			}
		}
	}

	return nil
}

// CreateChunk creates a chunk at the given chunk position.
func (s *System) CreateChunk(l *level.Level, chunkX, chunkY int) (*component.Chunk, error) {
	e, err := s.register.Create(s.terrainEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to create chunk in registry: %w", err)
	}

	// set position
	t, ok := s.register.Transform.First(e)
	if ok {
		t.X = (chunkX * ChunkSize * TileSize) * s.PositionResolution
		t.Y = (chunkY * ChunkSize * TileSize) * s.PositionResolution
	}

	// Create chunk
	c := component.NewChunk(0, e)
	c.X = chunkX
	c.Y = chunkY

	if err := s.register.Chunk.Add(c); err != nil {
		return nil, fmt.Errorf("failed to add chunk component to chunk store: %w", err)
	}

	// Add chunk to level
	l.SetChunk(c)
	s.logger.Info("created chunk", slog.Group("chunk", slog.Int("x", chunkX), slog.Int("y", chunkY)))

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

	return nil
}
