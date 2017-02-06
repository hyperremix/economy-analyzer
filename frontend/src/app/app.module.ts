import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { MonthlyDetailComponent } from './components/monthly/monthly-detail.component'

import { MonthlyService } from './components/monthly/monthly.service';

import { MapToIterablePipe } from './pipes/mapToIterablePipe'
import { CurrencyPipe } from './pipes/currencyPipe'
import { MonthPipe } from './pipes/monthPipe'
import { DatePipe } from './pipes/datePipe'

@NgModule({
  declarations: [
    AppComponent,
    MonthlyDetailComponent,
    MapToIterablePipe,
    CurrencyPipe,
    MonthPipe,
    DatePipe
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule
  ],
  exports: [
    MapToIterablePipe,
    CurrencyPipe,
    MonthPipe,
    DatePipe
  ],
  providers: [ MonthlyService ],
  bootstrap: [ AppComponent ]
})
export class AppModule { }
