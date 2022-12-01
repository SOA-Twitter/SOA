import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProfileRecoveryComponent } from './profile-recovery.component';

describe('ProfileRecoveryComponent', () => {
  let component: ProfileRecoveryComponent;
  let fixture: ComponentFixture<ProfileRecoveryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProfileRecoveryComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProfileRecoveryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
