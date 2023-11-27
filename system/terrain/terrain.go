package terrain

import (
	"fmt"
	"log/slog"
	"math"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/level"
	"github.com/dwethmar/judoka/system"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/exp/shiny/materialdesign/colornames"
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
	PositionResolution int // used to divide X and Y positions
	defaultGenerator   Generator
	level              *level.Level  // used to link chunks together
	terrainEntity      entity.Entity // used to group all all chunk entities
	debug              bool          // used to draw debug info
	waterBackground1   *ebiten.Image
	waterBackground2   *ebiten.Image
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
func (s *System) Init() error {
	if s.initialized {
		return nil
	}

	terrainEntity, err := s.register.Create(s.register.Root())
	if err != nil {
		return fmt.Errorf("failed to create terrain entity: %w", err)
	}

	s.terrainEntity = terrainEntity
	s.initialized = true

	// draw water background
	s.waterBackground1 = ebiten.NewImage(ChunkSize*TileSize, ChunkSize*TileSize)
	for i := 0; i < ChunkSize; i++ { // x
		x := i * TileSize
		for j := 0; j < ChunkSize; j++ { // y
			y := j * TileSize

			image := assets.WaterImg.SubImage(assets.WaterCells[0][0]).(*ebiten.Image)
			{
				w := image.Bounds().Dx()
				h := image.Bounds().Dy()

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(float64(TileSize)/float64(w), float64(TileSize)/float64(h))
				op.GeoM.Translate(float64(x), float64(y))
				s.waterBackground1.DrawImage(image, op)
			}
		}
	}

	s.waterBackground2 = ebiten.NewImage(ChunkSize*TileSize, ChunkSize*TileSize)
	for i := 0; i < ChunkSize; i++ { // x
		x := i * TileSize
		for j := 0; j < ChunkSize; j++ { // y
			y := j * TileSize

			image := assets.WaterImg.SubImage(assets.WaterCells[1][0]).(*ebiten.Image)
			{
				w := image.Bounds().Dx()
				h := image.Bounds().Dy()

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(float64(TileSize)/float64(w), float64(TileSize)/float64(h))
				op.GeoM.Translate(float64(x), float64(y))
				s.waterBackground2.DrawImage(image, op)
			}
		}
	}

	return nil
}

// Draw implements system.System.
func (s *System) Draw(screen *ebiten.Image) error {
	return nil
}

// Update implements system.System.
func (s *System) Update() error {
	for _, e := range s.register.Actor.Entities() {
		var x, y int
		if t, ok := s.register.Transform.First(e); ok {
			x = t.X
			y = t.Y
		}

		// Assign a value to chunkX
		chunkX := int(math.Floor(float64(x) / (float64(s.PositionResolution) * float64(ChunkSize*TileSize))))
		chunkY := int(math.Floor(float64(y) / (float64(s.PositionResolution) * float64(ChunkSize*TileSize))))

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

			// 3. Draw chunk, draws it to its sprite component
			if err := s.DrawChunk(c); err != nil {
				return fmt.Errorf("failed to draw chunk: %w", err)
			}

			// 4. Redraw surrounding chunks, so they connect to the new chunk
			if err := s.RedrawNeighboringChunks(chunkX, chunkY); err != nil {
				return fmt.Errorf("failed to redraw neighboring chunks: %w", err)
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

	// Create background sprite
	{
		spr := component.NewSprite(0, c.Entity(), 0, 0, s.waterBackground1)
		spr.Name = "background"
		if err := s.register.Sprite.Add(spr); err != nil {
			return fmt.Errorf("failed to add sprite: %w", err)
		}
	}

	// Create sprite
	{
		img := ebiten.NewImage(ChunkSize*TileSize, ChunkSize*TileSize)
		spr := component.NewSprite(0, c.Entity(), 0, 0, img)
		spr.Name = "terrain"
		if err := s.register.Sprite.Add(spr); err != nil {
			return fmt.Errorf("failed to add sprite: %w", err)
		}
	}

	s.logger.Info("created chunk", slog.Group("chunk", slog.Int("x", chunkX), slog.Int("y", chunkY)))

	return nil
}

// DrawChunk draws a chunk to its sprite component.
// it is required that the chunk has a sprite component.
func (s *System) DrawChunk(c *component.Chunk) error {
	var sprite *component.Sprite
	for _, spr := range s.register.Sprite.List(c.Entity()) {
		if spr.Name == "terrain" {
			sprite = spr
			break
		}
	}

	if sprite == nil {
		return fmt.Errorf("sprite not found")
	}

	if sprite.Image == nil {
		return fmt.Errorf("sprite image is nil")
	}

	sprite.Image.Clear()

	for i := 0; i < ChunkSize; i++ { // x
		x := i * TileSize
		for j := 0; j < ChunkSize; j++ { // y
			y := j * TileSize

			tile := c.Tiles.Get(i, j, -1)
			if tile <= 0 { // we skip drawing water
				continue
			}

			image := Shapes(i+(ChunkSize*c.X), j+(ChunkSize*c.Y), s.level)
			if image == nil {
				// skip drawing
				goto drawdebug
			}
			{
				w := image.Bounds().Dx()
				h := image.Bounds().Dy()

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(float64(TileSize)/float64(w), float64(TileSize)/float64(h))
				op.GeoM.Translate(float64(x), float64(y))
				sprite.Image.DrawImage(image, op)
			}
		drawdebug:
			if s.debug {
				wX := i + c.X*ChunkSize
				wY := j + c.Y*ChunkSize
				text.Draw(sprite.Image, fmt.Sprintf("T%d\nX%d\nY%d", tile, wX, wY), assets.GetVGAFonts(1), x+2, y+8, colornames.Red300)
			}
		}
	}

	s.logger.Info("created chunk", slog.Group("chunk", slog.Int("x", c.X), slog.Int("y", c.Y)))

	return nil
}

// RedrawNeighboringChunks redraws all existing chunks around the given chunk index.
func (s *System) RedrawNeighboringChunks(chunkX, chunkY int) error {
	for x := chunkX - 1; x <= chunkX+1; x++ {
		for y := chunkY - 1; y <= chunkY+1; y++ {
			if x == chunkX && y == chunkY {
				continue
			}

			c := s.level.Chunk(x, y)
			if c == nil {
				continue
			}

			if err := s.DrawChunk(c); err != nil {
				return fmt.Errorf("failed to draw chunk: %w", err)
			}
		}
	}

	return nil
}
