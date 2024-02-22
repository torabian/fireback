import Foundation
class DataNodeTypeEntity : Codable, Identifiable {
    var name: String? = nil
}
class DataNodeTypeEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> DataNodeTypeEntity {
      var dto = DataNodeTypeEntity()
    dto.name = self.name
      return dto
  }
}