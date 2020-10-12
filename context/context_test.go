package context_test

import (
	"fmt"
	"git.agiletech.de/AgileRCM/support-tools/context"
)

func ExampleNewContext() {
	ctx := context.NewContext()
	ctx.UserId = "Testuser"
	fmt.Println(ctx.UserId)
	// Output: Testuser
}
