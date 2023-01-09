# backend

# .vscode

`.vscode/settings.json`

```json
{
    "go.toolsEnvVars": {
        "GOFLAGS": "-tags=e2em,integration"
    }
}
```

# cURL

```bash
curl -H "Content-Type: application/json" -H "X-API-KEY: local" --data '{"url": "https://example.com"}' http://localhost:8080/article.v1.ArticleService/Share
```

```bash
curl -H "Content-Type: application/json" --data '{"id": "", "tag": "Go"}' http://localhost:8080/article.v1.ArticleService/AddTag
```

```bash
curl -H "Content-Type: application/json" --data '{}' http://localhost:8080/article.v1.ArticleService/ListTag
```

```bash
curl -H "Content-Type: application/json" -H "Authorization: " --data '{"max_page_size": 1, "page_token": ""}' http://localhost:8080/article.v1.ArticleService/List
```

```bash
curl -H "Content-Type: application/json" --data '{"id": ""}' http://localhost:8080/article.v1.ArticleService/Delete
```

```bash
curl -H "Content-Type: application/json" --data '{"email": "example@example.com", "login_id": "example", "password": "password"}' http://localhost:8080/auth.v1.AuthService/SignUp
```

```bash
curl -H "Content-Type: application/json" --data '{"login_id": "example", "password": "password"}' http://localhost:8080/auth.v1.AuthService/SignIn
```

```bash
curl -H "Content-Type: application/json" -H "Authorization: " --data '{"refresh_token": ""}' http://localhost:8080/auth.v1.AuthService/Refresh
```
