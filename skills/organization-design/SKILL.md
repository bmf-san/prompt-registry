---
name: organization-design
description: ソフトウェア開発の組織・チーム設計を支援するスキル。チーム構造の設計、チームの価値観（MVV）定義、共通基盤づくり、振り返りの改善をするときに使う。組織設計, チームトポロジー, MVV, プラットフォームエンジニアリング, レトロスペクティブ, 振り返り, チームづくり, 認知負荷, 逆コンウェイ で発火する。
---

# 組織・チーム設計

ソフトウェア開発の組織とチームを、価値のフローと認知負荷を軸に設計する。チームが何者かを定義し（MVV）、価値フローに沿って構造化し（チームトポロジー）、共通基盤で自律を支え（プラットフォーム・エンジニアリング）、振り返りで継続改善する（bmf）。技術的正解は環境で変わるため、拠り所となる設計原則を持つことが目的。

## いつ使うか

- 新しいチームを立ち上げる、既存チームの境界・責務を見直すとき
- チームの判断軸が揃わず議論が空転している、オーナーシップが薄いと感じるとき
- 共通基盤（プラットフォーム）を製品として投資すべきか判断したいとき
- 振り返りが形骸化し、実行に繋がっていないと感じるとき

## 4つの視点

| 視点 | 問い | 参照 |
|-----|-----|-----|
| **MVV** | このチームは誰に何を提供する専門集団か。何を正しいと信じるか | reference/team-mvv.md |
| **チームトポロジー** | 価値フローに沿った構造か。認知負荷は適切か | reference/team-topologies.md |
| **プラットフォーム・エンジニアリング** | 共通基盤を製品として投資する価値があるか | reference/platform-engineering.md |
| **振り返り（bmf）** | 成果を言語化し、次の集中対象を絞れているか | reference/retrospective-bmf.md |

共通する思想は「フローと認知負荷」。組織構造とアーキテクチャはセット（コンウェイの法則）で考え、認知負荷を制約条件として扱う。一度決めて終わりにせず、定期的に問い直す（組織センシング）。

## 参照

- [reference/team-mvv.md](reference/team-mvv.md) — チームの Mission / Vision / Values の定義と運用
- [reference/team-topologies.md](reference/team-topologies.md) — 4チームタイプ・3インタラクションモード・逆コンウェイ作戦
- [reference/platform-engineering.md](reference/platform-engineering.md) — 共通基盤を製品として提供する専門分野、IDP・ゴールデンパス
- [reference/retrospective-bmf.md](reference/retrospective-bmf.md) — Build / Miss / Focus による選択と集中の振り返り

## 一次ソース

- https://bmf-tech.com/posts/ソフトウェア開発チームがMVVを定めるべき理由
- https://bmf-tech.com/posts/チームトポロジーとは何か
- https://teamtopologies.com/
- https://bmf-tech.com/posts/プラットフォーム・エンジニアリングとは何か
- https://www.oreilly.co.jp/books/9784814401413/
- https://platformengineering.org/blog/what-is-platform-engineering
- https://bmf-tech.com/posts/選択と集中を促す振り返りフレームワーク「bmf」
