import Foundation
class UserProfileEntity : Codable, Identifiable {
    var firstName: String? = nil
    var lastName: String? = nil
}
class UserProfileEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var firstName: String? = nil
  @Published var firstNameErrorMessage: String? = nil
  @Published var lastName: String? = nil
  @Published var lastNameErrorMessage: String? = nil
  func getDto() -> UserProfileEntity {
      var dto = UserProfileEntity()
    dto.firstName = self.firstName
    dto.lastName = self.lastName
      return dto
  }
}