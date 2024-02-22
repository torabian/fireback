import Foundation
class DataNodeModeEntity : Codable, Identifiable {
    var name: String? = nil
}
class DataNodeModeEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> DataNodeModeEntity {
      var dto = DataNodeModeEntity()
    dto.name = self.name
      return dto
  }
}