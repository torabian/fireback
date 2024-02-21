import SwiftUI


struct OkayResponse : Codable {

}

struct ImportRequestDto : Codable {

}


struct IResponseErrorItem : Codable, Error {
    var domain: String?
    var reason: String?
    var message: String?
    var messageTranslated: String
    var location: String?
    var locationType: String?
    var extendedHelp: String?
    var sendReport: String?
}

struct IResponseError: Codable, Error {
    var code: Int?
    var message: String
    var messageTranslated: String
    var errors: [IResponseErrorItem]?
}

struct SingleResponse<T : Codable> : Codable {
    var data: T?
    var error: IResponseError?
}

struct ArrayResponse<T : Codable> : Codable {
    var data: ArrayResponseItems<T>?
}

struct ArrayResponseItems<T : Codable> : Codable {
    var items: [T]
}

extension Binding {
    func toUnwrapped<T>(defaultValue: T) -> Binding<T> where Value == Optional<T>  {
        Binding<T>(get: { self.wrappedValue ?? defaultValue }, set: { self.wrappedValue = $0 })
    }
}