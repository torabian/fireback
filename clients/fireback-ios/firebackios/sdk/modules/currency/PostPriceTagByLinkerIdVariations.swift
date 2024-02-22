import Promises
func PostPriceTagByLinkerIdVariations(dto: PriceTagVariations) -> Promise<PriceTagVariations?> {
    return Promise<PriceTagVariations?>(on: .main) { fulfill, reject in
        guard let encoded = try? JSONEncoder().encode(dto) else {
            print("Failed to encode login request")
            return
        }
        let url = URL(string: "http://localhost:61901/price-tag/:linkerId/variations")!
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
                    let result = try decoder.decode(SingleResponse<PriceTagVariations>.self, from: data)
                    fulfill(result.data)
                } catch {
                    print(error)
                }
            }
        }.resume()
    }
}