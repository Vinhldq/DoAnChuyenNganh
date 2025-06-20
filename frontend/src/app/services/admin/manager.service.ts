import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { Observable } from 'rxjs';
import { CreateManager, CreateManagerOutput } from '../../models/admin/manager.model';

@Injectable({
  providedIn: 'root'
})
export class ManagerService {
  private apiUrl = `${environment.apiUrl}/manager`;

  constructor(private http: HttpClient) { }

  createNewManager(manager: CreateManager): Observable<CreateManagerOutput> {
    return this.http.post<CreateManagerOutput>(this.apiUrl + '/register', manager);
  } 
}
