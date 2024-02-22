import Foundation
class SpaceTimeSnapshotEntity : Codable, Identifiable {
    var lat: Float64? = nil
    var lng: Float64? = nil
    var alt: Float64? = nil
    var movableObject: MovableObjectEntity? = nil
    // var movableObjectId: String? = nil
}
class SpaceTimeSnapshotEntityViewModel: ObservableObject {
  // improve the fields here
  func getDto() -> SpaceTimeSnapshotEntity {
      var dto = SpaceTimeSnapshotEntity()
      return dto
  }
}