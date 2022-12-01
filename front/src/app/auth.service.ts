import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { BuisnessUser } from "./model/businessUser";
import { User } from "./model/user";
import { UserLogin } from "./model/userLogin";

@Injectable({
    providedIn: 'root'
})

export class AuthService {
    
    constructor (private httpClient: HttpClient){}

    signUp(user: User): Observable<User>{
        return this.httpClient.post<User>('https://localhost:8081/auth/register', user);
    }

    signUpBusiness(user: BuisnessUser): Observable<BuisnessUser>{
        return this.httpClient.post<BuisnessUser>('https://localhost:8081/auth/register', user);
    }

    login(user: UserLogin): Observable<any>{
        return this.httpClient.post('https://localhost:8081/auth/login', user, {responseType:'text'});
    }

}