
  // Single Post
  #{{ .r.GetFuncName}}Mutation = injectMutation();
  {{ .r.GetFuncName}}() {
    const url = "{{ .r.Url}}".substr(1);
    let computedUrl = `${url}?${new URLSearchParams(
      queryBeforeSend(query)
    ).toString()}`;
    
    return this.#{{ .r.GetFuncName}}Mutation({
      mutationFn: (body: Partial<{{ .r.RequestEntityComputed}}>) =>
        this.#http.post<IResponse<{{ .r.ResponseEntityComputed}}>>(computedUrl, body),
    });
  }