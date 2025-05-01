//
//  firebackiosApp.swift
//  firebackios
//
//  Created by ali on 13/11/2023.
//

import SwiftUI

@main
struct firebackiosApp: App {
    @StateObject private var authService = AuthService()

    var body: some Scene {
        WindowGroup {
            ContentView().environmentObject(authService)
        }
    }
}
