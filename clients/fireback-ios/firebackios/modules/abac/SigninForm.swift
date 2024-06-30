//
//  SignupForm.swift
//  firebackios
//
//  Created by ali on 10/01/2024.
//

import Foundation
import SwiftUI
import Combine
 
struct SigninForm: View {
    @ObservedObject var viewModel = EmailAccountSigninDtoViewModel()
    @EnvironmentObject private var authService: AuthService
    @State private var birthDate = Date()
    @State private var isSecureEntry = true
 
    var body: some View {
   
        Form {
            TextField("Email address", text: $viewModel.email.toUnwrapped(defaultValue: ""))
            HStack {
                if isSecureEntry {
                    SecureField("Password", text: $viewModel.password.toUnwrapped(defaultValue: ""))
                } else {
                    TextField("Password", text: $viewModel.password.toUnwrapped(defaultValue: ""))
                }
                
                Button(action: {
                    // Show/hide password
                    isSecureEntry.toggle()
                }) {
                    Image(systemName: "eye")
                }
            }
            
            Button(action: {
                PostPassportSigninEmail(dto: viewModel.getDto()).then { res in
                    if res != nil && res?.token != nil {
                        authService.setSession(dto: res)
                    }
                }.catch{reject in
                    print(reject)
                }
                
            }) {
                HStack {
                    Spacer()
                    Text("Submit")
                        .foregroundColor(.white)
                    Spacer()
                }
            }
            .padding()
            .background(Color.blue)
            .cornerRadius(10)
        }.navigationBarTitle("Detail View", displayMode: .inline)

    }
}
