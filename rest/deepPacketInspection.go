package rest

import (
	"awesomeProject/domain"
	"awesomeProject/utils"
	"fmt"
	"net/http"
)

func DpiSwitchHandler(templates *domain.Templates, page *domain.PageData) http.HandlerFunc {
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
