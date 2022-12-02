import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpInterceptor,
  HttpEvent
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from '../auth.service';
// import { userService } from '../services/user.service';

@Injectable()
export class TokenInterceptor implements HttpInterceptor {
  constructor(public user: AuthService) { }
  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    console.log("Cookie: " + this.user.getCookieToken());

    if (this.user.tokenIsPresent()) {
      request = request.clone({
        setHeaders: {
          Authorization: `${this.user.getCookieToken()}` 
        }
      });
    }
    return next.handle(request);
  }
}