import Foundation
struct FakeIotEnvDto : Codable {
    var core1temperature: Float64? = nil
    var core2temperature: Float64? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class FakeIotEnvDtoViewModel: ObservableObject {
  // improve the fields here
  func getDto() -> FakeIotEnvDto {
      var dto = FakeIotEnvDto()
      return dto
  }
}
