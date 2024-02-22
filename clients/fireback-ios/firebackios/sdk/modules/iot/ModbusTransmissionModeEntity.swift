import Foundation
class ModbusTransmissionModeEntity : Codable, Identifiable {
    var name: String? = nil
}
class ModbusTransmissionModeEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> ModbusTransmissionModeEntity {
      var dto = ModbusTransmissionModeEntity()
    dto.name = self.name
      return dto
  }
}