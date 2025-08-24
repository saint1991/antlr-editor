# CodeMirror言語拡張機能 実装計画

## 概要
editor-appのspec.mdに記載された要件1〜3（シンタックスハイライト、括弧マッチング、折りたたみ機能）を実現するための実装計画。

# MUST要件

Expression.g4の定義に従って、以下の機能を提供できること。

- [ ] 1. シンタックスハイライト
- [ ] 2. Paren, Bracketの対応をハイライトできること
- [ ] 3. エラー箇所には下波線をいれる。ホバー時にエラーメッセージを表示する。
- [ ] 4. Parenの内部を折りたたみできること(関数引数も含む)
- [ ] 5. フォーマッタ
- [ ] 6. 関数名の入力候補表示
  - 関数名だけでなく関数の動作の説明や使用例を表示できること


## 現状分析

### 既存リソース
- **WASM Analyzer**: `tokenizeExpression`でトークン情報（type, text, start, end, line, column）を取得可能
- **ANTLR4文法定義**: Expression.g4で言語仕様が定義済み
- **依存関係**: `@codemirror/language`, `@lezer/common`, `@lezer/lr`がインストール済み

### 現在の実装状況
- ✅ 基本的なCodeMirrorエディタ
- ⚠️ 括弧マッチング（CodeMirrorのデフォルト機能のみ）
- ❌ シンタックスハイライト（未実装）
- ❌ 折りたたみ機能（未実装）

## 実装アプローチ

### 選択方針: ANTLR Token Stream + Lezer Parser Adapter

既存のWASM Analyzerのトークン情報を活用してLezer Parser Adapterを作成する方針を採用。

**選定理由**:
- 既存のANTLR文法とトークナイザーを再利用可能
- 一貫性のある構文解析を維持
- WASMによる高速な処理
- 二重管理を回避

## 実装設計

### 1. シンタックスハイライト機能

#### 1.1 カスタムParserクラスの作成
```typescript
// extensions/expression-language.ts
class ExpressionParser extends Parser {
  private analyzer: Analyzer;
  
  createParse(input: string) {
    const tokens = this.analyzer.tokenizeExpression(input);
    return this.buildTree(tokens);
  }
}
```

#### 1.2 トークンタイプとスタイルタグのマッピング
- `STRING_LITERAL` → `t.string`
- `INTEGER_LITERAL` → `t.integer`
- `FLOAT_LITERAL` → `t.float`
- `BOOLEAN_LITERAL` → `t.bool`
- `FUNCTION_NAME` → `t.function(t.variableName)`
- `IDENTIFIER` → `t.variableName`
- `LPAREN/RPAREN` → `t.paren`
- `LBRACKET/RBRACKET` → `t.squareBracket`
- 演算子 → `t.arithmeticOperator` / `t.logicOperator`

### 2. 括弧マッチングの強化

#### 2.1 カスタム括弧設定
- 対象: `()` と `[]`
- 視覚的強調: マッチした括弧と未マッチの括弧で異なるスタイル
- 括弧タイプの識別: paren（丸括弧）とbracket（角括弧）

#### 2.2 実装内容
```typescript
const expressionBrackets = bracketMatching({
  brackets: "()[]",
  renderMatch: カスタムレンダリング関数
});
```

### 3. 折りたたみ機能

#### 3.1 折りたたみ可能な要素
- **括弧内式**: `(expression)` の内部
- **関数引数**: `FUNCTION(arg1, arg2, ...)` の引数リスト
- **カラム参照**: `[column_name]` の内部

#### 3.2 折りたたみルールの定義
```typescript
const foldingRules = foldNodeProp.add({
  ParenExpression: foldInside,
  FunctionCall: カスタム折りたたみ範囲関数,
  ColumnReference: foldInside
});
```

#### 3.3 UI要素
- Gutter表示: `▼`（展開時）/ `▶`（折りたたみ時）
- キーボードショートカット: 標準のfoldKeymap

## ファイル構成

```
editor-app/
├── src/app/antlr-editor/
│   ├── extensions/
│   │   ├── analyzer.ts (既存)
│   │   ├── expression-language.ts (新規)
│   │   ├── expression-parser.ts (新規)
│   │   └── expression-folding.ts (新規)
│   └── antlr-editor.component.ts (更新)
```

## 実装手順

### Phase 1: Parser Adapter基盤
1. `ExpressionParser`クラスの実装
2. トークンからLezer Treeへの変換ロジック
3. NodeTypeの定義とマッピング

### Phase 2: シンタックスハイライト
1. スタイルタグのマッピング実装
2. `LRLanguage.define()`での言語定義
3. エディタへの統合

### Phase 3: 括弧マッチング強化
1. カスタムbracketMatching設定
2. 視覚的強調のCSS定義
3. エディタへの統合

### Phase 4: 折りたたみ機能
1. foldNodePropの設定
2. foldGutterの追加
3. キーマップの設定

### Phase 5: 統合とテスト
1. 全機能の統合
2. パフォーマンステスト
3. エラーハンドリングの実装

## パフォーマンス最適化

### 最適化方針
1. **インクリメンタル解析**: 変更部分のみ再トークン化
2. **デバウンス処理**: 入力中の過剰な解析を防ぐ（例: 300ms）
3. **キャッシュ機構**: トークン結果のキャッシュ
4. **Web Worker検討**: 重い処理をバックグラウンドで実行

### パフォーマンス目標
- 初回解析: 100ms以内（1000トークン）
- インクリメンタル更新: 50ms以内
- メモリ使用量: 10MB以内

## エラーハンドリング

### エラー表示機能
```typescript
linter((view) => {
  const result = analyzer.tokenizeExpression(view.state.doc.toString());
  return result.errors.map(err => ({
    from: err.start,
    to: err.end,
    severity: 'error',
    message: err.message
  }));
})
```

## 今後の拡張可能性

### 将来的な機能追加
- オートコンプリート（spec.md要件6）
- フォーマッタのUI統合（spec.md要件5）
- ホバー情報表示
- リファクタリング支援

### 拡張性の確保
- モジュール化された設計
- 各機能の独立性を維持
- テスト可能な構造

## リスクと対策

### 技術的リスク
1. **WASMとの通信オーバーヘッド**
   - 対策: バッチ処理とキャッシュ

2. **大規模ファイルでのパフォーマンス**
   - 対策: 仮想スクロールとビューポート最適化

3. **ANTLR文法の変更対応**
   - 対策: トークンマッピングの抽象化

## 成功基準

- [ ] Expression.g4の全トークンタイプに対応したハイライト
- [ ] 括弧の対応が視覚的に明確
- [ ] スムーズな折りたたみ/展開動作
- [ ] 100ms以内のレスポンスタイム
- [ ] エラー時の適切なフィードバック