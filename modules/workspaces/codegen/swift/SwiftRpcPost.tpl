import Promises

func {{ .r.GetFuncNameUpper}}(dto: {{ .r.RequestEntityComputed}}) -> Promise<{{ .r.ResponseEntityComputed}}?> {
    
    return Promise<{{ .r.ResponseEntityComputed}}?>(on: .main) { fulfill, reject in

        {{ template "rpcActionCommon" .r }}

        request.httpMethod = "POST"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")

        guard let encoded = try? JSONEncoder().encode(dto) else {
            print("Failed to encode login request")
            return
        }

        request.httpBody = encoded

        URLSession.shared.dataTask(with: request) { data, response, error in
            if let data = data {
                let decoder = JSONDecoder()
                
                if let str = String(bytes: data, encoding: .utf8) {
                    print(str)
                }
                
                do {

                    let result = try decoder.decode(SingleResponse<{{ .r.ResponseEntityComputed}}>.self, from: data)
                    if result.error != nil {
                        reject(result.error!)
                    } else {
                        fulfill(result.data)
                    }
                } catch {
                    let errorCast = IResponseError(message: "Unknown error", messageTranslated: "Unknown error")
                    reject(errorCast)
                }
            }

            if let error = error {
                let message = handleFailedRequestError(error: error)
                let errorCast = IResponseError(message: message, messageTranslated: message)
                reject(errorCast)
            }
        }.resume()
        
    }
     
}

