# Design Doc: prompt-registry 構造設計

## 背景と目的

ナレッジを抽象化→プロンプト化→評価→改善というサイクルを持続的に回すために、まずプロンプトを「整理・管理できる状態」にする。

現状は `agents/` に8つのペルソナ定義が存在するが、以下の課題がある：

- **Fat Agent 問題**: 各エージェントがペルソナ定義・知識・フォーマット・チェックリストをすべて内包しており肥大化している
- **知識の重複**: `app-architect.md` と `system-architect.md` が DDD・API 設計などの知識を重複して持っている
- **種別の混在**: ペルソナ定義・スキル知識・レビュー観点・成果物テンプレートが単一の `agents/` ディレクトリに混在している
- **管理ルールの不在**: どのディレクトリに何を置くかのルールが存在せず、成長とともにカオスになることが想定される

## 設計方針

### プロンプトの4分類

すべてのプロンプトを以下の4種類に分類する。分類の軸は「AIは何の役割を果たすか」。

| type | ディレクトリ | AIの役割 | 問い |
|------|------------|---------|-----|
| `persona` | `agents/` | "誰"として振る舞うか | このAIはどんな専門家か |
| `skill` | `skills/<domain>/` | "何を"知っているか | この知識をどう使うか |
| `review` | `reviews/` | "何を基準に"評価するか | 何を確認すべきか |
| `artifact` | `artifacts/` | "何を"生成するか | どんな成果物を作るか |

`templates/` はユーザーが記入する依頼フォームであり、AIへの指示ではないため上記分類の対象外とする。

### agents/ 分解の方針

Fat Agent のまま `skills/` を追加すると、重複した記述が agents と skills の両方に生まれて管理コストが増大する。そのため、**今回からエージェントをスリム化しながら skills/ を作成する**。

分解後のエージェント構成：
- **ペルソナ定義**: 役割・ミッション・対話フロー・対応外タスクの宣言
- **使用するスキルの参照**: `## 使用するスキル` セクションで `skills/` へのリンクを列挙

スキルはエージェント間で再利用・独立テスト・個別更新が可能になる。

### フロントマター仕様

`agents/`・`skills/`・`reviews/`・`artifacts/` 配下の全 `.md` ファイルに必須。

```yaml
---
id: <ファイル名（拡張子なし・ケバブケース）>   # ファイル名と一致させる
type: <persona | skill | review | artifact>
domain: <所属ドメイン（自由文字列）>
tags: []
sources: []                                  # ナレッジの出所（URL や相対パス）
---
```

`templates/` はユーザー記入フォームのためフロントマター不要・バリデーション対象外。

## ディレクトリ構成（Before/After）

### Before

```
prompt-registry/
├── AGENTS.md
├── README.md
├── agents/          # 8ファイル（Fat Agent——スキル・フォーマット等を内包）
└── templates/       # 8ファイル（依頼フォーム）
```

### After

```
prompt-registry/
├── AGENTS.md
├── README.md
├── docs/
│   └── design-doc.md
├── scripts/
│   └── validate/
│       └── main.go
├── go.mod
├── go.sum
├── agents/          # ペルソナ定義のみ（スリム化済み）
├── skills/          # 知識モジュール（ドメイン別サブディレクトリ）
│   └── architecture/
│       ├── ddd.md
│       ├── solid.md
│       ├── design-patterns.md
│       ├── api-design.md
│       ├── iso25010.md
│       └── tradeoff-analysis.md
├── reviews/         # レビュー観点
│   └── db-migration.md
├── artifacts/       # 成果物生成テンプレート
│   └── adr.md
├── templates/       # ユーザー記入フォーム（変更なし）
└── .github/
    └── workflows/
        └── validate.yml
```

## CI バリデーション設計

### 採用技術: Go + GitHub Actions

- 外部依存は `gopkg.in/yaml.v3` のみ（フロントマターの YAML パース用）
- `go run ./scripts/validate/` で単体実行可能
- GitHub Actions は `push` / `pull_request` で自動実行

### バリデーションルール

| ルール | 内容 |
|--------|------|
| ディレクトリ↔type 対応 | `agents/` は `persona`、`skills/` は `skill`、`reviews/` は `review`、`artifacts/` は `artifact` のみ許可 |
| 必須フィールド | `id`, `type`, `domain` が存在すること |
| id 一致 | `id` 値がファイル名（拡張子なし）と一致すること |
| type 値 | `persona/skill/review/artifact` 以外は拒否 |
| フロントマター存在 | `---` ブロックが存在しない場合はエラー |

エラー時はファイルパスと違反内容を標準出力し `os.Exit(1)` でブロック。全ファイルを走査してからまとめてエラーを出力する（1ファイルで止めない）。

## 今後のロードマップ

### Phase 1: 既存リソースからの抽出（今回）

- `agents/` から知識を `skills/architecture/` に切り出し
- 各リポジトリの `copilot-instructions.md` からレビュー観点を `reviews/` に変換
- `artifacts/adr.md` で ADR 生成テンプレートを整備

### Phase 2: techblog からの抽出

- `techblog/entries/` の技術記事を走査
- 判断基準・評価基準が含まれる記事を `skills/` に変換
- `related_agents` や `tags` でエージェントとのマッピングを整備

### Phase 3: 評価サイクルの導入

- Golden Dataset の整備（`tests/` ディレクトリ）
- LLM-as-Judge による自動評価
- スキル単体での動作検証

## 決定事項

- `prompt-registry-platform` との連携・考慮は行わない（独立管理）
- agents/ 分解は今回から着手する（後回しにしない）
- CI は Go スクリプト + GitHub Actions（外部依存最小）
- `templates/` はバリデーション対象外
