# Design Doc: prompt-registry 構造設計

## 背景

プロンプトを構造なく積み上げると、以下の課題が生じる：

- **Fat Agent 問題**: エージェントがペルソナ定義・知識・フォーマット・チェックリストをすべて内包する設計では、定義が肥大化しやすい
- **知識の重複**: 複数のエージェントが同じ知識を個別に持つと、DDD・API 設計などの記述が重複して管理コストが増大する
- **種別の混在**: ペルソナ定義・スキル知識・レビュー観点・成果物テンプレートが単一ディレクトリに混在すると、役割の区別が曖昧になる
- **管理ルールの不在**: どのディレクトリに何を置くかのルールがないと、ファイル数が増えるにつれてカオス化する

## 目的

ナレッジを抽象化→プロンプト化→評価→改善というサイクルを持続的に回すために、まずプロンプトを「整理・管理できる状態」にする。

## 設計方針

### プロンプトの4分類

すべてのプロンプトを以下の4種類に分類する。分類の軸は「AIは何の役割を果たすか」。

| type | ディレクトリ | AIの役割 | 問い |
|------|------------|---------|-----|
| `persona` | `personas/` | "誰"として振る舞うか | このAIはどんな専門家か |
| `skill` | `skills/` | "何を"知っているか | この知識をどう使うか |
| `review` | `reviews/` | "何を基準に"評価するか | 何を確認すべきか |
| `artifact` | `artifacts/` | "何を"生成するか | どんな成果物を作るか |

### personas/ の構成方針

ペルソナはペルソナ定義に専念させ、知識はスキルとして外部化する。

ペルソナ構成：
- **ペルソナ定義**: 役割・ミッション・対話フロー・対応外タスクの宣言
- **使用するスキルの参照**: `## 使用するスキル` セクションで `skills/` へのリンクを列挙

スキルはペルソナ間で再利用・独立テスト・個別更新が可能になる。

### フロントマター仕様

`personas/`・`skills/`・`reviews/`・`artifacts/` 配下の全 `.md` ファイルに必須。

```yaml
---
id: <ファイル名（拡張子なし・ケバブケース）>   # ファイル名と一致させる
type: <persona | skill | review | artifact>
domain: <domain>                             # 許可値は config.yaml の domains を参照
sources: []                                  # 任意。ナレッジの出所（URL や相対パス）
---
```

## ディレクトリ構成

```
prompt-registry/
├── AGENTS.md
├── README.md
├── config.yaml      # 許可 domain 値の定義（バリデーターが参照）
├── docs/
│   ├── DESIGN_DOC.md
│   └── WRITING_GUIDE.md
├── scripts/
│   └── validate/
│       └── main.go
├── go.mod
├── go.sum
├── personas/        # ペルソナ定義
├── skills/          # 知識モジュール
│   ├── non-functional-requirements.md
│   └── prompt-engineering.md
├── reviews/         # レビュー観点
│   └── db-migration.md
├── artifacts/       # 成果物生成テンプレート
│   └── adr.md
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
| フロントマター存在 | `---` ブロックが存在しない場合はエラー |
| 必須フィールド | `id`, `type`, `domain` が存在すること |
| type 値 | `persona/skill/review/artifact` 以外は拒否 |
| domain 値 | `config.yaml` の `domains` に定義された値以外は拒否 |
| id 一致 | `id` 値がファイル名（拡張子なし）と一致すること |
| ディレクトリ↔type 対応 | `personas/` は `persona`、`skills/` は `skill`、`reviews/` は `review`、`artifacts/` は `artifact` のみ許可 |

エラー時はファイルパスと違反内容を標準出力し `os.Exit(1)` でブロック。全ファイルを走査してからまとめてエラーを出力する（1ファイルで止めない）。

## 決定事項

- `prompt-registry-platform` との連携・考慮は行わない（独立管理）
- `AGENTS.md`（ルート）は `personas/` とは無関係のリポジトリ案内ファイルであり、バリデーション対象外
- CI は Go スクリプト + GitHub Actions（外部依存最小）
- `skills/`・`reviews/` はフラット構造とし、domain サブディレクトリは持たない。domain の管理はフロントマターで行う（ファイル数が少ない現状では二重管理になるため）
- 許可 domain 値は `config.yaml` の `domains:` に一元定義し、バリデーターが動的に参照する（ハードコード禁止）
