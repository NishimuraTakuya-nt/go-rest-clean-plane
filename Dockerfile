# ビルドステージ
FROM golang:1.23-alpine AS builder

# 作業ディレクトリの設定
WORKDIR /app

# 依存関係のコピーとダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードのコピー
COPY . .

# アプリケーションのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# 実行ステージ
FROM alpine:3.19

# セキュリティ強化: 非rootユーザーの作成
RUN adduser -D appuser

# タイムゾーンの設定
RUN apk --no-cache add tzdata

# 作業ディレクトリの設定
WORKDIR /app

# ビルドステージから実行可能ファイルをコピー
COPY --from=builder /app/main .

# 非rootユーザーに切り替え
USER appuser

# コンテナ起動時に実行されるコマンド
CMD ["./main"]