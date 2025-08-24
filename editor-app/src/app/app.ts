import { Component, type OnInit, ViewChild } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AntlrEditorComponent } from './antlr-editor/antlr-editor.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [AntlrEditorComponent, FormsModule],
  templateUrl: './app.html',
  styleUrls: ['./app.css'],
})
export class App implements OnInit {
  @ViewChild(AntlrEditorComponent) codeEditor!: AntlrEditorComponent;

  title = 'Expression Language Editor';
  selectedTheme: 'light' | 'dark' = 'dark';

  initialCode = `MAX(LEN("hello"), MIN([column1], 42.5)) >= 5 && TRUE || (SUM(1, [column3], 3.14) > 5)`;

  ngOnInit() {}

  protected handleValueChange(value: string) {
    console.debug('value changed:', value);
  }
}
