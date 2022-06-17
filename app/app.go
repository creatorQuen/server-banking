package app

import (
	"ashishi-banking/domain"
	"ashishi-banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func Start() {
	router := mux.NewRouter()
	//go http.ListenAndServe("localhost:3000", nil)
	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustommerRepositoryStub())}

	dbConnector := domain.NewCustomerRepositoryDb()
	ch := CustomerHandler{service.NewCustomerService(dbConnector)}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	//router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	//router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
