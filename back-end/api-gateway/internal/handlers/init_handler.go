package handlers

type PortList struct {
	AuthPort string
	PostPort string
}

type Handlers struct {
	AuthHandler *AuthHandler
	PostHandler *PostHandler
}

func NewHandlers(portList *PortList) *Handlers {
	return &Handlers{
		AuthHandler: NewAuthHandler(portList.AuthPort),
		PostHandler: NewPostHandler(portList.PostPort),
	}
}
