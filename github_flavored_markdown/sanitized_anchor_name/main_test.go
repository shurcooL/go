package sanitized_anchor_name_test

import (
	"fmt"

	"github.com/shurcooL/go/github_flavored_markdown/sanitized_anchor_name"
)

func ExampleCreate() {
	anchorName := sanitized_anchor_name.Create("This is a header")

	fmt.Println(anchorName)

	// Output:
	//this-is-a-header
}

func ExampleCollisions() {
	var anchorNames string

	context := sanitized_anchor_name.NewContext()
	anchorNames += context.Create("Header") + "\n"
	anchorNames += context.Create("Header") + "\n"
	anchorNames += context.Create("Header") + "\n"

	fmt.Println(anchorNames)

	// Output:
	//header
	//header-2
	//header-3
	//
}

func ExampleCollisions2() {
	var anchorNames string

	context := sanitized_anchor_name.NewContext()
	anchorNames += context.Create("Header") + "\n"
	anchorNames += context.Create("Header 2") + "\n"
	anchorNames += context.Create("Header") + "\n"

	fmt.Println(anchorNames)

	// Output:
	//header
	//header-2
	//header-3
	//
}

func ExampleCollisions3() {
	var anchorNames string

	context := sanitized_anchor_name.NewContext()
	anchorNames += context.Create("Header") + "\n"
	anchorNames += context.Create("Header") + "\n"
	anchorNames += context.Create("Header 2") + "\n"

	fmt.Println(anchorNames)

	// Output:
	//header
	//header-2
	//header-2-2
	//
}
