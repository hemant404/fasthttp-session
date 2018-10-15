package session

import (
	"sync"
	"time"

	"github.com/savsgio/dictpool"
	"github.com/savsgio/gotils/dao"
)

// Config config struct
type Config struct {

	// cookie name
	CookieName string

	// cookie domain
	Domain string

	// If you want to delete the cookie when the browser closes, set it to -1.
	//
	//  0 means no expire, (24 years)
	// -1 means when browser closes
	// >0 is the time.Duration which the session cookies should expire.
	Expires time.Duration

	// gc life time to execute it
	GCLifetime time.Duration

	// session life time in seconds
	SessionLifetime int64

	// set whether to pass this bar cookie only through HTTPS
	Secure bool

	// sessionID is in url query
	SessionIDInURLQuery bool

	// sessionName in url query
	SessionNameInURLQuery string

	// sessionID is in http header
	SessionIDInHTTPHeader bool

	// sessionName in http header
	SessionNameInHTTPHeader string

	// SessionIDGeneratorFunc should returns a random session id.
	SessionIDGeneratorFunc func() []byte

	// value cookie length
	cookieLen uint32
}

// Dict memory store
type Dict struct {
	dictpool.Dict
}

// Session session struct
type Session struct {
	provider Provider
	config   *Config
	cookie   *Cookie
}

// Dao database connection
type Dao struct {
	dao.Dao
}

// Store store
type Store struct {
	sessionID []byte
	data      *Dict
	lock      sync.RWMutex
}

// Encrypt encrypt struct
type Encrypt struct{}

// Cookie cookie struct
type Cookie struct{}

// Storer session store interface
type Storer interface {
	Save() error
	Get(key string) interface{}
	GetBytes(key []byte) interface{}
	GetAll() Dict
	Set(key string, value interface{})
	SetBytes(key []byte, value interface{})
	Delete(key string)
	DeleteBytes(key []byte)
	Flush()
	GetSessionID() []byte
}

// Provider provider interface
type Provider interface {
	Init(lifeTime int64, cfg ProviderConfig) error
	Get(id []byte) (Storer, error)
	Destroy(id []byte) error
	Regenerate(oldID, newID []byte) (Storer, error)
	Count() int
	NeedGC() bool
	GC()
}

// ProviderConfig provider config interface
type ProviderConfig interface {
	Name() string
}
