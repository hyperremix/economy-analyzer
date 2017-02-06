import { Component } from '@angular/core';

import { Monthly } from './components/monthly/monthly'
import { MonthlyService } from './components/monthly/monthly.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  monthlies: Monthly[];

  constructor(private monthlyService: MonthlyService) {
    this.getMonthlies();
  }

  getMonthlies(): void {
    this.monthlyService.getMonthlies().then(monthlies => this.monthlies = monthlies);
  }
}
