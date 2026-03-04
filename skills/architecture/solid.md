---
id: solid
type: skill
domain: architecture
tags: [SOLID, SRP, OCP, LSP, ISP, DIP, object-oriented, class-design]
sources: []
---

# SOLID 原則スキル

オブジェクト指向設計の5原則（SOLID）の定義と適用方法を定義するスキル。

## 使用するエージェント

- [app-architect](../../agents/app-architect.md)

---

## 5原則の定義

### S — 単一責任の原則（Single Responsibility Principle）

> 1つのクラスは1つの責務のみを持つ。変更理由が1つに限定される。

**判断方法:** 「このクラスはなぜ変更されるか？」の答えが複数あれば SRP 違反。

```
✗ UserService が「ユーザー登録」と「メール送信」と「ログ出力」を担当
◯ UserService はユーザーのビジネスロジックのみ
   EmailService はメール送信のみ
   Logger はログ出力のみ
```

### O — 開放閉鎖の原則（Open/Closed Principle）

> 拡張に対して開いている。修正に対して閉じている。

**適用方法:** 条件分岐で機能を追加するのではなく、インターフェースや抽象クラスで拡張する。

```
✗ if payment == "credit_card" ... elif payment == "paypal" ...
◯ interface PaymentStrategy { execute() }
   class CreditCardPayment implements PaymentStrategy { ... }
   class PayPalPayment implements PaymentStrategy { ... }
```

### L — リスコフの置換原則（Liskov Substitution Principle）

> 基底クラスは派生クラスで置換可能。派生クラスは基底クラスの契約（事前条件・事後条件）を破らない。

**違反の典型例:** サブクラスで例外を投げて親クラスの振る舞いを壊す。

### I — インターフェース分離の原則（Interface Segregation Principle）

> クライアントは使わないメソッドへの依存を強制されない。

**適用方法:** 巨大なインターフェースより、用途別の小さなインターフェースに分割する。

```
✗ interface IStorage { read(); write(); delete(); getMetadata(); compress(); }
◯ interface IReader { read(); }
   interface IWriter { write(); delete(); }
```

### D — 依存性逆転の原則（Dependency Inversion Principle）

> 上位モジュールは下位モジュールに依存しない。両者は抽象（インターフェース）に依存する。

**適用方法:** 具体的な実装クラスではなく、インターフェースに依存させる。DI（Dependency Injection）と組み合わせる。

```
✗ UserService が直接 MySQLUserRepository を new する
◯ UserService は UserRepository インターフェースに依存
   MySQLUserRepository が UserRepository を実装
   DI コンテナで実態を注入
```

---

## クラス設計のフォーマット

```
【クラス名】
<ClassName>

【責務】
<このクラスが担う唯一の責任>

【依存】
- <InterfaceName>（インターフェース）
- <InterfaceName>（インターフェース）

【主要メソッド】
- <methodName>(<params>): <return>
  - <処理の説明>
  - <ビジネスルール>

【設計判断】
- <なぜこの設計を選んだか、トレードオフは何か>
```

---

## コードスメル（設計レベル）

SOLID 違反を示唆するコードスメル:

| コードスメル | 関連原則 | 症状 |
|------------|---------|------|
| God Object（神クラス） | SRP | クラスが巨大で多くの責務を持つ |
| Divergent Change | SRP | 1クラスが複数の理由で変更される |
| Shotgun Surgery | OCP | 1つの変更が多くのクラスに波及 |
| Feature Envy | DIP | 他クラスのデータに過度に依存 |
| Large Interface | ISP | 使われないメソッドを持つ巨大インターフェース |
| Concrete Dependency | DIP | 具体クラスへの直接依存 |

---

## チェックリスト

- [ ] 各クラスの変更理由は1つか（SRP）
- [ ] 機能追加時に既存クラスを修正せず拡張できるか（OCP）
- [ ] サブクラスは親クラスの代わりに使えるか（LSP）
- [ ] インターフェースは必要最小限のメソッドに絞られているか（ISP）
- [ ] 具体クラスではなくインターフェースに依存しているか（DIP）
- [ ] DI コンテナや Constructor Injection を活用しているか（DIP）
