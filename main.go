package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "github.com/hajimehoshi/ebiten/v2/vector"
)

//go:embed assets/pngs/square_brown_dark_png_shadow_128px.png
var boardDarkData []byte

//go:embed assets/pngs/square_brown_light_png_shadow_128px.png
var boardLightData []byte

//go:embed assets/pngs/b_rook_png_shadow_128px.png
var blackRookData []byte

//go:embed assets/pngs/w_rook_png_shadow_128px.png
var whiteRookData []byte

//go:embed assets/pngs/b_knight_png_shadow_128px.png
var blackKnightData []byte

//go:embed assets/pngs/w_knight_png_shadow_128px.png
var whiteKnightData []byte

//go:embed assets/pngs/b_bishop_png_shadow_128px.png
var blackBishopData []byte

//go:embed assets/pngs/w_bishop_png_shadow_128px.png
var whiteBishopData []byte

//go:embed assets/pngs/b_pawn_png_shadow_128px.png
var blackPawnData []byte

//go:embed assets/pngs/w_pawn_png_shadow_128px.png
var whitePawnData []byte

//go:embed assets/pngs/b_king_png_shadow_128px.png
var blackKingData []byte

//go:embed assets/pngs/w_king_png_shadow_128px.png
var whiteKingData []byte

//go:embed assets/pngs/b_queen_png_shadow_128px.png
var blackQueenData []byte

//go:embed assets/pngs/w_queen_png_shadow_128px.png
var whiteQueenData []byte

type Sprites struct {
	name      string
	rank      int
	file      rune
	engineImg *ebiten.Image
	rawImg    *image.Image
}

