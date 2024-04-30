import { injectQuery } from "@ngneat/query";

@Injectable({ providedIn: "root" })
export class TodosService {
  #http = inject(HttpClient);
  #query = injectQuery();

  getTodos() {
    return this.#query({
      queryKey: ["todos"] as const,
      queryFn: () => {
        return this.#http.get<Todo[]>(
          "https://jsonplaceholder.typicode.com/todos"
        );
      },
    });
  }
}
