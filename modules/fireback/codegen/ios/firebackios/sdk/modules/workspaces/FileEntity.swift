import Foundation
class FileEntity : Codable, Identifiable {
    var name: String? = nil
    var diskPath: String? = nil
    var size: Int? = nil
    var virtualPath: String? = nil
    var type: String? = nil
}
class FileEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var diskPath: String? = nil
  @Published var diskPathErrorMessage: String? = nil
  @Published var size: Int? = nil
  @Published var sizeErrorMessage: Int? = nil
  @Published var virtualPath: String? = nil
  @Published var virtualPathErrorMessage: String? = nil
  @Published var type: String? = nil
  @Published var typeErrorMessage: String? = nil
  func getDto() -> FileEntity {
      var dto = FileEntity()
    dto.name = self.name
    dto.diskPath = self.diskPath
    dto.size = self.size
    dto.virtualPath = self.virtualPath
    dto.type = self.type
      return dto
  }
}