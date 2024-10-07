# go-rest-clean-plane
base project for go rest api

## このプロジェクトは以下のディレクトリ構造に基づいています：
```
.
├── cmd
│     └── api : アプリケーションのエントリーポイント
├── docs
│     └── swagger
├── internal : プロジェクト固有のパッケージ
│     ├── adapters : 外部システムとのインターフェース
│     │     ├── primary
│     │     │     └── http
│     │     │         ├── handlers
│     │     │         ├── middleware
│     │     │         └── routes
│     │     └── secondary
│     │         ├── aws
│     │         ├── db
│     │         └── graphql
│     ├── core : ビジネスロジック
│     │     ├── domain
│     │     ├── services
│     │     └── usecases
│     ├── errors
│     ├── infrastructure : 技術的な実装（ロギングなど）
│     │     ├── auth
│     │     ├── config
│     │     └── logger
│     └── utils
├── pkg
└── scripts
```
