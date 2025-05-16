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

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchRunResponse">SearchRunResponse</a>

Methods:

- <code title="post /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchRunParams">SearchRunParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#SearchRunResponse">SearchRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Object

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#BucketResponse">BucketResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectListObjectsResponse">ObjectListObjectsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectPutObjectResponse">ObjectPutObjectResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectGetObjectResponse">ObjectGetObjectResponse</a>

Methods:

- <code title="get /v1/object/{bucket_name}">client.Object.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectService.ListObjects">ListObjects</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bucketName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectListObjectsParams">ObjectListObjectsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectListObjectsResponse">ObjectListObjectsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/object/{bucket_name}/{object_key}">client.Object.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectService.PutObject">PutObject</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectPutObjectParams">ObjectPutObjectParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectPutObjectResponse">ObjectPutObjectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/object/{bucket_name}/{object_key}">client.Object.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectService.GetObject">GetObject</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectGetObjectParams">ObjectGetObjectParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go">raindrop</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/raindrop-go#ObjectGetObjectResponse">ObjectGetObjectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
