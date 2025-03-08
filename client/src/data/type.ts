export interface TabsInterface {
    id: number;
    name: string;
    value: string;
}

export interface Address {
    client_id: number
    province:string
    remain_address:string
}

export interface UserInfoInterface {
    addresses: Address[];
    client_id:number
    first_name :string
    last_name :string
    number_of_refererred : number
    phone_number:number
    referral_code : string
    wallet_balance:number
    timestamp:string
}

export interface UserDiscountInterface {
    dicount_code : number[]
    number_of_discount_code : number
}