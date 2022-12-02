import { Component, OnInit } from '@angular/core';
import { Post } from 'src/app/model/post';

@Component({
  selector: 'app-post-component',
  templateUrl: './post-component.component.html',
  styleUrls: ['./post-component.component.css']
})
export class PostComponentComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }
  posts!: Post[];

}
