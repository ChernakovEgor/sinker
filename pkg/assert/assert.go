package assert

func Assert(statement bool, msg string) {
	if !statement {
		panic(msg)
	}
}
