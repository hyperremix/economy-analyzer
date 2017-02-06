import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'currency'
})

export class CurrencyPipe implements PipeTransform {
  transform(amount: number): string {
    var d = Math.pow(10,2);
    return (amount*d/d).toFixed(2) + " €";
  }
}