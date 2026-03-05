---
id: mysql-transaction-anomaly
type: skill
domain: engineering
sources:
  - https://bmf-tech.com/posts/MySQLのトランザクションのアノマリーについて
  - https://dev.mysql.com/doc/refman/8.0/ja/innodb-transaction-isolation-levels.html
---

# MySQLトランザクションアノマリー評価スキル

## アノマリーとは

トランザクションの分離レベルや処理順序によって生じる**期待しない結果や不整合**のこと。
ANSI SQL標準やISO/IEC 9075によって定義されている。

以下、複数トランザクションはTX1、TX2と表記する。

---

## 5つのアノマリー

| アノマリー | 発生メカニズム | 対応分離レベル |
|-----------|--------------|-------------|
| **ダーティリード** | TX1がTX2のCOMMIT前のデータを読み取ってしまう | READ COMMITTED以上で防止 |
| **インコンシステントリード** | 読み取るデータに一貫性がない（COMMIT後の一貫性の崩れ） | ファジーリード/ファントムリードの上位概念的位置づけ |
| **ファジーリード**（ノンリピータブルリード） | TX1が他のTX2にて**更新**したデータを参照できてしまう | REPEATABLE READ以上で防止 |
| **ファントムリード** | TX2が新規**追加または削除**をCOMMITした場合にTX1が読み取るデータが変わってしまう | SERIALIZABLE（MySQLはREPEATABLE READでも防止） |
| **ロストアップデート** | TX1とTX2が同じデータを更新する際に競合が発生し、一部の更新が失われる | 楽観ロック・悲観ロックで対応 |

> **ファジーリード vs ファントムリード**: ファジーリードは「既存行の更新」が対象、ファントムリードは「行の追加または削除」が対象。

---

## MySQLのトランザクション分離レベルと防止できるアノマリー

MySQLのデフォルト分離レベルは **REPEATABLE READ**。

| 分離レベル | ダーティリード | ファジーリード | ファントムリード |
|-----------|-------------|-------------|--------------|
| READ UNCOMMITTED | ○（発生） | ○ | ○ |
| READ COMMITTED | × | ○ | ○ |
| REPEATABLE READ ※ | × | × | ○（MySQLでは×） |
| SERIALIZABLE | × | × | × |

※ MySQLのInnoDB では REPEATABLE READ においてもファントムリードが発生しないよう実装されている。

---

## アノマリー別の発生条件詳細

### ダーティリード（READ UNCOMMITTEDで発生）
- TX2がINSERTしてCOMMIT前にTX1がSELECTすると、TX2のデータが見える
- TX2がAbortした場合、TX1が読み取ったデータは不正なものになる

### ファジーリード（READ COMMITTEDで発生）
- TX1がSELECTした後、TX2がUPDATE+COMMITし、TX1が再度SELECTすると結果が変わる
- 1つのTX内で同じデータを複数回読み取った結果が変わる

### ファントムリード（READ COMMITTEDで発生）
- TX1がSELECTした後、TX2がINSERT+COMMITし、TX1が再度SELECTするとレコード数が変わる

### ロストアップデート（REPEATABLE READでも発生し得る）
- TX1とTX2が同じ行を読み取り、それぞれ更新してCOMMITすると、先のCOMMITが失われる
- 楽観ロック（更新前に取得時と同じ状態か検証）または悲観ロック（SELECT ... FOR UPDATE）で対処

---

## 評価に使える問い

- 使用している分離レベルはアプリケーションの要件に対して適切か？
- ロストアップデートが起こりうる更新処理で楽観ロックまたは悲観ロックが実装されているか？
- MySQLのデフォルト（REPEATABLE READ）を前提にした実装になっているか？
- READ COMMITTEDで運用している場合、ファジーリード・ファントムリードが問題にならないか？
- SERIALIZABLE を選択している場合、性能トレードオフが許容できる規模か？
