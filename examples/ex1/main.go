package main

import (
	"fmt"
	"time"

	"github.com/lcaballero/codefmt"
)

func main() {
	var example uses
	example.braceTemplate()
	example.mapDirectly()
	example.usingBuf()
}

type uses int

func (u uses) usingBuf() {
	codefmt.NewBuf().
		Write(
			"I'll have %d hamburgers with %s, %s, and %s.",
			4, "pickle", "ketchup", "mustard").
		NL().Stdout()

	codefmt.NewBuf().NL().
		Writeln(
			"Your order number is: %d, should be ready at: %v",
			42, time.Now().Add(time.Minute*15).Format(time.Kitchen)).
		Stdout()

	codefmt.NewBuf().NL().
		Expand(`Order for ${num}, "${item}" is ready!!!`,
			"num", 42,
			"item", "deluxe hamburger",
		).
		NL().Stdout()
}

func (u uses) mapDirectly() {
	works := map[string]string{
		"Robert Frost": "The Road Not Taken",
		"Maya Angelou": "Stil I Rise",
		"Dylan Thomas": "Do Not Go Gentle into that That Good Night",
	}

	buf := codefmt.NewBuf()
	buf.Write("Author/Poems").NL()

	for author, poem := range works {
		buf.Expand(
			"Author: ${author}, Poem: ${poem}",
			"author", author, "poem", poem,
		).NL()
	}

	fmt.Println(buf)
}

func (u uses) braceTemplate() {
	hw1 := codefmt.BraceTemplate("${greeting}, ${name}!").Replace(
		"greeting", "Hello",
		"name", "World",
	)
	fmt.Println(hw1)
	fmt.Println()
}
