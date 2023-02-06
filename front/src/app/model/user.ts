export class User{
    first_name:string;
    last_name:string;
    age:number;
    country:string;
    email:string;
    password:string;
    gender:string;
    role: string;

    constructor(obj?: any){
        this.first_name = obj && obj.first_name || null;
        this.last_name = obj && obj.last_name || null;
        this.age = obj && obj.age || null;
        this.country = obj && obj.country || null;
        this.email = obj && obj.email || null;
        this.password = obj && obj.password || null;
        this.gender = obj && obj.gender || null;
        this.role = obj && obj.role || null;
    }
}