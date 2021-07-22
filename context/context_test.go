package context_test

import (
	"fmt"
	"github.com/agile-rcm/support-tools/context"
)

func ExampleNewContext() {
	ctx := context.NewContext()
	ctx.UserId = "Testuser"
	fmt.Println(ctx.UserId)
	// Output: Testuser
}
