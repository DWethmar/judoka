package terrain

import (
	"fmt"
	"image"
	"log/slog"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/component"
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
	logger   *slog.Logger
	registry *registry.Registry
	level    *level.Level
}

func New(logger *slog.Logger, registry *registry.Registry) *System {
	return &System{
		logger:   logger.WithGroup("terrain"),
		registry: registry,
		level:    level.New(ChunkSize),
	}
}

// Draw implements system.System.
func (s *System) Draw(screen *ebiten.Image) error {
	for _, e := range s.registry.Chunk.Entities() {
		c, ok := s.registry.Chunk.First(e)
		if !ok {
			return fmt.Errorf("failed to get chunk")
		}

		x := c.X * ChunkSize * TileSize
		y := c.Y * ChunkSize * TileSize

		for i := 0; i < ChunkSize; i++ {
			for j := 0; j < ChunkSize; j++ {
				tile := c.Tiles.Get(i, j)

				image := Shapes(ChunkSize*c.X+i, ChunkSize*c.Y+j, s.level)
				if image == nil {
					// skip drawing
					goto drawdebug
				}

				{
					w := image.Bounds().Max.X - image.Bounds().Min.X
					h := image.Bounds().Max.Y - image.Bounds().Min.Y

					op := &ebiten.DrawImageOptions{}
					op.GeoM.Scale(float64(TileSize)/float64(w), float64(TileSize)/float64(h))
					op.GeoM.Translate(float64(x+i*TileSize), float64(y+j*TileSize))
					screen.DrawImage(image, op)
				}
			drawdebug:
				// tile number
				text.Draw(screen, fmt.Sprintf("%d", tile), assets.GetVGAFonts(2), x+i*TileSize+10, y+j*TileSize+12, colornames.Yellow700)
			}
		}

		// draw chunk x, y
		text.Draw(screen, fmt.Sprintf("%d, %d", c.X, c.Y), assets.GetVGAFonts(3), x+40, y+40, colornames.Blue100)
	}

	return nil
}

// Update implements system.System.
func (s *System) Update() error {
	for _, e := range s.registry.Actor.Entities() {
		x, y := transform.Position(s.registry, e)

		chunkX := (x / system.PositionResolution) / (ChunkSize * TileSize)
		chunkY := (y / system.PositionResolution) / (ChunkSize * TileSize)

		c := s.level.GetChunk(chunkX, chunkY)
		if c == nil {
			// create chunk
			e, err := s.registry.Create(s.registry.Root())
			if err != nil {
				return fmt.Errorf("failed to create chunk: %w", err)
			}

			c := component.NewChunk(0, e)
			c.X = chunkX
			c.Y = chunkY
			c.Tiles = Generate(image.Rect(chunkX*ChunkSize, chunkY*ChunkSize, chunkX*ChunkSize+ChunkSize, chunkY*ChunkSize+ChunkSize))
			// c.Tiles = TestChunk()

			if err := s.registry.Chunk.Add(c); err != nil {
				return fmt.Errorf("failed to add chunk: %w", err)
			}

			s.level.SetChunk(chunkX, chunkY, c.Tiles)
		}
	}

	return nil
}
