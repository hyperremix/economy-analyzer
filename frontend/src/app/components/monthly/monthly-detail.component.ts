import { Component, Input } from '@angular/core';

import { Monthly } from './monthly'

@Component({
  selector: 'monthly-detail',
  templateUrl: './monthly-detail.component.html'
})

export class MonthlyDetailComponent {
    @Input()
    monthly: Monthly;
}