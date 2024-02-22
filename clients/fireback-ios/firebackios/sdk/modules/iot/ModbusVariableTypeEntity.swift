import Foundation
class ModbusVariableTypeEntity : Codable, Identifiable {
    var name: String? = nil
}
class ModbusVariableTypeEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> ModbusVariableTypeEntity {
      var dto = ModbusVariableTypeEntity()
    dto.name = self.name
      return dto
  }
}