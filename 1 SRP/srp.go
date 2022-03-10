package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


// Single Responsibility Principle

type Journal struct {
	entries []string
}

func (j *Journal) entryCount() int {
	return len(j.entries)
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) {
	entry := fmt.Sprintf("%d: %s", j.entryCount()+1, text)
	j.entries = append(j.entries, entry)
}

func (j *Journal) RemoveEntry(index int) {
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
}

// breaks srp
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// ...
}

var lineSeparator = "\n"

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("Lorem ipsum dolor sit amet")
	j.AddEntry("Lorem ipsum dolor sit amet, consectetur adipiscing elit")
	fmt.Println(strings.Join(j.entries, "\n"))

	p := Persistence{lineSeparator}
	p.saveToFile(&j, "journal.txt")
}
