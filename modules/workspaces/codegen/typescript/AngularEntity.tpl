import { injectQuery, injectMutation } from "@ngneat/query";

{{ template "tsimport" . }}

@Injectable({ providedIn: "root" })
export class {{ .Group}}Rpc {
  #http = inject(HttpClient);
  
  {{ printf "%s" .content }}


}
