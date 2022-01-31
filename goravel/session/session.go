package session

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifeTime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
	DBPool         *sql.DB
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	// how long should session last?
	minutes, err := strconv.Atoi(c.CookieLifeTime)

	if err != nil {
		minutes = 60
	}

	// should cookie persist
	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	} else {
		persist = false
	}

	// must cookie be secure
	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	}

	// create session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// which session store
	switch strings.ToLower(c.SessionType) {
	case "redis":

	case "mysql", "mariadb":
		session.Store = mysqlstore.New(c.DBPool)

	case "postgres", "postgresql":
		session.Store = postgresstore.New(c.DBPool)

	default:
		// cookie

	}

	return session
}
