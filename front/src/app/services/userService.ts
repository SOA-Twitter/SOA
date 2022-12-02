import { Injectable } from "@angular/core";
import { Route, Router } from "@angular/router";
import { AuthService } from "../auth.service";


@Injectable({
    providedIn: 'root'
})
export class userService{

    constructor(
        private authService: AuthService,
        private router: Router
    ){}

    get isAuthenticated() {
        return localStorage.getItem('token');
    }

    logout() {
        localStorage.removeItem('token')
        this.router.navigateByUrl('/login');
    }
    
    getUsername(): string {
      let token = this.parseToken();
        if (token) {
          return this.parseToken()['sub']
        }
      return "";
    }
    
    getRoles() {
      let token = this.parseToken();
      if (token) {
        return this.parseToken()['role'];
      }
      return []
    }
    
    private parseToken() {
      let jwt = localStorage.getItem('token');
      if (jwt !== null) {
        let jwtData = jwt.split('.')[1];
        let decodedJwtJsonData = atob(jwtData);
        let decodedJwtData = JSON.parse(decodedJwtJsonData);
        return decodedJwtData;
      }
    }
    
    tokenIsPresent(): Boolean {
      let token = this.getToken()
      return token != null;
    }
    
    getToken() {
      let token = localStorage.getItem('token');
      return token
    }

}