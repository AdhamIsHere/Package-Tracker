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
  successMessage: string = '';
  errorMessage: string = ''; // Add an error message property
  items: { id: string; name: string; quantity: number; selected: boolean }[] = [
    { id: '1', name: 'Laptop', quantity: 1, selected: false },
    { id: '2', name: 'Phone', quantity: 1, selected: false },
    { id: '3', name: 'Computer', quantity: 1, selected: false },
    { id: '4', name: 'Camera', quantity: 1, selected: false }
  ];

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
      // Get selected items to include in the order
      const selectedItems = this.items
        .filter(item => item.selected)
        .map(item => ({
          id: item.id,
          name: item.name,
          quantity: item.quantity
        }));

      if (selectedItems.length === 0) {
        this.errorMessage = 'Please select at least one item.';
        return;
      }

      const orderDetails = {
        pickup_location: this.pickupLocation,
        dropoff_location: this.deliveryLocation,
        delivery_time: this.deliveryTime,
        items: selectedItems
      };

      console.log('Order Submitted:', orderDetails);

      // Submit the order and handle the response
      this.orderService.submitOrder(orderDetails).subscribe(
        response => {
          console.log('Order successfully submitted', response);
          this.successMessage = 'Order added successfully!';
          this.errorMessage = ''; // Clear any previous error

          // Reset form fields after submission
          this.pickupLocation = '';
          this.deliveryLocation = '';
          this.deliveryTime = '';
          this.items.forEach(item => (item.selected = false)); // Uncheck all items
        },
        error => {
          console.error('Error submitting order', error);
          this.errorMessage = 'Failed to submit the order. Please try again.';
        }
      );
    } else {
      this.errorMessage = 'Please fill in all fields';
    }
  }
}
