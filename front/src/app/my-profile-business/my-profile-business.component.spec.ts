import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MyProfileBusinessComponent } from './my-profile-business.component';

describe('MyProfileBusinessComponent', () => {
  let component: MyProfileBusinessComponent;
  let fixture: ComponentFixture<MyProfileBusinessComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MyProfileBusinessComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MyProfileBusinessComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
