import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, RouterStateSnapshot, UrlTree } from '@angular/router';
import { AuthService } from '../auth.service';

@Injectable({
  providedIn: 'root'
})
export class Guard implements CanActivate {

  constructor(private auth:AuthService){} 

  canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): boolean {
      if(this.auth.tokenIsPresent()){
        if(this.auth.getRoles() === "RegularUser" || "BusinessUser"){
          return true;
        }
      }
      return false;
  }
  
}
