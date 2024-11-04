import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminAssignComponent } from './admin-assign.component';

describe('AdminAssignComponent', () => {
  let component: AdminAssignComponent;
  let fixture: ComponentFixture<AdminAssignComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AdminAssignComponent]
    });
    fixture = TestBed.createComponent(AdminAssignComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
