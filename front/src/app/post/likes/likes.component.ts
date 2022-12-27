import { Component, Input, OnInit } from '@angular/core';
import { AuthService } from 'src/app/auth.service';
import { Like } from 'src/app/model/like';
import { Post } from 'src/app/model/post';

@Component({
  selector: 'app-likes',
  templateUrl: './likes.component.html',
  styleUrls: ['./likes.component.css']
})
export class LikesComponent implements OnInit {

  likes: Like[] = [];

  @Input()
  post!:Post;

  @Input()
  l!:Like;

  constructor(private authService:AuthService) { }

  ngOnInit(): void {}

  get post1(){
    return this.post;
  }

  get like1(){
    return this.l;
  }
}
