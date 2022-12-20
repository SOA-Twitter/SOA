import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-logged-navbar',
  templateUrl: './logged-navbar.component.html',
  styleUrls: ['./logged-navbar.component.css']
})
export class LoggedNavbarComponent implements OnInit {

  constructor(private authService: AuthService, private router: Router) { }

  ngOnInit(): void {
  }

  profile(){
    if(this.authService.getRoles() == "BusinessUser"){
      this.router.navigateByUrl("/my-profile-business")
    }else{
      this.router.navigateByUrl("/my-profile")
    }
  }

  logout() {
    localStorage.removeItem('jwt token')
    this.router.navigateByUrl('/login');
  }

}
