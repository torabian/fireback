import Foundation
class ModbusFunctionCodeEntity : Codable, Identifiable {
    var name: String? = nil
    var code: Int? = nil
}
class ModbusFunctionCodeEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var code: Int? = nil
  @Published var codeErrorMessage: Int? = nil
  func getDto() -> ModbusFunctionCodeEntity {
      var dto = ModbusFunctionCodeEntity()
    dto.name = self.name
    dto.code = self.code
      return dto
  }
}