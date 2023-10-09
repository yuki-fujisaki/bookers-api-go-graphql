# Goアプリケーション用のDockerイメージをビルド
FROM golang:latest

WORKDIR /app
COPY . .

# 生成したバイナリは /app ディレクトリに配置する
RUN go build -o /app/main ./cmd/todo

EXPOSE 8080

# 実行コマンドのパスを修正
CMD ["/app/main"]
