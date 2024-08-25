package model

const (
	OrderStatus_PreOrder   = "PreOrder"
	OrderStatus_Confirmed  = "Confirmed"
	OrderStatus_Canceled   = "Canceled"
	OrderStatus_InProgress = "InProgress"
	OrderStatus_Finished   = "Finished"
)

type Product struct {
	Title       string           `json:"title"`
	Price       float32          `json:"price"`
	Ingredients map[string]int32 `json:"ingredients"`
}

const Coffee = "COFFEE"
const Milk = "MILK"
const Water = "WATER"

const (
	ProductID_Espresso = iota
	ProductID_Americano
	ProductID_Cappuccino
	ProductID_Latte
	ProductID_TuCowsMilk
)

// Espresso is a shot of Coffee (100%)
var Espresso = Product{
	Title: "Espresso",
	Price: 2,
	Ingredients: map[string]int32{
		Coffee: 100,
	},
}

// Americano is a combination of Coffee(60%) and Water(40%)
var Americano = Product{
	Title: "Americano",
	Price: 2.5,
	Ingredients: map[string]int32{
		Coffee: 60,
		Water:  40,
	},
}

// Cappuccino is a combination of Coffee(50%), Milk(40%) and Water(10%)
var Cappuccino = Product{
	Title: "Cappuccino",
	Price: 4.5,
	Ingredients: map[string]int32{
		Coffee: 50,
		Milk:   40,
		Water:  10,
	},
}

// Latte is a combination of Coffee(20%) and Milk(80%)
var Latte = Product{
	Title: "Latte",
	Price: 3,
	Ingredients: map[string]int32{
		Coffee: 20,
		Milk:   80,
	},
}

// TuCowsMilk is the Best Hot milk in Town
var TuCowsMilk = Product{
	Title: "Hot Milk",
	Price: 5,
	Ingredients: map[string]int32{
		Milk: 100,
	},
}

var Menu = []Product{
	Espresso, Americano, Cappuccino, Latte, TuCowsMilk,
}
