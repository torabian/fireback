import Foundation
struct DnReadSerialPortConfigDto : Codable {
    var address: String? = nil
    var baudRate: Int? = nil
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
class DnReadSerialPortConfigDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var address: String? = nil
  @Published var addressErrorMessage: String? = nil
  @Published var baudRate: Int? = nil
  @Published var baudRateErrorMessage: Int? = nil
  func getDto() -> DnReadSerialPortConfigDto {
      var dto = DnReadSerialPortConfigDto()
    dto.address = self.address
    dto.baudRate = self.baudRate
      return dto
  }
}
