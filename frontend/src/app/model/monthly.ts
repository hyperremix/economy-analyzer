import {Transaction} from './transaction'

export class Monthly {
    ClassifiedTransactions: {[classificationType: string]: Transaction[];}
    Month: Date;
    LastMonthBalance: number;
}