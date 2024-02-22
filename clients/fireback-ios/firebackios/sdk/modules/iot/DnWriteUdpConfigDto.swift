import Foundation
struct DnWriteUdpConfigDto : Codable {
    var host: String? = nil
    var port: String? = nil
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
class DnWriteUdpConfigDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var host: String? = nil
  @Published var hostErrorMessage: String? = nil
  @Published var port: String? = nil
  @Published var portErrorMessage: String? = nil
  func getDto() -> DnWriteUdpConfigDto {
      var dto = DnWriteUdpConfigDto()
    dto.host = self.host
    dto.port = self.port
      return dto
  }
}
