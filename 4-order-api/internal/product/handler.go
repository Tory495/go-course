package product

import (
	"errors"
	"go/courses/pkg/request"
	"go/courses/pkg/response"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandler struct {
	ProductRepository *ProductRepository
}

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}

	router.HandleFunc("GET /products", handler.getAll())
	router.HandleFunc("GET /products/{id}", handler.getById())
	router.HandleFunc("POST /products", handler.create())
	router.HandleFunc("PATCH /products/{id}", handler.update())
	router.HandleFunc("DELETE /products/{id}", handler.delete())
}

func (handler *ProductHandler) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		products, err := handler.ProductRepository.GetAll(int(limit), int(page))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.SendJsonResponse(w, products, http.StatusOK)
	}
}

func (handler *ProductHandler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := handler.ProductRepository.GetById(uint(id))

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.SendJsonResponse(w, product, http.StatusOK)
	}
}

func (handler *ProductHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[ProductCreateRequest](&w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		createdProduct, err := handler.ProductRepository.Create(&Product{
			Model:       gorm.Model{},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.SendJsonResponse(w, createdProduct, http.StatusOK)
	}
}

func (handler *ProductHandler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, err := request.HandleBody[ProductUpdateRequest](&w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updatedProduct, err := handler.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.SendJsonResponse(w, updatedProduct, http.StatusOK)
	}
}

func (handler *ProductHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = handler.ProductRepository.GetById(uint(id))

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.ProductRepository.Delete(uint(id))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.SendJsonResponse(w, nil, http.StatusOK)
	}
}
