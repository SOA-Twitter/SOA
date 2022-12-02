export class recoveryMail{
    email:string;
    constructor(obj?: any){
        this.email = obj && obj.email || null;
    }
}