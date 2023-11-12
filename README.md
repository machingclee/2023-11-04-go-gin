# 2023-11-04-go-gin

More in depth study of doing transactin in raw sql, and bring that sql into go.

- Generate db mock:
  ```shell
  go install github.com/golang/mock/mockgen@v1.6.0
  go get github.com/golang/mock/mockgen@v1.6.0
  mockgen -destination internal/db/mock/store.go github.com/machingclee/2023-11-04-go-gin/internal/db Store
  ```
