# prompt-registry

AIエージェント・プロンプトの分類管理リポジトリ。
ナレッジをプロンプトに抽象化し、評価・改善を繰り返すための起点。

## ディレクトリ構成

```
prompt-registry/
├── agents/       # ペルソナ（type: persona） — 役割定義・対話フロー
├── skills/       # スキル（type: skill）     — 再利用可能な知識・技術
│   └── architecture/
├── reviews/      # レビュー（type: review）  — コードレビュー観点集
├── artifacts/    # 成果物（type: artifact）  — ドキュメントテンプレート
├── templates/    # 指示雛形                 — バリデーション対象外
├── docs/         # 設計ドキュメント
└── scripts/
    └── validate/ # フロントマター検証スクリプト（Go）
```

### 各ディレクトリの役割

| ディレクトリ | type | 説明 |
|------------|------|------|
| `agents/` | `persona` | エージェントの役割・ミッション・対話フローを定義する |
| `skills/` | `skill` | エージェントが参照する再利用可能な技術知識 |
| `reviews/` | `review` | コードレビュー・設計レビューのチェックリスト |
| `artifacts/` | `artifact` | ADRや設計書などのドキュメントテンプレート |
| `templates/` | — | 人間向けの指示雛形。バリデーション対象外 |

## フロントマター仕様

`agents/` `skills/` `reviews/` `artifacts/` 配下の全 `.md` ファイルには以下の YAML フロントマターが必須。

```yaml
---
id: {ファイル名（拡張子なし）と一致させる}
type: persona | skill | review | artifact
domain: {ドメイン名（例: architecture, engineering, design）}
tags: []        # 任意
sources: []     # 任意（参考URL）
---
```

### バリデーションルール

- `id` はファイル名（拡張子なし）と一致すること
- `type` はディレクトリと対応すること（`agents/` → `persona`、`skills/` → `skill` など）
- `domain` は必須

## バリデーション（CI）

```bash
# ローカル実行
go run ./scripts/validate/
```

プッシュ・プルリクエスト時に GitHub Actions で自動チェックされる。

## 設計ドキュメント

- [design-doc.md](docs/design-doc.md) — アーキテクチャ設計の背景と方針
