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

func (p Product) filterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.color == color {
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

// color specification
type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

// color specification

// size specification
type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

// size specification

// and specification
type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

// and specification

type Filter struct{}

func (f *Filter) Filter(products []Product, spec Specification) []*Product {
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

	filter := Filter{}

	fmt.Print("Green products:\n")
	greenSpec := ColorSpecification{green}
	for _, v := range filter.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Print("Green small products:\n")
	greenSmallSpec := AndSpecification{first: ColorSpecification{green}, second: SizeSpecification{small}}
	for _, v := range filter.Filter(products, greenSmallSpec) {
		fmt.Printf(" - %s is green and small\n", v.name)
	}
}
