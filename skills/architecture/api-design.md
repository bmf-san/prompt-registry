---
id: api-design
type: skill
domain: architecture
tags: [API, REST, GraphQL, HTTP, endpoint-design, status-code]
sources: []
---

# API設計スキル

RESTful API および GraphQL API の設計原則と実践的なフォーマットを定義するスキル。

## 使用するエージェント

- [app-architect](../../agents/app-architect.md)

---

## RESTful API 設計

### リソース設計の原則

- リソースは **名詞** で表現（動詞は使わない）
- 階層構造を適切に使う（`/users/123/orders`）
- **複数形** を使う（`/users` ではなく `/user` は NG）

```
✗ POST /createUser
✗ GET  /getUserList
◯ POST /users
◯ GET  /users
```

### HTTP メソッドの使い分け

| メソッド | 用途 | 冪等性 |
|---------|------|--------|
| GET | リソースの取得 | ◯ |
| POST | リソースの作成 | ✗ |
| PUT | リソースの完全更新 | ◯ |
| PATCH | リソースの部分更新 | △ |
| DELETE | リソースの削除 | ◯ |

### ステータスコードの選択

**2xx 成功:**

| コード | 説明 | 使用場面 |
|--------|------|---------|
| 200 OK | 成功 | GET, PUT, PATCH |
| 201 Created | 作成成功 | POST |
| 204 No Content | 成功（レスポンスボディなし） | DELETE |

**4xx クライアントエラー:**

| コード | 説明 | 使用場面 |
|--------|------|---------|
| 400 Bad Request | リクエストが不正 | 必須パラメータ欠落など |
| 401 Unauthorized | 認証が必要 | 未ログイン |
| 403 Forbidden | アクセス権限なし | 認証済みだが権限なし |
| 404 Not Found | リソースが存在しない | 存在しない ID |
| 422 Unprocessable Entity | バリデーションエラー | 入力値の制約違反 |
| 429 Too Many Requests | レート制限超過 | API 呼び出し回数制限 |

**5xx サーバーエラー:**

| コード | 説明 |
|--------|------|
| 500 Internal Server Error | サーバー内部エラー |
| 503 Service Unavailable | サービス一時停止 |

---

## API エンドポイント設計フォーマット

```
【エンドポイント】
GET /api/users/{userId}/orders

【説明】
指定されたユーザーの注文一覧を取得する

【リクエストパラメータ】
- Path: userId (number, required)
- Query:
  - status (string, optional): 注文ステータスでフィルタ（pending|completed|cancelled）
  - limit (number, optional, default: 20, max: 100)
  - offset (number, optional, default: 0)

【レスポンス】
200 OK:
{
  "orders": [
    {
      "orderId": 123,
      "status": "completed",
      "totalAmount": 1000,
      "createdAt": "2025-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "total": 100,
    "limit": 20,
    "offset": 0
  }
}

404 Not Found:
{
  "error": {
    "code": "USER_NOT_FOUND",
    "message": "User not found"
  }
}

【設計判断】
- <採用した設計と理由、トレードオフ>
```

---

## GraphQL API 設計

### 適用場面

- クライアントが必要なフィールドを柔軟に選択したい
- 複数リソースを1リクエストで取得したい（N+1 の回避）
- BFF（Backend for Frontend）として複数クライアントを統一したい

### スキーマ設計の基本

```graphql
type User {
  id: ID!
  name: String!
  email: String!
  orders: [Order!]!
}

type Order {
  id: ID!
  status: OrderStatus!
  totalAmount: Float!
  items: [OrderItem!]!
}

enum OrderStatus {
  PENDING
  COMPLETED
  CANCELLED
}

type Query {
  user(id: ID!): User
  users(limit: Int, offset: Int): UserConnection
}

type Mutation {
  createUser(input: CreateUserInput!): User!
  updateUser(id: ID!, input: UpdateUserInput!): User!
}
```

### REST vs GraphQL の選択基準

| 観点 | REST | GraphQL |
|------|------|---------|
| シンプルさ | ◎ | △ |
| 柔軟なクエリ | △ | ◎ |
| キャッシュ | ◎（HTTP キャッシュ） | △（クライアント側で工夫が必要） |
| 学習コスト | 低 | 中〜高 |
| 適用場面 | CRUD 中心のシンプルな API | 複雑なデータ取得ニーズ、BFF |

---

## チェックリスト

- [ ] リソースは名詞で表現されているか
- [ ] HTTP メソッドを適切に使い分けているか
- [ ] ステータスコードは適切か
- [ ] エラーレスポンスのフォーマットは統一されているか
- [ ] ページネーション設計はあるか（一覧取得 API）
- [ ] 認証・認可の方式は明記されているか
- [ ] バージョニング戦略（`/v1/`など）は検討済みか
