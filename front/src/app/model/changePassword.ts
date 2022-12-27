export class ChangePassword{
    old_password: string;
    new_password: string;
    repeated_password: string;

    constructor(obj?: any){
        this.old_password = obj && obj.old_password || null;
        this.new_password = obj && obj.new_password || null;
        this.repeated_password = obj && obj.repeated_password || null;
    }
}
