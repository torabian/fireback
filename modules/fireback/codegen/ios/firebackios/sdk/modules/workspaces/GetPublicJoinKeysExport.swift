import Promises
import Combine
import SwiftUI
func GetPublicJoinKeysExportFetcher() -> AnyPublisher<ArrayResponse<PublicJoinKeyEntity>, Error> {
  var prefix = ""
  if let api_url = ProcessInfo.processInfo.environment["api_url"] {
    prefix = api_url
  }
  let url = URL(string: prefix + "/public-join-keys/export")!
  var request = URLRequest(url: url)
    request.addValue(AuthService.shared.TokenSnapShot, forHTTPHeaderField: "Authorization")
    request.addValue("root", forHTTPHeaderField: "workspace-id")
    return URLSession.shared
        .dataTaskPublisher(for: request)
        .map(\.data)
        .decode(type: ArrayResponse<PublicJoinKeyEntity>.self, decoder: JSONDecoder())
        .receive(on: DispatchQueue.main)
        .eraseToAnyPublisher()
}
/*
I have commented this for now, because this is not returning correctly, as well as the function above is good enough
func GetPublicJoinKeysExport() -> Promise<[PublicJoinKeyEntity]?> {
    return Promise<[PublicJoinKeyEntity]?>(on: .main) { fulfill, reject in
  var prefix = ""
  if let api_url = ProcessInfo.processInfo.environment["api_url"] {
    prefix = api_url
  }
  let url = URL(string: prefix + "/public-join-keys/export")!
  var request = URLRequest(url: url)
        request.httpMethod = "GET"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        URLSession.shared.dataTask(with: request) { data, response, error in
            if let data = data {
                let decoder = JSONDecoder()
                if let str = String(bytes: data, encoding: .utf8) {
                    print(str)
                }
                do {
                    let result = try decoder.decode(SingleResponse<PublicJoinKeyEntity>.self, from: data)
                    if result.error != nil {
                        reject(result.error!)
                    } else {
                        fulfill(result.data)
                    }
                } catch {
                    print(error)
                }
            }
        }.resume()
    }
}
*/