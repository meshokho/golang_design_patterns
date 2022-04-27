package main

import "fmt"

// Функциональный способ. Возможно расширение Builder'а.

type Game struct {
	title, genre string
}

type gameMod func(*Game)
type GameBuilder struct {
	actions []gameMod
}

func (b *GameBuilder) SetTitle(title string) *GameBuilder {
	b.actions = append(b.actions, func(g *Game) {
		g.title = title
	})

	return b
}

// Extent builder

func (b *GameBuilder) SetGenre(genre string) *GameBuilder {
	b.actions = append(b.actions, func(game *Game) {
		game.genre = genre
	})

	return b
}

func (b *GameBuilder) Build() *Game {
	g := Game{}
	for _, a := range b.actions {
		a(&g)
	}

	return &g
}

func (g *Game) Describe() {
	fmt.Println("Game genre: " + g.genre + ". Game title: " + g.title)
}

func main() {
	b := GameBuilder{}
	game := b.SetTitle("Skyrim").SetGenre("RPG").Build()
	anotherGame := b.SetGenre("Racing").SetTitle("NFS").Build()

	game.Describe()
	anotherGame.Describe()
}
