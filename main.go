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
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))

	mux.HandleFunc("/router/stats", routerStatsHandler(templates, page))
	mux.HandleFunc("/dpi/switch", dpiSwitchHandler(templates, page))
	mux.HandleFunc("/domains/add", addDomainHandler(templates, page))
	mux.HandleFunc("/domains/update", updateDomainHandler(templates, page))
	mux.HandleFunc("/domains/delete", deleteDomainHandler(templates, page))
	mux.HandleFunc("/", indexHandler(templates, page))

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
		page.Domains = utils.GetDomains()

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

func addDomainHandler(templates *domain.Templates, page *domain.PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		domain := r.FormValue("domain")
		if domain == "" {
			http.Error(w, "Domain is required", http.StatusBadRequest)
			return
		}

		if err := utils.AddDomain(domain); err != nil {
			http.Error(w, fmt.Sprintf("error adding domain: %v", err), http.StatusInternalServerError)
			return
		}

		// Render only the new domain item
		if err := templates.Render(w, "oob-domain", domain); err != nil {
			http.Error(w, fmt.Sprintf("error rendering domain item: %v", err), http.StatusInternalServerError)
		}
	}
}

func updateDomainHandler(templates *domain.Templates, page *domain.PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		oldDomain := r.FormValue("oldDomain")
		newDomain := r.FormValue("newDomain")
		if oldDomain == "" || newDomain == "" {
			http.Error(w, "Both old and new domains are required", http.StatusBadRequest)
			return
		}

		if err := utils.UpdateDomain(oldDomain, newDomain); err != nil {
			http.Error(w, fmt.Sprintf("error updating domain: %v", err), http.StatusInternalServerError)
			return
		}

		// Render only the updated domain item
		if err := templates.Render(w, "domain-item", newDomain); err != nil {
			http.Error(w, fmt.Sprintf("error rendering domain item: %v", err), http.StatusInternalServerError)
		}
	}
}

func deleteDomainHandler(templates *domain.Templates, page *domain.PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		domain := r.FormValue("domain")
		if domain == "" {
			http.Error(w, "Domain is required", http.StatusBadRequest)
			return
		}

		if err := utils.DeleteDomain(domain); err != nil {
			http.Error(w, fmt.Sprintf("error deleting domain: %v", err), http.StatusInternalServerError)
			return
		}

		// Return an empty response to remove the domain element from the list
		w.WriteHeader(http.StatusOK)
	}
}
