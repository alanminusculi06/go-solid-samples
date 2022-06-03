package main

import "fmt"

// Open-Closed Principle
// Objetos ou entidades devem estar abertos para extensão, mas fechados para modificação

type Color int

const (
	green Color = iota
	blue
)

type Size int

const (
	small Size = iota
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}

func (f *Filter) filterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) filterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) filterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

//filterBySize, filterBySizeAndColor

// better approach
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	bf := BetterFilter{}

	fmt.Print("Green products:\n")
	greenSpec := ColorSpecification{green}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Print("Green small products:\n")
	greenSmallSpec := AndSpecification{
		first:  ColorSpecification{green},
		second: SizeSpecification{small},
	}
	for _, v := range bf.Filter(products, greenSmallSpec) {
		fmt.Printf(" - %s is green and small\n", v.name)
	}
}
