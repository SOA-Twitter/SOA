import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { TokenInterceptor } from './app/interceptor/TokenInterceptor';

export const httpInterceptorProviders = [

  { provide: HTTP_INTERCEPTORS, useClass: TokenInterceptor, multi: true },


];