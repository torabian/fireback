import Foundation
struct DnModbusRtuConfigDto : Codable {
    var baudRate: Int? = nil
    var dataBits: Int? = nil
    var parity: String? = nil
    var stopBits: Int? = nil
    var slaveId: Int? = nil
    var timeout: Int? = nil
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
class DnModbusRtuConfigDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var baudRate: Int? = nil
  @Published var baudRateErrorMessage: Int? = nil
  @Published var dataBits: Int? = nil
  @Published var dataBitsErrorMessage: Int? = nil
  @Published var parity: String? = nil
  @Published var parityErrorMessage: String? = nil
  @Published var stopBits: Int? = nil
  @Published var stopBitsErrorMessage: Int? = nil
  @Published var slaveId: Int? = nil
  @Published var slaveIdErrorMessage: Int? = nil
  @Published var timeout: Int? = nil
  @Published var timeoutErrorMessage: Int? = nil
  @Published var address: String? = nil
  @Published var addressErrorMessage: String? = nil
  func getDto() -> DnModbusRtuConfigDto {
      var dto = DnModbusRtuConfigDto()
    dto.baudRate = self.baudRate
    dto.dataBits = self.dataBits
    dto.parity = self.parity
    dto.stopBits = self.stopBits
    dto.slaveId = self.slaveId
    dto.timeout = self.timeout
    dto.address = self.address
      return dto
  }
}
