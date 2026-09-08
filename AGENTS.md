# AGENTS.md — prompt-registry

AI エージェントがこのリポジトリで作業する際のガイダンスです。

## リポジトリ概要

prompt-registry は、AI エージェント向けの skill を管理するレジストリです。[Agent Skills](https://agentskills.io) オープン標準に準拠します。

## 構造

```
skills/<name>/
  SKILL.md      # 必須。name + description のフロントマター＋本体
  reference/    # 任意。詳細・テンプレート・チェックリスト
docs/DESIGN_DOC.md
```

分類は skill 1種のみです。役割・チェックリスト・成果物テンプレートもすべて skill として表現します。

## SKILL.md のフロントマター

```yaml
---
name: <ディレクトリ名と一致>
description: <何をする skill か＋いつ使うか。冒頭に主用途、続けて発火トリガーとなる語>
---
```

- `name` と `description` の2つだけです。`id` / `type` / `domain` は使いません。
- `description` が自動発火のトリガーになります。冒頭に主用途を置き、トリガー語を含めます。

## skill を追加・編集するとき

- 1 skill = 1 ディレクトリ。`skills/<name>/SKILL.md` を必須とし、詳細・具体値・テンプレート・チェックリストは `reference/*.md` へ逃がします。
- SKILL.md は概要・いつ使うか・判断基準・手順・一次ソースで構成し、500 行以内を目安にします。
- 密接に関連する概念は1 skill に統合します（トピック単位）。1ファイル=1概念の細粒度にはしません。
- skill 名は単一の凝集した概念にします。`A-and-B` のような連結名は避けます。
- 出典はフロントマターではなく本文末尾の「一次ソース」節に書きます。
- 日本語は直書きします。

## 検証

専用のバリデータや CI は持ちません。標準書式（`name` + `description`、`skills/<name>/SKILL.md`）を守ってください。必要なら標準ツール（例: `claude plugin validate skills`）を都度使えます。

設計の背景は [docs/DESIGN_DOC.md](docs/DESIGN_DOC.md) を参照してください。
