import { type AfterViewInit, Component, type ElementRef, EventEmitter, Input, type OnInit, Output, ViewChild } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { indentWithTab } from '@codemirror/commands';
import { bracketMatching, foldGutter, foldKeymap, indentUnit } from '@codemirror/language';
import type { Extension } from '@codemirror/state';
import { oneDark } from '@codemirror/theme-one-dark';
import { keymap, lineNumbers } from '@codemirror/view';
import { basicSetup, EditorView } from 'codemirror';
import { loadAnalyzer } from './extensions/analyzer';
import { expressionLanguage } from './extensions/syntax-highlight/stream-language';

@Component({
  selector: 'antlr-editor',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './antlr-editor.html',
  styleUrls: ['./antlr-editor.css'],
})
export class AntlrEditorComponent implements OnInit, AfterViewInit {
  @ViewChild('editor', { static: true }) editorElement!: ElementRef<HTMLDivElement>;
  @Input() initialValue: string = '';
  @Input() theme: 'light' | 'dark' = 'dark';
  @Output() valueChange = new EventEmitter<string>();

  private editorView!: EditorView;

  ngOnInit() {}

  async ngAfterViewInit() {
    await this.initializeEditor();
  }

  private async initializeEditor() {
    // Load the WASM analyzer
    const analyzer = await loadAnalyzer();

    const extensions: Extension[] = [
      basicSetup,
      keymap.of([indentWithTab]),
      indentUnit.of('  '),
      expressionLanguage(analyzer, this.theme), // Add syntax highlighting with theme
      bracketMatching(),
      lineNumbers(),
      foldGutter(),
      keymap.of(foldKeymap),
      EditorView.updateListener.of((update) => {
        if (update.docChanged) {
          const value = update.state.doc.toString();
          this.valueChange.emit(value);
        }
      }),
    ];

    // Add dark theme if selected
    if (this.theme === 'dark') {
      extensions.push(oneDark);
    }

    this.editorView = new EditorView({
      doc: this.initialValue,
      parent: this.editorElement.nativeElement,
      extensions: extensions,
    });
  }

  getValue(): string {
    return this.editorView.state.doc.toString();
  }

  setValue(value: string) {
    const transaction = this.editorView.state.update({
      changes: {
        from: 0,
        to: this.editorView.state.doc.length,
        insert: value,
      },
    });
    this.editorView.dispatch(transaction);
  }
}
