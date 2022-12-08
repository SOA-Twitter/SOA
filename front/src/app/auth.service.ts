import { HttpClient, HttpStatusCode } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Observable } from "rxjs";
import { BuisnessUser } from "./model/businessUser";
import { ChangePassword } from "./model/changePassword";
import { recoverProfile } from "./model/recoverProfile";
import { recoverProfileWithUuid } from "./model/recoverProfileWithUuid";
import { recoveryMail } from "./model/recoveryMail";
import { User } from "./model/user";
import { UserLogin } from "./model/userLogin";
import { uuid } from "./model/uuid";
import jwt_decode, { JwtPayload } from "jwt-decode";
import { JsonPipe } from "@angular/common";
import { MyToken } from "./model/myToken";

@Injectable({
    providedIn: 'root'
})

export class AuthService {
    
    constructor (
        private httpClient: HttpClient,
        private router: Router
        ){}

    private readonly path:string = "https://localhost:8081/auth/";

    signUp(user: User): Observable<HttpStatusCode>{
        return this.httpClient.post<HttpStatusCode>(this.path + 'register', user);
    }

    signUpBusiness(user: BuisnessUser): Observable<BuisnessUser>{
        return this.httpClient.post<BuisnessUser>(this.path + 'register', user);
    }

    login(user: UserLogin): Observable<any>{
        return this.httpClient.post(this.path + 'login', user, {responseType:'text'});
    }

    recoveryMail(recoveryMail: recoveryMail): Observable<recoveryMail>{
        return this.httpClient.post<recoveryMail>(this.path + 'recoverEmail', recoveryMail)
    }

    recoverAccount(newPassword: recoverProfileWithUuid): Observable<any>{
        return this.httpClient.post(this.path + 'recover', newPassword)
    }

    changePassword(changePassword: ChangePassword): Observable<any>{
      return this.httpClient.post<ChangePassword>('https://localhost:8081/profile/changePassword', changePassword)
    }

    get isAuthenticated() {
        return this.getToken();
    }

    logout() {
        localStorage.removeItem('token');
        document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";

        this.router.navigateByUrl('/login');
    }
    
    getUsername() {
      let token = this.parseToken();
        if (token) {
          // return this.parseToken()['sub']
          return this.parseToken()?.email;
        }
      return "";
    }
    
    getRoles() {
      // let token = this.parseToken();
      // if (token) {
      //   return this.parseToken();
      // }
      // return []
      var rola = this.parseToken()?.role;
      return rola;
    }
    
    private parseToken() {
      // NPR. COOKIE JWT VALUE  token=aiwomdwoaimd.awdawdwadwadaw.awdawdwadawda;cookie=wa,da;d,;
      let jwt = this.getToken();
      // alert(jwt);
      if (jwt !== null) {
        var decoded = jwt_decode<MyToken>(jwt);
        console.log(decoded);
        return decoded;
      }
      return
      
    }

    setCookieToken(token: string){
      document.cookie = "token="+token;
    }
    
    tokenIsPresent(): Boolean {
      let token = this.getToken()
      return token != null && token != "" && token != undefined;
    }
    
    // TODO check if it takes token from cookie & if guards work consequently as expected
    getToken() {
      // let token = localStorage.getItem('token');
      let token = this.getCookieToken();
      
      return token
    }

    getCookieToken() {
      // alert("cookie value: " + document.cookie);

      let name = "token" + "=";
      let decodedCookie = decodeURIComponent(document.cookie);
      let ca = decodedCookie.split(';');
      for(let i = 0; i <ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) == ' ') {
          c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
          // alert("C var/ cookie whole: " + c);
          // return c.substring(name.length, c.length);
          return c;
        }
      }
      // DELETE THIS alert
      // alert(this.parseToken());
      return "";
    }

}