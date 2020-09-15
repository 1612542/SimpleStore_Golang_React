package api

// func GetProducts(s service.ProductService) http.Handler {
// 	return http.Handle(func(w http.ResponseWriter, r *http.Request) {
// 		products, err := s.FetchAll()
// 		if err != nil {
// 			utils.Respond(w, utils.Message(http.StatusInternalServerError, err.Error()))
// 			return
// 		}

// 		res := utils.Message(http.StatusOK, "Get data successfully")
// 		res["product"] = products
// 		utils.Respond(w, res)
// 	})
// }

// func MakeProductHandler(router *mux.Router, service service.ProductService) {
// 	router.Handle("/api/product", FetchAll(service).Methods("GET"))
// 	// router.Handle("/api/product", FetchAll(service))
// }
