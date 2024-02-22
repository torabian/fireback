import Foundation
struct SerialPortDto : Codable {
    var address: String? = nil
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
class SerialPortDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var address: String? = nil
  @Published var addressErrorMessage: String? = nil
  func getDto() -> SerialPortDto {
      var dto = SerialPortDto()
    dto.address = self.address
      return dto
  }
}
