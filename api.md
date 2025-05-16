# DocumentQuery

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#BucketLocatorUnionParam">BucketLocatorUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#DocumentQueryNewResponse">DocumentQueryNewResponse</a>

Methods:

- <code title="post /v1/document_query">client.DocumentQuery.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#DocumentQueryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#DocumentQueryNewParams">DocumentQueryNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#DocumentQueryNewResponse">DocumentQueryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ChunkSearch

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#TextResult">TextResult</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ChunkSearchExecuteResponse">ChunkSearchExecuteResponse</a>

Methods:

- <code title="post /v1/chunk_search">client.ChunkSearch.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ChunkSearchService.Execute">Execute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ChunkSearchExecuteParams">ChunkSearchExecuteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ChunkSearchExecuteResponse">ChunkSearchExecuteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SummarizePage

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SummarizePageNewSummaryResponse">SummarizePageNewSummaryResponse</a>

Methods:

- <code title="post /v1/summarize_page">client.SummarizePage.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SummarizePageService.NewSummary">NewSummary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SummarizePageNewSummaryParams">SummarizePageNewSummaryParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SummarizePageNewSummaryResponse">SummarizePageNewSummaryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Search

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#PaginationInfo">PaginationInfo</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchRunResponse">SearchRunResponse</a>

Methods:

- <code title="post /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchRunParams">SearchRunParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchRunResponse">SearchRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# LiquidmetalV1alpha1SearchAgentService

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsResponse">LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsResponse</a>

Methods:

- <code title="post /liquidmetal.v1alpha1.SearchAgentService/GetPaginatedResults">client.LiquidmetalV1alpha1SearchAgentService.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#LiquidmetalV1alpha1SearchAgentServiceService.GetPaginatedResults">GetPaginatedResults</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsParams">LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsResponse">LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
