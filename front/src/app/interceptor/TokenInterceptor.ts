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

    if (this.user.tokenIsPresent()) {
      request = request.clone({
        setHeaders: {
          'Authorization': `${this.user.getCookieToken()}`,
          'Accept': 'application/json',
          'Content-Type': 'application/json',
          "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept",
          'Access-Control-Allow-Origin': '*'
        }
      });
    }
    return next.handle(request);
  }
}