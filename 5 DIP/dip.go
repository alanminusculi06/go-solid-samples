package main

import "fmt"

// Dependency Inversion Principle
// Módulos de alto nível não devem depender de módulos de baixo nível
// Ambos devem depender de abastrações

type Relationship int

const (
	Parent Relationship = iota
	Child
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type Relationships struct {
	relations []Info
}

// low-level
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range rs.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}
	return result
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations, Info{parent, Parent, child})
	rs.relations = append(rs.relations, Info{child, Child, parent})
}

type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("Pedro") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := Person{"Pedro"}
	child1 := Person{"Maria"}
	child2 := Person{"João"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := Research{&relationships}
	research.Investigate()
}
