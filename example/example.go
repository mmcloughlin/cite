package example

import "fmt"

// Greet says hello to who.
//
// Reference: https://github.com/mmcloughlin/cite/blob/master/example/grinch.txt#L6-L8
//
//	Every Who Down in Whoville Liked Christmas a lot...
//	But the Grinch,Who lived just north of Whoville, Did NOT!
//	The Grinch hated Christmas! The whole Christmas season!
//
func Greet(who string) {
	fmt.Printf("Hello, %s!\n", who)
}
