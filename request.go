package pocpack

type response struct {
	success bool
	message string
}

func get() response {
	r := response{true, "Funfa"}
	return r
}

func post() response {
	r := response{false, "Não Funfa"}
	return r
}
