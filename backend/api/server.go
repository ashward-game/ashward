package api

import (
	"context"
	"crypto/tls"
	"net/http"
	"orbit_nft/api/middleware"
	"orbit_nft/api/route"
	"orbit_nft/api/util/validation"
	"orbit_nft/db"
	"os"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/acme/autocert"
)

type apiServer struct {
	router *gin.Engine
	server *http.Server

	// options/attributes

	// event handler
	onShutDown func(ctx context.Context)
}

func NewServer(env string, db *db.Database) *apiServer {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	router := gin.New()
	switch env {
	case "development":
		gin.SetMode(gin.DebugMode)
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: false,
			},
		)
	case "production":
		gin.SetMode(gin.ReleaseMode)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		gin.SetMode(gin.DebugMode)
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// register our custom validator
	if err := validation.Register(); err != nil {
		log.Fatal().Err(err).Msg("Cannot register custom validation")
	}

	// we dont use proxies for now
	_ = router.SetTrustedProxies(nil)

	// global middleware
	router.Use(logger.SetLogger())
	router.Use(middleware.Recovery())
	router.Use(middleware.Security())
	router.Use(middleware.CORS())
	router.Use(middleware.SetSigningServiceRPC(os.Getenv("SIGNING_SERVICE_ADDRESS")))

	route.Init(db, router)
	return &apiServer{router: router}
}

// Run requires a configuration for TLS: `domain`, `cert`, `key`.
// If `domain` is empty, self-signed `cert` and `key` are required.
// Otherwise, LetsEncrypt is used for `domain`.
// Currently, we support a single domain only.
func (s *apiServer) Run(port string, domain string, tlsConfig ...string) {
	srv := &http.Server{
		Addr:    port,
		Handler: s.router,
	}
	if len(domain) == 0 {
		cert, key := tlsConfig[0], tlsConfig[1]
		cer, err := tls.LoadX509KeyPair(cert, key)
		if err != nil {
			panic(err)
		}
		srv.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cer}}
	} else {
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(domain),
			Cache:      autocert.DirCache("certs"),
		}
		srv.TLSConfig = m.TLSConfig()
	}
	s.server = srv

	if err := s.server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
		log.Error().Err(err).Msg("server closed unexpect")
	}
}

// Shutdown gracefully shutdowns the server with
// a timeout of 5 seconds.
func (s *apiServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer func() {
		if s.onShutDown != nil {
			s.onShutDown(ctx)
		}
	}()
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("server forced to shutdown")
	}
}

func (s *apiServer) OnShutdown(f func(ctx context.Context)) {
	s.onShutDown = f
}
