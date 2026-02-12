import type {UserType} from "./user.ts";
import type {ExpenseType} from "./Expense.ts";

export type GroupType = {
    ID: number | null;
    name: string;
    description: string;
    members: UserType[];
    total_expense: number;
    icon: string;
    currency: string;
    expenses?: ExpenseType[];
}
