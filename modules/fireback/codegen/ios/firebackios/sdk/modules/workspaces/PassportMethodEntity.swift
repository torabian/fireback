import Foundation
class PassportMethodEntity : Codable, Identifiable {
    var name: String? = nil
    var type: String? = nil
    var region: String? = nil
}
class PassportMethodEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var type: String? = nil
  @Published var typeErrorMessage: String? = nil
  @Published var region: String? = nil
  @Published var regionErrorMessage: String? = nil
  func getDto() -> PassportMethodEntity {
      var dto = PassportMethodEntity()
    dto.name = self.name
    dto.type = self.type
    dto.region = self.region
      return dto
  }
}