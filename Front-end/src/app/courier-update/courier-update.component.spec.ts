import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CourierUpdateComponent } from './courier-update.component';

describe('CourierUpdateComponent', () => {
  let component: CourierUpdateComponent;
  let fixture: ComponentFixture<CourierUpdateComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CourierUpdateComponent]
    });
    fixture = TestBed.createComponent(CourierUpdateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
