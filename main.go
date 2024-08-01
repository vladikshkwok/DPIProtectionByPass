package main

import (
	"awesomeProject/domain"
	"awesomeProject/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	templates := domain.NewTemplates()
	page := domain.NewPageData()

	mux := http.NewServeMux()

	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	mux.HandleFunc("/", indexHandler(templates, page))
	mux.HandleFunc("GET /router/stats", routerStatsHandler(templates, page))
	mux.HandleFunc("POST /dpi/switch", dpiSwitchHandler(templates, page))

	log.Println("Server started on :8082")
	if err := http.ListenAndServe(":8082", logRequest(mux)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func indexHandler(templates *domain.Templates, page *domain.PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page.DpiProp.Status = utils.GetDpiProtectionStatus()
		page.Router.Stats = utils.GetRouterStats(false)

		log.Println("Return index page")

		if err := templates.Render(w, "index", page); err != nil {
			http.Error(w, fmt.Sprintf("error rendering template: %v", err), http.StatusInternalServerError)
		}
	}
}

func routerStatsHandler(templates *domain.Templates, page *domain.PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page.Router.Stats = utils.GetRouterStats(false)
		page.DpiProp.Status = utils.GetDpiProtectionStatus()

		if err := templates.Templates.ExecuteTemplate(w, "router-stats", page.Router); err != nil {
			http.Error(w, fmt.Sprintf("error rendering template: %v", err), http.StatusInternalServerError)
		}
	}
}

func dpiSwitchHandler(templates *domain.Templates, page *domain.PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := utils.SwitchProtection(); err != nil {
			http.Error(w, fmt.Sprintf("error switching dpi protection: %v", err), http.StatusInternalServerError)
			return
		}

		page.DpiProp.Status = utils.GetDpiProtectionStatus()

		if err := templates.Render(w, "dpi-prot", page.DpiProp); err != nil {
			http.Error(w, fmt.Sprintf("error rendering template: %v", err), http.StatusInternalServerError)
		}
	}
}
