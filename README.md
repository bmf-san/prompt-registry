# prompt-registry

AIエージェント・プロンプトの分類管理リポジトリ。
ナレッジをプロンプトに抽象化し、評価・改善を繰り返すための起点。

## ディレクトリ構成

```
prompt-registry/
├── personas/     # ペルソナ（type: persona） — 役割定義・対話フロー
├── skills/       # スキル（type: skill）     — 再利用可能な知識・技術
├── reviews/      # レビュー（type: review）  — コードレビュー観点集
├── artifacts/    # 成果物（type: artifact）  — ドキュメントテンプレート
├── docs/         # 設計ドキュメント・作成ガイド
└── scripts/
    └── validate/ # フロントマター検証スクリプト（Go）
```

各ディレクトリ直下にファイルをフラットに配置する（サブディレクトリ不可）。

### 各ディレクトリの役割

| ディレクトリ | type | 説明 |
|------------|------|------|
| `personas/` | `persona` | エージェントの役割・ミッション・対話フローを定義する |
| `skills/` | `skill` | エージェントが参照する再利用可能な技術知識 |
| `reviews/` | `review` | コードレビュー・設計レビューのチェックリスト |
| `artifacts/` | `artifact` | ADRや設計書などのドキュメントテンプレート |

## フロントマター仕様

`personas/` `skills/` `reviews/` `artifacts/` 配下の全 `.md` ファイルには以下の YAML フロントマターが必須。

```yaml
---
id: {ファイル名（拡張子なし）と一致させる}
type: persona | skill | review | artifact
domain: {config.yaml の domains に定義されたいずれかの値}
sources: []     # 任意（参考URL）
---
```

追加できる `domain` の値は `config.yaml` の `domains` リストで管理する。

### バリデーションルール

| ルール | 内容 |
|--------|------|
| フロントマター存在 | YAML フロントマターが存在すること |
| 必須フィールド | `id` / `type` / `domain` が存在すること |
| `type` 値 | `persona` / `skill` / `review` / `artifact` のいずれかであること |
| `domain` 値 | `config.yaml` の `domains` に定義された値であること |
| `id` 一致 | `id` がファイル名（拡張子なし）と一致すること |
| ディレクトリ対応 | ディレクトリと `type` が対応すること（`personas/` → `persona` など） |

## バリデーション（CI）

```bash
# ローカル実行
go run ./scripts/validate/ .
```

プッシュ・プルリクエスト時に GitHub Actions で自動チェックされる。

## ドキュメント

- [DESIGN_DOC.md](docs/DESIGN_DOC.md) — リポジトリ構造の設計背景・方針・決定事項
- [WRITING_GUIDE.md](docs/WRITING_GUIDE.md) — プロンプト作成規約・コンポーネント・フレームワーク
