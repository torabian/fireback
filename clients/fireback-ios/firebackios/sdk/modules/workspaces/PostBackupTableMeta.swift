import Promises
func PostBackupTableMeta(dto: BackupTableMetaEntity) -> Promise<BackupTableMetaEntity?> {
    return Promise<BackupTableMetaEntity?>(on: .main) { fulfill, reject in
  var prefix = ""
  if let api_url = ProcessInfo.processInfo.environment["api_url"] {
    prefix = api_url
  }
  let url = URL(string: prefix + "/backup-table-meta")!
  var request = URLRequest(url: url)
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
                    let result = try decoder.decode(SingleResponse<BackupTableMetaEntity>.self, from: data)
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