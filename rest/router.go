package rest

import (
	"awesomeProject/domain"
	"awesomeProject/utils"
	"fmt"
	"net/http"
)

func RouterStatsHandler(templates *domain.Templates, page *domain.PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page.Router.Stats = utils.GetRouterStats(false)
		page.DpiProp.Status = utils.GetDpiProtectionStatus()

		if err := templates.Templates.ExecuteTemplate(w, "router-stats", page.Router); err != nil {
			http.Error(w, fmt.Sprintf("error rendering template: %v", err), http.StatusInternalServerError)
		}
	}
}
