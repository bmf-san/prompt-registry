---
id: mysql-lock
type: skill
domain: engineering
sources:
  - https://bmf-tech.com/posts/MySQLのロックについて
  - https://dev.mysql.com/doc/refman/8.0/ja/innodb-locking.html
---

# MySQLロック評価スキル

## 内部レベルロック（ロック粒度）

MySQLには行レベルとテーブルレベルの2種類の粒度がある。

| ロック粒度 | 対象 | 特徴 |
|-----------|------|------|
| **行レベルロック** | テーブル内の個々の行 | ロック競合・ロールバック範囲が小さい。1行を長時間ロック可能 |
| **テーブルレベルロック** | テーブル全体 | メモリ消費が少ない。テーブルの大部分を対象とする操作（GROUP BY・全テーブルスキャン）は高速 |

---

## InnoDBロックの種類

### 共有（READ）ロック / 占有（排他・WRITE）ロック

| ロック種別 | SQLでの取得 | 他TXのREAD | 他TXのWRITE |
|-----------|------------|-----------|------------|
| **共有（READ）ロック** `S` | `SELECT ... LOCK IN SHARE MODE` | ○ | × |
| **占有（排他・WRITE）ロック** `X` | `SELECT ... FOR UPDATE` | ○（単純SELECTのみ） | × |

### インテンションロック

トランザクションがテーブルの行に必要とするロックタイプ（共有または排他）を示す**テーブルレベルのロック**。
行ロックとテーブルロックの共存をサポートするために用意されている。

- **インテンション共有ロック（IS）**: 行に共有ロックをかける前に設定
- **インテンション排他ロック（IX）**: 行に排他ロックをかける前に設定

SQLで明示的に操作するものではなく、DB内部で自動管理される。

### レコードロック

インデックスレコード（クラスタインデックスとセカンダリインデックス）に対するロック。スキャンしたインデックスに対して設定される。

### ギャップロック

インデックスレコード**間のギャップ**のロック。または、インデックスレコードの前後のギャップのロック。

- 行単位のロックに見えても**範囲でロックされる**ことに注意
- 例: `SELECT ... WHERE ID BETWEEN 1 AND 5 FOR UPDATE` → id=3 が存在しなくてもその位置へのINSERTがブロックされる

### ネクストキーロック

**レコードロック** + **そのレコードの前のギャップへのギャップロック**の組み合わせ。

- `SELECT ... WHERE ID < 5 FOR UPDATE` を実行すると、id=5 未満の行だけでなく末尾インデックス値の後のギャップもロックされる
- ファントムリードを防ぐ仕組み

### インテンションロックの挿入

INSERT前に設定されるギャップロックの一種。INSERTのインテンションロック。

### AUTO-INCロック

`AUTO_INCREMENT` カラムを含むテーブルへのINSERT時に取得されるテーブルロック。
TX1がAUTO_INCREMENTの値を取得している間は、TX2がAUTO_INCREMENTの値を取得できない。

---

## ロックの確認方法

```sql
-- ロックの状態確認
SELECT * FROM performance_schema.data_locks;

-- ロック件数確認＋スレッドID
SHOW ENGINE INNODB STATUS;

-- ロック件数確認
SELECT trx_id, trx_rows_locked, trx_mysql_thread_id FROM information_schema.INNODB_TRX;
```

**デッドロックの確認**: `SHOW ENGINE INNODB STATUS` を実行し `LATEST DETECTED DEADLOCK` セクションを確認。

---

## 評価に使える問い

- `SELECT ... FOR UPDATE`（占有ロック）と `LOCK IN SHARE MODE`（共有ロック）の使い分けは適切か？
- ギャップロックやネクストキーロックにより意図しない範囲のINSERTがブロックされないか？
- デッドロックが発生しうる複数ロック取得のパターンはないか（ロック取得順序の逆転）？
- AUTO_INCROEMENTロックが高トラフィックのINSERTでボトルネックになっていないか？
- インテンションロックの存在を踏まえ、行ロックとテーブル操作の競合が設計に含まれていないか？
- `EXPLAIN` や `performance_schema.data_locks` でロックの実際の範囲を確認しているか？
