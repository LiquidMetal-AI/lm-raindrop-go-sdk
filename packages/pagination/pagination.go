// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type SearchPagePagination struct {
	HasMore    bool  `json:"has_more"`
	Page       int64 `json:"page"`
	PageSize   int64 `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Page        respjson.Field
		PageSize    respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchPagePagination) RawJSON() string { return r.JSON.raw }
func (r *SearchPagePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchPage[T any] struct {
	Pagination SearchPagePagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r SearchPage[T]) RawJSON() string { return r.JSON.raw }
func (r *SearchPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SearchPage[T]) GetNextPage() (res *SearchPage[T], err error) {
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SearchPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SearchPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SearchPageAutoPager[T any] struct {
	page *SearchPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewSearchPageAutoPager[T any](page *SearchPage[T], err error) *SearchPageAutoPager[T] {
	return &SearchPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SearchPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.data) == 0 {
		return false
	}
	if r.idx >= len(r.page.data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.data) == 0 {
			return false
		}
	}
	r.cur = r.page.data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SearchPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SearchPageAutoPager[T]) Err() error {
	return r.err
}

func (r *SearchPageAutoPager[T]) Index() int {
	return r.run
}
