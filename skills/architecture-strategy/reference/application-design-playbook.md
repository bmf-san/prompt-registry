# アプリケーション設計プレイブック

アプリケーション層（クラス設計・モジュール構成・デザインパターン・ドメイン設計・API 設計）を、保守性・拡張性高く設計するための手順とパターン。設計方針を合意してから詳細化する。

## 1. 設計準備

- **機能要件**: 実装する機能、ユースケース、ビジネスルール
- **既存コードの確認**: 既存のアーキテクチャパターン（MVC・レイヤード・クリーン等）、コーディング規約、言語・FW
- **非機能要件**: パフォーマンス、テスタビリティの重要度、将来的な拡張予定
- **設計範囲**: 新規開発か既存改修か、影響範囲、変更できない部分・守るべきパターン

## 2. アーキテクチャパターンの選択

実装の前に必ずパターンを提案し、合意を得る。

| パターン | 適用場面 | メリット | デメリット |
|---|---|---|---|
| **レイヤード**（Presentation→Business→Data Access） | 中小規模、チームが慣れている | シンプル・理解しやすい | 層が増えると複雑化、ビジネスロジックが Data Access に依存しがち |
| **クリーン / ヘキサゴナル**（Domain←Application←Infra/Presentation） | ビジネスロジックが複雑、長期保守重視 | ドメイン層の独立性、テスタビリティ高い | 学習コスト・初期コスト高い |
| **CQRS**（Read と Write を分離） | 読み取りと書き込みの要件が大きく異なる | それぞれ最適化可能 | 複雑性増加 |

**提案フォーマット**: 提案するパターン／選択理由（要件適合・スキル・拡張性）／トレードオフ（利点・欠点と軽減策）／代替案。

## 3. ドメイン設計（DDD 適用時）

ドメインモデル（エンティティ・値オブジェクト・集約）、ユビキタス言語、境界づけられたコンテキストを整理する。

```
【エンティティ】User
  識別子: userId / 属性: name, email / 振る舞い: register(), updateProfile()
【値オブジェクト】Email
  属性: value / 制約: メールアドレス形式の検証
【集約】Order 集約
  集約ルート: Order / エンティティ: OrderItem / 値オブジェクト: Money, Quantity
  ビジネスルール: 注文は最低1つの注文アイテムを持つ
```

## 4. クラス設計

SOLID 原則（SRP・OCP・LSP・ISP・DIP）を遵守し、テスタビリティと依存性注入を意識する。

```
【クラス名】UserService
【責務】ユーザーのビジネスロジックを提供する
【依存】UserRepository（IF）／ EmailService（IF）
【主要メソッド】registerUser(name, email): User（メール重複チェック・ウェルカムメール送信）
【設計判断】Repository/Service をインターフェース化し、DI での差し替え・テスト時のモック化を可能にする
```

## 5. デザインパターンの適用

「なぜこのパターンか」を説明できること。**過剰適用を避ける**。

例: 支払い処理（クレジットカード・銀行振込・PayPal 等）に Strategy Pattern を適用する。

```
interface PaymentStrategy { execute(amount: Money): PaymentResult }
class CreditCardPayment implements PaymentStrategy { ... }
class BankTransferPayment implements PaymentStrategy { ... }
class PaymentProcessor {
  constructor(private strategy: PaymentStrategy) {}
  process(amount: Money): PaymentResult { return this.strategy.execute(amount) }
}
```

メリット: 支払い方法の追加が容易、各ロジックが独立、テストが容易。適用理由: 今後支払い方法が増える予定があり、拡張性を確保する。

## 6. モジュール・パッケージ構成

コードの物理配置を設計する。代表的な構成:

- **技術的な分割（レイヤー別）**: `controllers/` `services/` `repositories/` `models/`
- **機能的な分割（ドメイン別）**: `user/`（User.ts, UserService.ts, UserRepository.ts, UserController.ts）, `order/` ...
- **クリーンアーキテクチャ風**: `domain/`（entities, value-objects, repositories<IF>）, `application/`（use-cases）, `infrastructure/`（repositories 実装）, `presentation/`（controllers）

機能的な分割は、機能ごとにコードがまとまり理解しやすく、機能追加時の影響範囲が明確で、チームで分担しやすい。

## 7. API 設計

### RESTful API

エンドポイント・説明・リクエストパラメータ（Path/Query）・レスポンス（成功・エラーの統一フォーマット）・設計判断（ページネーション方式など）を定義する。

```
GET /api/users/{userId}/orders
Query: status?（フィルタ）, limit?（default 20）, offset?（default 0）
200: { "orders": [...], "pagination": { "total", "limit", "offset" } }
404: { "error": { "code": "USER_NOT_FOUND", "message": "User not found" } }
設計判断: ページネーションは limit/offset（カーソルベースは将来検討）、エラーは統一フォーマット
```

### GraphQL API（必要な場合）

```graphql
type User { id: ID!, name: String!, email: String!, orders: [Order!]! }
type Order { id: ID!, status: OrderStatus!, totalAmount: Float!, items: [OrderItem!]! }
type Query { user(id: ID!): User, users(limit: Int, offset: Int): UserConnection }
type Mutation { createUser(input: CreateUserInput!): User!, updateUser(id: ID!, input: UpdateUserInput!): User! }
```

## 8. 設計レビュー・改善

### 評価観点

SOLID 原則の遵守、適切な抽象化レベル、責務の分離、テスタビリティ、拡張性、可読性。

### コードスメルの指摘

- **クラス・メソッドレベル**: Long Method（メソッドが長すぎる）、Large Class（責務が多すぎる）、Long Parameter List、Duplicate Code
- **設計レベル**: Divergent Change（1クラスが複数の理由で変更される）、Shotgun Surgery（1変更が複数クラスに波及）、Feature Envy（他クラスのデータに過度に依存）、Inappropriate Intimacy（クラス間の結合が強すぎる）

### 改善提案の提示方法

```
【現在の設計】問題のある設計を説明
【問題点】違反している原則・コードスメル・テスト困難性
【改善案】改善後の設計
【改善理由】どのパターンで何が改善されるか
【リファクタリング手順】1. 抽出 2. インターフェース化 3. 注入
【影響範囲】この変更で修正が必要になる箇所
```

レビュー観点のチェックリストは [system-design-review.md](system-design-review.md) も併用する。
