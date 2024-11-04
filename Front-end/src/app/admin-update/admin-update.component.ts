import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { OrderService } from "../services/order.service";

@Component({
  selector: 'app-admin-update',
  templateUrl: './admin-update.component.html',
  styleUrls: ['./admin-update.component.css']
})
export class AdminUpdateComponent implements OnInit {
  // Order properties
  pickupLocation: string = '';
  deliveryLocation: string = '';
  deliveryTime: string = '';
  selectedItems: any[] = [];
  items: any[] = [];
  successMessage: string = '';
  errorMessage: string = '';
  orderId: number = 0;

  constructor(
    private orderService: OrderService,
    private route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    this.loadOrders();

    this.route.paramMap.subscribe(params => {
      this.orderId = Number(params.get('orderId'));
    });
  }
  loadOrders() {
    this.orderService.getAllOrders().subscribe(
      data => {
        // find order by id
        let order = data.find((order) => order.id == this.orderId);
        if (order) {
          this.pickupLocation = order.pickup_location;
          this.deliveryLocation = order.delivery_location;
          this.deliveryTime = order.delivery_time;
          this.selectedItems = order.items;
        } else {
          this.errorMessage = 'Order not found';
        }
      },
      error => {
        this.errorMessage = 'Error fetching order';
        console.error('Error fetching orders', error);
      }
    );
  }

  // Load order details along with available items for selection
  loadOrderDetails(orderId: number): void {
    this.orderService.getOrderById(orderId).subscribe(
      data => {
        // Assuming data contains 'order' with order details and 'availableItems' with all items
        this.pickupLocation = data.order.pickup_location;
        this.deliveryLocation = data.order.delivery_location;
        this.deliveryTime = data.order.delivery_time;
        this.selectedItems = data.order.items;
        this.items = data.availableItems; // Populate available items for the item selection dropdowns
      },
      error => {
        this.errorMessage = 'Error fetching order details';
        console.error('Error fetching order details', error);
      }
    );
  }

  // Add a new item slot to selectedItems array for adding another item to the order
  addItem(): void {
    this.selectedItems.push({ name: '' }); // Adds a new item slot with an empty name
  }

  // Remove an item from selectedItems array based on index
  removeItem(index: number): void {
    if (index > -1) {
      this.selectedItems.splice(index, 1);
    }
  }

  // Submit the updated order details
  submitOrder(): void {
    const updatedOrder = {
      pickup_location: this.pickupLocation,
      delivery_location: this.deliveryLocation,
      delivery_time: this.deliveryTime,
      items: this.selectedItems
    };

    // Call the service method to update the order
    this.orderService.updateOrder(this.orderId, updatedOrder).subscribe(
      response => {
        this.successMessage = 'Order updated successfully';
        this.errorMessage = '';
      },
      error => {
        this.errorMessage = 'Error updating order';
        this.successMessage = '';
        console.error('Error updating order', error);
      }
    );
  }
}
