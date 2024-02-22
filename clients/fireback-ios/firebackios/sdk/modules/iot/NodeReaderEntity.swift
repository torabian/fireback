import Foundation
class NodeReaderEntity : Codable, Identifiable {
    var name: String? = nil
    var nativeFn: String? = nil
    var config: String? = nil
}
class NodeReaderEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var nativeFn: String? = nil
  @Published var nativeFnErrorMessage: String? = nil
  @Published var config: String? = nil
  @Published var configErrorMessage: String? = nil
  func getDto() -> NodeReaderEntity {
      var dto = NodeReaderEntity()
    dto.name = self.name
    dto.nativeFn = self.nativeFn
    dto.config = self.config
      return dto
  }
}