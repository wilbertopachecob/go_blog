package channels

func OK(ch <-chan bool) bool {
	return <-ch
}
