---
id: design-patterns
type: skill
domain: architecture
tags: [design-patterns, GoF, strategy, observer, factory, decorator, adapter]
sources: []
---

# デザインパターンスキル

GoF（Gang of Four）デザインパターンの分類と適用方法を定義するスキル。

## 使用するエージェント

- [app-architect](../../agents/app-architect.md)

---

## パターン分類

### 生成に関するパターン（Creational）

| パターン | 目的 | 適用場面 |
|---------|------|---------|
| **Factory Method** | オブジェクト生成をサブクラスに委ねる | 生成するクラスを実行時に決定したい |
| **Abstract Factory** | 関連するオブジェクト群を生成するインターフェース | OS やテーマごとに UI 部品群を切り替えたい |
| **Builder** | 複雑なオブジェクトを段階的に構築 | 属性が多く、組み合わせが多様なオブジェクト |
| **Singleton** | インスタンスを1つに制限 | 設定オブジェクト、ログ（注: テスタビリティへの影響に注意） |
| **Prototype** | 既存オブジェクトをコピーして生成 | 生成コストが高いオブジェクトの複製 |

### 構造に関するパターン（Structural）

| パターン | 目的 | 適用場面 |
|---------|------|---------|
| **Adapter** | インターフェースの不一致を解消 | 外部ライブラリを既存インターフェースに合わせたい |
| **Decorator** | 動的に機能を追加 | 継承を使わずに振る舞いを拡張したい |
| **Facade** | 複雑なサブシステムへのシンプルなインターフェース | 複雑なライブラリを簡潔に使いたい |
| **Composite** | 木構造を均一に扱う | ファイルシステム、UI コンポーネントツリー |
| **Proxy** | 別のオブジェクトの代理 | 遅延初期化、アクセス制御、ログ |

### 振る舞いに関するパターン（Behavioral）

| パターン | 目的 | 適用場面 |
|---------|------|---------|
| **Strategy** | アルゴリズムをカプセル化し、交換可能にする | 支払い方法、ソートアルゴリズム、通知手段 |
| **Observer** | イベント通知の仕組み | UI イベント処理、ドメインイベント |
| **Template Method** | アルゴリズムの骨格を定義し、サブクラスで詳細を実装 | レポート生成の共通フロー |
| **Command** | 操作をオブジェクトとしてカプセル化 | Undo/Redo、キューへの積み込み |
| **Chain of Responsibility** | 要求を処理するオブジェクトのチェーン | ミドルウェア、バリデーションパイプライン |
| **State** | 状態ごとに振る舞いを変える | 注文状態管理（未払い/支払済み/発送済み） |

---

## パターン適用の提案フォーマット

```
【適用するパターン】
<パターン名>

【適用箇所】
<どのクラス・処理に適用するか>

【設計】
<インターフェース・クラス定義の概要>

【メリット】
- <利点1>
- <利点2>

【適用理由】
<なぜこのパターンが適切か、具体的な要件との対応>
```

---

## Strategy Pattern の実装例

```typescript
interface PaymentStrategy {
  execute(amount: Money): PaymentResult
}

class CreditCardPayment implements PaymentStrategy {
  execute(amount: Money): PaymentResult { ... }
}

class BankTransferPayment implements PaymentStrategy {
  execute(amount: Money): PaymentResult { ... }
}

class PaymentProcessor {
  constructor(private strategy: PaymentStrategy) {}

  process(amount: Money): PaymentResult {
    return this.strategy.execute(amount)
  }
}
```

**適用理由:** 支払い方法の追加が OCP（開放閉鎖の原則）に沿って可能。各方法が独立してテスト可能。

---

## パターン選択の判断基準

1. **まず問題を明確にする**: パターンありきで適用しない
2. **過剰設計を避ける**: 将来の拡張が明確でない場合はシンプルな実装を優先（YAGNI）
3. **チームの理解度を考慮**: パターンを知らないメンバーが保守できるか
4. **テスタビリティを確認**: パターン適用後にテストが書きやすくなるか

---

## チェックリスト

- [ ] 使用するパターンの意図と適用場面を説明できるか
- [ ] パターン適用の理由が要件と対応しているか
- [ ] 過剰設計になっていないか（YAGNI）
- [ ] テスタビリティが向上しているか
- [ ] チームメンバーが理解・保守できるか
