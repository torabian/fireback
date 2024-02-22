import Foundation
class ModbusConnectionTypeEntity : Codable, Identifiable {
    var name: String? = nil
}
class ModbusConnectionTypeEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> ModbusConnectionTypeEntity {
      var dto = ModbusConnectionTypeEntity()
    dto.name = self.name
      return dto
  }
}