package main

// цель - добиться уважения >100 + другие параметры !=0
import (
	"bufio"
	"fmt"
	"os"

	"github.com/RyabovNick/databasecourse_p2/tree/master/golang/tasks/console_game/all_functions"
)

func main() {

	fmt.Println("Do you know a cheat code?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	hero := all_functions.New(10, 100, 20, 30.0)
	if input == "wwssadadkj" {
		hero.Respect += 999
	} else {
		fmt.Println("Увы, но ты не знаешь чит код")
	}

	fmt.Print("Твои параметры:\n Длина норы = ", hero.NornLength, "\n Хп = ", hero.Hp, "\n Уважение = ", hero.Respect, "\n Вес = ", hero.Weight, "\n")

	// easy := 30.0
	// medium := 50.0
	// hard := 70.0

	for {

		fmt.Println("\nТвои действия\n 1 - Копать нору\n 2 - Поесть травки\n 3 - Пойти махаться\n 4 - Пойти баеньки\n ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()
		hero.DoDay(choice)
		hero.GoodNight()
		fmt.Print("\nТвои параметры:\n Длина норы = ", hero.NornLength, "\n Хп = ", hero.Hp, "\n Уважение = ", hero.Respect, "\n Вес = ", hero.Weight, "\n")
		if hero.NornLength <= 0 || hero.Hp <= 0 || hero.Weight <= 0 || hero.Respect <= 0 {
			fmt.Println("\n GAME OVER")
			break
		} else if hero.Respect > 100 {
			fmt.Println("\n POBEDA VMESTO OBEDA")
			break
		} else if hero.Weight >= 150 {
			fmt.Println("\n Ты умер от ожирения(без негатива)")
			break
		}

	}

}
