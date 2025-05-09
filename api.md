# Search

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchResponse">SearchResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#TextResult">TextResult</a>

Methods:

- <code title="get /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchGetParams">SearchGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go/packages/pagination#SearchPageQuery">SearchPageQuery</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#TextResult">TextResult</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchService.Find">Find</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchFindParams">SearchFindParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DocumentQuery

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#DocumentQueryAskResponse">DocumentQueryAskResponse</a>

Methods:

- <code title="post /v1/document_query">client.DocumentQuery.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#DocumentQueryService.Ask">Ask</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#DocumentQueryAskParams">DocumentQueryAskParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#DocumentQueryAskResponse">DocumentQueryAskResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ChunkSearch

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ChunkSearchFindResponse">ChunkSearchFindResponse</a>

Methods:

- <code title="post /v1/chunk_search">client.ChunkSearch.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ChunkSearchService.Find">Find</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ChunkSearchFindParams">ChunkSearchFindParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ChunkSearchFindResponse">ChunkSearchFindResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SummarizePage

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SummarizePageNewResponse">SummarizePageNewResponse</a>

Methods:

- <code title="post /v1/summarize_page">client.SummarizePage.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SummarizePageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SummarizePageNewParams">SummarizePageNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SummarizePageNewResponse">SummarizePageNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# StorageObject

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#StorageObjectDeleteResponse">StorageObjectDeleteResponse</a>

Methods:

- <code title="delete /v1/object/{bucket}/{key}">client.StorageObject.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#StorageObjectService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, key <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#StorageObjectDeleteParams">StorageObjectDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#StorageObjectDeleteResponse">StorageObjectDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
