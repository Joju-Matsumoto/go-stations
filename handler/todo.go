package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code, res := func() (code int, response interface{}) {
		switch r.Method {
		case http.MethodGet:
			req := model.ReadTODORequest{}

			q := r.URL.Query()
			req.PrevID, _ = strconv.ParseInt(q.Get("prev_id"), 10, 64)
			req.Size, _ = strconv.ParseInt(q.Get("size"), 10, 64)

			res, err := h.Read(r.Context(), &req)
			if err != nil {
				return http.StatusInternalServerError, nil
			}
			return http.StatusOK, res
		case http.MethodPost:
			req := model.CreateTODORequest{}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return http.StatusBadRequest, nil
			}
			res, err := h.Create(r.Context(), &req)
			if err != nil {
				return http.StatusBadRequest, nil
			}
			return http.StatusOK, res
		case http.MethodPut:
			req := model.UpdateTODORequest{}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				if _, ok := err.(*model.ErrNotFound); ok {
					return http.StatusNotFound, nil
				}
				return http.StatusBadRequest, nil
			}
			res, err := h.Update(r.Context(), &req)
			if err != nil {
				return http.StatusBadRequest, nil
			}
			return http.StatusOK, res
		case http.MethodDelete:
			req := model.DeleteTODORequest{}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return http.StatusBadRequest, nil
			}
			res, err := h.Delete(r.Context(), &req)
			if err != nil {
				if _, ok := err.(*model.ErrNotFound); ok {
					return http.StatusNotFound, nil
				}
				return http.StatusBadRequest, nil
			}
			return http.StatusOK, res
		}
		return http.StatusInternalServerError, nil
	}()

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	todo, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	if err != nil {
		return nil, err
	}
	return &model.CreateTODOResponse{
		TODO: *todo,
	}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	if req.Size == 0 {
		req.Size = 5
	}
	todos, err := h.svc.ReadTODO(ctx, req.PrevID, req.Size)
	if err != nil {
		return nil, err
	}
	return &model.ReadTODOResponse{
		TODOs: todos,
	}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	if req.ID == 0 {
		return nil, fmt.Errorf("request id is 0")
	} else if len(req.Subject) == 0 {
		return nil, fmt.Errorf("request subject is empty")
	}
	todo, err := h.svc.UpdateTODO(ctx, req.ID, req.Subject, req.Description)
	if err != nil {
		return nil, err
	}
	return &model.UpdateTODOResponse{
		TODO: *todo,
	}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	if len(req.IDs) == 0 {
		return nil, fmt.Errorf("request ids is empty")
	}
	if err := h.svc.DeleteTODO(ctx, req.IDs); err != nil {
		return nil, err
	}
	return &model.DeleteTODOResponse{}, nil
}
