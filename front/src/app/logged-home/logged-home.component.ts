import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Post } from '../model/post';

@Component({
  selector: 'app-logged-home',
  templateUrl: './logged-home.component.html',
  styleUrls: ['./logged-home.component.css']
})
export class LoggedHomeComponent implements OnInit {

  posts: Post[] = [];

  constructor(private service: AuthService) { }

  ngOnInit(): void {
    this.service.getHomePageForLoggedUser().subscribe(posts => { this.posts = posts })
  }

}
