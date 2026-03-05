# AGENTS.md — prompt-registry

AIエージェントがこのリポジトリで作業する際のガイダンスです。

## リポジトリ概要

prompt-registry は、AIエージェントに与えるプロンプトを4種類に分類して管理するレジストリです。

## ディレクトリ構造

```
personas/   # type: persona — AIに与える役割・ペルソナ定義
skills/     # type: skill   — 特定領域の評価基準・判断知識
reviews/    # type: review  — レビュー・承認時のチェックリスト
artifacts/  # type: artifact — 成果物テンプレート（文書・レポートなど）
docs/       # リポジトリドキュメント（WRITING_GUIDE.md など）
scripts/    # バリデーター（Go）
config.yaml # ドメイン定義
```

## ファイルフォーマット

各 Markdown ファイルは先頭に YAML フロントマターを持ちます。

```yaml
---
id: <ファイル名と一致させる>
type: persona | skill | review | artifact
domain: architecture | engineering | writing | design | product
sources: []   # 参考文献 URL のリスト（省略可）
---
```

`id` / `type` / `domain` は必須フィールドです。

## ドメイン一覧

`config.yaml` で管理しています。新しいドメインを追加する場合はここに追記してください。

- `architecture`
- `engineering`
- `writing`
- `design`
- `product`

## バリデーション

ファイル追加・編集後は必ず実行してください。

```bash
go run ./scripts/validate/ .
```

フロントマターの必須フィールド・`type` / `domain` の値を検証します。

## プロンプト作成規約

[docs/WRITING_GUIDE.md](docs/WRITING_GUIDE.md) を参照してください。

特に **personas** は以下のセクション構成（WRITING_GUIDE 型）に統一されています。

```
# Role
## Task
## Input
## Output Format
## Guidelines
## Prohibited Actions
## Example
## Knowledge Base
```
