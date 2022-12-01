import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SignupBusinessComponent } from './signup-business.component';

describe('SignupBusinessComponent', () => {
  let component: SignupBusinessComponent;
  let fixture: ComponentFixture<SignupBusinessComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SignupBusinessComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SignupBusinessComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
