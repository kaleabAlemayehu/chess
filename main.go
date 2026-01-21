package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)
//go:embed assets/PNGs/with_shadow/128px/square_brown_dark_png_shadow_128px.png
var boardDarkData []byte

//go:embed assets/PNGs/with_shadow/128px/square_brown_light_png_shadow_128px.png
var boardLightData []byte

type Chess struct {
	boardDarkTile  *ebiten.Image
	boardLightTile *ebiten.Image
	screenWidth    int
	screenHeight   int
	boardSize      int
	tileSize       int
}

func NewGame(screenW, screenH, boardSize int) *Chess {

	boardDarkImage, _, err := image.Decode(bytes.NewReader(boardDarkData))
	if err != nil {
		log.Fatalf("unable to load board dark tile image error : %w", err)
	}
	boardDarkImg := ebiten.NewImageFromImage(boardDarkImage)

	boardLightImage, _, err := image.Decode(bytes.NewReader(boardLightData))
	if err != nil {
		log.Fatalf("unable to load board light tile image error : %w", err)
	}
	boardLightImg := ebiten.NewImageFromImage(boardLightImage)

	return &Chess{
		screenWidth:    screenW,
		screenHeight:   screenH,
		boardSize:      boardSize,
		tileSize:       screenW / boardSize,
		boardDarkTile:  boardDarkImg,
		boardLightTile: boardLightImg,
	}
}

func (g *Chess) Update() error {
	return nil
}
func (g *Chess) Draw(screen *ebiten.Image) {
	scale := float64(g.tileSize) / float64(g.boardDarkTile.Bounds().Dx())
	for row := 0; row < g.boardSize; row++ {
		for col := 0; col < g.boardSize; col++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(scale, scale)
			x := float64(col * g.tileSize)
			y := float64(row * g.tileSize)
			op.GeoM.Translate(x, y)

			isDark := (row+col)%2 == 1

			if isDark {
				screen.DrawImage(g.boardDarkTile, op)
			} else {
				screen.DrawImage(g.boardLightTile, op)
			}
		}
	}
}
func (g *Chess) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 1024
}

func main() {
	game := NewGame(1024, 1024, 8)
	ebiten.SetWindowSize(1024, 1024)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
