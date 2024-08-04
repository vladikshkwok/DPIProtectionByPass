package main

import (
	"awesomeProject/domain"
	"awesomeProject/rest"
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

	mux.HandleFunc("/router/stats", rest.RouterStatsHandler(templates, page))
	mux.HandleFunc("/dpi/switch", rest.DpiSwitchHandler(templates, page))
	mux.HandleFunc("/domains/add", rest.AddDomainHandler(templates, page))
	mux.HandleFunc("/domains/update", rest.UpdateDomainHandler(templates, page))
	mux.HandleFunc("/domains/delete", rest.DeleteDomainHandler(templates, page))
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
