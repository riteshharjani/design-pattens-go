package main

import "fmt"

var entryCount = 0

type Journal struct {
	entries []string
}

// pointer is necessary here otherwise the struct won't get modified
func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d : %s\n", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

func (j Journal) PrintJournal() {
	for _, entry := range j.entries {
		fmt.Println(entry)
	}
}

func main() {
	j := Journal{}
	j.AddEntry("hello World")
	j.AddEntry("I wrote go SRP today")
	j.PrintJournal()

	//output:=
	// 1 : hello World
	// 2 : I wrote go SRP today

}
