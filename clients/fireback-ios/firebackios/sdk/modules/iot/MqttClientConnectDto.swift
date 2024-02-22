import Foundation
struct MqttClientConnectDto : Codable {
    var connect: Bool? = nil
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
class MqttClientConnectDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var connect: Bool? = nil
  @Published var connectErrorMessage: Bool? = nil
  func getDto() -> MqttClientConnectDto {
      var dto = MqttClientConnectDto()
    dto.connect = self.connect
      return dto
  }
}
