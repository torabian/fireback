import Foundation
class MemoryStatEntity : Codable, Identifiable {
    var heapSize: Int? = nil
}
class MemoryStatEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var heapSize: Int? = nil
  @Published var heapSizeErrorMessage: Int? = nil
  func getDto() -> MemoryStatEntity {
      var dto = MemoryStatEntity()
    dto.heapSize = self.heapSize
      return dto
  }
}