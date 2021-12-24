package goravel

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/namnguyen191/goravel/render"
	"github.com/namnguyen191/goravel/session"
)

const version = "1.0.0"

type Goravel struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	Session  *scs.SessionManager
	JetViews *jet.Set
	config   config
}

type config struct {
	// the port the server will listen on
	port string
	// the renderer engine that the app will be using (jet or go)
	renderer    string
	cookie      cookieConfig
	sessionType string
}

func (grv *Goravel) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := grv.Init(pathConfig)

	if err != nil {
		return err
	}

	err = grv.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// read .env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	// create logger
	infoLog, errorLog := grv.startLoggers()
	grv.InfoLog = infoLog
	grv.ErrorLog = errorLog

	grv.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	grv.Version = version
	grv.RootPath = rootPath

	grv.Routes = grv.routes().(*chi.Mux)

	grv.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
		cookie: cookieConfig{
			name:     os.Getenv("COOKIE_NAME"),
			lifetime: os.Getenv("COOKIE_LIFETIME"),
			persist:  os.Getenv("COOKIE_PERSISTS"),
			secure:   os.Getenv("COOKIE_SECURE"),
			domain:   os.Getenv("COOKIE_DOMAIN"),
		},
		sessionType: os.Getenv("SESSION_TYPE"),
	}

	// create a Session
	sess := session.Session{
		CookieLifeTime: grv.config.cookie.lifetime,
		CookiePersist:  grv.config.cookie.persist,
		CookieName:     grv.config.cookie.name,
		SessionType:    grv.config.sessionType,
		CookieDomain:   grv.config.cookie.domain,
	}

	grv.Session = sess.InitSession()

	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	grv.JetViews = views

	grv.createRenderer()

	return nil
}

func (grv *Goravel) Init(p initPaths) error {
	root := p.rootPath

	for _, path := range p.folderNames {
		// create folder if it does not exist
		err := grv.CreateDirIfNotExist(root + "/" + path)

		if err != nil {
			return err
		}
	}

	return nil
}

// ListenAndServe starts web server
func (grv *Goravel) ListenAndServe() {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     grv.ErrorLog,
		Handler:      grv.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	grv.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	grv.ErrorLog.Fatal(err)
}

func (grv *Goravel) checkDotEnv(path string) error {
	err := grv.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return err
	}

	return nil
}

func (grv *Goravel) CreateFileIfNotExist(path string) error {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)

		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}

	return nil
}

func (grv *Goravel) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (grv *Goravel) createRenderer() {
	myRenderer := render.Render{
		Renderer: grv.config.renderer,
		RootPath: grv.RootPath,
		Port:     grv.config.port,
		JetViews: grv.JetViews,
	}

	grv.Render = &myRenderer
}
