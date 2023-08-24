package main

import (
	"io/ioutil"
	"strings"
)

// Single Responsibility Principle
// Uma classe deve ter um, e somente um, motivo para mudar.

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
	j.entries = append(j.entries, text)
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

// better approach
type JournalRepository struct{}

func (p *JournalRepository) Save(j *Journal, filename string) {
	lineSeparator := "\n"
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, lineSeparator)), 0644)
}

func (p *JournalRepository) Load(filename string) {
	// ...
}

func main() {
	j := Journal{}
	j.AddEntry("Lorem ipsum dolor sit amet")
	j.AddEntry("Consectetur adipiscing elit")

	p := JournalRepository{}
	p.Save(&j, "journal.txt")
}
