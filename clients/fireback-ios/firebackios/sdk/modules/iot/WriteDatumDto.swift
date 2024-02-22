import Foundation
struct WriteDatumDto : Codable {
    var uniqueId: String? = nil
    var key: String? = nil
    var value: String? = nil
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
class WriteDatumDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var uniqueId: String? = nil
  @Published var uniqueIdErrorMessage: String? = nil
  @Published var key: String? = nil
  @Published var keyErrorMessage: String? = nil
  func getDto() -> WriteDatumDto {
      var dto = WriteDatumDto()
    dto.uniqueId = self.uniqueId
    dto.key = self.key
      return dto
  }
}
