import Foundation
struct DnInterpolateConfigDto : Codable {
    var sources: [Float64]? = nil
    var targets: [Float64]? = nil
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
class DnInterpolateConfigDtoViewModel: ObservableObject {
  // improve the fields here
  func getDto() -> DnInterpolateConfigDto {
      var dto = DnInterpolateConfigDto()
      return dto
  }
}
