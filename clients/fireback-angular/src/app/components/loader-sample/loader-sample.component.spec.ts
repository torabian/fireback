import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoaderSampleComponent } from './loader-sample.component';

describe('LoaderSampleComponent', () => {
  let component: LoaderSampleComponent;
  let fixture: ComponentFixture<LoaderSampleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LoaderSampleComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(LoaderSampleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
