import Foundation
class ModbusTaskEntity : Codable, Identifiable {
    var name: String? = nil
    var modbusId: Int? = nil
    var device: DeviceEntity? = nil
    // var deviceId: String? = nil
    var connectionType: ModbusConnectionTypeEntity? = nil
    // var connectionTypeId: String? = nil
    var functionCode: ModbusFunctionCodeEntity? = nil
    // var functionCodeId: String? = nil
    var register: Int? = nil
    var writeInterval: Int? = nil
    var readInterval: Int? = nil
    var range: Int? = nil
    var length: Int? = nil
    var variableType: ModbusVariableTypeEntity? = nil
    // var variableTypeId: String? = nil
}
class ModbusTaskEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var modbusId: Int? = nil
  @Published var modbusIdErrorMessage: Int? = nil
  @Published var register: Int? = nil
  @Published var registerErrorMessage: Int? = nil
  @Published var writeInterval: Int? = nil
  @Published var writeIntervalErrorMessage: Int? = nil
  @Published var readInterval: Int? = nil
  @Published var readIntervalErrorMessage: Int? = nil
  @Published var range: Int? = nil
  @Published var rangeErrorMessage: Int? = nil
  @Published var length: Int? = nil
  @Published var lengthErrorMessage: Int? = nil
  func getDto() -> ModbusTaskEntity {
      var dto = ModbusTaskEntity()
    dto.name = self.name
    dto.modbusId = self.modbusId
    dto.register = self.register
    dto.writeInterval = self.writeInterval
    dto.readInterval = self.readInterval
    dto.range = self.range
    dto.length = self.length
      return dto
  }
}