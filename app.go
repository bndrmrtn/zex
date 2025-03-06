package zex

import "net/http"

const Version = "1.0.1"

// App is the main struct of the application
type App struct {
	CompleteRouter
	http.Handler

	conf        *Config
	middlewares []MiddlewareFunc
	public      map[string]string
}

// New creates a new App instance
func New(conf ...*Config) *App {
	app := &App{
		CompleteRouter: newRouter(),
		middlewares:    make([]MiddlewareFunc, 0),
		public:         make(map[string]string),
	}

	if len(conf) > 0 {
		app.conf = conf[0]
	} else {
		app.conf = defaultConfig()
	}

	app.conf.make()
	app.Handler = NewServer(app)
	registerDefaultRouteValidators(app)
	return app
}

// Config returns the configuration of the application
func (a *App) Config() *Config {
	return a.conf
}

// Use registers a middleware to be used by the application
func (a *App) Use(middleware MiddlewareFunc) {
	a.middlewares = append(a.middlewares, middleware)
}

func (a *App) Public(prefix, path string) {
	a.public[prefix] = path
}

// Serve starts the server on the given address
func (a *App) Serve(listenAddr string) error {
	displayServeInfo(listenAddr, a.conf.Development)
	return http.ListenAndServe(listenAddr, a)
}

// ServeTLS starts the server on the given address with TLS
func (a *App) ServeTLS(listenAddr, certFile, keyFile string) error {
	displayServeInfo(listenAddr, a.conf.Development)
	return http.ListenAndServeTLS(listenAddr, certFile, keyFile, a)
}
