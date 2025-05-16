# Search

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchFindResponse">SearchFindResponse</a>

Methods:

- <code title="post /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchService.Find">Find</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchFindParams">SearchFindParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#SearchFindResponse">SearchFindResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DocumentQuery

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

# Chat

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChatInteractResponse">ChatInteractResponse</a>

Methods:

- <code title="post /v1/chat">client.Chat.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChatService.Interact">Interact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChatInteractParams">ChatInteractParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ChatInteractResponse">ChatInteractResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Object

Response Types:

- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectListObjectsResponse">ObjectListObjectsResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectPutObjectResponse">ObjectPutObjectResponse</a>
- <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectGetObjectResponse">ObjectGetObjectResponse</a>

Methods:

- <code title="get /v1/object/{bucket_name}">client.Object.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectService.ListObjects">ListObjects</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bucketName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectListObjectsParams">ObjectListObjectsParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectListObjectsResponse">ObjectListObjectsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/object/{bucket_name}/{object_key}">client.Object.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectService.PutObject">PutObject</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectPutObjectParams">ObjectPutObjectParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectPutObjectResponse">ObjectPutObjectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/object/{bucket_name}/{object_key}">client.Object.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectService.GetObject">GetObject</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, objectKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectGetObjectParams">ObjectGetObjectParams</a>) (<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk">raindrop</a>.<a href="https://pkg.go.dev/github.com/LiquidMetal-AI/lm-raindrop-go-sdk#ObjectGetObjectResponse">ObjectGetObjectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
