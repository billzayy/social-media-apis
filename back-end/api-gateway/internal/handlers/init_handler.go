package handlers

type PortList struct {
	AuthPort         string
	PostPort         string
	UserPort         string
	NotificationPort string
}

type Handlers struct {
	AuthHandler         *AuthHandler
	PostHandler         *PostHandler
	InteractHandler     *InteractHandler
	UserHandler         *UserHandler
	NotificationHandler *NotificationHandler
}

func NewHandlers(portList *PortList) *Handlers {
	return &Handlers{
		AuthHandler:         NewAuthHandler(portList.AuthPort),
		PostHandler:         NewPostHandler(portList.PostPort),
		InteractHandler:     NewInteractHandler(portList.PostPort),
		UserHandler:         NewUserHandler(portList.UserPort),
		NotificationHandler: NewNotificationHandler(portList.NotificationPort),
	}
}
