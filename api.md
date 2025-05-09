# Search

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchResponse">SearchResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#TextResult">TextResult</a>

Methods:

- <code title="get /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchGetParams">SearchGetParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchService.Find">Find</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchFindParams">SearchFindParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DocumentQuery

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DocumentQueryAskResponse">DocumentQueryAskResponse</a>

Methods:

- <code title="post /v1/document_query">client.DocumentQuery.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DocumentQueryService.Ask">Ask</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DocumentQueryAskParams">DocumentQueryAskParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DocumentQueryAskResponse">DocumentQueryAskResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ChunkSearch

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChunkSearchFindResponse">ChunkSearchFindResponse</a>

Methods:

- <code title="post /v1/chunk_search">client.ChunkSearch.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChunkSearchService.Find">Find</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChunkSearchFindParams">ChunkSearchFindParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChunkSearchFindResponse">ChunkSearchFindResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SummarizePage

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizePageNewResponse">SummarizePageNewResponse</a>

Methods:

- <code title="post /v1/summarize_page">client.SummarizePage.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizePageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizePageNewParams">SummarizePageNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizePageNewResponse">SummarizePageNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# StorageObject

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectListResponse">StorageObjectListResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectDeleteResponse">StorageObjectDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectUploadResponse">StorageObjectUploadResponse</a>

Methods:

- <code title="get /v1/object/{bucket}">client.StorageObject.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bucket <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectListResponse">StorageObjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/object/{bucket}/{key}">client.StorageObject.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, key <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectDeleteParams">StorageObjectDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectDeleteResponse">StorageObjectDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/object/{bucket}/{key}">client.StorageObject.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, key <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectDownloadParams">StorageObjectDownloadParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/object/{bucket}/{key}">client.StorageObject.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, key <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectUploadParams">StorageObjectUploadParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#StorageObjectUploadResponse">StorageObjectUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
