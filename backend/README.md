# backend

```bash
curl -H "Content-Type: application/json" -H "X-API-KEY: local" --data '{"url": "https://example.com"}' http://localhost:8080/article.v1.ArticleService/Share
```

```bash
curl -H "Content-Type: application/json" --data '{"max_page_size": 1, "page_token": ""}' http://localhost:8080/article.v1.ArticleService/List
```

```bash
curl -H "Content-Type: application/json" --data '{"id": ""}' http://localhost:8080/article.v1.ArticleService/Delete
```
