import { Component, Directive, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { strings } from './strings/translations';
import { LocaleService } from './locale.service';
import { CommonModule } from '@angular/common';
import { BaseComponent } from './components/base.component';
import { LoaderSampleComponent } from './components/loader-sample/loader-sample.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, CommonModule, LoaderSampleComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent extends BaseComponent {
  override s = strings;
  title = 'fireback-angular';

  constructor(protected override localeService: LocaleService) {
    super(localeService, strings);
  }
}
