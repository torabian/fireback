import Promises
func PostWorkspaceInvite(dto: WorkspaceInviteEntity) -> Promise<WorkspaceInviteEntity?> {
    return Promise<WorkspaceInviteEntity?>(on: .main) { fulfill, reject in
  guard let encoded = try? JSONEncoder().encode(dto) else {
    print("Failed to encode login request")
    return
  }
  var prefix = ""
  if let api_url = ProcessInfo.processInfo.environment["api_url"] {
    prefix = api_url
  }
  let url = URL(string: prefix + "//workspace/invite")!
  var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        request.httpBody = encoded
        URLSession.shared.dataTask(with: request) { data, response, error in
            if let data = data {
                let decoder = JSONDecoder()
                if let str = String(bytes: data, encoding: .utf8) {
                    print(str)
                }
                do {
                    let result = try decoder.decode(SingleResponse<WorkspaceInviteEntity>.self, from: data)
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