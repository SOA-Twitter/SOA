import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UnloggedHomeComponent } from './unlogged-home.component';

describe('UnloggedHomeComponent', () => {
  let component: UnloggedHomeComponent;
  let fixture: ComponentFixture<UnloggedHomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UnloggedHomeComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UnloggedHomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
