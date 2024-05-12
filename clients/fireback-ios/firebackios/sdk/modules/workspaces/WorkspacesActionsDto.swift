import Combine
import Foundation 
struct SendEmailActionReqDto : Codable {
    var toAddress: String? = nil
    var body: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class SendEmailActionReqDtoVm: ObservableObject {
  @Published var toAddress: String? = nil
  @Published var toAddressErrorMessage: String? = nil
  @Published var body: String? = nil
  @Published var bodyErrorMessage: String? = nil
    func getDto() -> SendEmailActionReqDto {
        var dto = SendEmailActionReqDto()
    dto.toAddress = self.toAddress
    dto.body = self.body
        return dto
    }
}
struct SendEmailActionResDto : Codable {
    var queueId: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
struct SendEmailWithProviderActionReqDto : Codable {
    var emailProvider: EmailProviderEntity? = nil
    // var emailProviderId: String? = nil
    var toAddress: String? = nil
    var body: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class SendEmailWithProviderActionReqDtoVm: ObservableObject {
  @Published var toAddress: String? = nil
  @Published var toAddressErrorMessage: String? = nil
  @Published var body: String? = nil
  @Published var bodyErrorMessage: String? = nil
    func getDto() -> SendEmailWithProviderActionReqDto {
        var dto = SendEmailWithProviderActionReqDto()
    dto.toAddress = self.toAddress
    dto.body = self.body
        return dto
    }
}
struct SendEmailWithProviderActionResDto : Codable {
    var queueId: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
struct GsmSendSmsActionReqDto : Codable {
    var toNumber: String? = nil
    var body: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class GsmSendSmsActionReqDtoVm: ObservableObject {
  @Published var toNumber: String? = nil
  @Published var toNumberErrorMessage: String? = nil
  @Published var body: String? = nil
  @Published var bodyErrorMessage: String? = nil
    func getDto() -> GsmSendSmsActionReqDto {
        var dto = GsmSendSmsActionReqDto()
    dto.toNumber = self.toNumber
    dto.body = self.body
        return dto
    }
}
struct GsmSendSmsActionResDto : Codable {
    var queueId: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
struct GsmSendSmsWithProviderActionReqDto : Codable {
    var gsmProvider: GsmProviderEntity? = nil
    // var gsmProviderId: String? = nil
    var toNumber: String? = nil
    var body: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class GsmSendSmsWithProviderActionReqDtoVm: ObservableObject {
  @Published var toNumber: String? = nil
  @Published var toNumberErrorMessage: String? = nil
  @Published var body: String? = nil
  @Published var bodyErrorMessage: String? = nil
    func getDto() -> GsmSendSmsWithProviderActionReqDto {
        var dto = GsmSendSmsWithProviderActionReqDto()
    dto.toNumber = self.toNumber
    dto.body = self.body
        return dto
    }
}
struct GsmSendSmsWithProviderActionResDto : Codable {
    var queueId: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
struct ClassicSigninActionReqDto : Codable {
    var value: String? = nil
    var password: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class ClassicSigninActionReqDtoVm: ObservableObject {
  @Published var value: String? = nil
  @Published var valueErrorMessage: String? = nil
  @Published var password: String? = nil
  @Published var passwordErrorMessage: String? = nil
    func getDto() -> ClassicSigninActionReqDto {
        var dto = ClassicSigninActionReqDto()
    dto.value = self.value
    dto.password = self.password
        return dto
    }
}
      enum ClassicSignupActionReqDtoType : Codable {
        phonenumber
        email
      }
struct ClassicSignupActionReqDto : Codable {
    var value: String? = nil
    var type: String? = nil
    var password: String? = nil
    var firstName: String? = nil
    var lastName: String? = nil
    var inviteId: String? = nil
    var publicJoinKeyId: String? = nil
    var workspaceTypeId: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class ClassicSignupActionReqDtoVm: ObservableObject {
  @Published var value: String? = nil
  @Published var valueErrorMessage: String? = nil
  @Published var password: String? = nil
  @Published var passwordErrorMessage: String? = nil
  @Published var firstName: String? = nil
  @Published var firstNameErrorMessage: String? = nil
  @Published var lastName: String? = nil
  @Published var lastNameErrorMessage: String? = nil
  @Published var inviteId: String? = nil
  @Published var inviteIdErrorMessage: String? = nil
  @Published var publicJoinKeyId: String? = nil
  @Published var publicJoinKeyIdErrorMessage: String? = nil
  @Published var workspaceTypeId: String? = nil
  @Published var workspaceTypeIdErrorMessage: String? = nil
    func getDto() -> ClassicSignupActionReqDto {
        var dto = ClassicSignupActionReqDto()
    dto.value = self.value
    dto.password = self.password
    dto.firstName = self.firstName
    dto.lastName = self.lastName
    dto.inviteId = self.inviteId
    dto.publicJoinKeyId = self.publicJoinKeyId
    dto.workspaceTypeId = self.workspaceTypeId
        return dto
    }
}
struct CreateWorkspaceActionReqDto : Codable {
    var name: String? = nil
    var workspace: WorkspaceEntity? = nil
    // var workspaceId: String? = nil
    var workspaceId: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class CreateWorkspaceActionReqDtoVm: ObservableObject {
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var workspaceId: String? = nil
  @Published var workspaceIdErrorMessage: String? = nil
    func getDto() -> CreateWorkspaceActionReqDto {
        var dto = CreateWorkspaceActionReqDto()
    dto.name = self.name
    dto.workspaceId = self.workspaceId
        return dto
    }
}
struct CheckClassicPassportActionReqDto : Codable {
    var value: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class CheckClassicPassportActionReqDtoVm: ObservableObject {
  @Published var value: String? = nil
  @Published var valueErrorMessage: String? = nil
    func getDto() -> CheckClassicPassportActionReqDto {
        var dto = CheckClassicPassportActionReqDto()
    dto.value = self.value
        return dto
    }
}
struct CheckClassicPassportActionResDto : Codable {
    var exists: Bool? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
struct ClassicPassportOtpActionReqDto : Codable {
    var value: String? = nil
    var otp: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class ClassicPassportOtpActionReqDtoVm: ObservableObject {
  @Published var value: String? = nil
  @Published var valueErrorMessage: String? = nil
  @Published var otp: String? = nil
  @Published var otpErrorMessage: String? = nil
    func getDto() -> ClassicPassportOtpActionReqDto {
        var dto = ClassicPassportOtpActionReqDto()
    dto.value = self.value
    dto.otp = self.otp
        return dto
    }
}
struct ClassicPassportOtpActionResDto : Codable {
    var suspendUntil: Int? = nil
    var session: UserSessionDto? = nil
    // var sessionId: String? = nil
    var validUntil: Int? = nil
    var blockedUntil: Int? = nil
    var secondsToUnblock: Int? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}