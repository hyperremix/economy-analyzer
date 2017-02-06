import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'monthly-doughnut-chart',
  templateUrl: './monthly-doughnut-chart.component.html'
})

export class MonthlyDoughnutChartComponent implements OnInit {
  @Input() classificationSums: Array<any>;
  @Input() balance: number;

  labels: string[];
  data: number[];
  type: string;

  ngOnInit() {
    this.labels = ["Balance"];
    this.data = [this.balance];
    this.type = 'doughnut';

    for (var classifiedSum of this.classificationSums) {
        if (this.isNotIncome(classifiedSum.Classification)) {
            this.labels.push(classifiedSum.Classification);
            this.data.push(Math.abs(classifiedSum.Sum));
        }
    }
  }

  private isNotIncome(classification: string): boolean {
    return classification != "Income" &&
        classification != "Salary" &&
        classification != "Unclassified" &&
        classification != "TAX"
  }
}