import Promises
func PostHmiByLinkerIdComponents(dto: HmiComponents) -> Promise<HmiComponents?> {
    return Promise<HmiComponents?>(on: .main) { fulfill, reject in
        guard let encoded = try? JSONEncoder().encode(dto) else {
            print("Failed to encode login request")
            return
        }
        let url = URL(string: "http://localhost:61901/hmi/:linkerId/components")!
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
                    let result = try decoder.decode(SingleResponse<HmiComponents>.self, from: data)
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