import { Component, OnInit } from '@angular/core';
import { OrderService } from '../services/order.service';

@Component({
  selector: 'app-order-page',
  templateUrl: './order-page.component.html',
  styleUrls: ['./order-page.component.css']
})
export class OrderPageComponent implements OnInit {
  orders: any[] = [];
  pickupLocation: string = '';
  deliveryLocation: string = '';
  deliveryTime: string = '';
  successMessage: string = ''; // Add this property to hold success message

  constructor(private orderService: OrderService) {}

  ngOnInit(): void {
    this.loadOrders();
  }

  loadOrders(): void {
    this.orderService.getOrders().subscribe(
      data => {
        this.orders = data;
      },
      error => {
        console.error('Error fetching orders', error);
      }
    );
  }

  submitOrder(): void {
    if (this.pickupLocation && this.deliveryLocation && this.deliveryTime) {
      const orderDetails = {
        pickup: this.pickupLocation,
        location: this.deliveryLocation,
        time: this.deliveryTime
      };

      console.log('Order Submitted:', orderDetails);

      // Simulating order submission and setting success message
      this.orderService.submitOrder(orderDetails).subscribe(response => {
        console.log('Order successfully submitted', response);
        this.successMessage = 'Order added successfully!'; // Set success message here

        // Optionally, reset form fields after submission
        this.pickupLocation = '';
        this.deliveryLocation = '';
        this.deliveryTime = '';
      }, error => {
        console.error('Error submitting order', error);
      });
    } else {
      console.error('Please fill in all fields');
    }
  }
}
