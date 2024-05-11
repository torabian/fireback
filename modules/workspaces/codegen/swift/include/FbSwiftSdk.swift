import SwiftUI


struct OkayResponse : Codable {

}
struct OkayResponseDto : Codable {

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

func handleFailedRequestError(error: Error) -> String {
    var errorDescription = error.localizedDescription
    
    if let nsError = error as NSError? {
        if nsError.domain == NSURLErrorDomain {
            if let urlError = nsError.userInfo[NSURLErrorKey] as? URLError {
                switch urlError.code {
                case .cancelled:
                    errorDescription = "Request was cancelled"
                case .timedOut:
                    errorDescription = "Request timed out"
                case .notConnectedToInternet:
                    errorDescription = "No internet connection"
                default:
                    errorDescription = "Unknown error: \(urlError.localizedDescription)"
                }
            }
        } else {
            errorDescription = "Error domain: \(nsError.domain)\nLocalized description: \(nsError.localizedDescription)"
        }
    } else {
        errorDescription = "Unknown error"
    }
    
    return errorDescription
}