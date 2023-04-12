package main

import "fmt"

type CoffeeMachine struct {
	availableWater  int
	availableMilk   int
	availableCoffee int
	availableCups   int
	availableMoney  int
}

type Drink struct {
	water  int
	milk   int
	coffee int
	price  int
}

// -> 250 ml of water and 16 g of coffee beans. It costs $4.
var espresso = Drink{250, 0, 16, 4}

// -> 350 ml of water, 75 ml of milk, and 20 g of coffee beans. It costs $7
var latte = Drink{350, 75, 20, 7}

// -> 200 ml of water, 100 ml of milk, and 12 g of coffee beans. It costs $6.
var cappuccino = Drink{200, 100, 12, 6}

func (coffeeMachine *CoffeeMachine) displayCurrentState() {
	fmt.Printf(`The coffee machine has:
%d ml of water
%d ml of milk
%d g of coffee beans
%d disposable cups
$%d of money

`,
		coffeeMachine.availableWater,
		coffeeMachine.availableMilk,
		coffeeMachine.availableCoffee,
		coffeeMachine.availableCups,
		coffeeMachine.availableMoney,
	)
}

func (coffeeMachine *CoffeeMachine) run() {
	for {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		var choice string
		_, _ = fmt.Scan(&choice)

		switch choice {
		case "buy":
			coffeeMachine.choiceBuy()
		case "fill":
			coffeeMachine.choiceFill()
		case "take":
			coffeeMachine.choiceTake()
		case "remaining":
			coffeeMachine.displayCurrentState()
		case "exit":
			return
		}
	}
}

func (coffeeMachine *CoffeeMachine) choiceTake() {
	fmt.Printf("I gave you $%d\n\n", coffeeMachine.availableMoney)
	coffeeMachine.availableMoney = 0
}

func (coffeeMachine *CoffeeMachine) choiceFill() {
	var water, milk, coffee, cups int

	fmt.Println("Write how many ml of water you want to add:")
	_, _ = fmt.Scan(&water)

	fmt.Println("Write how many ml of milk you want to add:")
	_, _ = fmt.Scan(&milk)

	fmt.Println("Write how many grams of coffee beans you want to add:")
	_, _ = fmt.Scan(&coffee)

	fmt.Println("Write how many disposable cups you want to add:")
	_, _ = fmt.Scan(&cups)

	coffeeMachine.availableWater += water
	coffeeMachine.availableMilk += milk
	coffeeMachine.availableCoffee += coffee
	coffeeMachine.availableCups += cups
}

func (coffeeMachine *CoffeeMachine) makeDrink(drink Drink) {
	coffeeMachine.availableWater -= drink.water
	coffeeMachine.availableMilk -= drink.milk
	coffeeMachine.availableCoffee -= drink.coffee
	coffeeMachine.availableMoney += drink.price
	coffeeMachine.availableCups -= 1
}

func (coffeeMachine *CoffeeMachine) checkMissing(drink Drink) string {
	missing := ""
	if coffeeMachine.availableWater < drink.water {
		missing = "water"
	} else if coffeeMachine.availableMilk < drink.milk {
		missing = "milk"
	} else if coffeeMachine.availableCoffee < drink.coffee {
		missing = "coffee"
	} else if coffeeMachine.availableCups < 1 {
		missing = "cups"
	}
	return missing
}

func (coffeeMachine *CoffeeMachine) canPrepareDrink(drink Drink) bool {
	missing := coffeeMachine.checkMissing(drink)

	if missing != "" {
		fmt.Printf("Sorry, not enough %s!\n", missing)
		return false
	}

	fmt.Println("I have enough resources, making you a coffee!")
	return true
}

func (coffeeMachine *CoffeeMachine) choiceBuy() {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")

	var drinkChoice string
	_, _ = fmt.Scan(&drinkChoice)
	if drinkChoice == "back" {
		return
	}

	var drink Drink
	switch drinkChoice {
	case "1":
		drink = espresso
	case "2":
		drink = latte
	case "3":
		drink = cappuccino
	}
	if coffeeMachine.canPrepareDrink(drink) {
		coffeeMachine.makeDrink(drink)
	}
}

func main() {
	coffeeMachine := CoffeeMachine{400, 540, 120, 9, 550}
	coffeeMachine.run()
}
