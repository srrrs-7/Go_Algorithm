package utils

import (
	"fmt"
	"strings"
)

type Category struct {
	ID       int
	Name     string
	ParentID int
	Children []*Category
}

func ParentChild() {
	categories := []*Category{
		{ID: 1, Name: "Electronics", ParentID: 0},
		{ID: 2, Name: "Computers", ParentID: 1},
		{ID: 3, Name: "Laptops", ParentID: 2},
		{ID: 4, Name: "Smartphones", ParentID: 1},
		{ID: 5, Name: "Tablets", ParentID: 1},
		{ID: 6, Name: "iPad", ParentID: 5},
	}

	categoryTree := BuildCategoryTree(categories)
	PrintCategoryTree(categoryTree, 0)
}

func BuildCategoryTree(categories []*Category) []*Category {
	categoryMap := make(map[int]*Category)

	// First pass: create all categories without children
	for _, category := range categories {
		categoryMap[category.ID] = &Category{
			ID:       category.ID,
			Name:     category.Name,
			ParentID: category.ParentID,
		}
	}

	// Second pass: add child categories to their parent's Children slice
	var rootCategories []*Category
	for _, category := range categoryMap {
		parentID := category.ParentID
		if parentID == 0 {
			rootCategories = append(rootCategories, category)
		} else {
			parentCategory := categoryMap[parentID]
			parentCategory.Children = append(parentCategory.Children, category)
		}
	}

	return rootCategories
}

func PrintCategoryTree(categories []*Category, indentLevel int) {
	for _, category := range categories {
		fmt.Printf("%s- %s\n", strings.Repeat("  ", indentLevel), category.Name)
		PrintCategoryTree(category.Children, indentLevel+1)
	}
}
