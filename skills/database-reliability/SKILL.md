---
name: database-reliability
description: リレーショナルデータベース（主に MySQL/InnoDB）の信頼性を評価・レビューする。トランザクション設計、ロック競合とデッドロック、分離レベルとアノマリー（整合性異常）、ACID・BASE の整合性モデル選定、インデックス設計、DB マイグレーションの安全性を検討するときに使う。発火トリガ: トランザクション, ロック, デッドロック, 分離レベル, アノマリー, ACID, BASE, インデックス, EXPLAIN, マイグレーションのレビュー, 後方互換, ゼロダウンタイム, DB, MySQL, InnoDB。
---

# データベース信頼性

リレーショナルデータベース（主に MySQL/InnoDB）の信頼性に関わる主要トピック——トランザクション、ロック、整合性異常、ACID/BASE、インデックス、マイグレーション安全性——を横断的に評価・レビューするための索引 skill。各トピックの詳細は `reference/` に置く。まず本ファイルで論点を掴み、必要な詳細を該当 reference で確認する。

## いつ使うか

- トランザクション・ロック・分離レベルまわりの実装や不具合を評価するとき
- 強整合か最終的整合性かなど、整合性モデルを選定するとき
- クエリ性能・インデックス設計を検討するとき
- DB マイグレーション（スキーマ変更）を安全性の観点でレビューするとき
- 「どの分離レベルにすべきか」「デッドロックの原因は何か」「このインデックスは効くか」といった問いに答えるとき

## トピック索引

| 検討したいこと | 参照 |
|---|---|
| トランザクションの基礎理論（ACID、同時実行制御、防ぐべき異常、分離レベル、楽観/悲観ロック） | reference/transaction.md |
| MySQL/InnoDB のロック（粒度・種類・ギャップ/ネクストキー・確認方法） | reference/mysql-lock.md |
| MySQL のアノマリーと分離レベルの対応（REPEATABLE READ の挙動） | reference/transaction-anomaly.md |
| 強整合（ACID）か最終的整合性（BASE）かの選定 | reference/acid-and-base.md |
| インデックス設計と EXPLAIN による効果測定 | reference/indexing.md |
| スキーマ変更（マイグレーション）の安全性レビュー | reference/migration-review.md |

## 判断の要点

- **分離レベルはトレードオフ**。高いほど異常（ダーティ/ノンリピータブル/ファントムリード）を防げるが並列性能は落ちる。要件から必要最小の分離レベルを選ぶ。詳細は reference/transaction.md・reference/transaction-anomaly.md。
- **MySQL/InnoDB のデフォルトは REPEATABLE READ**。InnoDB では REPEATABLE READ でもファントムリードが（ネクストキーロックにより）発生しない。他 DB の常識をそのまま持ち込まない。
- **Atomicity は「成功保証」ではなく「失敗時の全取消保証」**。Abort 時のリトライ・エラー処理を実装していないトランザクションは扱いが誤っている。
- **ロストアップデートは分離レベルだけでは防げない**。楽観ロック（更新前に取得時と同状態か検証）または悲観ロック（`SELECT ... FOR UPDATE`）で対処する。
- **行ロックでも範囲がロックされることがある**。ギャップロック/ネクストキーロックにより、存在しない行への INSERT までブロックされうる。デッドロックはロック取得順序の逆転を疑う。
- **整合性モデルは要件で選ぶ**。金融・在庫・予約など二重処理が許されない領域は ACID、SNS・CDN・検索インデックスなど最終的整合性で足りる領域は BASE。分散環境で ACID を無理に強制するとコストが跳ねる。
- **インデックスは必ず EXPLAIN で計測**。演算・関数・暗黙型変換・後方一致 LIKE・否定形はインデックスを無効化する。書き込み頻度の高いテーブルでの過剰なインデックスは書き込みコストを増やす。
- **マイグレーションは後方互換・ゼロダウンタイム・ロールバックを最優先で見る**。参照中のカラム/テーブル削除、既存 NULL がある状態での NOT NULL 追加、大規模テーブルのロックを伴う操作は MUST FIX。詳細と合否基準は reference/migration-review.md。
- **外部キー的に機能するカラム（`_id` 等）が複数あるテーブルは、業務上のユニーク制約の要否を必ず確認する**。

## 参照

- [トランザクション基礎](reference/transaction.md) — ACID、同時実行制御、防ぐべき異常、ロックとスケジュール、デッドロック、楽観/悲観ロック、分離レベル。
- [MySQL ロック](reference/mysql-lock.md) — ロック粒度、共有/占有・インテンション・レコード・ギャップ・ネクストキー・AUTO-INC ロック、確認クエリ。
- [トランザクションアノマリー（MySQL）](reference/transaction-anomaly.md) — 5 つのアノマリーと分離レベルの対応、MySQL/InnoDB 固有の挙動。
- [ACID と BASE](reference/acid-and-base.md) — 2 つの整合性モデル、実装技術、設計思想の違いと選定。
- [インデックス](reference/indexing.md) — インデックスの仕組み、パターン、EXPLAIN、効かないパターン。
- [マイグレーションのレビュー観点](reference/migration-review.md) — 後方互換・ゼロダウンタイム・インデックス・ロールバックのチェックリストと合否基準。

## 一次ソース

- https://bmf-tech.com/posts/トランザクション概観
- https://gihyo.jp/book/2015/978-4-7741-7197-5
- https://bmf-tech.com/posts/MySQLのロックについて
- https://dev.mysql.com/doc/refman/8.0/ja/innodb-locking.html
- https://bmf-tech.com/posts/MySQLのトランザクションのアノマリーについて
- https://dev.mysql.com/doc/refman/8.0/ja/innodb-transaction-isolation-levels.html
- https://bmf-tech.com/posts/ACIDとBASEについて
- https://bmf-tech.com/posts/Indexとはなにか
- https://dev.mysql.com/doc/refman/8.0/ja/glossary.html#glos_covering_index
- https://planetscale.com/blog/backward-compatible-database-migrations
- https://www.braintreepayments.com/blog/safe-operations-for-high-volume-postgresql
