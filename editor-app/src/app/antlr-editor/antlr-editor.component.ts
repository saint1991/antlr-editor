import { type AfterViewInit, Component, type ElementRef, EventEmitter, Input, type OnInit, Output, ViewChild } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { indentWithTab } from '@codemirror/commands';
import { bracketMatching, indentUnit } from '@codemirror/language';
import type { Extension } from '@codemirror/state';
import { keymap } from '@codemirror/view';
import { basicSetup, EditorView } from 'codemirror';

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

  ngAfterViewInit() {
    this.initializeEditor();
  }

  private initializeEditor() {
    const extensions: Extension[] = [
      basicSetup,
      keymap.of([indentWithTab]),
      indentUnit.of('  '),
      bracketMatching(),
      EditorView.updateListener.of((update) => {
        if (update.docChanged) {
          const value = update.state.doc.toString();
          this.valueChange.emit(value);
        }
      }),
    ];

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
