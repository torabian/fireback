import Foundation
class CommonProfileEntity : Codable, Identifiable {
    var firstName: String? = nil
    var lastName: String? = nil
    var phoneNumber: String? = nil
    var email: String? = nil
    var company: String? = nil
    var street: String? = nil
    var houseNumber: String? = nil
    var zipCode: String? = nil
    var city: String? = nil
    var gender: String? = nil
}
class CommonProfileEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var firstName: String? = nil
  @Published var firstNameErrorMessage: String? = nil
  @Published var lastName: String? = nil
  @Published var lastNameErrorMessage: String? = nil
  @Published var phoneNumber: String? = nil
  @Published var phoneNumberErrorMessage: String? = nil
  @Published var email: String? = nil
  @Published var emailErrorMessage: String? = nil
  @Published var company: String? = nil
  @Published var companyErrorMessage: String? = nil
  @Published var street: String? = nil
  @Published var streetErrorMessage: String? = nil
  @Published var houseNumber: String? = nil
  @Published var houseNumberErrorMessage: String? = nil
  @Published var zipCode: String? = nil
  @Published var zipCodeErrorMessage: String? = nil
  @Published var city: String? = nil
  @Published var cityErrorMessage: String? = nil
  @Published var gender: String? = nil
  @Published var genderErrorMessage: String? = nil
  func getDto() -> CommonProfileEntity {
      var dto = CommonProfileEntity()
    dto.firstName = self.firstName
    dto.lastName = self.lastName
    dto.phoneNumber = self.phoneNumber
    dto.email = self.email
    dto.company = self.company
    dto.street = self.street
    dto.houseNumber = self.houseNumber
    dto.zipCode = self.zipCode
    dto.city = self.city
    dto.gender = self.gender
      return dto
  }
}