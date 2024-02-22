import Foundation
struct IntervalNodeConfigDto : Codable {
    var interval: Int? = nil
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
class IntervalNodeConfigDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var interval: Int? = nil
  @Published var intervalErrorMessage: Int? = nil
  func getDto() -> IntervalNodeConfigDto {
      var dto = IntervalNodeConfigDto()
    dto.interval = self.interval
      return dto
  }
}
