import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'month'
})

export class MonthPipe implements PipeTransform {
    private months = [
        "January", "February", "March",
        "April", "May", "June", "July",
        "August", "September", "October",
        "November", "December"
    ];

    transform(dateString: string): string {
        var date = new Date(dateString)
        return this.months[date.getMonth()] + " " + date.getFullYear()
    }
}