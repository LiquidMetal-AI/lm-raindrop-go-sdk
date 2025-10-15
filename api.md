# Query

Params Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketLocatorUnionParam">BucketLocatorUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryChunkSearchResponse">QueryChunkSearchResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryDocumentQueryResponse">QueryDocumentQueryResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryGetPaginatedSearchResponse">QueryGetPaginatedSearchResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QuerySearchResponse">QuerySearchResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QuerySumarizePageResponse">QuerySumarizePageResponse</a>

Methods:

- <code title="post /v1/chunk_search">client.Query.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryService.ChunkSearch">ChunkSearch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryChunkSearchParams">QueryChunkSearchParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryChunkSearchResponse">QueryChunkSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/document_query">client.Query.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryService.DocumentQuery">DocumentQuery</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryDocumentQueryParams">QueryDocumentQueryParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryDocumentQueryResponse">QueryDocumentQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/search_get_page">client.Query.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryService.GetPaginatedSearch">GetPaginatedSearch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryGetPaginatedSearchParams">QueryGetPaginatedSearchParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryGetPaginatedSearchResponse">QueryGetPaginatedSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/search">client.Query.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QuerySearchParams">QuerySearchParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QuerySearchResponse">QuerySearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/summarize_page">client.Query.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryService.SumarizePage">SumarizePage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QuerySumarizePageParams">QuerySumarizePageParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QuerySumarizePageResponse">QuerySumarizePageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Memory

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryMemorySearchResponse">QueryMemorySearchResponse</a>

Methods:

- <code title="post /v1/search_memory">client.Query.Memory.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryMemoryService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryMemorySearchParams">QueryMemorySearchParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryMemorySearchResponse">QueryMemorySearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## EpisodicMemory

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryEpisodicMemorySearchResponse">QueryEpisodicMemorySearchResponse</a>

Methods:

- <code title="post /v1/search_episodic_memory">client.Query.EpisodicMemory.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryEpisodicMemoryService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryEpisodicMemorySearchParams">QueryEpisodicMemorySearchParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryEpisodicMemorySearchResponse">QueryEpisodicMemorySearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Procedures

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryProcedureSearchResponse">QueryProcedureSearchResponse</a>

Methods:

- <code title="post /v1/search_procedures">client.Query.Procedures.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryProcedureService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryProcedureSearchParams">QueryProcedureSearchParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QueryProcedureSearchResponse">QueryProcedureSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## SemanticMemory

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QuerySemanticMemorySearchResponse">QuerySemanticMemorySearchResponse</a>

Methods:

- <code title="post /v1/search_semantic_memory">client.Query.SemanticMemory.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QuerySemanticMemoryService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QuerySemanticMemorySearchParams">QuerySemanticMemorySearchParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#QuerySemanticMemorySearchResponse">QuerySemanticMemorySearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Bucket

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketListResponse">BucketListResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketDeleteResponse">BucketDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketGetResponse">BucketGetResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketPutResponse">BucketPutResponse</a>

Methods:

- <code title="post /v1/list_objects">client.Bucket.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketListParams">BucketListParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketListResponse">BucketListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/delete_object">client.Bucket.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketDeleteParams">BucketDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketDeleteResponse">BucketDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/get_object">client.Bucket.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketGetParams">BucketGetParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketGetResponse">BucketGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/put_object">client.Bucket.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketService.Put">Put</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketPutParams">BucketPutParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketPutResponse">BucketPutResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PutMemory

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutMemoryNewResponse">PutMemoryNewResponse</a>

Methods:

- <code title="post /v1/put_memory">client.PutMemory.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutMemoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutMemoryNewParams">PutMemoryNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutMemoryNewResponse">PutMemoryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# GetMemory

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetMemoryGetResponse">GetMemoryGetResponse</a>

Methods:

