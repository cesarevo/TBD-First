package main

// цель - добиться уважения >100 + другие параметры !=0
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type attrs struct {
	NornLength int64
	Hp         int64
	Respect    int64
	Weight     float64
}

func Roll(hero *attrs, EnemyWeight float64) float64 {
	propability := hero.Weight / (EnemyWeight + hero.Weight)
	return propability
}

func Defeat(hero *attrs, EnemyWeight float64) {
	fmt.Println("Боже, ты просто нулёвый чел")
	fmt.Println("К твоему сожалению, ты проиграл")
	if hero.Weight < EnemyWeight {
		hero.Hp -= 50
	} else if hero.Weight == EnemyWeight {
		hero.Hp -= 35
	} else if hero.Weight > EnemyWeight {
		hero.Hp -= 20
	}

}
func Pobeda(hero *attrs, EnemyWeight float64) {
	fmt.Println("Вау, какой ты крутой и классный")
	fmt.Println("Ты победил, вп!")

	if hero.Weight < EnemyWeight {
		hero.Respect += 40
	} else if hero.Weight == EnemyWeight {
		hero.Respect += 20
	} else if hero.Weight > EnemyWeight {
		hero.Respect += 5
	}

}

func GoodNight(hero *attrs) {
	hero.NornLength -= 2
	hero.Hp += 20
	hero.Respect -= 2
	hero.Weight -= 5

}

func main() {

	fmt.Println("Do you know a cheat code?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	hero := &attrs{10, 100, 20, 30.0}
	if input == "wwssadadkj" {
		hero.Respect += 999
	} else {
		fmt.Println("Увы, но ты не знаешь чит код")
	}

	fmt.Print("Твои параметры:\n Длина норы = ", hero.NornLength, "\n Хп = ", hero.Hp, "\n Уважение = ", hero.Respect, "\n Вес = ", hero.Weight, "\n")

	easy := 30.0
	medium := 50.0
	hard := 70.0

	for {
		fmt.Println("\nТвои действия\n 1 - Копать нору\n 2 - Поесть травки\n 3 - Пойти махаться\n 4 - Пойти баеньки\n ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			fmt.Println("Your choice is: Копать нору")
			fmt.Println("А как ты хочешь копать нору?")
			fmt.Println("Твои действия:\n 1 - Интенсивно\n 2 - Лениво")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			choice := scanner.Text()
			switch choice {
			case "1":
				hero.NornLength += 5
				hero.Hp -= 30
			case "2":
				hero.NornLength += 2
				hero.Hp -= 10
			default:
				fmt.Println("Не тупи, в наказание ты идёшь в кровать")
			}

		case "2":
			fmt.Println("Your choice is: Поесть травки")
			fmt.Println("А какую травку ты хочешь испробовать?")
			fmt.Println("Твой выбор:\n 1 - Жухлой\n 2 - Зелёной")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			choice := scanner.Text()
			switch choice {
			case "1":
				hero.Weight += 15
				hero.Hp += 10
			case "2":
				if hero.Respect < 30 {
					hero.Hp -= 30
				} else if hero.Respect >= 30 {
					hero.Hp += 30
					hero.Weight += 30
				} else {
					fmt.Println("Не тупи, в наказание ты идёшь в кровать")
				}
			default:
				fmt.Println("Не тупи, в наказание ты идёшь в кровать")
			}
		case "3":
			result := rand.Float64()
			fmt.Println("Your choice is: Пойти махаться")
			fmt.Println("С кем пойдёшь драться?")
			fmt.Println("Соперник:\n 1 - Лёгенький чел\n 2 - Среднячок\n 3 - На массе")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			choice := scanner.Text()
			if choice == "1" {
				fmt.Println("Ты выбрал лёгкую жертву")
				propability := Roll(hero, easy)
				if result <= propability {
					Defeat(hero, easy)
				} else if result > propability {
					Pobeda(hero, easy)
				}
			} else if choice == "2" {
				fmt.Println("Ты выбрал среднячковую жертву")
				propability := Roll(hero, easy)
				if result <= propability*100 {
					Defeat(hero, medium)
				} else if result > propability*100 {
					Pobeda(hero, medium)
				}
			} else if choice == "3" {
				fmt.Println("Ты выбрал массивного паренька")
				propability := Roll(hero, easy)
				if result <= propability*100 {
					Defeat(hero, hard)
				} else if result > propability*100 {
					Pobeda(hero, hard)
				}
			} else {
				fmt.Println("Не тупи, в наказание ты идёшь в кровать")
			}
		case "4":
			fmt.Println("Ты выбрал поспать, я тебя поздравляю, ты наверняка выспишься")
			GoodNight(hero)
		default:
			fmt.Println("Не тупи, в наказание ты идёшь в кровать")
		}

		GoodNight(hero)

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
