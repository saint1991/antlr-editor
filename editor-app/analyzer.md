# Analyzer API設計仕様書

## 概要

CodeMirrorのLezerパーサーシステムとの完全な統合を実現するために必要なAnalyzer APIの拡張設計。
現在のトークンベースのAPIから、構文木（AST）ベースのAPIへの移行を定義する。

## 現状の課題

現在のAnalyzer APIは以下の制限がある：
- フラットなトークンリストのみを返す（構造情報なし）
- 親子関係や階層構造の情報がない
- インクリメンタル解析をサポートしていない
- セマンティック情報（型、スコープ等）を提供していない

## 新API設計

### 1. SyntaxNode定義

```typescript
interface SyntaxNode {
  // 基本情報
  id: number;                    // ユニークなノードID
  type: NodeType;                // ノードタイプ（詳細は後述）
  
  // 位置情報
  range: {
    start: number;               // 開始位置（0-indexed）
    end: number;                 // 終了位置
    startLine: number;           // 開始行（1-indexed）
    startColumn: number;         // 開始列（1-indexed）
    endLine: number;             // 終了行
    endColumn: number;           // 終了列
  };
  
  // 構造情報
  parent?: number;               // 親ノードのID
  children: number[];            // 子ノードのIDリスト
  nextSibling?: number;          // 次の兄弟ノードのID
  prevSibling?: number;          // 前の兄弟ノードのID
  
  // コンテンツ
  text: string;                  // ノードのテキスト内容
  value?: any;                   // パース済みの値（リテラルの場合）
  
  // セマンティック情報
  semantic?: {
    dataType?: DataType;         // データ型（NUMBER, STRING, BOOLEAN, DATE等）
    isConstant?: boolean;        // 定数式かどうか
    constantValue?: any;         // 定数の場合の計算済み値
    referencedColumns?: string[]; // 参照しているカラム名
    referencedFunctions?: string[]; // 使用している関数名
  };
  
  // 編集支援情報
  features?: {
    foldable?: boolean;          // 折りたたみ可能か
    foldRange?: {                // 折りたたみ範囲
      start: number;
      end: number;
    };
    completionContext?: CompletionContext; // オートコンプリート用コンテキスト
  };
  
  // エラー情報
  error?: {
    severity: 'error' | 'warning' | 'info';
    message: string;
    code?: string;               // エラーコード
    quickFixes?: QuickFix[];    // 修正候補
  };
}
```

### 2. NodeType定義

```typescript
enum NodeType {
  // ルート
  Program = 'Program',
  
  // 式
  Expression = 'Expression',
  BinaryExpression = 'BinaryExpression',
  UnaryExpression = 'UnaryExpression',
  TernaryExpression = 'TernaryExpression',
  ParenthesizedExpression = 'ParenthesizedExpression',
  
  // 関数
  FunctionCall = 'FunctionCall',
  FunctionName = 'FunctionName',
  ArgumentList = 'ArgumentList',
  Argument = 'Argument',
  
  // リテラル
  IntegerLiteral = 'IntegerLiteral',
  FloatLiteral = 'FloatLiteral',
  StringLiteral = 'StringLiteral',
  BooleanLiteral = 'BooleanLiteral',
  NullLiteral = 'NullLiteral',
  DateLiteral = 'DateLiteral',
  
  // 識別子
  Identifier = 'Identifier',
  ColumnReference = 'ColumnReference',
  
  // 演算子
  BinaryOperator = 'BinaryOperator',
  UnaryOperator = 'UnaryOperator',
  ComparisonOperator = 'ComparisonOperator',
  LogicalOperator = 'LogicalOperator',
  
  // 区切り文字
  LeftParen = 'LeftParen',
  RightParen = 'RightParen',
  LeftBracket = 'LeftBracket',
  RightBracket = 'RightBracket',
  Comma = 'Comma',
  
  // エラーノード
  ErrorNode = 'ErrorNode',
  MissingNode = 'MissingNode',  // 期待されるが存在しないノード
}
```

### 3. 主要API関数

```typescript
interface EnhancedAnalyzer {
  /**
   * 式を解析して構文木を返す
   */
  parseExpression(expression: string): ParseResult;
  
  /**
   * インクリメンタル解析（差分のみ再解析）
   */
  parseIncremental(
    expression: string,
    previousTree: SyntaxTree,
    changes: TextChange[]
  ): ParseResult;
  
  /**
   * 特定位置のノードを取得
   */
  getNodeAtPosition(
    tree: SyntaxTree,
    position: number
  ): SyntaxNode | null;
  
  /**
   * オートコンプリート候補を取得
   */
  getCompletions(
    tree: SyntaxTree,
    position: number
  ): CompletionItem[];
  
  /**
   * ホバー情報を取得
   */
  getHoverInfo(
    tree: SyntaxTree,
    position: number
  ): HoverInfo | null;
  
  /**
   * 折りたたみ可能な範囲を取得
   */
  getFoldingRanges(tree: SyntaxTree): FoldingRange[];
  
  /**
   * セマンティックトークンを取得（より詳細なハイライト用）
   */
  getSemanticTokens(tree: SyntaxTree): SemanticToken[];
  
  /**
   * リファクタリング候補を取得
   */
  getRefactorings(
    tree: SyntaxTree,
    selection: Range
  ): Refactoring[];
}
```

### 4. データ構造定義

