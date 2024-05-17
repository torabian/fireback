import Foundation
class UserEntity : Codable, Identifiable {
    var person: PersonEntity? = nil
    // var personId: String? = nil
    var avatar: String? = nil
}
class UserEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var avatar: String? = nil
  @Published var avatarErrorMessage: String? = nil
  func getDto() -> UserEntity {
      var dto = UserEntity()
    dto.avatar = self.avatar
      return dto
  }
}