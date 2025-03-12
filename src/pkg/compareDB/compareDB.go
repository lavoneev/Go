package DBcomparer

import (
	DBreader "Day-01/pkg/readDB"
	"fmt"
	"os"
)

func LoadRecipes(filename string) DBreader.Recipes {
	reader, format := DBreader.GetReader(filename)
	if reader == nil {
		fmt.Println("invalid input:", format)
		os.Exit(1)
	}
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return reader.ReadData(file)
}

func CompareCakes(oldRecipes, newRecipes DBreader.Recipes) {
	oldCakes := make(map[string]DBreader.Cake)
	newCakes := make(map[string]DBreader.Cake)

	for _, cake := range oldRecipes.Cakes {
		oldCakes[cake.Name] = cake
	}

	for _, cake := range newRecipes.Cakes {
		newCakes[cake.Name] = cake
	}

	for name, newCake := range newCakes {
		oldCake, exists := oldCakes[name]
		if !exists {
			fmt.Printf("ADDED cake \"%s\"\n", name)
			continue
		}

		if newCake.CookTime != oldCake.CookTime {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, newCake.CookTime, oldCake.CookTime)
		}

		CompareIngredients(name, oldCake, newCake)
	}

	for name := range oldCakes {
		_, exists := newCakes[name]
		if !exists {
			fmt.Printf("REMOVED cake \"%s\"\n", name)
		}
	}
}

func CompareIngredients(cakeName string, oldCake, newCake DBreader.Cake) {
	oldIngrs := make(map[string]DBreader.Ingredient)
	newIngrs := make(map[string]DBreader.Ingredient)

	for _, ingr := range oldCake.Ingredients {
		oldIngrs[ingr.Name] = ingr
	}

	for _, ingr := range newCake.Ingredients {
		newIngrs[ingr.Name] = ingr
	}

	for name, newIngr := range newIngrs {
		oldIngr, exists := oldIngrs[name]
		if !exists {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", name, cakeName)
			continue
		}

		if oldIngr.Count != newIngr.Count {
			fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, cakeName, newIngr.Count, oldIngr.Count)
		}

		if oldIngr.Unit != newIngr.Unit {
			if oldIngr.Unit == "" {
				fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\", for cake \"%s\"\n", newIngr.Unit, name, cakeName)
			} else if newIngr.Unit == "" {
				fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", oldIngr.Unit, name, cakeName)
			} else {
				fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, cakeName, oldIngr.Unit, newIngr.Unit)
			}
		}
	}

	for name := range oldIngrs {
		_, exists := newIngrs[name]
		if !exists {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", name, cakeName)
		}
	}
}

func Run(oldFilename, newFilename string) {
	oldRecipes := LoadRecipes(oldFilename)
	newRecipes := LoadRecipes(newFilename)

	CompareCakes(oldRecipes, newRecipes)
}
