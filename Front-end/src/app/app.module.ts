import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { SignUpComponent } from './sign-up/sign-up.component';
import { AuthService } from './services/auth.service';
import { SplashScreenComponent } from './splash-screen/splash-screen.component';
import { HomePageComponent } from './home-page/home-page.component';
import { OrderPageComponent } from './order-page/order-page.component';
import { MyOrdersComponent } from './my-orders/my-orders.component';
import { OrderDetailsComponent } from './order-details/order-details.component';
import { CourierAssignComponent } from './courier-assign/courier-assign.component';
import { CourierUpdateComponent } from './courier-update/courier-update.component';
import { AdminAssignComponent } from './admin-assign/admin-assign.component';
import { AdminUpdateComponent } from './admin-update/admin-update.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    SignUpComponent,
    SplashScreenComponent,
    HomePageComponent,
    OrderPageComponent,
    MyOrdersComponent,
    OrderDetailsComponent,
    CourierAssignComponent,
    CourierUpdateComponent,
    AdminAssignComponent,
    AdminUpdateComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [AuthService],
  bootstrap: [AppComponent]
})
export class AppModule {}
