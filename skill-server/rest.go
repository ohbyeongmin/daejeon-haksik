package skillserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ohbyeongmin/daejeon-haksik/constants"
	"github.com/ohbyeongmin/daejeon-haksik/utils"
)

const port string = ":3000"

func TodayHandler(rw http.ResponseWriter, r *http.Request) {
	menu := HRCService.Today(constants.LUNCH)
	json.NewEncoder(rw).Encode(GetOneMenuReasponse(menu))
}

func TomorrowHandler(rw http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(rw).Encode(HRCService.Tomorrow())
}

func AllHandler(rw http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(rw).Encode(HRCService.AllWeeks())
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func ServerStart() {
	r := mux.NewRouter()
	r.Use(jsonContentTypeMiddleware)
	r.HandleFunc("/today", TodayHandler).Methods("GET")
	r.HandleFunc("/tomorrow", TomorrowHandler).Methods("GET")
	r.HandleFunc("/all", AllHandler).Methods("GET")

	fmt.Printf("Listen on http://localhost%s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		utils.HandleErr(err)
	}
}