func (g *Chess) loadSpirtes() *[]Sprites {

	var sprites []Sprites

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

	for row := 1; row <= g.BoardSize; row++ {
		for col := 1; col <= g.BoardSize; col++ {
			x := float64(col * g.TileSize)
			y := float64(row * g.TileSize)
			log.Printf("x: %f, y: %f", x, y)
			isDark := (row+col)%2 == 1
			if isDark {
				sprites = append(sprites, Sprites{
					name:      "dark_tile",
					rank:      9 - row,
					file:      rune(96 + col),
					engineImg: boardDarkImg,
					rawImg:    &boardDarkImage,
				})
			} else {
				sprites = append(sprites, Sprites{
					name:      "light_tile",
					rank:      9 - row,
					file:      rune(96 + col),
					engineImg: boardLightImg,
					rawImg:    &boardLightImage,
				})
			}
		}
	}

	bRookImage, _, err := image.Decode(bytes.NewReader(blackRookData))
	if err != nil {
		log.Fatalf("unable to load black rook image error : %w", err)
	}
	bRookImg := ebiten.NewImageFromImage(bRookImage)
	sprites = append(sprites, Sprites{
		name:      "black_rook",
		rank:      8,
		file:      'a',
		engineImg: bRookImg,
		rawImg:    &bRookImage,
	})
	sprites = append(sprites, Sprites{
		name:      "black_rook",
		rank:      8,
		file:      'h',
		engineImg: bRookImg,
		rawImg:    &bRookImage,
	})

	wRookImage, _, err := image.Decode(bytes.NewReader(whiteRookData))
	if err != nil {
		log.Fatalf("unable to load black rook image error : %w", err)
	}
	wRookImg := ebiten.NewImageFromImage(wRookImage)
	sprites = append(sprites, Sprites{
		name:      "white_rook",
		rank:      1,
		file:      'a',
		engineImg: wRookImg,
		rawImg:    &wRookImage,
	})
	sprites = append(sprites, Sprites{
		name:      "white_rook",
		rank:      1,
		file:      'h',
		engineImg: wRookImg,
		rawImg:    &wRookImage,
	})

	wKnightImage, _, err := image.Decode(bytes.NewReader(whiteKnightData))
	if err != nil {
		log.Fatalf("unable to load black rook image error : %w", err)
	}
	wKnightImg := ebiten.NewImageFromImage(wKnightImage)
	sprites = append(sprites, Sprites{
		name:      "white_knight",
		rank:      1,
		file:      'b',
		engineImg: wKnightImg,
		rawImg:    &wKnightImage,
	})
	sprites = append(sprites, Sprites{
		name:      "white_knight",
		rank:      1,
		file:      'g',
		engineImg: wKnightImg,
		rawImg:    &wKnightImage,
	})

	bKnightImage, _, err := image.Decode(bytes.NewReader(blackKnightData))
	if err != nil {
		log.Fatalf("unable to load black rook image error : %w", err)
	}
	bKnightImg := ebiten.NewImageFromImage(bKnightImage)
	sprites = append(sprites, Sprites{
		name:      "black_knight",
		rank:      8,
		file:      'b',
		engineImg: bKnightImg,
		rawImg:    &bKnightImage,
	})
	sprites = append(sprites, Sprites{
		name:      "black_knight",
		rank:      8,
		file:      'g',
		engineImg: bKnightImg,
		rawImg:    &bKnightImage,
	})

	bBishopImage, _, err := image.Decode(bytes.NewReader(blackBishopData))
	if err != nil {
		log.Fatalf("unable to load black rook image error : %w", err)
	}
	bBishopImg := ebiten.NewImageFromImage(bBishopImage)
	sprites = append(sprites, Sprites{
		name:      "black_bishop",
		rank:      8,
		file:      'c',
		engineImg: bBishopImg,
		rawImg:    &bBishopImage,
	})
	sprites = append(sprites, Sprites{
		name:      "black_bishop",
		rank:      8,
		file:      'f',
		engineImg: bBishopImg,
		rawImg:    &bBishopImage,
	})

	wBishopImage, _, err := image.Decode(bytes.NewReader(whiteBishopData))
	if err != nil {
		log.Fatalf("unable to load black rook image error : %w", err)
	}
	wBishopImg := ebiten.NewImageFromImage(wBishopImage)
	sprites = append(sprites, Sprites{
		name:      "white_bishop",
		rank:      1,
		file:      'c',
		engineImg: wBishopImg,
		rawImg:    &wBishopImage,
	})
	sprites = append(sprites, Sprites{
		name:      "white_bishop",
		rank:      1,
		file:      'f',
		engineImg: wBishopImg,
		rawImg:    &wBishopImage,
	})

	wPawnImage, _, err := image.Decode(bytes.NewReader(whitePawnData))
	if err != nil {
		log.Fatalf("unable to load white pawn image error : %w", err)
	}
	wPawnImg := ebiten.NewImageFromImage(wPawnImage)

	bPawnImage, _, err := image.Decode(bytes.NewReader(blackPawnData))
	if err != nil {
		log.Fatalf("unable to load white pawn image error : %w", err)
	}
	bPawnImg := ebiten.NewImageFromImage(bPawnImage)
	for i := range 8 {
		sprites = append(sprites, Sprites{
			name:      "white_pawn",
			rank:      2,
			file:      rune(97 + i),
			engineImg: wPawnImg,
			rawImg:    &wPawnImage,
		})
		sprites = append(sprites, Sprites{
			name:      "black_pawn",
			rank:      7,
			file:      rune(97 + i),
			engineImg: bPawnImg,
			rawImg:    &bPawnImage,
		})
	}

	wKingImage, _, err := image.Decode(bytes.NewReader(whiteKingData))
	if err != nil {
		log.Fatalf("unable to load white pawn image error : %w", err)
	}
	wKingImg := ebiten.NewImageFromImage(wKingImage)
	sprites = append(sprites, Sprites{
		name:      "white_king",
		rank:      1,
		file:      'e',
		engineImg: wKingImg,
		rawImg:    &wKingImage,
	})

	bKingImage, _, err := image.Decode(bytes.NewReader(blackKingData))
	if err != nil {
		log.Fatalf("unable to load white pawn image error : %w", err)
	}
	bKingImg := ebiten.NewImageFromImage(bKingImage)
	sprites = append(sprites, Sprites{
		name:      "black_king",
		rank:      8,
		file:      'e',
		engineImg: bKingImg,
		rawImg:    &bKingImage,
	})

	bQueenImage, _, err := image.Decode(bytes.NewReader(blackQueenData))
	if err != nil {
		log.Fatalf("unable to load white pawn image error : %w", err)
	}
	bQueenImg := ebiten.NewImageFromImage(bQueenImage)
	sprites = append(sprites, Sprites{
		name:      "black_queen",
		rank:      8,
		file:      'd',
		engineImg: bQueenImg,
		rawImg:    &bQueenImage,
	})

	wQueenImage, _, err := image.Decode(bytes.NewReader(whiteQueenData))
	if err != nil {
		log.Fatalf("unable to load white pawn image error : %w", err)
	}
	wQueenImg := ebiten.NewImageFromImage(wQueenImage)
	sprites = append(sprites, Sprites{
		name:      "white_queen",
		rank:      1,
		file:      'd',
		engineImg: wQueenImg,
		rawImg:    &wQueenImage,
	})

	return &sprites
}

type Chess struct {
	GameSpirtes  []Sprites
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
	g.GameSpirtes = *g.loadSpirtes()
	return g
}

func (g *Chess) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()

		// Convert pixels to board coordinates (1-8)
		file := (mx / g.TileSize)
		rank := -(my / g.TileSize) + 8
		fileRune := rune(97 + file)

		log.Printf("Clicked on File: %c, file:%d, Rank: %d\n", fileRune, file, rank)
	}
	return nil
}

func getPosition(file rune, rank int) (float64, float64) {
	fileInt := file - 97
	rankInt := 8 - rank
	return float64(fileInt), float64(rankInt)
}

func (g *Chess) Draw(screen *ebiten.Image) {
	scale := float64(g.TileSize) / float64(g.GameSpirtes[0].engineImg.Bounds().Dx())
	for _, s := range g.GameSpirtes {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		posX, posY := getPosition(s.file, s.rank)
		op.GeoM.Translate(float64(g.TileSize)*posX, float64(g.TileSize)*posY)
		screen.DrawImage(s.engineImg, op)
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
