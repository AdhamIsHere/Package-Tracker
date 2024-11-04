import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class OrderService {
  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  getOrders(): Observable<any[]> {
    // add token from local storage to the request header
    const token = localStorage.getItem('token');
    return this.http.get<any[]>(`${this.apiUrl}/order/myorders`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
  }

  getItems(): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/items`);
  }

  submitOrder(orderDetails: any): Observable<any> {
    const token = localStorage.getItem('token');
    console.log('Order Details:', orderDetails);
    console.log('Token:', token);
    return this.http.post<any>(`${this.apiUrl}/order/create`, orderDetails, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
  }

  getOrderById(orderId: number): Observable<any> {
    const token = localStorage.getItem('token');
    return this.http.get<any>(`${this.apiUrl}/order/view?id=${orderId}`,{headers: {
        Authorization: `Bearer ${token}`
      }});
  }

  cancelOrder(orderId: number): Observable<any> {
    const token = localStorage.getItem('token');
    return this.http.delete<any>(`${this.apiUrl}/order/delete?id=${orderId}`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
  }

  getAllOrders(): Observable<any[]> {
    const token = localStorage.getItem('token');
    return this.http.get<any[]>(`${this.apiUrl}/order/viewall`,{headers: {
      Authorization: `Bearer ${token}`
    }}
    );
  }

}
