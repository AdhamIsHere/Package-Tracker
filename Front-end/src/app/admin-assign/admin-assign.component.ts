import {Component, OnInit} from '@angular/core';
import {OrderService} from "../services/order.service";

@Component({
  selector: 'app-admin-assign',
  templateUrl: './admin-assign.component.html',
  styleUrls: ['./admin-assign.component.css']
})
export class AdminAssignComponent implements OnInit {
  orders: any[] = [];
  successMessage: string = '';
  errorMessage: string = '';
  constructor(private orderService: OrderService) {}

  ngOnInit(): void {
    this.loadOrders();
  }
  assignOrder(orderId: number): void {
    this.orderService.assignOrder(orderId).subscribe(
      data => {
        this.successMessage = 'Order assigned successfully';
        this.loadOrders();
      },
      error => {
        this.errorMessage = 'Error assigning order';
        console.error('Error assigning order', error);
      }
    );
  }
  loadOrders(): void {
    this.orderService.getAllOrders().subscribe(
      data => {
        this.orders = data;
      },
      error => {
        this.errorMessage = 'Error fetching orders';
        console.error('Error fetching orders', error);
      }
    );
  }
}
