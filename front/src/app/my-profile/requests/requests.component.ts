import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
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
  @Output()
  handledFollowRequest: EventEmitter<any>;

  constructor(private service: AuthService) {
    this.handledFollowRequest = new EventEmitter<string>();
  }

  ngOnInit(): void {
  }

  get username1() {
    let username2 = this.username;
    return username2;
  }

  protected accept(username: string) {
    this.service.acceptReq(username).subscribe(
      () => {
        this.handledFollowRequest.emit(username);
      },
      (error) => {

      }
    )
  }

  protected decline(username: string) {
    this.service.decline(username).subscribe(
      () => {
        this.handledFollowRequest.emit(username);
      },
      (error) => {

      })
  }

}
