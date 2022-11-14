package main

import (
	"testing"

	"github.com/RyabovNick/databasecourse_p2/tree/master/golang/tasks/console_game/all_functions"
	"github.com/stretchr/testify/assert"
)

func Test_night_and_sleep(t *testing.T) {
	hero := all_functions.New(10, 10, 10, 30.0)
	hero.DoDay("4")
	hero.Sleep()
	got := hero.NornLength
	want := int64(6)
	assert.Equal(t, got, want)

}
func Test_pobeda(t *testing.T) {
	hero := all_functions.New(10, 10, 10, 30.0)
	all_functions.Pobeda(hero, 50.0)
	want := int64(50)
	assert.Equal(t, hero.Respect, want)
}

func Test_defeat(t *testing.T) {
	hero := all_functions.New(10, 10, 10, 30.0)
	all_functions.Defeat(hero, 50.0)
	want := int64(-40)
	assert.Equal(t, hero.Hp, want)
}

func Test_night(t *testing.T) {
	hero := all_functions.New(10, 10, 10, 30.0)
	hero.GoodNight()
	want := all_functions.New(8, 30, 8, 25.0)
	assert.Equal(t, hero, want)
}

func Test_roll(t *testing.T) {
	hero := all_functions.New(10, 10, 10, 30.0)
	got := all_functions.Roll(hero, 50.0)
	want := 3.0 / 8.0
	assert.Equal(t, got, want)

}
