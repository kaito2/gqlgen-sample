# gqlgen Sample

[gqlgen](https://gqlgen.com/) の素振り

## Run

```
go run cmd/main.go
```

## MEMO

- 明示的に Resolver を実装するためにはアノテーションを付与する
  - e.g. https://github.com/kaito2/gqlgen-sample/blob/main/graph/schema.graphqls#L13
- Resolver が別途定義されている model は `nil` を返しておけばよしなに解決してくれる
  - e.g. https://github.com/kaito2/gqlgen-sample/blob/main/graph/schema.resolvers.go#L27

## N+1 問題

該当リリース: https://github.com/kaito2/gqlgen-sample/releases/tag/v0.0.2

### 再現

```
$ go run cmd/main.go
...
```

`localhost:8080` を開いて以下の Query を実行

```graphql
query Users {
  todos {
    user {
      name
    }
  }
}
```

実行結果

```json
{
  "data": {
    "todos": [
      {
        "user": {
          "name": "Name of a"
        }
      },
      {
        "user": {
          "name": "Name of a"
        }
      },
      {
        "user": {
          "name": "Name of b"
        }
      }
    ]
  }
}
```

サーバーログ

```
...
2022/08/28 19:11:55 GetTodos is called.
2022/08/28 19:11:55 GetUser is called (id: b)
2022/08/28 19:11:55 GetUser is called (id: a)
2022/08/28 19:11:55 GetUser is called (id: a)
```

すべての User 解決のたびに DB Query が発行されているのがわかる。

### 改善

該当リリース: https://github.com/kaito2/gqlgen-sample/releases/tag/v0.0.3

Source: https://github.com/99designs/gqlgen/blob/master/docs/content/reference/dataloaders.md

```
go run cmd/main.go
...
```

`localhost:8080` を開いて以下の Query を実行

```graphql
query Users {
  todos {
    user {
      name
    }
  }
}
```

実行結果

```json
{
  "data": {
    "todos": [
      {
        "user": {
          "name": "Name of a"
        }
      },
      {
        "user": {
          "name": "Name of a"
        }
      },
      {
        "user": {
          "name": "Name of b"
        }
      }
    ]
  }
}
```

サーバーログ

```
...
2022/08/28 19:56:57 GetTodos is called.
2022/08/28 19:56:57 FindUsersByIDs is called (ids: [b a])
```

3回実行されていた `GetUser` アクセスが、1回の `FindUsersByIDs` に置き換えられた。
