package http_api

import (
	"sync"

	"github.com/RetakerMZ/backend-tiktok-video-player/domain"
	"github.com/RetakerMZ/backend-tiktok-video-player/library/fiber_response"
	"github.com/gofiber/fiber/v2"
)

type httpSearchApiDelivery struct {
	searchService domain.SearchService
}

func NewSearchHttpApiDelivery(app fiber.Router, searchService domain.SearchService) {
	handler := httpSearchApiDelivery{
		searchService: searchService,
	}

	app.Get("/search", handler.Search)

}

func (h *httpSearchApiDelivery) Search(ctx *fiber.Ctx) error {
	wg := ctx.Locals("wg").(*sync.WaitGroup)
	wg.Add(1)
	defer wg.Done()

	query := ctx.Query("query")
	cursor := ctx.Query("cursor")
	searchRes, err := h.searchService.GetSearch(query, cursor)
	if err != nil {
		return fiber_response.ReturnStatusUnprocessableEntity(ctx, err.Error(), err)
	}
	return fiber_response.ReturnStatusOk(ctx, "Success", searchRes)
}
