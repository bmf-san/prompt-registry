---
name: prompt-engineering
description: プロンプトを構造化し、タスクの性質に応じて適切なフレームワークとパターンを選ぶためのスキル。プロンプトを新規に書くとき、精度や出力の安定性を上げたいとき、既存プロンプトを改善するときに使う。発火トリガー: プロンプトエンジニアリング, プロンプト設計, COSTAR, Few-shot, Chain-of-Thought, プロンプトの書き方。
---

# プロンプトエンジニアリング

プロンプト設計のフレームワークとパターンカタログをまとめたスキル。タスクの性質に応じて、プロンプト全体の骨格（フレームワーク）と個々の手法（パターン）を選ぶための参照資料。プロンプトを新規に書くとき、出力精度・安定性を上げたいとき、既存プロンプトを改善するときに使う。

まずこの SKILL.md で全体の骨格とタスク別の選び方を掴み、詳細は `reference/` を参照する。

- プロンプトの構成要素と書き方の指針（コンポーネント／全体指針）→ [reference/components.md](reference/components.md)
- 個々のプロンプティング手法の一覧と選び方 → [reference/patterns.md](reference/patterns.md)

## プロンプト設計フレームワーク

プロンプト全体を構造化するためのテンプレート。個別パターンを組み合わせる際の骨格として使う。

| フレームワーク | 構成要素 | 向いているケース |
|---|---|---|
| **COSTAR** | Context / Objective / Style / Tone / Audience / Response | 汎用。出力の制御性を高めたいとき |
| **CREATE** | Character / Request / Examples / Adjustment / Type / Extras | 専門家ペルソナを使った分析・創造的タスク |
| **KERNEL** | Keep it simple / Easy to verify / Reproducible / Narrow scope / Explicit constraints / Logical structure | トークン削減と精度を両立したいとき |

## タスク別パターンの選び方

まずタスクの性質から使う手法の当たりをつける。各パターンの詳細は [reference/patterns.md](reference/patterns.md) を参照する。

| タスクの性質 | 推奨パターン |
|---|---|
| シンプルな質問・汎用タスク | ゼロショット |
| 出力フォーマットをそろえたい | Few-shot |
| 専門知識を引き出したい | 役割プロンプティング + COSTAR |
| 多段階の算術・論理問題 | Chain-of-Thought（CoT） |
| 精度をとことん高めたい算術問題 | CoT + Self-Consistency |
| 複雑な構造を持つ難問 | Least-to-Most（LtM） |
| 外部ツール・検索が必要なタスク | ReAct |
| 戦略策定・オープンエンドな問題 | Tree of Thoughts（ToT） |
| 事実の正確性が最優先 | Chain of Verification（CoVe） |
| コード・文章の品質改善 | Self-Refine |
| 推論エラーの自己修正 | Reflexion |

## 使う順序の目安

1. **骨格を決める**: 上のフレームワーク（COSTAR など）でプロンプト全体の構造を決める。
2. **要素を埋める**: [reference/components.md](reference/components.md) のコンポーネント（Role / Task / Input / Output Format など）を、タスクに応じて取捨選択して埋める。
3. **手法を足す**: [reference/patterns.md](reference/patterns.md) から Few-shot や CoT などのパターンを必要に応じて組み合わせる。
4. **指針で仕上げる**: [reference/components.md](reference/components.md) の「全体指針」（曖昧語の排除、否定形回避、指示は末尾など）で見直す。

## 参照

- [reference/components.md](reference/components.md) — プロンプトの基本構造（コンポーネント）と作成の全体指針
- [reference/patterns.md](reference/patterns.md) — プロンプティング・パターンカタログ（基本・推論・検証）とタスク別の選び方

## 一次ソース

- https://www.promptingguide.ai/
- https://arxiv.org/abs/2201.11903
