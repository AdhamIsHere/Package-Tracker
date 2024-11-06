import {Component, OnInit} from '@angular/core';
import {OrderService} from "../services/order.service";

@Component({
  selector: 'app-courier-update',
  templateUrl: './courier-update.component.html',
  styleUrls: ['./courier-update.component.css']
})
export class CourierUpdateComponent implements OnInit {

  constructor(private orderService:OrderService) { }

  ngOnInit(): void {
  }

  logout(){
    localStorage.removeItem('token');
    window.location.href='/';
  }

}
