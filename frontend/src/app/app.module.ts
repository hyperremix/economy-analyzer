import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { ChartsModule } from 'ng2-charts';

import { AppComponent } from './app.component';
import { MonthlyDetailComponent } from './components/monthly/monthly-detail.component'
import { MonthlyDoughnutChartComponent } from './components/monthly/monthly-doughnut-chart.component'

import { MonthlyService } from './components/monthly/monthly.service';

import { CurrencyPipe } from './pipes/currencyPipe'
import { MonthPipe } from './pipes/monthPipe'
import { DatePipe } from './pipes/datePipe'

@NgModule({
  declarations: [
    AppComponent,
    MonthlyDetailComponent,
    MonthlyDoughnutChartComponent,
    CurrencyPipe,
    MonthPipe,
    DatePipe
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    ChartsModule
  ],
  exports: [
    CurrencyPipe,
    MonthPipe,
    DatePipe
  ],
  providers: [ MonthlyService ],
  bootstrap: [ AppComponent ]
})
export class AppModule { }
