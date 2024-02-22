import Foundation
class ExpanderFunctionEntity : Codable, Identifiable {
    var name: String? = nil
    var nativeFn: String? = nil
}
class ExpanderFunctionEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var nativeFn: String? = nil
  @Published var nativeFnErrorMessage: String? = nil
  func getDto() -> ExpanderFunctionEntity {
      var dto = ExpanderFunctionEntity()
    dto.name = self.name
    dto.nativeFn = self.nativeFn
      return dto
  }
}