package server

import (
	"net/http"

	"github.com/gira-games/client/pkg/client"

	"github.com/golangcollege/sessions"
	"github.com/sirupsen/logrus"
)

var (
	homePage       = "home.page.tmpl"
	addGamePage    = "add.page.tmpl"
	listGamesPage  = "list.page.tmpl"
	createGamePage = "create.page.tmpl"
	signupUserPage = "signup.page.tmpl"
	loginUserPage  = "login.page.tmpl"

	emptyTemplateData = TemplateData{}
)

// TemplateData is the struct that holds all the data that can be passed to the template renderer to render
type TemplateData struct {
	Game       *client.Game
	User       *client.User
	Games      []*client.Game
	UserGames  []*client.UserGame
	Statuses   []client.Status
	Franchises []*client.Franchise

	SelectedFranchiseID string
	Error               string
	Flash               string
}

// Renderer is the interface that will be used to interact with the part of the program
// that is responsible for rendering the web pages
type Renderer interface {
	Render(w http.ResponseWriter, r *http.Request, d TemplateData, p string) error
}

// APIClient is the interface that interacts with the API
type APIClient interface {
	GetFranchises(*client.GetFranchisesRequest) (*client.GetFranchisesResponse, error)
	CreateFranchise(*client.CreateFranchiseRequest) (*client.CreateFranchiseResponse, error)

	GetGames(*client.GetGamesRequest) (*client.GetGamesResponse, error)
	CreateGame(*client.CreateGameRequest) (*client.CreateGameResponse, error)

	GetUserGames(*client.GetUserGamesRequest) (*client.GetUserGamesResponse, error)
	LinkGameToUser(*client.LinkGameToUserRequest) error
	UpdateGameProgress(*client.UpdateGameProgressRequest) error
	DeleteUserGame(*client.DeleteUserGameRequest) error

	LoginUser(*client.LoginUserRequest) (*client.UserLoginResponse, error)
	CreateUser(*client.CreateUserRequest) (*client.CreateUserResponse, error)
	GetUser(*client.GetUserRequest) (*client.GetUserResponse, error)
	LogoutUser(*client.LogoutUserRequest) error
}

// Server is the struct that holds all the dependencies
// needed to run the application
type Server struct {
	Log      *logrus.Logger
	Session  *sessions.Session
	Client   APIClient
	Renderer Renderer
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.routes().ServeHTTP(w, r)
}
