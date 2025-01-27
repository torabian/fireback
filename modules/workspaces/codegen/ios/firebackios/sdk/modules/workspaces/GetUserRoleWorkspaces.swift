import Promises
import Combine
import SwiftUI
func GetUserRoleWorkspacesFetcher() -> AnyPublisher<ArrayResponse<UserRoleWorkspaceEntity>, Error> {
    var computedUrl = "http://localhost:61901/user-role-workspaces"
    var request = URLRequest(url: URL(string: computedUrl)!)
    print("Token:", AuthService.shared.TokenSnapShot)
    request.addValue(AuthService.shared.TokenSnapShot, forHTTPHeaderField: "Authorization")
    request.addValue("root", forHTTPHeaderField: "workspace-id")
    return URLSession.shared
        .dataTaskPublisher(for: request)
        .map(\.data)
        .decode(type: ArrayResponse<UserRoleWorkspaceEntity>.self, decoder: JSONDecoder())
        .receive(on: DispatchQueue.main)
        .eraseToAnyPublisher()
}
/*
I have commented this for now, because this is not returning correctly, as well as the function above is good enough
func GetUserRoleWorkspaces() -> Promise<[UserRoleWorkspaceEntity]?> {
    return Promise<[UserRoleWorkspaceEntity]?>(on: .main) { fulfill, reject in
    var computedUrl = "http://localhost:61901/user-role-workspaces"
        var request = URLRequest(url: URL(string: computedUrl)!)
        request.httpMethod = "GET"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        URLSession.shared.dataTask(with: request) { data, response, error in
            if let data = data {
                let decoder = JSONDecoder()
                if let str = String(bytes: data, encoding: .utf8) {
                    print(str)
                }
                do {
                    let result = try decoder.decode(SingleResponse<UserRoleWorkspaceEntity>.self, from: data)
                    fulfill(result.data)
                } catch {
                    print(error)
                }
            }
        }.resume()
    }
}
*/