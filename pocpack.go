package pocpack

type response struct {
	success bool
	message string
}

func get() response {
	r := response{true, "Funfa"}
	return r
}