- <code title="post /v1/get_memory">client.GetMemory.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetMemoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetMemoryGetParams">GetMemoryGetParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetMemoryGetResponse">GetMemoryGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DeleteMemory

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteMemoryNewResponse">DeleteMemoryNewResponse</a>

Methods:

- <code title="post /v1/delete_memory">client.DeleteMemory.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteMemoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteMemoryNewParams">DeleteMemoryNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteMemoryNewResponse">DeleteMemoryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SummarizeMemory

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizeMemoryNewResponse">SummarizeMemoryNewResponse</a>

Methods:

- <code title="post /v1/summarize_memory">client.SummarizeMemory.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizeMemoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizeMemoryNewParams">SummarizeMemoryNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizeMemoryNewResponse">SummarizeMemoryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# StartSession

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StartSessionNewResponse">StartSessionNewResponse</a>

Methods:

- <code title="post /v1/start_session">client.StartSession.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StartSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StartSessionNewParams">StartSessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StartSessionNewResponse">StartSessionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# EndSession

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#EndSessionNewResponse">EndSessionNewResponse</a>

Methods:

- <code title="post /v1/end_session">client.EndSession.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#EndSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#EndSessionNewParams">EndSessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#EndSessionNewResponse">EndSessionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RehydrateSession

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#RehydrateSessionRehydrateResponse">RehydrateSessionRehydrateResponse</a>

Methods:

- <code title="post /v1/rehydrate_session">client.RehydrateSession.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#RehydrateSessionService.Rehydrate">Rehydrate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#RehydrateSessionRehydrateParams">RehydrateSessionRehydrateParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#RehydrateSessionRehydrateResponse">RehydrateSessionRehydrateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PutProcedure

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutProcedureNewResponse">PutProcedureNewResponse</a>

Methods:

- <code title="post /v1/put_procedure">client.PutProcedure.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutProcedureService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutProcedureNewParams">PutProcedureNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutProcedureNewResponse">PutProcedureNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# GetProcedure

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetProcedureNewResponse">GetProcedureNewResponse</a>

Methods:

- <code title="post /v1/get_procedure">client.GetProcedure.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetProcedureService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetProcedureNewParams">GetProcedureNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetProcedureNewResponse">GetProcedureNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DeleteProcedure

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteProcedureNewResponse">DeleteProcedureNewResponse</a>

Methods:

- <code title="post /v1/delete_procedure">client.DeleteProcedure.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteProcedureService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteProcedureNewParams">DeleteProcedureNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteProcedureNewResponse">DeleteProcedureNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ListProcedures

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ListProcedureNewResponse">ListProcedureNewResponse</a>

Methods:

- <code title="post /v1/list_procedures">client.ListProcedures.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ListProcedureService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ListProcedureNewParams">ListProcedureNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ListProcedureNewResponse">ListProcedureNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PutSemanticMemory

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutSemanticMemoryNewResponse">PutSemanticMemoryNewResponse</a>

Methods:

- <code title="post /v1/put_semantic_memory">client.PutSemanticMemory.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutSemanticMemoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutSemanticMemoryNewParams">PutSemanticMemoryNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutSemanticMemoryNewResponse">PutSemanticMemoryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# GetSemanticMemory

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetSemanticMemoryNewResponse">GetSemanticMemoryNewResponse</a>

Methods:

- <code title="post /v1/get_semantic_memory">client.GetSemanticMemory.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetSemanticMemoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetSemanticMemoryNewParams">GetSemanticMemoryNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetSemanticMemoryNewResponse">GetSemanticMemoryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DeleteSemanticMemory

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteSemanticMemoryDeleteResponse">DeleteSemanticMemoryDeleteResponse</a>

Methods:

- <code title="post /v1/delete_semantic_memory">client.DeleteSemanticMemory.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteSemanticMemoryService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteSemanticMemoryDeleteParams">DeleteSemanticMemoryDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteSemanticMemoryDeleteResponse">DeleteSemanticMemoryDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
