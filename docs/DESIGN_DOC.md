# Design Doc: prompt-registry 構造設計

## この文書について

prompt-registry の構造設計と、その決定の背景を記録する。2026-09 に独自4分類（v1）から Agent Skills 標準（v2）へ全面移行した。以下は v2 の設計であり、末尾に v1 からの変更点を残す。

## 目的

ナレッジを抽象化 → プロンプト化 → 評価 → 改善のサイクルを持続的に回すために、プロンプトを「整理・管理でき、標準ツールでそのまま使える状態」にする。

## 採用する標準

[Agent Skills](https://agentskills.io)（Anthropic が策定したオープン標準）を唯一のスキーマとして採用する。理由:

- **1 skill = 1 ディレクトリ + `SKILL.md`** という単純な単位で、役割・知識・チェックリスト・テンプレートをすべて表現できる。
- **プログレッシブ・ディスクロージャ** により、メタ情報 → 本体 → 詳細の3段で必要な分だけ読み込める。コンテキスト消費を抑えられる。
- **ポータビリティ**。Claude Code / claude.ai / Agent SDK など標準対応ツールでそのまま読み込める。独自スキーマの囲い込みを避ける。

## 設計方針

### 分類は skill 1種のみ

v1 の persona / skill / review / artifact は廃止し、すべて skill に統一する。分類軸を「AI のモード」から「能力・トピック」へ移す。

- 役割（旧 persona）→ その役割が行うワークフローを skill として表現する。役割宣言そのものは残さない。
- チェックリスト（旧 review）→ 該当トピック skill の `reference/checklist.md`。
- 成果物テンプレート（旧 artifact）→ 該当トピック skill の `reference/template.md`。

### 1 skill = 1 ディレクトリ

```
skills/
  <skill-name>/
    SKILL.md          # 必須。概要・いつ使うか・判断基準・手順・一次ソース
    reference/        # 任意。詳細・具体値・テンプレート・チェックリスト
```

### トピック単位でまとめる

密接に関連する概念は1 skill に統合し、`reference/` で分割する。1ファイル=1概念の細粒度は採らない。極小 skill の乱立と、常時コンテキストに載るメタ情報の肥大を避けるためである。統合しても、`description` に内包トピックをキーワードとして列挙することで自動発火の精度を保つ。

### skill 名の付け方

skill 名は単一の凝集した概念にする。`A-and-B` のように2概念を連結した名前は避ける。連結したくなったら、統合しすぎか、トピックの切り出し方を誤っているサインとして見直す。

### フロントマター仕様

```yaml
---
name: <ディレクトリ名と一致>
description: <何をする skill か＋いつ使うか。冒頭に主用途、続けてトリガーとなる語を書く>
---
```

- v1 の `id` / `type` / `domain` は廃止する。
- 出典は本文末尾の「一次ソース」節に置く。標準の許可フィールドに `sources` は無いため、フロントマターには入れない。将来機械可読が必要になれば標準の `metadata:` を使う。

### SKILL.md の書き方

- 冒頭に概要と「いつ使うか」。続けて判断基準・手順。末尾に「一次ソース」。
- **500 行以内**を目安とし、詳細・具体値・テンプレート・チェックリストは `reference/*.md` へ逃がす。
- 手順は「やること」を番号付き、「守ること」を箇条書きで書く。

## ディレクトリ構成

```
prompt-registry/
├── AGENTS.md
├── README.md
├── docs/
│   └── DESIGN_DOC.md
├── skills/
│   └── <skill-name>/
│       ├── SKILL.md
│       └── reference/
└── .github/
```

- `personas/` `reviews/` `artifacts/` は廃止する。
- `config.yaml`（domain enum）は廃止する。
- `docs/WRITING_GUIDE.md` のプロンプト・パターンカタログは `skills/prompt-engineering/` に統合し、リポジトリ固有の「skill の書き方」は AGENTS.md / README.md に集約する。

## 現行 → v2 マッピング

移行時のトピック統合案。persona の内容は移行時に各ファイルを読み、該当トピック skill に手順・知識を salvage する。役割宣言そのものは残さない。pm / designer など単独ドメインの persona は内容を確認したうえで配置を最終決定する。

| v2 skill | 統合する現行ファイル |
|---|---|
| `prompt-engineering` | skills/prompt-engineering ＋ docs/WRITING_GUIDE のパターンカタログ |
| `database-reliability` | skills/{transaction, mysql-lock, mysql-transaction-anomaly, acid-and-base, database-index} ＋ reviews/db-migration |
| `web-scalability` | skills/{cap-and-pacelc, cache-strategy, web-app-processing-model, load-testing, non-functional-requirements} |
| `software-design` | skills/cohesion-and-coupling ＋ reviews/app-design |
| `code-review` | reviews/code-review ＋ personas/engineer |
| `adr` | skills/adr-writing ＋ artifacts/adr |
| `design-documentation` | skills/{architecture-document, design-expiry} ＋ reviews/design-review ＋ artifacts/design-doc |
| `requirements-engineering` | skills/requirements-vs-constraints ＋ reviews/requirement-review ＋ artifacts/system-spec |
| `architecture-strategy` | skills/{architecture-strategy, tech-decision-constraints-tradeoffs} ＋ reviews/{architecture-strategy-review, system-design-review} ＋ personas/{app-architect, architecture-strategist, system-architect} |
| `decision-under-uncertainty` | skills/{cynefin-story-points, project-uncertainty} |
| `organization-design` | skills/{team-mvv, team-topologies, platform-engineering, retrospective-bmf} |
| `technical-writing` | reviews/document-review ＋ personas/{writer, editor} |
| `product-management` | personas/pm |
| `product-design` | personas/designer（内容が薄ければ廃止） |

## ツール構成

独自スキーマを廃止し、それを強制していた検証・設定・ビルドを撤去する。Agent Skills 標準は `name` + `description` のみを要求し、検証すべき独自ルールが無くなるため、専用ツールを持たない。

- **削除**: `config.yaml`（domain enum）、`scripts/`（Go バリデータ）、`Makefile`、`go.mod` / `go.sum`、`.github/workflows/validate.yml`。
- **検証ゲートは持たない**: CI での自動検証は行わない。標準準拠は SKILL.md の書式（`name` + `description`）を人手・エージェントで守る。将来必要になれば標準ツール（例: `claude plugin validate`）を都度使う。
- **GitHub テンプレート**: Issue の type / domain ドロップダウン、PR の type チェック・id/type/domain チェック・stale な検証パスを skill 前提に更新する。

## 移行計画

1. 本 DESIGN_DOC の合意。
2. ツール撤去（`config.yaml` / `scripts/` / `Makefile` / `go.mod` / `go.sum` / `.github/workflows/validate.yml` の削除、GitHub テンプレート更新）。docs（README / AGENTS）は skills/ 構築後に改訂する。
3. トピック単位で `skills/<name>/` を作成し、現行ファイルを `SKILL.md` + `reference/` に統合する。大きい主題（database, architecture 系, prompt-engineering）から着手する。
4. 統合済みの旧ファイル（`personas/` `reviews/` `artifacts/` と旧 `skills/*.md`）を削除する。
5. 最終確認。全 skill が標準準拠、バリデータ green。

## 決定事項

- Agent Skills 標準（agentskills.io）を唯一のスキーマとする。
- skill 1種・1ディレクトリ・トピック統合・`name` + `description`。
- persona / review / artifact を廃止する。役割宣言は残さず、ワークフロー・知識を skill に統合する。
- `config.yaml` と Go ツール一式（`scripts/` / `Makefile` / `go.mod` / `go.sum` / CI 検証）を撤去する。検証ゲートは持たず、標準書式は人手・エージェントで守る。
- 既存の設計判断ログ（本 DESIGN_DOC）は残す。既存 skill の内容は移行時に統合し、価値ある知識は失わない。

## v1 からの変更点

v1 は独自4分類（persona / skill / review / artifact）＋独自フロントマター（id / type / domain）で管理していた。運用の結果、次の課題が判明したため v2 へ移行した。

- **主題の分散**: 1つの主題が type 違いで複数ファイルに割れていた。例: ADR が `skills/adr-writing.md` と `artifacts/adr.md` に重複。`reviews/code-review.md` の観点が `personas/engineer.md` と重複。
- **persona の低い費用対効果**: エージェント型ツールでは役割付与の効果が薄く、価値はワークフロー部分にあった。その多くは既存 skill / review と重複していた。
- **独自スキーマの非互換**: id / type / domain はこのリポジトリ専用で、標準ツールでそのまま使えなかった。
- **細粒度によるリスト肥大**: 1ファイル=1プロンプトのため、極小の知識片が乱立していた。
