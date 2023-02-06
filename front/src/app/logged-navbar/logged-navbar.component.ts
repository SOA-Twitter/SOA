import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-logged-navbar',
  templateUrl: './logged-navbar.component.html',
  styleUrls: ['./logged-navbar.component.css']
})
export class LoggedNavbarComponent implements OnInit {

  username!: string;

  constructor(private authService: AuthService, private router: Router) { }

  ngOnInit(): void {
  }

  profile(){
    this.username = this.authService.getUsername();

    this.router.navigateByUrl("/profile/" + this.username);
  }

  logout() {
    localStorage.removeItem('jwt token')
    this.router.navigateByUrl('/login');
  }

}
