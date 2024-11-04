import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CourierAssignComponent } from './courier-assign.component';

describe('CourierAssignComponent', () => {
  let component: CourierAssignComponent;
  let fixture: ComponentFixture<CourierAssignComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CourierAssignComponent]
    });
    fixture = TestBed.createComponent(CourierAssignComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
