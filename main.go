package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yagoernandes/testes-ebitengine/entities"
)

type Game struct {
	player      *entities.Player
	enemies     []*entities.Enemy
	potions     []*entities.Potion
	tilemapJSON *TilemapJSON
	tilemapImg  *ebiten.Image
}

func (g *Game) Update() error {

	// react to key presses
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.Y -= 2
	}

	for _, enemy := range g.enemies {
		if enemy.X > g.player.X {
			enemy.X -= 1
		} else if enemy.X < g.player.X {
			enemy.X += 1
		}
		if enemy.Y > g.player.Y {
			enemy.Y -= 1
		} else if enemy.Y < g.player.Y {
			enemy.Y += 1
		}
	}

	for _, potion := range g.potions {
		if potion.X > g.player.X {
			g.player.Health += potion.AmountHeal
			fmt.Printf("Picked up potion. Player healed by %d, new health: %d\n", potion.AmountHeal, g.player.Health)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}

	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	// loop over the layers of the tilemap
	for _, layer := range g.tilemapJSON.Layers {
		// loop over the tiles in the layer
		for index, id := range layer.Data {

			// calculate the position of the tile
			x := index % int(layer.Width)
			y := index / int(layer.Width)

			// multiply the position by the tile size
			x *= 16
			y *= 16

			// calculate the source position of the tile in the tileset
			srcX := (id - 1) % 28
			srcY := (id - 1) / 28

			// multiply the source position by the tile size
			srcX *= 16
			srcY *= 16

			// draw the tile
			opts.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(
				// croping out the tile from the tileset
				g.tilemapImg.SubImage(
					image.Rect(
						int(srcX),
						int(srcY),
						int(srcX+16),
						int(srcY+16),
					),
				).(*ebiten.Image),
				&opts,
			)
			opts.GeoM.Reset()
		}
	}

	// ebitenutil.DebugPrint(screen, "Hello, World!")
	opts.GeoM.Reset()
	opts.GeoM.Translate(g.player.X, g.player.Y)

	screen.DrawImage(
		g.player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)

	for _, enemy := range g.enemies {
		opts.GeoM.Reset()
		opts.GeoM.Translate(enemy.X, enemy.Y)

		screen.DrawImage(enemy.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image), &opts)
	}

	for _, potion := range g.potions {
		opts.GeoM.Reset()
		opts.GeoM.Translate(potion.X, potion.Y)

		screen.DrawImage(potion.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image), &opts)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	ninjaImage, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")
	if err != nil {
		log.Fatal(err)
	}

	skeletonImage, _, err := ebitenutil.NewImageFromFile("assets/images/skeleton.png")
	if err != nil {
		log.Fatal(err)
	}

	potionImage, _, err := ebitenutil.NewImageFromFile("assets/images/potion.png")
	if err != nil {
		log.Fatal(err)
	}

	tilemapImage, _, err := ebitenutil.NewImageFromFile("assets/maps/tileset.png")
	if err != nil {
		log.Fatal(err)
	}

	tilemapJSON, err := NewTilemapJSON("assets/maps/map.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{
		player: &entities.Player{
			Sprite: &entities.Sprite{
				Img: ninjaImage,
				X:   100,
				Y:   100,
			},
			Health: 100,
		},
		enemies: []*entities.Enemy{
			{
				Sprite: &entities.Sprite{
					Img: skeletonImage,
					X:   200,
					Y:   200,
				},
				FollowsPlayer: true,
			},
			{
				Sprite: &entities.Sprite{
					Img: skeletonImage,
					X:   150,
					Y:   100,
				},
				FollowsPlayer: false,
			},
			{
				Sprite: &entities.Sprite{
					Img: skeletonImage,
					X:   50,
					Y:   200,
				},
				FollowsPlayer: true,
			},
		},
		potions: []*entities.Potion{
			{
				Sprite: &entities.Sprite{
					Img: potionImage,
					X:   100,
					Y:   75,
				},
				AmountHeal: 10,
			},
			{
				Sprite: &entities.Sprite{
					Img: potionImage,
					X:   50,
					Y:   100,
				},
				AmountHeal: 20,
			},
		},
		tilemapJSON: tilemapJSON,
		tilemapImg:  tilemapImage,
	}); err != nil {
		log.Fatal(err)
	}
}
