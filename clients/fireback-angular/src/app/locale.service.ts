import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class LocaleService {
  private localeSubject = new BehaviorSubject<string>('en');
  locale$ = this.localeSubject.asObservable();

  constructor() {}

  setLocale(locale: string) {
    this.localeSubject.next(locale);
  }
}
