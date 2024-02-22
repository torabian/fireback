import Foundation
class NodeWriterEntity : Codable, Identifiable {
    var name: String? = nil
    var nativeFn: String? = nil
    var config: String? = nil
}
class NodeWriterEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var nativeFn: String? = nil
  @Published var nativeFnErrorMessage: String? = nil
  func getDto() -> NodeWriterEntity {
      var dto = NodeWriterEntity()
    dto.name = self.name
    dto.nativeFn = self.nativeFn
      return dto
  }
}