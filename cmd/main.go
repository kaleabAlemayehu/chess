package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/kaleabAlemayehu/chess/assets"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "github.com/hajimehoshi/ebiten/v2/vector"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Sprites struct {
	name      string
	rank      int
	file      rune
	engineImg *ebiten.Image
}

func (g *Chess) loadBoard() [8][8]*Sprites {

	var board [8][8]*Sprites

	for rank := 8; rank >= 1; rank-- {
		for file := 97; file <= 104; file++ {
			var sprite *Sprites
			switch rune(file) {
			case 'a', 'h':
				switch rank {
				case 8:
					sprite = &Sprites{
						name:      "black_rook",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.BlackRookData),
					}
				case 7:
					sprite = &Sprites{
						name:      "black_pawn",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.BlackPawnData),
					}
				case 2:
					sprite = &Sprites{
						name:      "white_pawn",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.WhitePawnData),
					}
				case 1:
					sprite = &Sprites{
						name:      "white_rook",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.WhiteRookData),
					}
				}
			case 'b', 'g':
				switch rank {
				case 8:
					sprite = &Sprites{
						name:      "black_knight",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.BlackKnightData),
					}
				case 7:
					sprite = &Sprites{
						name:      "black_pawn",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.BlackPawnData),
					}
				case 2:
					sprite = &Sprites{
						name:      "white_pawn",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.WhitePawnData),
					}
				case 1:
					sprite = &Sprites{
						name:      "white_knight",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.WhiteKnightData),
					}
				}
			case 'c', 'f':
				switch rank {
				case 8:
					sprite = &Sprites{
						name:      "black_bishop",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.BlackBishopData),
					}
				case 7:
					sprite = &Sprites{
						name:      "black_pawn",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.BlackPawnData),
					}
				case 2:
					sprite = &Sprites{
						name:      "white_pawn",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.WhitePawnData),
					}
				case 1:
					sprite = &Sprites{
						name:      "white_bishop",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.WhiteBishopData),
					}
				}
			case 'd':
				switch rank {
				case 8:
					sprite = &Sprites{
						name:      "black_queen",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.BlackQueenData),
					}
				case 7:
					sprite = &Sprites{
						name:      "black_pawn",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.BlackPawnData),
					}
				case 2:
					sprite = &Sprites{
						name:      "white_pawn",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.WhitePawnData),
					}
				case 1:
					sprite = &Sprites{
						name:      "white_queen",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.WhiteQueenData),
					}
				}
			case 'e':
				switch rank {
				case 8:
					sprite = &Sprites{
						name:      "black_king",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.BlackKingData),
					}
				case 7:
					sprite = &Sprites{
						name:      "black_pawn",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.BlackPawnData),
					}
				case 2:
					sprite = &Sprites{
						name:      "white_pawn",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.WhitePawnData),
					}
				case 1:
					sprite = &Sprites{
						name:      "white_king",
						rank:      rank,
						file:      rune(file),
						engineImg: parseImage(assets.WhiteKingData),
					}
				}
			}
			board[rank-1][file-97] = sprite
		}
	}
	return board
}

type Chess struct {
	Board        [8][8]*Sprites
	ScreenWidth  int
	ScreenHeight int
	BoardSize    int
	TileSize     int
}

func NewGame(screenW, screenH, boardSize int) *Chess {

	g := &Chess{
		ScreenWidth:  screenW,
		ScreenHeight: screenH,
		BoardSize:    boardSize,
		TileSize:     screenW / boardSize,
	}
	g.Board = g.loadBoard()
	return g
}

func (g *Chess) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()

		// Convert pixels to board coordinates (1-8)
		file := (mx / g.TileSize)
		rank := -(my / g.TileSize) + 9
		fileRune := rune(96 + file)

		log.Printf("Clicked on File: %c, file:%d, Rank: %d\n", fileRune, file, rank)
	}
	return nil
}

func parseImage(imgData []byte) *ebiten.Image {
	Image, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		log.Fatalf("unable to load white pawn image error : %w", err)
	}
	return ebiten.NewImageFromImage(Image)
}

func getPosition(file rune, rank int) (float64, float64) {
	fileInt := file - 97
	rankInt := 8 - rank
	return float64(fileInt), float64(rankInt)
}

func (g *Chess) Draw(screen *ebiten.Image) {
	boardDarkImg := parseImage(assets.BoardDarkData)
	boardLightImg := parseImage(assets.BoardLightData)

	scale := float64(g.TileSize) / float64(boardDarkImg.Bounds().Dx())

	for row := 0; row < g.BoardSize; row++ {
		for col := 0; col < g.BoardSize; col++ {
			if row == g.BoardSize-1 || row == 0 {
				// text.Draw(screen, "Hi", nil, nil)
				// ebitenutil.DebugPrint(screen, "hi")
				continue
			}
			if col == 0 || col == g.BoardSize-1 {
				// ebitenutil.DebugPrint(screen, "bye")
				continue
			}
			isDark := (row+col)%2 == 1
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(-float64(boardDarkImg.Bounds().Dx()/2), -float64(boardDarkImg.Bounds().Dy()/2))
			op.GeoM.Scale(scale, scale)
			targetX := float64(col*g.TileSize) + float64(g.TileSize)/2
			targetY := float64(row*g.TileSize) + float64(g.TileSize)/2
			op.GeoM.Translate(targetX, targetY)
			if isDark {
				screen.DrawImage(boardDarkImg, op)
			} else {
				screen.DrawImage(boardLightImg, op)
			}
		}
	}
	for _, spriteSlice := range g.Board {
		for _, s := range spriteSlice {
			if s != nil {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(-float64(s.engineImg.Bounds().Dx()/2), -float64(s.engineImg.Bounds().Dy()/2))
				op.GeoM.Scale(scale, scale)
				posX, posY := getPosition(s.file, s.rank)
				targetX := (posX+1)*float64(g.TileSize) + float64(g.TileSize)/2
				targetY := (posY+1)*float64(g.TileSize) + float64(g.TileSize)/2
				op.GeoM.Translate(targetX, targetY)
				screen.DrawImage(s.engineImg, op)
			}
		}
	}
}

func (g *Chess) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 1024
}

func main() {
	game := NewGame(1024, 1024, 10)
	ebiten.SetWindowSize(1024, 1024)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
