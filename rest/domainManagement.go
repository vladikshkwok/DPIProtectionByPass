package rest

import (
	"awesomeProject/domain"
	"awesomeProject/utils"
	"fmt"
	"net/http"
)

func AddDomainHandler(templates *domain.Templates, page *domain.PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

func UpdateDomainHandler(templates *domain.Templates, page *domain.PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

func DeleteDomainHandler(templates *domain.Templates, page *domain.PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
