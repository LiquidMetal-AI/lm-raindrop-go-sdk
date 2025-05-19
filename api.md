# DocumentQuery

Params Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#BucketLocatorUnionParam">BucketLocatorUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DocumentQueryAskResponse">DocumentQueryAskResponse</a>

Methods:

- <code title="post /v1/document_query">client.DocumentQuery.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DocumentQueryService.Ask">Ask</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DocumentQueryAskParams">DocumentQueryAskParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DocumentQueryAskResponse">DocumentQueryAskResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ChunkSearch

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#TextResult">TextResult</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChunkSearchFindResponse">ChunkSearchFindResponse</a>

Methods:

- <code title="post /v1/chunk_search">client.ChunkSearch.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChunkSearchService.Find">Find</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChunkSearchFindParams">ChunkSearchFindParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChunkSearchFindResponse">ChunkSearchFindResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SummarizePage

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizePageSumarizePageResponse">SummarizePageSumarizePageResponse</a>

Methods:

- <code title="post /v1/summarize_page">client.SummarizePage.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizePageService.SumarizePage">SumarizePage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizePageSumarizePageParams">SummarizePageSumarizePageParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SummarizePageSumarizePageResponse">SummarizePageSumarizePageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Search

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchFindResponse">SearchFindResponse</a>

Methods:

- <code title="post /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchService.Find">Find</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchFindParams">SearchFindParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchFindResponse">SearchFindResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PutObject

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutObjectUploadResponse">PutObjectUploadResponse</a>

Methods:

- <code title="post /v1/put_object">client.PutObject.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutObjectService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutObjectUploadParams">PutObjectUploadParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#PutObjectUploadResponse">PutObjectUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# GetObject

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetObjectDownloadResponse">GetObjectDownloadResponse</a>

Methods:

- <code title="post /v1/get_object">client.GetObject.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetObjectService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetObjectDownloadParams">GetObjectDownloadParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#GetObjectDownloadResponse">GetObjectDownloadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DeleteObject

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteObjectDeleteResponse">DeleteObjectDeleteResponse</a>

Methods:

- <code title="post /v1/delete_object">client.DeleteObject.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteObjectService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteObjectDeleteParams">DeleteObjectDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#DeleteObjectDeleteResponse">DeleteObjectDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ListObjects

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ListObjectNewResponse">ListObjectNewResponse</a>

Methods:

- <code title="post /v1/list_objects">client.ListObjects.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ListObjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ListObjectNewParams">ListObjectNewParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ListObjectNewResponse">ListObjectNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
