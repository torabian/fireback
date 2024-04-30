
  // Query
  #{{ .r.GetFuncName}}Query = injectQuery();
  {{ .r.GetFuncName}}() {

    const url = "{{ .r.Url}}".substr(1);
    let computedUrl = `${url}?${new URLSearchParams(
      queryBeforeSend(query)
    ).toString()}`;
    
    return this.#{{ .r.GetFuncName}}Query({
      queryKey: ["{{ .r.EntityKey }}"] as const,
      queryFn: () => {
        return this.#http.get<IResponseList<{{ .r.ResponseEntityComputed}}>>(
          computedUrl
        );
      },
    });
  }