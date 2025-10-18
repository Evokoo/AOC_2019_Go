package day14

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// RECIPE
// ========================
type Recipe struct {
	quantity    int
	ingredients map[string]int
	base        bool
}

func NewRecipe() *Recipe {
	return &Recipe{
		quantity:    0,
		ingredients: make(map[string]int),
	}
}

func (r *Recipe) SetQuantity(n int) {
	r.quantity = n
}
func (r *Recipe) GetQuantity() int {
	return r.quantity
}
func (r *Recipe) AddIngrident(name string, amount int) {
	r.ingredients[name] = amount
}
func (r *Recipe) IsBase() bool {
	return r.base
}

// ========================
// INVENTORY
// ========================
type Stock map[string]int

func (i *Stock) Prune(reactions Reactions) {
	for key := range *i {
		recipe := reactions.Get(key)

		if !recipe.IsBase() {
			delete(*i, key)
		}
	}
}

// ========================
// REACTIONS
// ========================

type Reactions map[string]Recipe

func (r *Reactions) Get(name string) Recipe {
	if recipe, found := (*r)[name]; found {
		return recipe
	}

	panic("Recipe Not found")
}

func (r Reactions) RequiredOre(id string, amount int, stock Stock, ore *int) {
	if id == "ORE" {
		*ore += amount
		return
	}

	// Use surplus if available
	if stock[id] >= amount {
		stock[id] -= amount
		return
	}

	// Remaining amount after using stock
	amount -= stock[id]
	stock[id] = 0

	recipe := r.Get(id)
	batches := (amount + recipe.quantity - 1) / recipe.quantity

	// Recurse for all ingredients
	for ingredient, quantity := range recipe.ingredients {
		r.RequiredOre(ingredient, quantity*batches, stock, ore)
	}

	// Store leftovers in stock
	stock[id] += recipe.quantity*batches - amount
}

// ========================
// PARSER
// ========================
func ParseInput(file string) Reactions {
	data := utils.ReadFile(file)

	reactions := make(Reactions)

	for line := range strings.SplitSeq(data, "\n") {
		recipe := NewRecipe()
		sections := strings.Split(line, "=> ")

		//Input Side
		for _, item := range strings.Split(sections[0], ", ") {
			ingrident := strings.Split(item, " ")
			quantity, _ := strconv.Atoi(ingrident[0])
			recipe.AddIngrident(ingrident[1], quantity)

			if ingrident[1] == "ORE" {
				recipe.base = true
			}
		}

		//Output Side
		product := strings.Split(sections[1], " ")
		quantity, _ := strconv.Atoi(product[0])
		recipe.SetQuantity(quantity)

		reactions[product[1]] = *recipe
	}

	return reactions
}
