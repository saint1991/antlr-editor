import { Component, ViewChild, OnInit } from '@angular/core';
import { AntlrEditorComponent } from './antlr-editor/antlr-editor.component';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [AntlrEditorComponent, FormsModule],
  templateUrl: './app.html',
  styleUrls: ['./app.css']
})
export class App implements OnInit {
  
  @ViewChild(AntlrEditorComponent) codeEditor!: AntlrEditorComponent;
  
  title = 'Expression Language Editor';
  selectedTheme: 'light' | 'dark' = 'dark';
  
  initialCode = `[column1] + [column2] * 2`;

  ngOnInit() {
    
  }

  protected handleValueChange(value: string) {
    console.log('Editor value changed');
  }
}