package http

import (
	"encoding/json"
	"fmt"
	"github.com/BohdanBoriak/boilerplate-go-back/config"
	"github.com/BohdanBoriak/boilerplate-go-back/config/container"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
)

func Router(cont container.Container) http.Handler {

	router := chi.NewRouter()

	router.Use(middleware.RedirectSlashes, middleware.Logger, cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "capacitor://localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"service": "IoT Device Management API",
			"version": "v1",
			"endpoints": []string{
				"GET  /api/ping",
				"POST /api/v1/auth/register",
				"POST /api/v1/auth/login",
				"GET  /api/v1/organizations",
				"POST /api/v1/organizations",
				"GET  /api/v1/organizations/{orgId}/rooms",
				"POST /api/v1/organizations/{orgId}/rooms",
				"GET  /api/v1/organizations/{orgId}/devices",
				"POST /api/v1/organizations/{orgId}/devices",
				"GET  /api/v1/devices/{devId}/measurements",
				"POST /api/v1/devices/{devId}/measurements",
				"GET  /api/v1/devices/{devId}/events",
				"POST /api/v1/devices/{devId}/events",
			},
		})
	})

	router.Route("/api", func(apiRouter chi.Router) {
		// Health
		apiRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())
			healthRouter.Handle("/*", NotFoundJSON())
		})

		apiRouter.Route("/v1", func(apiRouter chi.Router) {
			// Public routes
			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouter.Route("/auth", func(apiRouter chi.Router) {
					AuthRouter(apiRouter, cont.AuthController, cont.AuthMw)
				})
			})

			// Protected routes
			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouter.Use(cont.AuthMw)

				UserRouter(apiRouter, cont.UserController)
				OrganizationRouter(apiRouter, cont.OrganizationController)
				RoomRouter(apiRouter, cont.RoomController)
				DeviceRouter(apiRouter, cont.DeviceController)
				MeasurementRouter(apiRouter, cont.MeasurementController)
				EventRouter(apiRouter, cont.EventController)
				apiRouter.Handle("/*", NotFoundJSON())
			})
		})
	})

	router.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		workDir, _ := os.Getwd()
		filesDir := http.Dir(filepath.Join(workDir, config.GetConfiguration().FileStorageLocation))
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))
		fs.ServeHTTP(w, r)
	})

	return router
}

func AuthRouter(r chi.Router, ac controllers.AuthController, amw func(http.Handler) http.Handler) {
	r.Route("/", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/register",
			ac.Register(),
		)
		apiRouter.Post(
			"/login",
			ac.Login(),
		)
		apiRouter.With(amw).Post(
			"/logout",
			ac.Logout(),
		)
	})
}

func UserRouter(r chi.Router, uc controllers.UserController) {
	r.Route("/users", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			uc.FindMe(),
		)
		apiRouter.Put(
			"/",
			uc.Update(),
		)
		apiRouter.Delete(
			"/",
			uc.Delete(),
		)
	})
}

func OrganizationRouter(r chi.Router, oc controllers.OrganizationController) {
	r.Route("/organizations", func(apiRouter chi.Router) {
		apiRouter.Get("/", oc.FindAll())
		apiRouter.Post("/", oc.Save())
		apiRouter.Route("/{id}", func(orgRouter chi.Router) {
			orgRouter.Get("/", oc.Find())
			orgRouter.Put("/", oc.Update())
			orgRouter.Delete("/", oc.Delete())
		})
	})
}

func RoomRouter(r chi.Router, rc controllers.RoomController) {
	r.Route("/organizations/{orgId}/rooms", func(apiRouter chi.Router) {
		apiRouter.Get("/", rc.FindByOrg())
		apiRouter.Post("/", rc.Save())
		apiRouter.Route("/{id}", func(roomRouter chi.Router) {
			roomRouter.Get("/", rc.Find())
			roomRouter.Put("/", rc.Update())
			roomRouter.Delete("/", rc.Delete())
		})
	})
}

func DeviceRouter(r chi.Router, dc controllers.DeviceController) {
	r.Route("/organizations/{orgId}/devices", func(apiRouter chi.Router) {
		apiRouter.Get("/", dc.FindByOrg())
		apiRouter.Post("/", dc.Save())
		apiRouter.Get("/rooms/{roomId}", dc.FindByRoom())
		apiRouter.Route("/{id}", func(devRouter chi.Router) {
			devRouter.Get("/", dc.Find())
			devRouter.Put("/", dc.Update())
			devRouter.Delete("/", dc.Delete())
		})
	})
}

func MeasurementRouter(r chi.Router, mc controllers.MeasurementController) {
	r.Route("/devices/{devId}/measurements", func(apiRouter chi.Router) {
		apiRouter.Get("/", mc.FindByDevice())
		apiRouter.Post("/", mc.Save())
		apiRouter.Route("/{id}", func(mRouter chi.Router) {
			mRouter.Get("/", mc.Find())
			mRouter.Put("/", mc.Update())
			mRouter.Delete("/", mc.Delete())
		})
	})
}

func EventRouter(r chi.Router, ec controllers.EventController) {
	r.Route("/devices/{devId}/events", func(apiRouter chi.Router) {
		apiRouter.Get("/", ec.FindByDevice())
		apiRouter.Post("/", ec.Save())
		apiRouter.Route("/{id}", func(eRouter chi.Router) {
			eRouter.Get("/", ec.Find())
			eRouter.Put("/", ec.Update())
			eRouter.Delete("/", ec.Delete())
		})
	})
}

func NotFoundJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode("Resource Not Found")
		if err != nil {
			fmt.Printf("writing response: %s", err)
		}
	}
}

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("Ok")
		if err != nil {
			fmt.Printf("writing response: %s", err)
		}
	}
}
