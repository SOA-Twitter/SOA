export class recoverProfileWithUuid{
    new_password:string;
    repeated_password: string;
    recovery_uuid: string;
    constructor(obj?: any){
        this.new_password = obj && obj.new_password || null;
        this.repeated_password = obj && obj.repeated_password || null;
        this.recovery_uuid = obj && obj.recovery_uuid || null;
    }
}