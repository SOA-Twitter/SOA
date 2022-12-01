export class User{
    firstName:string;
    lastName:string;
    age:string;
    country:string;
    email:string;
    password:string;
    gender:string;

    constructor(obj?: any){
        this.firstName = obj && obj.firstName || null;
        this.lastName = obj && obj.lastName || null;
        this.age = obj && obj.age || null;
        this.country = obj && obj.country || null;
        this.email = obj && obj.email || null;
        this.password = obj && obj.password || null;
        this.gender = obj && obj.gender || null;
    }
}