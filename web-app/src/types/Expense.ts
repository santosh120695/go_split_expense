import type {UserType} from "./user.ts";

export type ExpenseType = {
    group_id: number | string | null;
    title: string;
    amount: number;
    user_ids: number[];
    paid_by?: UserType;
    created_at?: Date;
}