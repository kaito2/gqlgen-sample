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
