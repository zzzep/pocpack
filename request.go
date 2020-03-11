package pocpack

type Response struct {
	success bool
	message string
}

func Get() Response {
	r := Response{true, "Funfa"}
	return r
}

func Post() Response {
	r := Response{false, "Não Funfa"}
	return r
}
