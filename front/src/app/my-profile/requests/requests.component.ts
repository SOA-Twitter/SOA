import { Component, Input, OnInit } from '@angular/core';
import { AuthService } from 'src/app/auth.service';
import { Request } from 'src/app/model/request';

@Component({
  selector: 'app-requests',
  templateUrl: './requests.component.html',
  styleUrls: ['./requests.component.css']
})
export class RequestsComponent implements OnInit {

  @Input()
  username!: Request;

  constructor(private service: AuthService) { }

  ngOnInit(): void {
  }

  get username1(){
    console.log(this.username)
    return this.username;
  }

  accept(username: string){
    this.service.acceptReq(username).subscribe()
  }

  decline(username: string){
    this.service.decline(username).subscribe()
  }

}
