//
//  AuthService.swift
//  firebackios
//
//  Created by ali on 12/01/2024.
//

import Foundation

class AuthService : ObservableObject {
    static let shared = AuthService()

    public var TokenSnapShot = ""
    
    private let userDefaults = UserDefaults.standard
    private let authTokenKey = "authToken"
    private let sessionKey = "sessionKey"
    @Published var session: UserSessionDto?
    
    enum Screen {
        case signin
        case welcoome
    }
    @Published var screen: Screen = .welcoome

    func navigate(_ scr: Screen) {
        screen = scr
    }
    
    init() {
          session = getSessionFromUserDefaults()
        print("Session", session)
    }
    
    /*
     Use this function to get the session
     */
    private func getSessionFromUserDefaults() -> UserSessionDto? {
        let savedPerson = userDefaults.string(forKey: sessionKey)
        if savedPerson != nil {
            let content = savedPerson!.data(using: .utf8)
            if content == nil {
                return nil
            }
            print(20)
            print(content)
            let decoder = JSONDecoder()
            if let loadedPerson = try? decoder.decode(UserSessionDto.self, from: content!) {
                return loadedPerson
            }
        } else {
            print("Cannot get anything")
        }
        
        return nil
    }
 
    func setSession(dto: UserSessionDto?) {
        
        userDefaults.set(dto?.toJson(), forKey: sessionKey)
        session = dto
        
        if dto != nil {
            AuthService.shared.TokenSnapShot = dto!.token!
        }
    }
    
    var authToken: String? {
        get {
            return session!.token
        }
    }
     
    func signOut() {
        userDefaults.removeObject(forKey: sessionKey)
        session = nil
    }
}
