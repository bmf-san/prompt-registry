## 変更内容

<!-- どのファイルを追加・更新・削除したか簡潔に記述 -->

## 変更種別

- [ ] 新規ファイル追加
- [ ] 既存ファイルの更新
- [ ] ファイル削除
- [ ] その他（リファクタリング、ドキュメント修正など）

## 対象ファイルの type

- [ ] `persona` — personas/
- [ ] `skill` — skills/
- [ ] `review` — reviews/
- [ ] `artifact` — artifacts/
- [ ] 対象外（ドキュメント・設定ファイルなど）

## チェックリスト

- [ ] `go run ./scripts/validate/ .` でエラーなし
- [ ] フロントマターの `id` がファイル名と一致している
- [ ] `type` と配置ディレクトリが対応している（例: `persona` → `personas/`）
- [ ] `domain` が `config.yaml` に定義済みの値になっている
- [ ] **persona を追加・更新した場合**: [WRITING_GUIDE](../docs/WRITING_GUIDE.md) のセクション構成（Task / Input / Output Format / Guidelines / Prohibited Actions / Example / Knowledge Base）に準拠している
