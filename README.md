# engineering-prompts

AI エージェント向けスキルのレジストリ。ナレッジを skill に抽象化し、評価・改善を繰り返すための起点。

[Agent Skills](https://agentskills.io) オープン標準に準拠する。1 skill = 1 ディレクトリ + `SKILL.md`。

## ディレクトリ構成

```
engineering-prompts/
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
| [technical-review](skills/technical-review/SKILL.md) | コード・設計文書を CTO・アーキテクト視点でレビューする観点と進め方 |
| [adr](skills/adr/SKILL.md) | アーキテクチャ決定記録の原則とテンプレート |
| [requirements-engineering](skills/requirements-engineering/SKILL.md) | 要件と制約の区別・要件レビュー・仕様テンプレート |
| [architecture-strategy](skills/architecture-strategy/SKILL.md) | アーキテクチャ戦略・技術選定・戦略レビュー |
| [organization-design](skills/organization-design/SKILL.md) | チームの価値観（MVV）設計 |
| [technical-writing](skills/technical-writing/SKILL.md) | 文章・ドキュメントの執筆とレビュー |

## 規約

構造・命名・フロントマターの設計と規約は [docs/DESIGN_DOC.md](docs/DESIGN_DOC.md)。skill を追加・編集するエージェント向けの手順は [AGENTS.md](AGENTS.md)。
