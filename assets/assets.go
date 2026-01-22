package assets

import (
	_ "embed"
)

//go:embed pngs/square_brown_dark_png_shadow_128px.png
var BoardDarkData []byte

//go:embed pngs/square_brown_light_png_shadow_128px.png
var BoardLightData []byte

//go:embed pngs/b_rook_png_shadow_128px.png
var BlackRookData []byte

//go:embed pngs/w_rook_png_shadow_128px.png
var WhiteRookData []byte

//go:embed pngs/b_knight_png_shadow_128px.png
var BlackKnightData []byte

//go:embed pngs/w_knight_png_shadow_128px.png
var WhiteKnightData []byte

//go:embed pngs/b_bishop_png_shadow_128px.png
var BlackBishopData []byte

//go:embed pngs/w_bishop_png_shadow_128px.png
var WhiteBishopData []byte

//go:embed pngs/b_pawn_png_shadow_128px.png
var BlackPawnData []byte

//go:embed pngs/w_pawn_png_shadow_128px.png
var WhitePawnData []byte

//go:embed pngs/b_king_png_shadow_128px.png
var BlackKingData []byte

//go:embed pngs/w_king_png_shadow_128px.png
var WhiteKingData []byte

//go:embed pngs/b_queen_png_shadow_128px.png
var BlackQueenData []byte

//go:embed pngs/w_queen_png_shadow_128px.png
var WhiteQueenData []byte
