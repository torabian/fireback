import Foundation
class DeviceTypeEntity : Codable, Identifiable {
    var name: String? = nil
}
class DeviceTypeEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> DeviceTypeEntity {
      var dto = DeviceTypeEntity()
    dto.name = self.name
      return dto
  }
}