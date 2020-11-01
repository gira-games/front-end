package fixtures

//go:generate mockgen -destination renderer_mock.go  -package fixtures -mock_names Renderer=RendererMock github.com/asankov/gira/cmd/front-end/server Renderer
//go:generate mockgen -destination api_client_mock.go  -package fixtures -mock_names APIClient=APIClientMock github.com/asankov/gira/cmd/front-end/server APIClient
