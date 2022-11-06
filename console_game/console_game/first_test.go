package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pobeda(t *testing.T) {
	hero := &attrs{10, 10, 10, 30.0}
	Pobeda(hero, 50.0)
	want := int64(50)
	assert.Equal(t, hero.Respect, want)
}

func Test_defeat(t *testing.T) {
	hero := &attrs{10, 50, 10, 30.0}
	Defeat(hero, 50.0)
	want := int64(0)
	assert.Equal(t, hero.Hp, want)
}

func Test_night(t *testing.T) {
	hero := &attrs{10, 10, 10, 30.0}
	GoodNight(hero)
	want := &attrs{8, 30, 8, 25.0}
	assert.Equal(t, hero, want)
}

func Test_roll(t *testing.T) {
	hero := &attrs{10, 10, 10, 30.0}
	got := Roll(hero, 50.0)
	want := 3.0 / 8.0
	assert.Equal(t, got, want)

}
