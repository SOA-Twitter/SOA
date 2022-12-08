export class BuisnessUser{
    email:string;
    password:string;
    company_name:string;
    company_website:string;
    username:string;

    constructor(obj?: any){
        this.email = obj && obj.email || null;
        this.password = obj && obj.password || null;
        this.company_name = obj && obj.company_name || null;
        this.company_website = obj && obj.company_website || null;
        this.username = obj && obj.username || null;
    }
}