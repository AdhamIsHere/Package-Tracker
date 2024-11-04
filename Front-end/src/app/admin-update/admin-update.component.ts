import {Component, OnInit} from '@angular/core';
import {OrderService} from "../services/order.service";

@Component({
  selector: 'app-admin-update',
  templateUrl: './admin-update.component.html',
  styleUrls: ['./admin-update.component.css']
})
export class AdminUpdateComponent implements OnInit {
  orders: any[] = [];
  successMessage: string = '';
  errorMessage: string = '';
  constructor(private orderService: OrderService) {}

  ngOnInit(): void {
    this.loadOrders();
  }
  updateOrder(orderId: number): void {
    this.orderService.updateOrder(orderId).subscribe(
      data => {
        this.successMessage = 'Order updated successfully';
        this.loadOrders();
      },
      error => {
        this.errorMessage = 'Error updating order';
        console.error('Error updating order', error);
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
