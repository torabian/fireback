import Foundation
struct DataNodeReaderConfigDto : Codable {
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
class DataNodeReaderConfigDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var interval: Int? = nil
  @Published var intervalErrorMessage: Int? = nil
  func getDto() -> DataNodeReaderConfigDto {
      var dto = DataNodeReaderConfigDto()
    dto.interval = self.interval
      return dto
  }
}
