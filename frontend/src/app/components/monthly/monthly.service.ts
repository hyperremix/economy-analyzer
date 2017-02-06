import { Injectable } from '@angular/core';
import { Monthly } from './../../model/monthly'
import { Headers, Http } from '@angular/http'

import 'rxjs/add/operator/toPromise';

@Injectable()
export class MonthlyService {

    private monthlyurl = 'http://localhost:3000/api/monthlies';

    constructor(private http: Http) { }

    getMonthlies(): Promise<Monthly[]> {
        return this.http.get(this.monthlyurl).toPromise().then(this.handleSuccess).catch(this.handleError);
    }

    private handleSuccess(response: any): Monthly[] {
        return response.json() as Monthly[];
    }

    private handleError(error: any): Promise<any> {
        console.error('An error occurred', error); // for demo purposes only
        return Promise.reject(error.message || error);
    }
}