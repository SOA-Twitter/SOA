export class recoverProfile{
    new_password:string;
    repeated_password: string;
    constructor(obj?: any){
        this.new_password = obj && obj.new_password || null;
        this.repeated_password = obj && obj.repeated_password || null;
    }
}