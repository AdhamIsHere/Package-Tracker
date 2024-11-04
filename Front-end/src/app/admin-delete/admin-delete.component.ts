import {Component, OnInit} from '@angular/core';
import {OrderService} from "../services/order.service";

@Component({
  selector: 'app-admin-delete',
  templateUrl: './admin-delete.component.html',
  styleUrls: ['./admin-delete.component.css']
})
export class AdminDeleteComponent implements OnInit {
  orders: any[] = [];
  successMessage: string = '';
  errorMessage: string = '';
  constructor(private orderService: OrderService) {}

  ngOnInit(): void {
    this.loadOrders();
  }
  deleteOrder(orderId: number): void {
    this.orderService.deleteOrder(orderId).subscribe(
      data => {
        this.successMessage = 'Order deleted successfully';
        this.loadOrders();
      },
      error => {
        this.errorMessage = 'Error deleting order';
        console.error('Error deleting order', error);
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
