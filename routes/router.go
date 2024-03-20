package routes

type router struct{}

type Router interface{}

func NewRouter() Router {
	return &router{}
}
