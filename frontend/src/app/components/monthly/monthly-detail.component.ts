import { Component, Input, OnInit } from '@angular/core';

import { Monthly } from './../../model/monthly'
import { Transaction } from './../../model/transaction'

@Component({
  selector: 'monthly-detail',
  templateUrl: './monthly-detail.component.html'
})

export class MonthlyDetailComponent implements OnInit {
  @Input()
  monthly: Monthly;

  classifiedTransactions: Array<any>;
  classificationSums: Array<any>;
  balance: number;

  ngOnInit() {
    this.classifiedTransactions = [];
    this.classificationSums = [];
    this.balance = this.monthly.LastMonthBalance;

    for (var key in this.monthly.ClassifiedTransactions) {
      if (this.monthly.ClassifiedTransactions.hasOwnProperty(key)) {
        var transactions = this.monthly.ClassifiedTransactions[key] as Transaction[]

        var sum = 0;
        for (var transaction of transactions) {
          sum += transaction.Amount;

          this.classifiedTransactions.push({
            Classification: key,
            Amount: transaction.Amount,
            Client: transaction.Client,
            Purpose: transaction.Purpose,
            Date: transaction.Date
          });
        }

        this.balance += sum;
        this.classificationSums.push({Classification: key, Sum: sum})
      }
    }
  }
}