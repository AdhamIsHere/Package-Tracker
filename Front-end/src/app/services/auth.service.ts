import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private apiUrl = 'YOUR_API_ENDPOINT';  // Define your backend API

  constructor(private http: HttpClient) {}

  login(email: string, password: string): Observable<any> {
    return this.http.post(`${this.apiUrl}/login`, { email, password });
  }

  signUp(name: string, email: string, password: string, phone: string, role: string): Observable<any> {
    return this.http.post(`${this.apiUrl}/signup`, { name, email, password, phone, role });
  }
}
