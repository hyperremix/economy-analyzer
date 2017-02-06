import { Pipe, PipeTransform } from '@angular/core';
import { Transaction } from './transaction'

@Pipe({
  name: 'mapToIterable'
})

export class MapToIterablePipe implements PipeTransform {
  transform(dict: Object): Array<any> {
    var a = [];
    for (var key in dict) {
      if (dict.hasOwnProperty(key)) {
        var transactions = dict[key] as Transaction[]
        for (var transaction of transactions) {
          var classifiedTransaction = {
            Classification: key,
            Amount: transaction.Amount,
            Client: transaction.Client,
            Purpose: transaction.Purpose,
            Date: transaction.Date
          };
          a.push(classifiedTransaction);
        }
        
      }
    }
    return a;
  }
}