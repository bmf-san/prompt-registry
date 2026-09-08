# prompt-registry

AI エージェント向けスキルのレジストリ。ナレッジを skill に抽象化し、評価・改善を繰り返すための起点。

[Agent Skills](https://agentskills.io) オープン標準に準拠する。1 skill = 1 ディレクトリ + `SKILL.md`。

## ディレクトリ構成

```
prompt-registry/
├── skills/            # 各 skill を格納
│   └── <name>/
│       ├── SKILL.md   # 必須
│       └── reference/ # 任意
├── docs/
│   └── DESIGN_DOC.md
├── AGENTS.md
└── README.md
```

## skill 一覧

| skill | 概要 |
|---|---|
| [prompt-engineering](skills/prompt-engineering/SKILL.md) | プロンプト設計のフレームワークとパターン |
| [database-reliability](skills/database-reliability/SKILL.md) | トランザクション・ロック・整合性・インデックス・マイグレーション |
| [web-scalability](skills/web-scalability/SKILL.md) | CAP/PACELC・キャッシュ・処理モデル・負荷試験・非機能要件 |
| [software-design](skills/software-design/SKILL.md) | 凝集度・結合度とアプリ設計レビュー |
| [technical-review](skills/technical-review/SKILL.md) | コード・設計文書を CTO・アーキテクト視点でレビューする観点と進め方 |
| [adr](skills/adr/SKILL.md) | アーキテクチャ決定記録の原則とテンプレート |
| [design-documentation](skills/design-documentation/SKILL.md) | 設計ドキュメントの書き方・鮮度・レビュー |
| [requirements-engineering](skills/requirements-engineering/SKILL.md) | 要件と制約の区別・要件レビュー・仕様テンプレート |
| [architecture-strategy](skills/architecture-strategy/SKILL.md) | アーキテクチャ戦略・技術選定・システム設計レビュー |
| [decision-under-uncertainty](skills/decision-under-uncertainty/SKILL.md) | 不確実性の見極めと見積もり |
| [organization-design](skills/organization-design/SKILL.md) | チーム構造・MVV・基盤・振り返り |
| [technical-writing](skills/technical-writing/SKILL.md) | 文章・ドキュメントの執筆とレビュー |
| [product-management](skills/product-management/SKILL.md) | 何を・なぜ作るかの明確化 |
| [product-design](skills/product-design/SKILL.md) | プロダクトの UX/UI 設計 |

## 規約

構造・命名・フロントマターの設計と規約は [docs/DESIGN_DOC.md](docs/DESIGN_DOC.md)。skill を追加・編集するエージェント向けの手順は [AGENTS.md](AGENTS.md)。
