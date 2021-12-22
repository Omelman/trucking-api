package http

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	swagger "github.com/swaggo/http-swagger"

	"github.com/Omelman/trucking-api/src/config"
	"github.com/Omelman/trucking-api/src/server/handlers"
	healthcheck "github.com/Omelman/trucking-api/src/server/health-check"
	middleware "github.com/Omelman/trucking-api/src/server/http/middlewares"
	"github.com/Omelman/trucking-api/src/server/http/middlewares/policy"
)

const (
	version1 = "/v1"
)

type Server struct {
	http      *http.Server
	runErr    error
	readiness bool

	config *config.HTTP

	// handlers
	auh *handlers.AuthHandler
	itm *handlers.ItemHandler
	veh *handlers.VehicleHandler
}

func New(cfg *config.HTTP,
	authHandler *handlers.AuthHandler,
	itmHandler *handlers.ItemHandler,
	vehHandler *handlers.VehicleHandler,
) (*Server, error) {
	httpSrv := http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
	}

	// build Server
	srv := Server{
		config: cfg,
		auh:    authHandler,
		itm:    itmHandler,
		veh:    vehHandler,
	}

	if err := srv.setupHTTP(&httpSrv); err != nil {
		return nil, err
	}

	return &srv, nil
}

func (s *Server) setupHTTP(srv *http.Server) error {
	handler, err := s.buildHandler()
	if err != nil {
		return err
	}

	srv.Handler = handler
	s.http = srv

	return nil
}

// nolint: funlen,lll
func (s *Server) buildHandler() (http.Handler, error) {
	var (
		router        = mux.NewRouter()
		serviceRouter = router.PathPrefix(s.config.URLPrefix).Subrouter()
		v1Router      = serviceRouter.PathPrefix(version1).Subrouter()

		publicChain  = alice.New()
		privateChain = publicChain.
				Append(middleware.Auth)
		external = publicChain.
				Append(middleware.ExternalServiceAuth)

		owner    = privateChain.Append(policy.Owner)
		customer = privateChain.Append(policy.Customer)
	)

	// public routes
	v1Router.Handle("/health", publicChain.ThenFunc(healthcheck.Health)).Methods(http.MethodGet)
	v1Router.Handle("/login", publicChain.ThenFunc(s.auh.Login)).Methods(http.MethodPost)
	v1Router.Handle("/token", publicChain.ThenFunc(s.auh.TokenRefresh)).Methods(http.MethodPost)
	v1Router.Handle("/users/register", publicChain.ThenFunc(s.auh.Create)).Methods(http.MethodPost)

	// customer routes
	v1Router.Handle("/item", customer.ThenFunc(s.itm.CreateItem)).Methods(http.MethodPost)
	v1Router.Handle("/item", customer.ThenFunc(s.itm.UpdateItem)).Methods(http.MethodPut)
	v1Router.Handle("/item", customer.ThenFunc(s.itm.GetAllItems)).Methods(http.MethodGet)

	// owner routes
	v1Router.Handle("/vehicle", owner.ThenFunc(s.veh.CreateVehicle)).Methods(http.MethodPost)
	v1Router.Handle("/vehicle", owner.ThenFunc(s.veh.UpdateVehicle)).Methods(http.MethodPut)
	v1Router.Handle("/users/vehicle", external.ThenFunc(s.veh.GetAllOwnerVehicle)).Methods(http.MethodGet)
	v1Router.Handle("/users/vehicle/{vehicle_id:[0-9]+}", owner.ThenFunc(s.veh.DeleteUserVehicle)).Methods(http.MethodDelete)

	// external
	v1Router.Handle("/vehicle", external.ThenFunc(s.veh.GetVehicle)).Methods(http.MethodGet)
	v1Router.Handle("/item", external.ThenFunc(s.itm.GetAllItems)).Methods(http.MethodGet)

	// statistics
	v1Router.Handle("/statistics", privateChain.ThenFunc(s.veh.GetStatistics)).Methods(http.MethodGet)

	// private routes
	v1Router.Handle("/logout", privateChain.ThenFunc(s.auh.Logout)).Methods(http.MethodDelete)

	// ================================= Swagger =================================================

	if s.config.SwaggerEnable {
		router.
			PathPrefix("/swagger/static").
			Handler(http.StripPrefix("/swagger/static", http.FileServer(http.Dir(s.config.SwaggerServeDir))))
		router.
			PathPrefix("/swagger").
			Handler(swagger.Handler(swagger.URL("/swagger/static/swagger.json")))
	}

	return cors.New(cors.Options{
		AllowedOrigins: s.config.CORSAllowedHost,
		AllowedMethods: []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut,
			http.MethodDelete, http.MethodOptions, http.MethodPatch},
		AllowedHeaders:     []string{"*"},
		AllowCredentials:   true,
		OptionsPassthrough: false,
	}).Handler(router), nil
}

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	log.Info("http service: begin run")

	go func() {
		defer wg.Done()
		log.Debug("http service: addr=", s.http.Addr)
		err := s.http.ListenAndServe()
		s.runErr = err
		log.Info("http service: end run > ", err)
	}()

	go func() {
		<-ctx.Done()
		sdCtx, _ := context.WithTimeout(context.Background(), 5*time.Second) // nolint
		err := s.http.Shutdown(sdCtx)

		if err != nil {
			log.Info("http service shutdown (", err, ")")
		}
	}()

	s.readiness = true
}
