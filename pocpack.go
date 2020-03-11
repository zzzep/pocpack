package pocpack

type response struct {
	success bool
	message string
}

func main() response {
	r := response{true, "Funfa"}
	return r
}
