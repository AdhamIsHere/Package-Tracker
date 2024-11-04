import { Component, OnInit } from '@angular/core';
import { OrderService } from '../services/order.service';

@Component({
  selector: 'app-orders-admin',
  templateUrl: './orders-admin.component.html',
  styleUrls: ['./orders-admin.component.css']
})
export class OrdersAdminComponent implements OnInit {

  orders: any[] = [];
  currentOrderIndex: number = 0;

  constructor(private orderService: OrderService) {}

  ngOnInit(): void {
    // Fetch all orders from the service
    this.orderService.getAllOrders().subscribe(
      data => {
        this.orders = data;
      },
      error => {
        console.error('Error fetching all orders', error);
      }
    );
  }

  nextOrder(): void {
    if (this.currentOrderIndex < this.orders.length - 1) {
      this.currentOrderIndex++;
    }
  }

  previousOrder(): void {
    if (this.currentOrderIndex > 0) {
      this.currentOrderIndex--;
    }
  }

  getCurrentOrder() {
    return this.orders[this.currentOrderIndex];
  }
}