```typescript
interface ParseResult {
  tree: SyntaxTree;
  errors: ParseError[];
  warnings: ParseWarning[];
  metrics: {
    parseTime: number;          // 解析時間（ms）
    nodeCount: number;          // ノード数
    maxDepth: number;           // 最大深度
  };
}

interface SyntaxTree {
  root: SyntaxNode;
  nodes: Map<number, SyntaxNode>; // ID -> Node のマップ（高速アクセス用）
  version: number;                // ツリーのバージョン（インクリメンタル解析用）
}

interface TextChange {
  range: {
    start: number;
    end: number;
  };
  text: string;
}

interface CompletionItem {
  label: string;                  // 表示ラベル
  kind: CompletionKind;           // 種別（Function, Column, Keyword等）
  detail?: string;                // 詳細説明
  documentation?: string;         // ドキュメント
  insertText: string;             // 挿入テキスト
  range: Range;                   // 置換範囲
  additionalTextEdits?: TextEdit[]; // 追加の編集（import追加等）
}

interface HoverInfo {
  contents: string | MarkupContent;
  range: Range;
  actions?: CodeAction[];         // クイックアクション
}

interface FoldingRange {
  start: number;
  end: number;
  kind?: 'comment' | 'imports' | 'region';
  collapsedText?: string;         // 折りたたみ時の表示テキスト
}

interface SemanticToken {
  line: number;
  column: number;
  length: number;
  tokenType: SemanticTokenType;
  tokenModifiers?: SemanticTokenModifier[];
}
```

### 5. Go実装側の構造体

```go
package analyzer

// SyntaxNode represents a node in the syntax tree
type SyntaxNode struct {
    ID           int                 `json:"id"`
    Type         NodeType            `json:"type"`
    Range        Range               `json:"range"`
    Parent       *int                `json:"parent,omitempty"`
    Children     []int               `json:"children"`
    NextSibling  *int                `json:"nextSibling,omitempty"`
    PrevSibling  *int                `json:"prevSibling,omitempty"`
    Text         string              `json:"text"`
    Value        interface{}         `json:"value,omitempty"`
    Semantic     *SemanticInfo       `json:"semantic,omitempty"`
    Features     *NodeFeatures       `json:"features,omitempty"`
    Error        *NodeError          `json:"error,omitempty"`
}

// ANTLRのParseTreeからSyntaxNodeツリーを構築
func buildSyntaxTree(ctx antlr.ParserRuleContext) *SyntaxTree {
    builder := &treeBuilder{
        nodes:    make(map[int]*SyntaxNode),
        nextID:   1,
    }
    
    root := builder.visitNode(ctx)
    
    return &SyntaxTree{
        Root:    root,
        Nodes:   builder.nodes,
        Version: 1,
    }
}

// ビジターパターンでParseTreeを走査
func (b *treeBuilder) visitNode(node antlr.ParseTree) *SyntaxNode {
    syntaxNode := &SyntaxNode{
        ID:   b.nextID,
        Text: node.GetText(),
    }
    b.nextID++
    
    // ノードタイプの判定
    switch n := node.(type) {
    case *parser.ExpressionContext:
        syntaxNode.Type = NodeTypeExpression
        // セマンティック情報の追加
        syntaxNode.Semantic = b.analyzeExpression(n)
        
    case *parser.FunctionCallContext:
        syntaxNode.Type = NodeTypeFunctionCall
        // 折りたたみ情報の追加
        syntaxNode.Features = &NodeFeatures{
            Foldable: true,
            FoldRange: &Range{
                Start: n.GetStart().GetStart(),
                End:   n.GetStop().GetStop(),
            },
        }
        
    // ... 他のノードタイプ
    }
    
    // 子ノードの処理
    for i := 0; i < node.GetChildCount(); i++ {
        child := node.GetChild(i)
        childNode := b.visitNode(child)
        childNode.Parent = &syntaxNode.ID
        syntaxNode.Children = append(syntaxNode.Children, childNode.ID)
    }
    
    b.nodes[syntaxNode.ID] = syntaxNode
    return syntaxNode
}
```

## 実装優先順位

### Phase 1: 基本構文木API（必須）
1. `parseExpression` - 構文木の生成
2. `getNodeAtPosition` - 位置からノード取得
3. `getFoldingRanges` - 折りたたみ範囲

### Phase 2: エディタ支援機能（重要）
1. `getCompletions` - オートコンプリート
2. `getHoverInfo` - ホバー情報
3. `getSemanticTokens` - セマンティックハイライト

### Phase 3: 高度な機能（オプション）
1. `parseIncremental` - インクリメンタル解析
2. `getRefactorings` - リファクタリング支援
3. 型推論とデータフロー解析

## 移行戦略

1. **後方互換性の維持**: 既存のトークンベースAPIは残す
2. **段階的な移行**: 新旧APIを並行提供
3. **パフォーマンス最適化**: 構文木のキャッシュ機構
4. **WASM最適化**: シリアライズの効率化

## パフォーマンス考慮事項

- **メモリ使用量**: 大規模な式でのメモリ効率
- **シリアライズ**: Go構造体からJSONへの変換コスト
- **キャッシュ**: 頻繁にアクセスされるノードのキャッシュ
- **遅延評価**: セマンティック情報の遅延計算

## テスト戦略

1. **単体テスト**: 各API関数の個別テスト
2. **統合テスト**: CodeMirrorとの統合テスト
3. **パフォーマンステスト**: 大規模式での性能測定
4. **回帰テスト**: 既存機能の互換性確認