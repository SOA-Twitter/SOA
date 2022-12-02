import { TestBed } from '@angular/core/testing';
import { Guard } from './guard';


describe('GuardGuard', () => {
  let guard: Guard;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    guard = TestBed.inject(Guard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});
