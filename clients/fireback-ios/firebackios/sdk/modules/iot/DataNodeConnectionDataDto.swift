import Foundation
struct DataNodeConnectionDataDto : Codable {
    var subKey: String? = nil
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
class DataNodeConnectionDataDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var subKey: String? = nil
  @Published var subKeyErrorMessage: String? = nil
  func getDto() -> DataNodeConnectionDataDto {
      var dto = DataNodeConnectionDataDto()
    dto.subKey = self.subKey
      return dto
  }
}
