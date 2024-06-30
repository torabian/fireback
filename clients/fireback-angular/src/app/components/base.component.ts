import { Directive } from '@angular/core';
import { LocaleService } from '../locale.service';

@Directive({})
export class BaseComponent {
  protected s: unknown;
  constructor(protected localeService: LocaleService, private strings: any) {
    this.handleLocale();
  }

  handleLocale() {
    this.localeService.locale$.subscribe((v) => {
      if (v === 'en') {
        this.s = this.strings;
      } else if ((this.strings as any)['$' + v]) {
        this.s = (this.strings as any)['$' + v];
      }
    });
  }
}
