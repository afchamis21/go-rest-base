package product

import (
	"alura-go-base/types"
	"alura-go-base/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductRouter struct {
	productService types.IProductService
}

func (c *ProductRouter) RegisterHandlers(router *mux.Router) {
	subRouter := router.PathPrefix("/product").Subrouter()

	subRouter.HandleFunc("/create", c.HandleCreateProduct).Methods(http.MethodPost)
	subRouter.HandleFunc("/all", c.HandleGetAllProducts).Methods(http.MethodGet)
	subRouter.HandleFunc("/", c.HandleGetProductBySKU).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", c.HandleDeleteProduct).Methods(http.MethodDelete)
}

// HandleCreateProduct implements types.IProductRouter.
func (c *ProductRouter) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	req, e := util.ReadJson[types.CreateProductPayload](r)
	if e != nil {
		util.WriteError(w, http.StatusBadRequest, e.Error())
		return
	}

	res, err := c.productService.CreateProduct(req)
	if err != nil {
		util.WriteError(w, err.StatusCd, err.Message)
		return
	}

	util.WriteJson(w, http.StatusCreated, res)
}

// HandleDeleteProduct implements types.IProductRouter.
func (c *ProductRouter) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	e := c.productService.DeleteProduct(id)
	if e != nil {
		util.WriteError(w, e.StatusCd, e.Message)
	}

	util.WriteJson(w, http.StatusOK, map[string]string{"message": "product deleted with success"})
}

// HandleGetAllProducts implements types.IProductRouter.
func (c *ProductRouter) HandleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	res, err := c.productService.GetAllProducts()

	if err != nil {
		util.WriteError(w, err.StatusCd, err.Message)
	}

	util.WriteJson(w, http.StatusOK, res)
}

// HandleGetProductBySKU implements types.IProductRouter.
func (c *ProductRouter) HandleGetProductBySKU(w http.ResponseWriter, r *http.Request) {
	sku := r.URL.Query().Get("sku")
	if len(sku) == 0 {
		util.WriteError(w, http.StatusBadRequest, "missing query param 'sku'")
		return
	}

	product, e := c.productService.GetProductBySKU(sku)
	if e != nil {
		util.WriteError(w, e.StatusCd, e.Message)
		return
	}

	util.WriteJson(w, http.StatusOK, product)
}

func NewProductRouter(productService types.IProductService, router *mux.Router) *ProductRouter {
	productRouter := &ProductRouter{
		productService: productService,
	}

	productRouter.RegisterHandlers(router)

	return productRouter
}
