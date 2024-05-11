import { Component } from '@angular/core';
import { BaseComponent } from '../base.component';
import { LocaleService } from '../../locale.service';
import { strings } from './strings/translations';

@Component({
  selector: 'app-loader-sample',
  standalone: true,
  imports: [],
  templateUrl: './loader-sample.component.html',
  styleUrl: './loader-sample.component.scss',
})
export class LoaderSampleComponent extends BaseComponent {
  override s = strings;
  constructor(private locale: LocaleService) {
    super(locale, strings);
  }
}
