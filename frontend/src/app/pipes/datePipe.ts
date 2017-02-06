import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'date'
})

export class DatePipe implements PipeTransform {
    transform(dateString: string): string {
        var date = new Date(dateString)
        return date.toLocaleDateString('de')
    }
}