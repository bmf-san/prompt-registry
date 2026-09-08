---
name: design-documentation
description: 設計・アーキテクチャのドキュメントを書き、鮮度を保ち、レビューするときに使う。Design Doc / アーキテクチャドキュメントのテンプレートと書き方、設計の賞味期限（陳腐化）の考え方、設計レビューのチェックリストを束ねる。トリガー: 設計ドキュメント, デザインドック, Design Doc, アーキテクチャドキュメント, 設計の陳腐化, 設計の賞味期限, 設計レビュー。
---

# 設計ドキュメンテーション

設計・アーキテクチャに関するドキュメントを書き、鮮度を保ち、レビューするための実践知の集約。設計情報の羅列ではなく、読者の理解と納得を導き、設計の妥当性について合意を形成するための文書を作ることを目的とする。

## いつ使うか

- Design Doc（「どう作るか」の意思決定）を書くとき → [reference/template.md](reference/template.md)
- アーキテクチャドキュメント（設計の妥当性を説明し合意形成する資料）を書くとき → [reference/architecture-document.md](reference/architecture-document.md)
- 設計にどれだけの寿命・堅牢性を持たせるか判断するとき → [reference/design-expiry.md](reference/design-expiry.md)
- 成果物としての設計（UI/UX など）をレビューするとき → [reference/review.md](reference/review.md)

## 基本姿勢

- **目的を最初に決める**: 誰に向けて、何を伝える文書かを最初に定義する。読者（ステークホルダー）が誰かが内容を左右する変数になる。
- **判断理由を残す**: なぜその構成を選んだか、代替案との比較・却下理由を書く。図だけで終わらせない。
- **簡潔性から始める**: まず重要な設計判断に焦点を絞り、その後に他の観点や詳細を足す。最初から詳細を詰め込むと認知負荷が高くなる。
- **鮮度を保つ**: 設計の進化に合わせて更新する。更新方針を明記し、陳腐化を放置しない。
- **寿命を意識する**: 「この設計は何年もてばよいか」を問い、過剰設計を避ける。

## Design Doc を書くかどうか

スコープが大きい・不確実性が高い・複数人が関わる実装が対象のときに書く。小規模で明確な実装は Design Doc を書かずに進めた方が速い。「どう作るか」の意思決定フェーズには Design Doc を、「どうなっているか」の記録フェーズには System Spec（requirements-engineering skill の reference/template.md）を使い分ける。

## 参照

- [reference/template.md](reference/template.md) — Design Doc のテンプレートと書き方（背景・目的・設計・テスト方針・リリース計画・モニタリングなど）
- [reference/architecture-document.md](reference/architecture-document.md) — アーキテクチャドキュメントを書くときの注意（6つの特性・推奨構成・評価のための問い）
- [reference/design-expiry.md](reference/design-expiry.md) — 設計の賞味期限（設計に寿命の観点を与え、制約を受容する）
- [reference/review.md](reference/review.md) — 設計レビューのチェックリスト（ユーザビリティ・可読性・インタラクション・アクセシビリティ・レスポンシブ・エッジケース）

## 一次ソース

- https://bmf-tech.com/posts/アーキテクチャドキュメントを書くときに気をつけること
- https://bmf-tech.com/posts/設計の賞味期限を考える
- https://www.industrialempathy.com/posts/design-docs-at-google/
- https://medium.com/machine-words/writing-technical-design-docs-71f446e42173
